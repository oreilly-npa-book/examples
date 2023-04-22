#!/usr/bin/env python

import argparse
import os

from jinja2 import Environment
from jinja2 import FileSystemLoader
import yaml

ENV = Environment(loader=FileSystemLoader(
    "%s/templates/" % os.path.dirname(os.path.abspath(__file__))
))


def render_template(datafile):
    """Renders a Jinja2 template with the provided YAML data file
    """

    with open(datafile) as datafile:
        configdata = yaml.load(datafile)

        df_name = os.path.basename(str(datafile)).split(".")[0]

        template = ENV.get_template("%s.j2" % df_name)
        return template.render(config=configdata)


def main():
    """Gathers arguments, runs render_template
    """

    parser = argparse.ArgumentParser(
        description="Render a network configuration template"
    )
    parser.add_argument("datafile", type=str, action="store",
                        help="YAML file to use as data")

    args = parser.parse_args()

    print(render_template(args.datafile))
