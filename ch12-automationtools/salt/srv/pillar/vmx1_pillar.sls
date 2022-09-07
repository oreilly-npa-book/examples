proxy:
  proxytype: napalm
  driver: junos
  username: ntc
  host: vmx1
  password: ntc123
'openconfig-bgp':
  bgp:
    global:
      config:
        as: 65001
        router_id: 172.18.18.1
beacons:
  inotify:
    - files:
        /srv/pillar/ntp_peers.sls: {}
