/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storage

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/storage/utils"
	imageutils "k8s.io/kubernetes/test/utils/image"
)

type localTestConfig struct {
	ns           string
	nodes        []v1.Node
	nodeExecPods map[string]*v1.Pod
	node0        *v1.Node
	client       clientset.Interface
	scName       string
	ssTester     *framework.StatefulSetTester
	discoveryDir string
}

type localVolumeType string

const (
	// default local volume type, aka a directory
	DirectoryLocalVolumeType localVolumeType = "dir"
	// like DirectoryLocalVolumeType but it's a symbolic link to directory
	DirectoryLinkLocalVolumeType localVolumeType = "dir-link"
	// like DirectoryLocalVolumeType but bind mounted
	DirectoryBindMountedLocalVolumeType localVolumeType = "dir-bindmounted"
	// like DirectoryLocalVolumeType but it's a symbolic link to self bind mounted directory
	// Note that bind mounting at symbolic link actually mounts at directory it
	// links to.
	DirectoryLinkBindMountedLocalVolumeType localVolumeType = "dir-link-bindmounted"
	// creates a tmpfs and mounts it
	TmpfsLocalVolumeType localVolumeType = "tmpfs"
	// tests based on local ssd at /mnt/disks/by-uuid/
	GCELocalSSDVolumeType localVolumeType = "gce-localssd-scsi-fs"
	// Creates a local file, formats it, and maps it as a block device.
	BlockLocalVolumeType localVolumeType = "block"
	// Creates a local file serving as the backing for block device., formats it,
	// and mounts it to use as FS mode local volume.
	BlockFsWithFormatLocalVolumeType localVolumeType = "blockfswithformat"
	// Creates a local file serving as the backing for block device. do not format it manually,
	// and mounts it to use as FS mode local volume.
	BlockFsWithoutFormatLocalVolumeType localVolumeType = "blockfswithoutformat"
)

var setupLocalVolumeMap = map[localVolumeType]func(*localTestConfig, *v1.Node) *localTestVolume{
	GCELocalSSDVolumeType:                   setupLocalVolumeGCELocalSSD,
	TmpfsLocalVolumeType:                    setupLocalVolumeTmpfs,
	DirectoryLocalVolumeType:                setupLocalVolumeDirectory,
	DirectoryLinkLocalVolumeType:            setupLocalVolumeDirectoryLink,
	DirectoryBindMountedLocalVolumeType:     setupLocalVolumeDirectoryBindMounted,
	DirectoryLinkBindMountedLocalVolumeType: setupLocalVolumeDirectoryLinkBindMounted,
	BlockLocalVolumeType:                    setupLocalVolumeBlock,
	BlockFsWithFormatLocalVolumeType:        setupLocalVolumeBlockFsWithFormat,
	BlockFsWithoutFormatLocalVolumeType:     setupLocalVolumeBlockFsWithoutFormat,
}

var cleanupLocalVolumeMap = map[localVolumeType]func(*localTestConfig, *localTestVolume){
	GCELocalSSDVolumeType:                   cleanupLocalVolumeGCELocalSSD,
	TmpfsLocalVolumeType:                    cleanupLocalVolumeTmpfs,
	DirectoryLocalVolumeType:                cleanupLocalVolumeDirectory,
	DirectoryLinkLocalVolumeType:            cleanupLocalVolumeDirectoryLink,
	DirectoryBindMountedLocalVolumeType:     cleanupLocalVolumeDirectoryBindMounted,
	DirectoryLinkBindMountedLocalVolumeType: cleanupLocalVolumeDirectoryLinkBindMounted,
	BlockLocalVolumeType:                    cleanupLocalVolumeBlock,
	BlockFsWithFormatLocalVolumeType:        cleanupLocalVolumeBlockFsWithFormat,
	BlockFsWithoutFormatLocalVolumeType:     cleanupLocalVolumeBlockFsWithoutFormat,
}

type localTestVolume struct {
	// Node that the volume is on
	node *v1.Node
	// Path to the volume on the host node
	hostDir string
	// PVC for this volume
	pvc *v1.PersistentVolumeClaim
	// PV for this volume
	pv *v1.PersistentVolume
	// Type of local volume
	localVolumeType localVolumeType
	// Path to the loop block device on the host node.
	// Used during cleanup after block tests.
	loopDevDir string
}

const (
	// TODO: This may not be available/writable on all images.
	hostBase = "/tmp"
	// Path to the first volume in the test containers
	// created via createLocalPod or makeLocalPod
	// leveraging pv_util.MakePod
	volumeDir = "/mnt/volume1"
	// testFile created in setupLocalVolume
	testFile = "test-file"
	// testFileContent written into testFile
	testFileContent = "test-file-content"
	testSCPrefix    = "local-volume-test-storageclass"

	// A sample request size
	testRequestSize = "10Mi"

	// Max number of nodes to use for testing
	maxNodes = 5
)

var (
	// storage class volume binding modes
	waitMode      = storagev1.VolumeBindingWaitForFirstConsumer
	immediateMode = storagev1.VolumeBindingImmediate

	// Common selinux labels
	selinuxLabel = &v1.SELinuxOptions{
		Level: "s0:c0,c1"}
)

