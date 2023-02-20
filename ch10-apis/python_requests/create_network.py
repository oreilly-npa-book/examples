#!/usr/bin/env python3

# The json module is imported so that we can encode and decode JSON
# objects over the wire. While we work with JSON objects in Python as
# dictionaries, they are sent over the wire as JSON strings in API calls.
# This means we need a way to convert dictionaries to strings so they are
# understood by the network device (web server).  The converse is true too.
# We use the json module for these actions.
import json
import requests

if __name__ == "__main__":
    # Same as "get_networks" script
    my_token = "6bec40cf957de430a6f1f2baa056b99a4fac9ea0"
    headers = {
        "Accept": "application/json",
        "Content-Type": "application/json",
        "X-Cisco-Meraki-API-Key": my_token,
    }
    base_url = "https://api.meraki.com/api/v1"
    response = requests.get(f"{base_url}/organizations", headers=headers)
    organizations = response.json()
    first_organization_id = organizations[0]["id"]

    # Payload contains the data necessary to define the expected
    # data to create a new network in the API
    payload = {
        "name": "my brand new automated network",
        "productTypes": ["switch"],
    }

    # Using the `post` method instead of `get` to create an object
    response = requests.post(
        f"{base_url}/organizations/{first_organization_id}/networks",
        headers=headers,
        data=json.dumps(payload),
    )

    print(response.json())
