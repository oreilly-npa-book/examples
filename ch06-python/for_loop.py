#!/usr/bin/env python3

COMMANDS = {
    "description": "description {}",
    "speed": "speed {}",
    "duplex": "duplex {}",
}

CONFIG_PARAMS = {
    "description": "auto description by Python",
    "speed": "10000",
    "duplex": "auto",
}

commands_list = []
for feature, value in CONFIG_PARAMS.items():
    command = COMMANDS.get(feature).format(value)
    commands_list.append(command)

commands_list.insert(0, "interface Eth1/1")
print(commands_list)