var _ = utils.SIGDescribe("PersistentVolumes-local ", func() {
	f := framework.NewDefaultFramework("persistent-local-volumes-test")

	var (
		config *localTestConfig
		scName string
	)

	BeforeEach(func() {
		// Get all the schedulable nodes
		nodes := framework.GetReadySchedulableNodesOrDie(f.ClientSet)
		Expect(len(nodes.Items)).NotTo(BeZero(), "No available nodes for scheduling")

		// Cap max number of nodes
		maxLen := len(nodes.Items)
		if maxLen > maxNodes {
			maxLen = maxNodes
		}

		scName = fmt.Sprintf("%v-%v", testSCPrefix, f.Namespace.Name)
		// Choose the first node
		node0 := &nodes.Items[0]

		ssTester := framework.NewStatefulSetTester(f.ClientSet)
		config = &localTestConfig{
			ns:           f.Namespace.Name,
			client:       f.ClientSet,
			nodes:        nodes.Items[:maxLen],
			nodeExecPods: make(map[string]*v1.Pod, maxLen),
			node0:        node0,
			scName:       scName,
			ssTester:     ssTester,
			discoveryDir: filepath.Join(hostBase, f.Namespace.Name),
		}
	})

	for tempTestVolType := range setupLocalVolumeMap {

		// New variable required for gingko test closures
		testVolType := tempTestVolType
		serialStr := ""
		if testVolType == GCELocalSSDVolumeType {
			serialStr = " [Serial]"
		}
		ctxString := fmt.Sprintf("[Volume type: %s]%v", testVolType, serialStr)
		testMode := immediateMode

		Context(ctxString, func() {
			var testVol *localTestVolume

			BeforeEach(func() {
				if testVolType == GCELocalSSDVolumeType {
					SkipUnlessLocalSSDExists(config, "scsi", "fs", config.node0)
				}
				setupStorageClass(config, &testMode)
				testVols := setupLocalVolumesPVCsPVs(config, testVolType, config.node0, 1, testMode)
				testVol = testVols[0]
			})

			AfterEach(func() {
				cleanupLocalVolumes(config, []*localTestVolume{testVol})
				cleanupStorageClass(config)
			})

			Context("One pod requesting one prebound PVC", func() {
				var (
					pod1    *v1.Pod
					pod1Err error
				)

				BeforeEach(func() {
					By("Creating pod1")
					pod1, pod1Err = createLocalPod(config, testVol, nil)
					Expect(pod1Err).NotTo(HaveOccurred())
					verifyLocalPod(config, testVol, pod1, config.node0.Name)

					writeCmd := createWriteCmd(volumeDir, testFile, testFileContent, testVol.localVolumeType)

					By("Writing in pod1")
					podRWCmdExec(pod1, writeCmd)
				})

				AfterEach(func() {
					By("Deleting pod1")
					framework.DeletePodOrFail(config.client, config.ns, pod1.Name)
				})

				It("should be able to mount volume and read from pod1", func() {
					By("Reading in pod1")
					// testFileContent was written in BeforeEach
					testReadFileContent(volumeDir, testFile, testFileContent, pod1, testVolType)
				})

				It("should be able to mount volume and write from pod1", func() {
					// testFileContent was written in BeforeEach
					testReadFileContent(volumeDir, testFile, testFileContent, pod1, testVolType)

					By("Writing in pod1")
					writeCmd := createWriteCmd(volumeDir, testFile, testVol.hostDir /*writeTestFileContent*/, testVolType)
					podRWCmdExec(pod1, writeCmd)
				})
			})

			Context("Two pods mounting a local volume at the same time", func() {
				It("should be able to write from pod1 and read from pod2", func() {
					twoPodsReadWriteTest(config, testVol)
				})
			})

			Context("Two pods mounting a local volume one after the other", func() {
				It("should be able to write from pod1 and read from pod2", func() {
					twoPodsReadWriteSerialTest(config, testVol)
				})
			})

			Context("Set fsGroup for local volume", func() {
				BeforeEach(func() {
					if testVolType == BlockLocalVolumeType {
						framework.Skipf("We don't set fsGroup on block device, skipped.")
					}
				})

				It("should set fsGroup for one pod", func() {
					By("Checking fsGroup is set")
					pod := createPodWithFsGroupTest(config, testVol, 1234, 1234)
					By("Deleting pod")
					framework.DeletePodOrFail(config.client, config.ns, pod.Name)
				})

				It("should set same fsGroup for two pods simultaneously", func() {
					fsGroup := int64(1234)
					By("Create first pod and check fsGroup is set")
					pod1 := createPodWithFsGroupTest(config, testVol, fsGroup, fsGroup)
					By("Create second pod with same fsGroup and check fsGroup is correct")
					pod2 := createPodWithFsGroupTest(config, testVol, fsGroup, fsGroup)
					By("Deleting first pod")
					framework.DeletePodOrFail(config.client, config.ns, pod1.Name)
					By("Deleting second pod")
					framework.DeletePodOrFail(config.client, config.ns, pod2.Name)
				})

				It("should set different fsGroup for second pod if first pod is deleted", func() {
					fsGroup1, fsGroup2 := int64(1234), int64(4321)
					By("Create first pod and check fsGroup is set")
					pod1 := createPodWithFsGroupTest(config, testVol, fsGroup1, fsGroup1)
					By("Deleting first pod")
					err := framework.DeletePodWithWait(f, config.client, pod1)
					Expect(err).NotTo(HaveOccurred(), "while deleting first pod")
					By("Create second pod and check fsGroup is the new one")
					pod2 := createPodWithFsGroupTest(config, testVol, fsGroup2, fsGroup2)
					By("Deleting second pod")
					framework.DeletePodOrFail(config.client, config.ns, pod2.Name)
				})
			})

		})
	}

	Context("Local volume that cannot be mounted [Slow]", func() {
		// TODO:
		// - check for these errors in unit tests instead
		It("should fail due to non-existent path", func() {
			testVol := &localTestVolume{
				node:            config.node0,
				hostDir:         "/non-existent/location/nowhere",
				localVolumeType: DirectoryLocalVolumeType,
			}
			By("Creating local PVC and PV")
			createLocalPVCsPVs(config, []*localTestVolume{testVol}, immediateMode)
			pod, err := createLocalPod(config, testVol, nil)
			Expect(err).To(HaveOccurred())
			err = framework.WaitTimeoutForPodRunningInNamespace(config.client, pod.Name, pod.Namespace, framework.PodStartShortTimeout)
			Expect(err).To(HaveOccurred())
			cleanupLocalPVCsPVs(config, []*localTestVolume{testVol})
		})

		It("should fail due to wrong node", func() {
			if len(config.nodes) < 2 {
				framework.Skipf("Runs only when number of nodes >= 2")
			}

			testVols := setupLocalVolumesPVCsPVs(config, DirectoryLocalVolumeType, config.node0, 1, immediateMode)
			testVol := testVols[0]

			pod := makeLocalPodWithNodeName(config, testVol, config.nodes[1].Name)
			pod, err := config.client.CoreV1().Pods(config.ns).Create(pod)
			Expect(err).NotTo(HaveOccurred())

			err = framework.WaitTimeoutForPodRunningInNamespace(config.client, pod.Name, pod.Namespace, framework.PodStartShortTimeout)
			Expect(err).To(HaveOccurred())

			cleanupLocalVolumes(config, []*localTestVolume{testVol})
		})
	})

	Context("Pod with node different from PV's NodeAffinity", func() {
		var (
			testVol    *localTestVolume
			volumeType localVolumeType
		)

		BeforeEach(func() {
			if len(config.nodes) < 2 {
				framework.Skipf("Runs only when number of nodes >= 2")
			}

			volumeType = DirectoryLocalVolumeType
			setupStorageClass(config, &immediateMode)
			testVols := setupLocalVolumesPVCsPVs(config, volumeType, config.node0, 1, immediateMode)
			testVol = testVols[0]
		})

		AfterEach(func() {
			cleanupLocalVolumes(config, []*localTestVolume{testVol})
			cleanupStorageClass(config)
		})

		It("should fail scheduling due to different NodeAffinity", func() {
			testPodWithNodeConflict(config, volumeType, config.nodes[1].Name, makeLocalPodWithNodeAffinity, immediateMode)
		})

		It("should fail scheduling due to different NodeSelector", func() {
			testPodWithNodeConflict(config, volumeType, config.nodes[1].Name, makeLocalPodWithNodeSelector, immediateMode)
		})
	})

	Context("StatefulSet with pod affinity [Slow]", func() {
		var testVols map[string][]*localTestVolume
		const (
			ssReplicas  = 3
			volsPerNode = 6
		)

		BeforeEach(func() {
			setupStorageClass(config, &waitMode)

			testVols = map[string][]*localTestVolume{}
			for i, node := range config.nodes {
				// The PVCs created here won't be used
				By(fmt.Sprintf("Setting up local volumes on node %q", node.Name))
				vols := setupLocalVolumesPVCsPVs(config, DirectoryLocalVolumeType, &config.nodes[i], volsPerNode, waitMode)
				testVols[node.Name] = vols
			}
		})

		AfterEach(func() {
			for _, vols := range testVols {
				cleanupLocalVolumes(config, vols)
			}
			cleanupStorageClass(config)
		})

		It("should use volumes spread across nodes when pod has anti-affinity", func() {
			if len(config.nodes) < ssReplicas {
				framework.Skipf("Runs only when number of nodes >= %v", ssReplicas)
			}
			By("Creating a StatefulSet with pod anti-affinity on nodes")
			ss := createStatefulSet(config, ssReplicas, volsPerNode, true, false)
			validateStatefulSet(config, ss, true)
		})

		It("should use volumes on one node when pod has affinity", func() {
			By("Creating a StatefulSet with pod affinity on nodes")
			ss := createStatefulSet(config, ssReplicas, volsPerNode/ssReplicas, false, false)
			validateStatefulSet(config, ss, false)
		})

		It("should use volumes spread across nodes when pod management is parallel and pod has anti-affinity", func() {
			if len(config.nodes) < ssReplicas {
				framework.Skipf("Runs only when number of nodes >= %v", ssReplicas)
			}
			By("Creating a StatefulSet with pod anti-affinity on nodes")
			ss := createStatefulSet(config, ssReplicas, 1, true, true)
			validateStatefulSet(config, ss, true)
		})

		It("should use volumes on one node when pod management is parallel and pod has affinity", func() {
			By("Creating a StatefulSet with pod affinity on nodes")
			ss := createStatefulSet(config, ssReplicas, 1, false, true)
			validateStatefulSet(config, ss, false)
		})
	})

})

