from netmiko import ConnectHandler
from jinja2 import Environment, FileSystemLoader

device = ConnectHandler(
    host='nxos-spine1', 
    username='admin', 
    password='admin',
    device_type='cisco_nxos'
)

interface_dict = {
    "name": "Ethernet1/2",
    "description": "Server Port",
    "vlan": 10,
    "uplink": False
}

ENV = Environment(loader=FileSystemLoader('.'))
template = ENV.get_template("config.j2")
commands = template.render(interface=interface_dict)

filename = 'nxos.conf'
with open(filename, 'w') as config_file:
    config_file.writelines(commands)

output = device.send_config_from_file(filename)

verification = device.send_command(f'show run interface {interface_dict["name"]}')

print(verification)

device.disconnect()
