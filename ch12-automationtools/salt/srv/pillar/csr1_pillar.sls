---
proxy:
  proxytype: napalm
  driver: ios
  username: ntc
  host: csr1
  password: ntc123
'openconfig-bgp':
  bgp:
    global:
      config:
        as: 65001
        router_id: 172.17.17.1