type makeLocalPodWith func(config *localTestConfig, volume *localTestVolume, nodeName string) *v1.Pod

func testPodWithNodeConflict(config *localTestConfig, testVolType localVolumeType, nodeName string, makeLocalPodFunc makeLocalPodWith, bindingMode storagev1.VolumeBindingMode) {
	By(fmt.Sprintf("local-volume-type: %s", testVolType))
	testVols := setupLocalVolumesPVCsPVs(config, testVolType, config.node0, 1, bindingMode)
	testVol := testVols[0]

	pod := makeLocalPodFunc(config, testVol, nodeName)
	pod, err := config.client.CoreV1().Pods(config.ns).Create(pod)
	Expect(err).NotTo(HaveOccurred())

	err = framework.WaitForPodNameUnschedulableInNamespace(config.client, pod.Name, pod.Namespace)
	Expect(err).NotTo(HaveOccurred())
}

// The tests below are run against multiple mount point types

// Test two pods at the same time, write from pod1, and read from pod2
func twoPodsReadWriteTest(config *localTestConfig, testVol *localTestVolume) {
	By("Creating pod1 to write to the PV")
	pod1, pod1Err := createLocalPod(config, testVol, nil)
	Expect(pod1Err).NotTo(HaveOccurred())
	verifyLocalPod(config, testVol, pod1, config.node0.Name)

	writeCmd := createWriteCmd(volumeDir, testFile, testFileContent, testVol.localVolumeType)

	By("Writing in pod1")
	podRWCmdExec(pod1, writeCmd)

	// testFileContent was written after creating pod1
	testReadFileContent(volumeDir, testFile, testFileContent, pod1, testVol.localVolumeType)

	By("Creating pod2 to read from the PV")
	pod2, pod2Err := createLocalPod(config, testVol, nil)
	Expect(pod2Err).NotTo(HaveOccurred())
	verifyLocalPod(config, testVol, pod2, config.node0.Name)

	// testFileContent was written after creating pod1
	testReadFileContent(volumeDir, testFile, testFileContent, pod2, testVol.localVolumeType)

	writeCmd = createWriteCmd(volumeDir, testFile, testVol.hostDir /*writeTestFileContent*/, testVol.localVolumeType)

	By("Writing in pod2")
	podRWCmdExec(pod2, writeCmd)

	By("Reading in pod1")
	testReadFileContent(volumeDir, testFile, testVol.hostDir, pod1, testVol.localVolumeType)

	By("Deleting pod1")
	framework.DeletePodOrFail(config.client, config.ns, pod1.Name)
	By("Deleting pod2")
	framework.DeletePodOrFail(config.client, config.ns, pod2.Name)
}

// Test two pods one after other, write from pod1, and read from pod2
func twoPodsReadWriteSerialTest(config *localTestConfig, testVol *localTestVolume) {
	By("Creating pod1")
	pod1, pod1Err := createLocalPod(config, testVol, nil)
	Expect(pod1Err).NotTo(HaveOccurred())
	verifyLocalPod(config, testVol, pod1, config.node0.Name)

	writeCmd := createWriteCmd(volumeDir, testFile, testFileContent, testVol.localVolumeType)

	By("Writing in pod1")
	podRWCmdExec(pod1, writeCmd)

	// testFileContent was written after creating pod1
	testReadFileContent(volumeDir, testFile, testFileContent, pod1, testVol.localVolumeType)

	By("Deleting pod1")
	framework.DeletePodOrFail(config.client, config.ns, pod1.Name)

	By("Creating pod2")
	pod2, pod2Err := createLocalPod(config, testVol, nil)
	Expect(pod2Err).NotTo(HaveOccurred())
	verifyLocalPod(config, testVol, pod2, config.node0.Name)

	By("Reading in pod2")
	testReadFileContent(volumeDir, testFile, testFileContent, pod2, testVol.localVolumeType)

	By("Deleting pod2")
	framework.DeletePodOrFail(config.client, config.ns, pod2.Name)
}

