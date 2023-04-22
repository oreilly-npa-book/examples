#!/usr/bin/env python

import os
import sys
import yaml

# YAML_DIR is the location of the directory where the YAML files are kept
YAML_DIR = "%s/../datafiles/" % os.path.dirname(os.path.abspath(__file__))

# Let's loop over the YAML files and try to load them
for filename in os.listdir(YAML_DIR):
    yaml_file = "%s%s" % (YAML_DIR, filename)

    if os.path.isfile(yaml_file) and ".yml" in yaml_file:
        try:
            with open(yaml_file) as yamlfile:
                configdata = yaml.load(yamlfile)

        # If there was a problem importing the YAML, we can print
        # an error message, and quit with a non-zero error code
        # (which will trigger our CI system to indicate failure)
        except Exception:
            print("%s failed YAML import" % yaml_file)
            sys.exit(1)

sys.exit(0)
