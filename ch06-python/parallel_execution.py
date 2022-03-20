#!/usr/bin/env python3

import time


def get_commands(vlan, name):
    time.sleep(5)
    commands = []
    commands.append(f"vlan {vlan}")
    commands.append(f"name {name}")
    return commands


if __name__ == "__main__":

    vlans = [
        {"id": "10", "name": "USERS"},
        {"id": "20", "name": "VOICE"},
        {"id": "30", "name": "WLAN"},
    ]

    import concurrent.futures

    with concurrent.futures.ThreadPoolExecutor() as executor:
        start_time = time.time()
        tasks = []
        for vlan in vlans:
            tasks.append(
                executor.submit(get_commands, vlan=vlan["id"], name=vlan["name"])
            )
        for task in concurrent.futures.as_completed(tasks):
            print(task.result())
        print(f"Time spent: {time.time() - start_time} seconds")