// Test creating pod with fsGroup, and check fsGroup is expected fsGroup.
func createPodWithFsGroupTest(config *localTestConfig, testVol *localTestVolume, fsGroup int64, expectedFsGroup int64) *v1.Pod {
	pod, err := createLocalPod(config, testVol, &fsGroup)
	framework.ExpectNoError(err)
	_, err = framework.LookForStringInPodExec(config.ns, pod.Name, []string{"stat", "-c", "%g", volumeDir}, strconv.FormatInt(expectedFsGroup, 10), time.Second*3)
	Expect(err).NotTo(HaveOccurred(), "failed to get expected fsGroup %d on directory %s in pod %s", fsGroup, volumeDir, pod.Name)
	return pod
}

func setupStorageClass(config *localTestConfig, mode *storagev1.VolumeBindingMode) {
	sc := &storagev1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: config.scName,
		},
		Provisioner:       "kubernetes.io/no-provisioner",
		VolumeBindingMode: mode,
	}

	sc, err := config.client.StorageV1().StorageClasses().Create(sc)
	Expect(err).NotTo(HaveOccurred())
}

func cleanupStorageClass(config *localTestConfig) {
	framework.ExpectNoError(config.client.StorageV1().StorageClasses().Delete(config.scName, nil))
}

// podNode wraps RunKubectl to get node where pod is running
func podNodeName(config *localTestConfig, pod *v1.Pod) (string, error) {
	runtimePod, runtimePodErr := config.client.CoreV1().Pods(pod.Namespace).Get(pod.Name, metav1.GetOptions{})
	return runtimePod.Spec.NodeName, runtimePodErr
}

// setupLocalVolumes sets up directories to use for local PV
func setupLocalVolumes(config *localTestConfig, localVolumeType localVolumeType, node *v1.Node, count int) []*localTestVolume {
	vols := []*localTestVolume{}
	for i := 0; i < count; i++ {
		setupLocalVolume, ok := setupLocalVolumeMap[localVolumeType]
		Expect(ok).To(BeTrue())
		testVol := setupLocalVolume(config, node)
		vols = append(vols, testVol)
	}
	return vols
}

func cleanupLocalPVCsPVs(config *localTestConfig, volumes []*localTestVolume) {
	for _, volume := range volumes {
		By("Cleaning up PVC and PV")
		errs := framework.PVPVCCleanup(config.client, config.ns, volume.pv, volume.pvc)
		if len(errs) > 0 {
			framework.Failf("Failed to delete PV and/or PVC: %v", utilerrors.NewAggregate(errs))
		}
	}
}

// Deletes the PVC/PV, and launches a pod with hostpath volume to remove the test directory
func cleanupLocalVolumes(config *localTestConfig, volumes []*localTestVolume) {
	cleanupLocalPVCsPVs(config, volumes)

	for _, volume := range volumes {
		cleanup := cleanupLocalVolumeMap[volume.localVolumeType]
		cleanup(config, volume)
	}
}

func generateLocalTestVolume(hostDir string, config *localTestConfig, localVolumeType localVolumeType, node *v1.Node) *localTestVolume {
	if localVolumeType != BlockLocalVolumeType && localVolumeType != BlockFsWithoutFormatLocalVolumeType {
		mkdirCmd := fmt.Sprintf("mkdir -p %s", hostDir)
		err := issueNodeCommand(config, mkdirCmd, node)
		Expect(err).NotTo(HaveOccurred())
	}

	return &localTestVolume{
		node:            node,
		hostDir:         hostDir,
		localVolumeType: localVolumeType,
	}
}

func setupLocalVolumeTmpfs(config *localTestConfig, node *v1.Node) *localTestVolume {
	testDirName := "local-volume-test-" + string(uuid.NewUUID())
	hostDir := filepath.Join(hostBase, testDirName)
	createAndMountTmpfsLocalVolume(config, hostDir, node)
	return generateLocalTestVolume(hostDir, config, TmpfsLocalVolumeType, node)
}

func setupLocalVolumeGCELocalSSD(config *localTestConfig, node *v1.Node) *localTestVolume {
	res, err := issueNodeCommandWithResult(config, "ls /mnt/disks/by-uuid/google-local-ssds-scsi-fs/", node)
	Expect(err).NotTo(HaveOccurred())
	dirName := strings.Fields(res)[0]
	hostDir := "/mnt/disks/by-uuid/google-local-ssds-scsi-fs/" + dirName
	// gce local ssd does not need to create a directory
	return &localTestVolume{
		node:            node,
		hostDir:         hostDir,
		localVolumeType: GCELocalSSDVolumeType,
	}
}

func setupLocalVolumeDirectory(config *localTestConfig, node *v1.Node) *localTestVolume {
	testDirName := "local-volume-test-" + string(uuid.NewUUID())
	hostDir := filepath.Join(hostBase, testDirName)
	return generateLocalTestVolume(hostDir, config, DirectoryLocalVolumeType, node)
}

// launchNodeExecPodForLocalPV launches a hostexec pod for local PV and waits
// until it's Running.
func launchNodeExecPodForLocalPV(client clientset.Interface, ns, node string) *v1.Pod {
	hostExecPod := framework.NewHostExecPodSpec(ns, fmt.Sprintf("hostexec-%s", node))
	hostExecPod.Spec.NodeName = node
	hostExecPod.Spec.Volumes = []v1.Volume{
		{
			// Required to enter into host mount namespace via nsenter.
			Name: "rootfs",
			VolumeSource: v1.VolumeSource{
				HostPath: &v1.HostPathVolumeSource{
					Path: "/",
				},
			},
		},
	}
	hostExecPod.Spec.Containers[0].VolumeMounts = []v1.VolumeMount{
		{
			Name:      "rootfs",
			MountPath: "/rootfs",
			ReadOnly:  true,
		},
	}
	hostExecPod.Spec.Containers[0].SecurityContext = &v1.SecurityContext{
		Privileged: func(privileged bool) *bool {
			return &privileged
		}(true),
	}
	pod, err := client.CoreV1().Pods(ns).Create(hostExecPod)
	framework.ExpectNoError(err)
	err = framework.WaitForPodRunningInNamespace(client, pod)
	framework.ExpectNoError(err)
	return pod
}

