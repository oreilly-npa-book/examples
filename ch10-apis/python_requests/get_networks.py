#!/usr/bin/env python3

# The Python requests library is used to issue and work HTTP-based systems.
import requests

# This executes if our script is being run directly.
if __name__ == "__main__":
    # This token is taken from Cisco Developer Hub to experiment with the API
    my_token = "6bec40cf957de430a6f1f2baa056b99a4fac9ea0"

    # This statement creates a Python dictionary for the HTTP request
    # headers that are going to use in the API calls.  The first two
    # headers we are setting are Content-Type and Accept.
    # The last one uses a custom header used in Meraki for authentication
    headers = {
        "Accept": "application/json",
        "Content-Type": "application/json",
        "X-Cisco-Meraki-API-Key": my_token,
    }

    # The URL is saved as a variable called base_url to modularize our
    # code and simplify the next statement.
    base_url = "https://api.meraki.com/api/v1"

    # In the requests library, there is a function per HTTP verb and in
    # this example we are issuing a GET request, so we are therefore
    # using the `get` function. We pass two objects into the `get` function.
    # The first  object passed in must be the URL, and the others should be
    # keyword arguments (key=value pairs). Then, we pass the proper headers.
    response = requests.get(f"{base_url}/organizations", headers=headers)

    # Extract the content of the response deserializing as JSON, and return
    # a Python object, list or dict
    organizations = response.json()

    # We pick the "id" from the first organization to later gather the related networks
    first_organization_id = organizations[0]["id"]

    # Similar get request, composing the url with the organization id
    response = requests.get(
        f"{base_url}/organizations/{first_organization_id}/networks", headers=headers
    )

    networks = response.json()
