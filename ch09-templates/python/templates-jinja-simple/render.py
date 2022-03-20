from jinja2 import Environment, FileSystemLoader

ENV = Environment(loader=FileSystemLoader('.'))
template = ENV.get_template("template.jinja")

interface_dict = {
    "name": "GigabitEthernet0/1",
    "description": "Server Port",
    "vlan": 10,
    "uplink": False
}

print(template.render(interface=interface_dict))