base:
  ntc:
    - napalm
  csr1:
    - csr1_pillar
  vmx1:
    - vmx1_pillar
  'nxos-spine1':
    - nxos_spine1_pillar
  eos-spine1:
    - eos_spine1_pillar
  '*':
    - ntp_peers
