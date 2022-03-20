#!/usr/bin/env python3

import time


def get_commands(vlan, name):
    time.sleep(5)
    commands = []
    commands.append(f"vlan {vlan}")
    commands.append(f"name {name}")
    return commands


def run_task(vlans):
    start_time = time.time()
    for vlan in vlans:
        result = get_commands(vlan=vlan["id"], name=vlan["name"])
        print(result)
    print(f"Time spent: {time.time() - start_time} seconds")


if __name__ == "__main__":

    vlans = [
        {"id": "10", "name": "USERS"},
        {"id": "20", "name": "VOICE"},
        {"id": "30", "name": "WLAN"},
    ]

    run_task(vlans)
