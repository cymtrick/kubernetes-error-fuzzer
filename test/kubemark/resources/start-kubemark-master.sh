#!/bin/bash

# Copyright 2015 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Script that starts kubelet on kubemark-master as a supervisord process
# and then runs the master components as pods using kubelet.

# Define key path variables.
KUBE_ROOT="/home/kubernetes"
KUBE_BINDIR="${KUBE_ROOT}/kubernetes/server/bin"

function config-ip-firewall {
  echo "Configuring IP firewall rules"
  # The GCI image has host firewall which drop most inbound/forwarded packets.
  # We need to add rules to accept all TCP/UDP/ICMP packets.
  if iptables -L INPUT | grep "Chain INPUT (policy DROP)" > /dev/null; then
    echo "Add rules to accept all inbound TCP/UDP/ICMP packets"
    iptables -A INPUT -w -p TCP -j ACCEPT
    iptables -A INPUT -w -p UDP -j ACCEPT
    iptables -A INPUT -w -p ICMP -j ACCEPT
  fi
  if iptables -L FORWARD | grep "Chain FORWARD (policy DROP)" > /dev/null; then
    echo "Add rules to accept all forwarded TCP/UDP/ICMP packets"
    iptables -A FORWARD -w -p TCP -j ACCEPT
    iptables -A FORWARD -w -p UDP -j ACCEPT
    iptables -A FORWARD -w -p ICMP -j ACCEPT
  fi
}

function create-dirs {
	echo "Creating required directories"
	mkdir -p /var/lib/kubelet
	mkdir -p /etc/kubernetes/manifests
	mkdir -p /etc/kubernetes/addons
}

# Setup working directory for kubelet.
function setup-kubelet-dir {
	echo "Making /var/lib/kubelet executable for kubelet"
	mount -B /var/lib/kubelet /var/lib/kubelet/
	mount -B -o remount,exec,suid,dev /var/lib/kubelet
}

# Remove any default etcd config dirs/files.
function delete-default-etcd-configs {
	if [[ -d /etc/etcd ]]; then
		rm -rf /etc/etcd
	fi
	if [[ -e /etc/default/etcd ]]; then
		rm -f /etc/default/etcd
	fi
	if [[ -e /etc/systemd/system/etcd.service ]]; then
		rm -f /etc/systemd/system/etcd.service
	fi
	if [[ -e /etc/init.d/etcd ]]; then
		rm -f /etc/init.d/etcd
	fi
}

# Compute etcd related variables.
function compute-etcd-variables {
	ETCD_IMAGE="${ETCD_IMAGE:-}"
	ETCD_QUOTA_BYTES=""
	if [ "${ETCD_VERSION:0:2}" == "3." ]; then
		# TODO: Set larger quota to see if that helps with
		# 'mvcc: database space exceeded' errors. If so, pipe
		# though our setup scripts.
		ETCD_QUOTA_BYTES=" --quota-backend-bytes=4294967296 "
	fi
}

# Formats the given device ($1) if needed and mounts it at given mount point
# ($2).
function safe-format-and-mount() {
	device=$1
	mountpoint=$2

	# Format only if the disk is not already formatted.
	if ! tune2fs -l "${device}" ; then
		echo "Formatting '${device}'"
		mkfs.ext4 -F -E lazy_itable_init=0,lazy_journal_init=0,discard "${device}"
	fi

	mkdir -p "${mountpoint}"
	echo "Mounting '${device}' at '${mountpoint}'"
	mount -o discard,defaults "${device}" "${mountpoint}"
}

