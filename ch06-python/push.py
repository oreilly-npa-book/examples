#!/usr/bin/env python3


def get_commands(vlan, name):
    commands = []
    commands.append(f"vlan {vlan}")
    commands.append(f"name {name}")
    return commands


def push_commands(device, commands):
    print(f"Connecting to device: {device}")
    for cmd in commands:
        print(f"Sending command: {cmd}")


if __name__ == "__main__":

    devices = ["switch1", "switch2", "switch3"]

    vlans = [
        {"id": "10", "name": "USERS"},
        {"id": "20", "name": "VOICE"},
        {"id": "30", "name": "WLAN"},
    ]

    for vlan in vlans:
        vid = vlan.get("id")
        name = vlan.get("name")
        print()
        print(f"CONFIGURING VLAN: {vid}")
        commands = get_commands(vid, name)
        for device in devices:
            push_commands(device, commands)
            print()
