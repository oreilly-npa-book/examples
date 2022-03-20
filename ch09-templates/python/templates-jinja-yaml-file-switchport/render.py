from jinja2 import Environment, FileSystemLoader
import yaml

ENV = Environment(loader=FileSystemLoader('.'))

template = ENV.get_template("template.jinja")

with open("data.yml") as f:
    interfaces = yaml.safe_load(f)
    print(template.render(interface_list=interfaces))