# Finds a PD device with name '$1' attached to the master.
function find-attached-pd() {
	local -r pd_name=$1
	if [[ ! -e /dev/disk/by-id/${pd_name} ]]; then
		echo ""
	fi
	device_info=$(ls -l /dev/disk/by-id/${pd_name})
	relative_path=${device_info##* }
	echo "/dev/disk/by-id/${relative_path}"
}

# Mounts a persistent disk (formatting if needed) to store the persistent data
# on the master. safe-format-and-mount only formats an unformatted disk, and
# mkdir -p will leave a directory be if it already exists.
function mount-pd() {
	local -r pd_name=$1
	local -r mount_point=$2

	if [[ -z "${find-attached-pd ${pd_name}}" ]]; then
		echo "Can't find ${pd_name}. Skipping mount."
		return
	fi

	local -r pd_path="/dev/disk/by-id/${pd_name}"
	echo "Mounting PD '${pd_path}' at '${mount_point}'"
	# Format and mount the disk, create directories on it for all of the master's
	# persistent data, and link them to where they're used.
	mkdir -p "${mount_point}"
	safe-format-and-mount "${pd_path}" "${mount_point}"
	echo "Mounted PD '${pd_path}' at '${mount_point}'"

	# NOTE: These locations on the PD store persistent data, so to maintain
	# upgradeability, these locations should not change.  If they do, take care
	# to maintain a migration path from these locations to whatever new
	# locations.
}

# Create kubeconfig for controller-manager's service account authentication.
function create-kubecontrollermanager-kubeconfig {
	echo "Creating kube-controller-manager kubeconfig file"
	mkdir -p /etc/srv/kubernetes/kube-controller-manager
	cat <<EOF >/etc/srv/kubernetes/kube-controller-manager/kubeconfig
apiVersion: v1
kind: Config
users:
- name: kube-controller-manager
  user:
    token: ${KUBE_CONTROLLER_MANAGER_TOKEN}
clusters:
- name: local
  cluster:
    insecure-skip-tls-verify: true
    server: https://localhost:443
contexts:
- context:
    cluster: local
    user: kube-controller-manager
  name: service-account-context
current-context: service-account-context
EOF
}

function assemble-docker-flags {
	echo "Assemble docker command line flags"
	local docker_opts="-p /var/run/docker.pid --iptables=false --ip-masq=false"
	docker_opts+=" --log-level=debug"  # Since it's a test cluster
	# TODO(shyamjvs): Incorporate network plugin options, etc later.
	echo "DOCKER_OPTS=\"${docker_opts}\"" > /etc/default/docker
	echo "DOCKER_NOFILE=65536" >> /etc/default/docker  # For setting ulimit -n
	systemctl restart docker
}

# A helper function for loading a docker image. It keeps trying up to 5 times.
#
# $1: Full path of the docker image
function try-load-docker-image {
	local -r img=$1
	echo "Try to load docker image file ${img}"
	# Temporarily turn off errexit, because we don't want to exit on first failure.
	set +e
	local -r max_attempts=5
	local -i attempt_num=1
	until timeout 30 docker load -i "${img}"; do
		if [[ "${attempt_num}" == "${max_attempts}" ]]; then
			echo "Fail to load docker image file ${img} after ${max_attempts} retries. Exit!!"
			exit 1
		else
			attempt_num=$((attempt_num+1))
			sleep 5
		fi
	done
	# Re-enable errexit.
	set -e
}

# Loads kube-system docker images. It is better to do it before starting kubelet,
# as kubelet will restart docker daemon, which may interfere with loading images.
function load-docker-images {
	echo "Start loading kube-system docker images"
	local -r img_dir="${KUBE_BINDIR}"
	try-load-docker-image "${img_dir}/kube-apiserver.tar"
	try-load-docker-image "${img_dir}/kube-controller-manager.tar"
	try-load-docker-image "${img_dir}/kube-scheduler.tar"
}

# Computes command line arguments to be passed to kubelet.
function compute-kubelet-params {
	local params="${KUBELET_TEST_ARGS:-}"
	params+=" --allow-privileged=true"
	params+=" --babysit-daemons=true"
	params+=" --cgroup-root=/"
	params+=" --cloud-provider=gce"
	params+=" --config=/etc/kubernetes/manifests"
	if [[ -n "${KUBELET_PORT:-}" ]]; then
		params+=" --port=${KUBELET_PORT}"
	fi
	params+=" --enable-debugging-handlers=false"
	params+=" --hairpin-mode=none"
	echo "${params}"
}

# Creates the systemd config file for kubelet.service.
function create-kubelet-conf() {
	local -r kubelet_bin="$1"
	local -r kubelet_env_file="/etc/default/kubelet"
	local -r flags=$(compute-kubelet-params)
	echo "KUBELET_OPTS=\"${flags}\"" > "${kubelet_env_file}"

	# Write the systemd service file for kubelet.
	cat <<EOF >/etc/systemd/system/kubelet.service
[Unit]
Description=Kubermark kubelet
Requires=network-online.target
After=network-online.target

[Service]
Restart=always
RestartSec=10
EnvironmentFile=${kubelet_env_file}
ExecStart=${kubelet_bin} \$KUBELET_OPTS

[Install]
WantedBy=multi-user.target
EOF
}

# This function assembles the kubelet systemd service file and starts it using
# systemctl, on the kubemark master.
function start-kubelet {
	# Create systemd config.
	local -r kubelet_bin="/usr/bin/kubelet"
	create-kubelet-conf "${kubelet_bin}"

	# Flush iptables nat table
  	iptables -t nat -F || true

	# Start the kubelet service.
	systemctl start kubelet.service
}

# Create the log file and set its properties.
#
# $1 is the file to create.
function prepare-log-file {
	touch $1
	chmod 644 $1
	chown root:root $1
}

# A helper function for copying addon manifests and set dir/files
# permissions.
#
# $1: addon category under /etc/kubernetes
# $2: manifest source dir
function setup-addon-manifests {
  local -r src_dir="${KUBE_ROOT}/$2"
  local -r dst_dir="/etc/kubernetes/$1/$2"
  if [[ ! -d "${dst_dir}" ]]; then
    mkdir -p "${dst_dir}"
  fi
  local files=$(find "${src_dir}" -maxdepth 1 -name "*.yaml")
  if [[ -n "${files}" ]]; then
    cp "${src_dir}/"*.yaml "${dst_dir}"
  fi
  chown -R root:root "${dst_dir}"
  chmod 755 "${dst_dir}"
  chmod 644 "${dst_dir}"/*
}

# Computes command line arguments to be passed to etcd.
function compute-etcd-params {
	local params="${ETCD_TEST_ARGS:-}"
	params+=" --listen-peer-urls=http://127.0.0.1:2380"
	params+=" --advertise-client-urls=http://127.0.0.1:2379"
	params+=" --listen-client-urls=http://0.0.0.0:2379"
	params+=" --data-dir=/var/etcd/data"
	params+=" ${ETCD_QUOTA_BYTES}"
	echo "${params}"
}

# Computes command line arguments to be passed to etcd-events.
function compute-etcd-events-params {
	local params="${ETCD_TEST_ARGS:-}"
	params+=" --listen-peer-urls=http://127.0.0.1:2381"
	params+=" --advertise-client-urls=http://127.0.0.1:4002"
	params+=" --listen-client-urls=http://0.0.0.0:4002"
	params+=" --data-dir=/var/etcd/data-events"
	params+=" ${ETCD_QUOTA_BYTES}"
	echo "${params}"
}

# Computes command line arguments to be passed to apiserver.
function compute-kube-apiserver-params {
	local params="${APISERVER_TEST_ARGS:-}"
	params+=" --insecure-bind-address=0.0.0.0"
	params+=" --etcd-servers=http://127.0.0.1:2379"
	params+=" --etcd-servers-overrides=/events#${EVENT_STORE_URL}"
	params+=" --tls-cert-file=/etc/srv/kubernetes/server.cert"
	params+=" --tls-private-key-file=/etc/srv/kubernetes/server.key"
	params+=" --client-ca-file=/etc/srv/kubernetes/ca.crt"
	params+=" --token-auth-file=/etc/srv/kubernetes/known_tokens.csv"
	params+=" --secure-port=443"
	params+=" --basic-auth-file=/etc/srv/kubernetes/basic_auth.csv"
	params+=" --target-ram-mb=$((${NUM_NODES} * 60))"
	params+=" --storage-backend=${STORAGE_BACKEND}"
	params+=" --service-cluster-ip-range=${SERVICE_CLUSTER_IP_RANGE}"
	params+=" --admission-control=${CUSTOM_ADMISSION_PLUGINS}"
	params+=" --authorization-mode=RBAC"
	echo "${params}"
}

# Computes command line arguments to be passed to controller-manager.
function compute-kube-controller-manager-params {
	local params="${CONTROLLER_MANAGER_TEST_ARGS:-}"
	params+=" --use-service-account-credentials"
	params+=" --kubeconfig=/etc/srv/kubernetes/kube-controller-manager/kubeconfig"
	params+=" --service-account-private-key-file=/etc/srv/kubernetes/server.key"
	params+=" --root-ca-file=/etc/srv/kubernetes/ca.crt"
	params+=" --allocate-node-cidrs=${ALLOCATE_NODE_CIDRS}"
	params+=" --cluster-cidr=${CLUSTER_IP_RANGE}"
	params+=" --service-cluster-ip-range=${SERVICE_CLUSTER_IP_RANGE}"
	params+=" --terminated-pod-gc-threshold=${TERMINATED_POD_GC_THRESHOLD}"
	echo "${params}"
}

# Computes command line arguments to be passed to scheduler.
function compute-kube-scheduler-params {
	local params="${SCHEDULER_TEST_ARGS:-}"
	params+=" --master=127.0.0.1:8080"
	echo "${params}"
}

# Computes command line arguments to be passed to addon-manager.
function compute-kube-addon-manager-params {
	echo ""
}

# Start a kubernetes master component '$1' which can be any of the following:
# 1. etcd
# 2. etcd-events
# 3. kube-apiserver
# 4. kube-controller-manager
# 5. kube-scheduler
# 6. kube-addon-manager
#
# It prepares the log file, loads the docker tag, calculates variables, sets them
# in the manifest file, and then copies the manifest file to /etc/kubernetes/manifests.
#
# Assumed vars:
#   DOCKER_REGISTRY
function start-kubemaster-component() {
	echo "Start master component $1"
	local -r component=$1
	prepare-log-file /var/log/"${component}".log
	local -r src_file="${KUBE_ROOT}/${component}.yaml"
	local -r params=$(compute-${component}-params)

	# Evaluate variables.
	sed -i -e "s@{{params}}@${params}@g" "${src_file}"
	sed -i -e "s@{{kube_docker_registry}}@${DOCKER_REGISTRY}@g" "${src_file}"
	sed -i -e "s@{{instance_prefix}}@${INSTANCE_PREFIX}@g" "${src_file}"
	if [ "${component:0:4}" == "etcd" ]; then
		sed -i -e "s@{{etcd_image}}@${ETCD_IMAGE}@g" "${src_file}"
	elif [ "${component}" == "kube-addon-manager" ]; then
		setup-addon-manifests "addons" "kubemark-rbac-bindings"
	else
		local -r component_docker_tag=$(cat ${KUBE_BINDIR}/${component}.docker_tag)
		sed -i -e "s@{{${component}_docker_tag}}@${component_docker_tag}@g" "${src_file}"
	fi
	cp "${src_file}" /etc/kubernetes/manifests
}

############################### Main Function ########################################
echo "Start to configure master instance for kubemark"

# Extract files from the server tar and setup master env variables.
cd "${KUBE_ROOT}"
tar xzf kubernetes-server-linux-amd64.tar.gz
source "${KUBE_ROOT}/kubemark-master-env.sh"

# Setup IP firewall rules, required directory structure and etcd config.
config-ip-firewall
create-dirs
setup-kubelet-dir
delete-default-etcd-configs
compute-etcd-variables

# Setup authentication token and kubeconfig for controller-manager.
KUBE_CONTROLLER_MANAGER_TOKEN=$(dd if=/dev/urandom bs=128 count=1 2>/dev/null | base64 | tr -d "=+/" | dd bs=32 count=1 2>/dev/null)
echo "${KUBE_CONTROLLER_MANAGER_TOKEN},system:kube-controller-manager,uid:system:kube-controller-manager" >> /etc/srv/kubernetes/known_tokens.csv
create-kubecontrollermanager-kubeconfig

# Mount master PD for etcd and create symbolic links to it.
{
	main_etcd_mount_point="/mnt/disks/master-pd"
	mount-pd "google-master-pd" "${main_etcd_mount_point}"
	# Contains all the data stored in etcd.
	mkdir -m 700 -p "${main_etcd_mount_point}/var/etcd"
	ln -s -f "${main_etcd_mount_point}/var/etcd" /var/etcd
	mkdir -p /etc/srv
	# Contains the dynamically generated apiserver auth certs and keys.
	mkdir -p "${main_etcd_mount_point}/srv/kubernetes"
	ln -s -f "${main_etcd_mount_point}/srv/kubernetes" /etc/srv/kubernetes
	# Directory for kube-apiserver to store SSH key (if necessary).
	mkdir -p "${main_etcd_mount_point}/srv/sshproxy"
	ln -s -f "${main_etcd_mount_point}/srv/sshproxy" /etc/srv/sshproxy
}

# Mount master PD for event-etcd (if required) and create symbolic links to it.
{
	EVENT_STORE_IP="${EVENT_STORE_IP:-127.0.0.1}"
	EVENT_STORE_URL="${EVENT_STORE_URL:-http://${EVENT_STORE_IP}:4002}"
	EVENT_PD="${EVENT_PD:-false}"
	if [ "${EVENT_PD:-false}" == "true" ]; then
		event_etcd_mount_point="/mnt/disks/master-event-pd"
		mount-pd "google-master-event-pd" "${event_etcd_mount_point}"
		# Contains all the data stored in event etcd.
		mkdir -m 700 -p "${event_etcd_mount_point}/var/etcd/events"
		ln -s -f "${event_etcd_mount_point}/var/etcd/events" /var/etcd/events
	fi
}

# Setup docker flags and load images of the master components.
assemble-docker-flags
DOCKER_REGISTRY="gcr.io/google_containers"
load-docker-images

# Start kubelet as a supervisord process and master components as pods.
start-kubelet
start-kubemaster-component "etcd"
if [ "${EVENT_STORE_IP:-}" == "127.0.0.1" ]; then
	start-kubemaster-component "etcd-events"
fi
start-kubemaster-component "kube-apiserver"
start-kubemaster-component "kube-controller-manager"
start-kubemaster-component "kube-scheduler"
start-kubemaster-component "kube-addon-manager"

# Wait till apiserver is working fine.
until [ "$(curl 127.0.0.1:8080/healthz 2> /dev/null)" == "ok" ]; do
	sleep 1
done

echo "Done for the configuration for kubermark master"