// issueNodeCommandWithResult issues command on given node and returns stdout.
func issueNodeCommandWithResult(config *localTestConfig, cmd string, node *v1.Node) (string, error) {
	var pod *v1.Pod
	pod, ok := config.nodeExecPods[node.Name]
	if !ok {
		pod = launchNodeExecPodForLocalPV(config.client, config.ns, node.Name)
		if pod == nil {
			return "", fmt.Errorf("failed to create hostexec pod for node %q", node)
		}
		config.nodeExecPods[node.Name] = pod
	}
	args := []string{
		"exec",
		fmt.Sprintf("--namespace=%v", pod.Namespace),
		pod.Name,
		"--",
		"nsenter",
		"--mount=/rootfs/proc/1/ns/mnt",
		"--",
		"sh",
		"-c",
		cmd,
	}
	return framework.RunKubectl(args...)
}

// issueNodeCommand works like issueNodeCommandWithResult, but discards result.
func issueNodeCommand(config *localTestConfig, cmd string, node *v1.Node) error {
	_, err := issueNodeCommandWithResult(config, cmd, node)
	return err
}

func setupLocalVolumeDirectoryLink(config *localTestConfig, node *v1.Node) *localTestVolume {
	testDirName := "local-volume-test-" + string(uuid.NewUUID())
	hostDir := filepath.Join(hostBase, testDirName)
	hostDirBackend := hostDir + "-backend"
	cmd := fmt.Sprintf("mkdir %s && sudo ln -s %s %s", hostDirBackend, hostDirBackend, hostDir)
	_, err := issueNodeCommandWithResult(config, cmd, node)
	Expect(err).NotTo(HaveOccurred())
	return generateLocalTestVolume(hostDir, config, DirectoryLinkLocalVolumeType, node)
}

func setupLocalVolumeDirectoryBindMounted(config *localTestConfig, node *v1.Node) *localTestVolume {
	testDirName := "local-volume-test-" + string(uuid.NewUUID())
	hostDir := filepath.Join(hostBase, testDirName)
	cmd := fmt.Sprintf("mkdir %s && sudo mount --bind %s %s", hostDir, hostDir, hostDir)
	_, err := issueNodeCommandWithResult(config, cmd, node)
	Expect(err).NotTo(HaveOccurred())
	return generateLocalTestVolume(hostDir, config, DirectoryBindMountedLocalVolumeType, node)
}

func setupLocalVolumeDirectoryLinkBindMounted(config *localTestConfig, node *v1.Node) *localTestVolume {
	testDirName := "local-volume-test-" + string(uuid.NewUUID())
	hostDir := filepath.Join(hostBase, testDirName)
	hostDirBackend := hostDir + "-backend"
	cmd := fmt.Sprintf("mkdir %s && sudo mount --bind %s %s && sudo ln -s %s %s",
		hostDirBackend, hostDirBackend, hostDirBackend, hostDirBackend, hostDir)
	_, err := issueNodeCommandWithResult(config, cmd, node)
	Expect(err).NotTo(HaveOccurred())
	return generateLocalTestVolume(hostDir, config, DirectoryLinkBindMountedLocalVolumeType, node)
}

func setupLocalVolumeBlock(config *localTestConfig, node *v1.Node) *localTestVolume {
	testDirName := "local-volume-test-" + string(uuid.NewUUID())
	hostDir := filepath.Join(hostBase, testDirName)
	createAndMapBlockLocalVolume(config, hostDir, node)
	loopDev := getBlockLoopDev(config, hostDir, node)
	volume := generateLocalTestVolume(loopDev, config, BlockLocalVolumeType, node)
	volume.hostDir = loopDev
	volume.loopDevDir = hostDir
	return volume
}

func setupLocalVolumeBlockFsWithFormat(config *localTestConfig, node *v1.Node) *localTestVolume {
	testDirName := "local-volume-test-" + string(uuid.NewUUID())
	hostDir := filepath.Join(hostBase, testDirName)
	createAndMapBlockLocalVolume(config, hostDir, node)
	loopDev := getBlockLoopDev(config, hostDir, node)
	// format and mount at hostDir
	// give others rwx for read/write testing
	cmd := fmt.Sprintf("sudo mkfs -t ext4 %s && sudo mount -t ext4 %s %s && sudo chmod o+rwx %s", loopDev, loopDev, hostDir, hostDir)
	_, err := issueNodeCommandWithResult(config, cmd, node)
	Expect(err).NotTo(HaveOccurred())
	volume := generateLocalTestVolume(hostDir, config, BlockFsWithFormatLocalVolumeType, node)
	volume.hostDir = hostDir
	volume.loopDevDir = loopDev
	return volume
}

func setupLocalVolumeBlockFsWithoutFormat(config *localTestConfig, node *v1.Node) *localTestVolume {
	testDirName := "local-volume-test-" + string(uuid.NewUUID())
	hostDir := filepath.Join(hostBase, testDirName)
	createAndMapBlockLocalVolume(config, hostDir, node)
	loopDev := getBlockLoopDev(config, hostDir, node)
	volume := generateLocalTestVolume(loopDev, config, BlockFsWithoutFormatLocalVolumeType, node)
	// we do this in order to set block device path to local PV spec path directly
	// and test local volume plugin FileSystem mode on block device
	volume.hostDir = loopDev
	volume.loopDevDir = hostDir
	return volume
}

// Determine the /dev/loopXXX device associated with this test, via its hostDir.
func getBlockLoopDev(config *localTestConfig, hostDir string, node *v1.Node) string {
	loopDevCmd := fmt.Sprintf("E2E_LOOP_DEV=$(sudo losetup | grep %s/file | awk '{ print $1 }') 2>&1 > /dev/null && echo ${E2E_LOOP_DEV}", hostDir)
	loopDevResult, err := issueNodeCommandWithResult(config, loopDevCmd, node)
	Expect(err).NotTo(HaveOccurred())
	return strings.TrimSpace(loopDevResult)
}

func verifyLocalVolume(config *localTestConfig, volume *localTestVolume) {
	framework.ExpectNoError(framework.WaitOnPVandPVC(config.client, config.ns, volume.pv, volume.pvc))
}

func verifyLocalPod(config *localTestConfig, volume *localTestVolume, pod *v1.Pod, expectedNodeName string) {
	podNodeName, err := podNodeName(config, pod)
	Expect(err).NotTo(HaveOccurred())
	framework.Logf("pod %q created on Node %q", pod.Name, podNodeName)
	Expect(podNodeName).To(Equal(expectedNodeName))
}

