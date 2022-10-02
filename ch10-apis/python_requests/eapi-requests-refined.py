import json
import sys
import requests
from requests.auth import HTTPBasicAuth


if __name__ == "__main__":

    auth = HTTPBasicAuth('ntc', 'ntc123')

    url = 'http://eos-spine1/command-api'

    payload = {
        "jsonrpc": "2.0",
        "method": "runCmds",
        "params": {
            "format": "json",
            "timestamps": False,
            "cmds": [
                "show vlan brief"
            ],
            "version": 1
        },
        "id": "EapiExplorer-1"
    }

    response = requests.post(url, data=json.dumps(payload), auth=auth)
    rsp = response.json()
    vlans = rsp['result'][0]['vlans']
    print('{:12}{:<10}'.format('VLAN ID', 'NAME'))
    for vlan_id, config in vlans.items():
        print('{:<12}{:<10}'.format(vlan_id,  config['name']))
