flannel-tar:
  archive:
    - extracted
    - user: root
    - name: /usr/local/src
    - makedirs: True
    - source: https://github.com/coreos/flannel/releases/download/v0.5.5/flannel-0.5.5-linux-amd64.tar.gz
    - tar_options: v
    - source_hash: md5=972c717254775bef528f040af804f2cc
    - archive_format: tar
    - if_missing: /usr/local/src/flannel/flannel-0.5.5/

flannel-symlink:
  file.symlink:
    - name: /usr/local/bin/flanneld
    - target: /usr/local/src/flannel-0.5.5/flanneld
    - force: true
    - watch:
        - archive: flannel-tar

/etc/init.d/flannel:
  file.managed:
    - source: salt://flannel/initd
    - user: root
    - group: root
    - mode: 755

/var/run/flannel/network.json:
  file.managed:
    - source: salt://flannel/network.json
    - makedirs: True
    - user: root
    - group: root
    - mode: 755

flannel:
  service.running:
    - enable: True
    - watch:
      - file: /usr/local/bin/flanneld
      - file: /etc/init.d/flannel