// Deletes the PVC/PV, and launches a pod with hostpath volume to remove the test directory.
func cleanupLocalVolumeGCELocalSSD(config *localTestConfig, volume *localTestVolume) {
	By("Removing the test directory")
	file := volume.hostDir + "/" + testFile
	removeCmd := fmt.Sprintf("if [ -f %s ]; then rm %s; fi", file, file)
	err := issueNodeCommand(config, removeCmd, volume.node)
	Expect(err).NotTo(HaveOccurred())
}

// Deletes the PVC/PV, and launches a pod with hostpath volume to remove the test directory.
func cleanupLocalVolumeTmpfs(config *localTestConfig, volume *localTestVolume) {
	unmountTmpfsLocalVolume(config, volume.hostDir, volume.node)

	By("Removing the test directory")
	removeCmd := fmt.Sprintf("rm -r %s", volume.hostDir)
	err := issueNodeCommand(config, removeCmd, volume.node)
	Expect(err).NotTo(HaveOccurred())
}

// Deletes the PVC/PV, and launches a pod with hostpath volume to remove the test directory.
func cleanupLocalVolumeDirectory(config *localTestConfig, volume *localTestVolume) {
	By("Removing the test directory")
	removeCmd := fmt.Sprintf("rm -r %s", volume.hostDir)
	err := issueNodeCommand(config, removeCmd, volume.node)
	Expect(err).NotTo(HaveOccurred())
}

// Deletes the PVC/PV, and launches a pod with hostpath volume to remove the test directory.
func cleanupLocalVolumeDirectoryLink(config *localTestConfig, volume *localTestVolume) {
	By("Removing the test directory")
	hostDir := volume.hostDir
	hostDirBackend := hostDir + "-backend"
	removeCmd := fmt.Sprintf("sudo rm -r %s && rm -r %s", hostDir, hostDirBackend)
	err := issueNodeCommand(config, removeCmd, volume.node)
	Expect(err).NotTo(HaveOccurred())
}

// Deletes the PVC/PV, and launches a pod with hostpath volume to remove the test directory.
func cleanupLocalVolumeDirectoryBindMounted(config *localTestConfig, volume *localTestVolume) {
	By("Removing the test directory")
	hostDir := volume.hostDir
	removeCmd := fmt.Sprintf("sudo umount %s && rm -r %s", hostDir, hostDir)
	err := issueNodeCommand(config, removeCmd, volume.node)
	Expect(err).NotTo(HaveOccurred())
}

// Deletes the PVC/PV, and launches a pod with hostpath volume to remove the test directory.
func cleanupLocalVolumeDirectoryLinkBindMounted(config *localTestConfig, volume *localTestVolume) {
	By("Removing the test directory")
	hostDir := volume.hostDir
	hostDirBackend := hostDir + "-backend"
	removeCmd := fmt.Sprintf("sudo rm %s && sudo umount %s && rm -r %s", hostDir, hostDirBackend, hostDirBackend)
	err := issueNodeCommand(config, removeCmd, volume.node)
	Expect(err).NotTo(HaveOccurred())
}

// Deletes the PVC/PV and removes the test directory holding the block file.
func cleanupLocalVolumeBlock(config *localTestConfig, volume *localTestVolume) {
	volume.hostDir = volume.loopDevDir
	unmapBlockLocalVolume(config, volume.hostDir, volume.node)
	By("Removing the test directory")
	removeCmd := fmt.Sprintf("rm -r %s", volume.hostDir)
	err := issueNodeCommand(config, removeCmd, volume.node)
	Expect(err).NotTo(HaveOccurred())
}

// Deletes the PVC/PV and removes the test directory holding the block file.
func cleanupLocalVolumeBlockFsWithFormat(config *localTestConfig, volume *localTestVolume) {
	// umount first
	By("Umount blockfs mountpoint")
	umountCmd := fmt.Sprintf("sudo umount %s", volume.hostDir)
	err := issueNodeCommand(config, umountCmd, volume.node)
	unmapBlockLocalVolume(config, volume.hostDir, volume.node)
	By("Removing the test directory")
	removeCmd := fmt.Sprintf("rm -r %s", volume.hostDir)
	err = issueNodeCommand(config, removeCmd, volume.node)
	Expect(err).NotTo(HaveOccurred())
}

func cleanupLocalVolumeBlockFsWithoutFormat(config *localTestConfig, volume *localTestVolume) {
	volume.hostDir = volume.loopDevDir
	unmapBlockLocalVolume(config, volume.hostDir, volume.node)
	By("Removing the test directory")
	removeCmd := fmt.Sprintf("rm -r %s", volume.hostDir)
	err := issueNodeCommand(config, removeCmd, volume.node)
	Expect(err).NotTo(HaveOccurred())
}

func makeLocalPVCConfig(config *localTestConfig, volumeType localVolumeType) framework.PersistentVolumeClaimConfig {
	pvcConfig := framework.PersistentVolumeClaimConfig{
		AccessModes:      []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
		StorageClassName: &config.scName,
	}
	if volumeType == BlockLocalVolumeType {
		pvcVolumeMode := v1.PersistentVolumeBlock
		pvcConfig.VolumeMode = &pvcVolumeMode
	}
	return pvcConfig
}

func makeLocalPVConfig(config *localTestConfig, volume *localTestVolume) framework.PersistentVolumeConfig {
	// TODO: hostname may not be the best option
	nodeKey := "kubernetes.io/hostname"
	if volume.node.Labels == nil {
		framework.Failf("Node does not have labels")
	}
	nodeValue, found := volume.node.Labels[nodeKey]
	if !found {
		framework.Failf("Node does not have required label %q", nodeKey)
	}

	pvConfig := framework.PersistentVolumeConfig{
		PVSource: v1.PersistentVolumeSource{
			Local: &v1.LocalVolumeSource{
				Path: volume.hostDir,
			},
		},
		NamePrefix:       "local-pv",
		StorageClassName: config.scName,
		NodeAffinity: &v1.VolumeNodeAffinity{
			Required: &v1.NodeSelector{
				NodeSelectorTerms: []v1.NodeSelectorTerm{
					{
						MatchExpressions: []v1.NodeSelectorRequirement{
							{
								Key:      nodeKey,
								Operator: v1.NodeSelectorOpIn,
								Values:   []string{nodeValue},
							},
						},
					},
				},
			},
		},
	}

	if volume.localVolumeType == BlockLocalVolumeType {
		pvVolumeMode := v1.PersistentVolumeBlock
		pvConfig.VolumeMode = &pvVolumeMode
	}
	return pvConfig
}

