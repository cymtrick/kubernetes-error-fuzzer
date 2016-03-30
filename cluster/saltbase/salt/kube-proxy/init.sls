/var/lib/kube-proxy/kubeconfig:
  file.managed:
    - source: salt://kube-proxy/kubeconfig
    - user: root
    - group: root
    - mode: 400
    - makedirs: true

# kube-proxy in a static pod
/etc/kubernetes/manifests/kube-proxy.manifest:
  file.managed:
    - source: salt://kube-proxy/kube-proxy.manifest
    - template: jinja
    - user: root
    - group: root
    - mode: 644
    - makedirs: true
    - dir_mode: 755
    - context:
        # 20m might cause kube-proxy CPU starvation on full nodes, resulting in
        # delayed service updates. But, giving it more would be a breaking change 
        # to the overhead requirements for existing clusters.
        # Any change here should be accompanied by a proportional change in CPU
        # requests of other per-node add-ons (e.g. fluentd).
        cpurequest: '20m'
    - require:
      - service: docker
      - service: kubelet

/var/log/kube-proxy.log:
  file.managed:
    - user: root
    - group: root
    - mode: 644

#stop legacy kube-proxy service
stop_kube-proxy:
  service.dead:
    - name: kube-proxy
    - enable: None
