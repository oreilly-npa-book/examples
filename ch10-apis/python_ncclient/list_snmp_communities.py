from ncclient import manager

device = manager.connect(
    host='vmx1', port=830, username='ntc',
    password='ntc123', hostkey_verify=False,
    device_params={"name": "junos"},
)

get_filter = """
<configuration>
  <snmp>
  </snmp>
</configuration>
"""

nc_get_reply = device.get(('subtree', get_filter))

snmp_list = []

communities = nc_get_reply.findall('.//community')

for community in communities:
    temp = {}
    temp['name'] = community.find('.//name').text
    temp['auth'] = community.find('.//authorization').text
    snmp_list.append(temp)

print(snmp_list)