// Creates a PVC and PV with prebinding
func createLocalPVCsPVs(config *localTestConfig, volumes []*localTestVolume, mode storagev1.VolumeBindingMode) {
	var err error

	for _, volume := range volumes {
		pvcConfig := makeLocalPVCConfig(config, volume.localVolumeType)
		pvConfig := makeLocalPVConfig(config, volume)

		volume.pv, volume.pvc, err = framework.CreatePVPVC(config.client, pvConfig, pvcConfig, config.ns, false)
		framework.ExpectNoError(err)
	}

	if mode == storagev1.VolumeBindingImmediate {
		for _, volume := range volumes {
			verifyLocalVolume(config, volume)
		}
	} else {
		// Verify PVCs are not bound
		// There isn't really a great way to verify this without making the test be slow...
		err = wait.PollImmediate(time.Second, 10*time.Second, func() (done bool, err error) {
			for _, volume := range volumes {
				pvc, err := config.client.CoreV1().PersistentVolumeClaims(volume.pvc.Namespace).Get(volume.pvc.Name, metav1.GetOptions{})
				framework.ExpectNoError(err)
				Expect(pvc.Status.Phase).To(Equal(v1.ClaimPending))
			}
			return false, nil
		})
		Expect(err).To(HaveOccurred())
	}
}

func makeLocalPodWithNodeAffinity(config *localTestConfig, volume *localTestVolume, nodeName string) (pod *v1.Pod) {
	pod = framework.MakeSecPod(config.ns, []*v1.PersistentVolumeClaim{volume.pvc}, false, "", false, false, selinuxLabel, nil)
	if pod == nil {
		return
	}
	affinity := &v1.Affinity{
		NodeAffinity: &v1.NodeAffinity{
			RequiredDuringSchedulingIgnoredDuringExecution: &v1.NodeSelector{
				NodeSelectorTerms: []v1.NodeSelectorTerm{
					{
						MatchExpressions: []v1.NodeSelectorRequirement{
							{
								Key:      "kubernetes.io/hostname",
								Operator: v1.NodeSelectorOpIn,
								Values:   []string{nodeName},
							},
						},
					},
				},
			},
		},
	}
	pod.Spec.Affinity = affinity
	return
}

func makeLocalPodWithNodeSelector(config *localTestConfig, volume *localTestVolume, nodeName string) (pod *v1.Pod) {
	pod = framework.MakeSecPod(config.ns, []*v1.PersistentVolumeClaim{volume.pvc}, false, "", false, false, selinuxLabel, nil)
	if pod == nil {
		return
	}
	ns := map[string]string{
		"kubernetes.io/hostname": nodeName,
	}
	pod.Spec.NodeSelector = ns
	return
}

func makeLocalPodWithNodeName(config *localTestConfig, volume *localTestVolume, nodeName string) (pod *v1.Pod) {
	pod = framework.MakeSecPod(config.ns, []*v1.PersistentVolumeClaim{volume.pvc}, false, "", false, false, selinuxLabel, nil)
	if pod == nil {
		return
	}
	pod.Spec.NodeName = nodeName
	return
}

func createLocalPod(config *localTestConfig, volume *localTestVolume, fsGroup *int64) (*v1.Pod, error) {
	By("Creating a pod")
	return framework.CreateSecPod(config.client, config.ns, []*v1.PersistentVolumeClaim{volume.pvc}, false, "", false, false, selinuxLabel, fsGroup, framework.PodStartShortTimeout)
}

func createAndMountTmpfsLocalVolume(config *localTestConfig, dir string, node *v1.Node) {
	By(fmt.Sprintf("Creating tmpfs mount point on node %q at path %q", node.Name, dir))
	err := issueNodeCommand(config, fmt.Sprintf("mkdir -p %q && sudo mount -t tmpfs -o size=1m tmpfs-%q %q", dir, dir, dir), node)
	Expect(err).NotTo(HaveOccurred())
}

func unmountTmpfsLocalVolume(config *localTestConfig, dir string, node *v1.Node) {
	By(fmt.Sprintf("Unmount tmpfs mount point on node %q at path %q", node.Name, dir))
	err := issueNodeCommand(config, fmt.Sprintf("sudo umount %q", dir), node)
	Expect(err).NotTo(HaveOccurred())
}

func createAndMapBlockLocalVolume(config *localTestConfig, dir string, node *v1.Node) {
	By(fmt.Sprintf("Creating block device on node %q using path %q", node.Name, dir))
	mkdirCmd := fmt.Sprintf("mkdir -p %s", dir)
	// Create 10MB file that will serve as the backing for block device.
	ddCmd := fmt.Sprintf("dd if=/dev/zero of=%s/file bs=512 count=20480", dir)
	losetupCmd := fmt.Sprintf("sudo losetup -f  %s/file", dir)
	err := issueNodeCommand(config, fmt.Sprintf("%s && %s && %s", mkdirCmd, ddCmd, losetupCmd), node)
	Expect(err).NotTo(HaveOccurred())
}

func unmapBlockLocalVolume(config *localTestConfig, dir string, node *v1.Node) {
	loopDev := getBlockLoopDev(config, dir, node)
	By(fmt.Sprintf("Unmap block device %q on node %q at path %s/file", loopDev, node.Name, dir))
	losetupDeleteCmd := fmt.Sprintf("sudo losetup -d %s", loopDev)
	err := issueNodeCommand(config, losetupDeleteCmd, node)
	Expect(err).NotTo(HaveOccurred())
}

func createWriteCmd(testDir string, testFile string, writeTestFileContent string, volumeType localVolumeType) string {
	if volumeType == BlockLocalVolumeType {
		// testDir is the block device.
		testFileDir := filepath.Join("/tmp", testDir)
		testFilePath := filepath.Join(testFileDir, testFile)
		// Create a file containing the testFileContent.
		writeTestFileCmd := fmt.Sprintf("mkdir -p %s; echo %s > %s", testFileDir, writeTestFileContent, testFilePath)
		// sudo is needed when using ssh exec to node.
		// sudo is not needed and does not exist in some containers (e.g. busybox), when using pod exec.
		sudoCmd := fmt.Sprintf("SUDO_CMD=$(which sudo); echo ${SUDO_CMD}")
		// Write the testFileContent into the block device.
		writeBlockCmd := fmt.Sprintf("${SUDO_CMD} dd if=%s of=%s bs=512 count=100", testFilePath, testDir)
		// Cleanup the file containing testFileContent.
		deleteTestFileCmd := fmt.Sprintf("rm %s", testFilePath)
		return fmt.Sprintf("%s && %s && %s && %s", writeTestFileCmd, sudoCmd, writeBlockCmd, deleteTestFileCmd)
	} else {
		testFilePath := filepath.Join(testDir, testFile)
		return fmt.Sprintf("mkdir -p %s; echo %s > %s", testDir, writeTestFileContent, testFilePath)
	}
}
func createReadCmd(testFileDir string, testFile string, volumeType localVolumeType) string {
	if volumeType == BlockLocalVolumeType {
		// Create the command to read the beginning of the block device and print it in ascii.
		return fmt.Sprintf("hexdump -n 100 -e '100 \"%%_p\"' %s | head -1", testFileDir)
	} else {
		// Create the command to read (aka cat) a file.
		testFilePath := filepath.Join(testFileDir, testFile)
		return fmt.Sprintf("cat %s", testFilePath)
	}
}

