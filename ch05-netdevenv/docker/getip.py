import requests
ip = requests.get('https://api.ipify.org').text
print(f"Hello from Docker! Your IP address is {ip}")