// Read testFile and evaluate whether it contains the testFileContent
func testReadFileContent(testFileDir string, testFile string, testFileContent string, pod *v1.Pod, volumeType localVolumeType) {
	readCmd := createReadCmd(testFileDir, testFile, volumeType)
	readOut := podRWCmdExec(pod, readCmd)
	Expect(readOut).To(ContainSubstring(testFileContent))
}

// Execute a read or write command in a pod.
// Fail on error
func podRWCmdExec(pod *v1.Pod, cmd string) string {
	out, err := utils.PodExec(pod, cmd)
	framework.Logf("podRWCmdExec out: %q err: %v", out, err)
	Expect(err).NotTo(HaveOccurred())
	return out
}

// Initialize test volume on node
// and create local PVC and PV
func setupLocalVolumesPVCsPVs(
	config *localTestConfig,
	localVolumeType localVolumeType,
	node *v1.Node,
	count int,
	mode storagev1.VolumeBindingMode) []*localTestVolume {

	By("Initializing test volumes")
	testVols := setupLocalVolumes(config, localVolumeType, node, count)

	By("Creating local PVCs and PVs")
	createLocalPVCsPVs(config, testVols, mode)

	return testVols
}

// newLocalClaim creates a new persistent volume claim.
func newLocalClaimWithName(config *localTestConfig, name string) *v1.PersistentVolumeClaim {
	claim := v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: config.ns,
		},
		Spec: v1.PersistentVolumeClaimSpec{
			StorageClassName: &config.scName,
			AccessModes: []v1.PersistentVolumeAccessMode{
				v1.ReadWriteOnce,
			},
			Resources: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceName(v1.ResourceStorage): resource.MustParse(testRequestSize),
				},
			},
		},
	}

	return &claim
}

func createStatefulSet(config *localTestConfig, ssReplicas int32, volumeCount int, anti, parallel bool) *appsv1.StatefulSet {
	mounts := []v1.VolumeMount{}
	claims := []v1.PersistentVolumeClaim{}
	for i := 0; i < volumeCount; i++ {
		name := fmt.Sprintf("vol%v", i+1)
		pvc := newLocalClaimWithName(config, name)
		mounts = append(mounts, v1.VolumeMount{Name: name, MountPath: "/" + name})
		claims = append(claims, *pvc)
	}

	podAffinityTerms := []v1.PodAffinityTerm{
		{
			LabelSelector: &metav1.LabelSelector{
				MatchExpressions: []metav1.LabelSelectorRequirement{
					{
						Key:      "app",
						Operator: metav1.LabelSelectorOpIn,
						Values:   []string{"local-volume-test"},
					},
				},
			},
			TopologyKey: "kubernetes.io/hostname",
		},
	}

	affinity := v1.Affinity{}
	if anti {
		affinity.PodAntiAffinity = &v1.PodAntiAffinity{
			RequiredDuringSchedulingIgnoredDuringExecution: podAffinityTerms,
		}
	} else {
		affinity.PodAffinity = &v1.PodAffinity{
			RequiredDuringSchedulingIgnoredDuringExecution: podAffinityTerms,
		}
	}

	labels := map[string]string{"app": "local-volume-test"}
	spec := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "local-volume-statefulset",
			Namespace: config.ns,
		},
		Spec: appsv1.StatefulSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "local-volume-test"},
			},
			Replicas: &ssReplicas,
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:         "nginx",
							Image:        imageutils.GetE2EImage(imageutils.Nginx),
							VolumeMounts: mounts,
						},
					},
					Affinity: &affinity,
				},
			},
			VolumeClaimTemplates: claims,
			ServiceName:          "test-service",
		},
	}

	if parallel {
		spec.Spec.PodManagementPolicy = appsv1.ParallelPodManagement
	}

	ss, err := config.client.AppsV1().StatefulSets(config.ns).Create(spec)
	Expect(err).NotTo(HaveOccurred())

	config.ssTester.WaitForRunningAndReady(ssReplicas, ss)
	return ss
}

func validateStatefulSet(config *localTestConfig, ss *appsv1.StatefulSet, anti bool) {
	pods := config.ssTester.GetPodList(ss)

	nodes := sets.NewString()
	for _, pod := range pods.Items {
		nodes.Insert(pod.Spec.NodeName)
	}

	if anti {
		// Verify that each pod is on a different node
		Expect(nodes.Len()).To(Equal(len(pods.Items)))
	} else {
		// Verify that all pods are on same node.
		Expect(nodes.Len()).To(Equal(1))
	}

	// Validate all PVCs are bound
	for _, pod := range pods.Items {
		for _, volume := range pod.Spec.Volumes {
			pvcSource := volume.VolumeSource.PersistentVolumeClaim
			if pvcSource != nil {
				err := framework.WaitForPersistentVolumeClaimPhase(
					v1.ClaimBound, config.client, config.ns, pvcSource.ClaimName, framework.Poll, time.Second)
				Expect(err).NotTo(HaveOccurred())
			}
		}
	}
}

// SkipUnlessLocalSSDExists takes in an ssdInterface (scsi/nvme) and a filesystemType (fs/block)
// and skips if a disk of that type does not exist on the node
func SkipUnlessLocalSSDExists(config *localTestConfig, ssdInterface, filesystemType string, node *v1.Node) {
	ssdCmd := fmt.Sprintf("ls -1 /mnt/disks/by-uuid/google-local-ssds-%s-%s/ | wc -l", ssdInterface, filesystemType)
	res, err := issueNodeCommandWithResult(config, ssdCmd, node)
	Expect(err).NotTo(HaveOccurred())
	num, err := strconv.Atoi(strings.TrimSpace(res))
	Expect(err).NotTo(HaveOccurred())
	if num < 1 {
		framework.Skipf("Requires at least 1 %s %s localSSD ", ssdInterface, filesystemType)
	}
}
