class Device:
    def __init__(self, ip, username, password):
        self.host_ip = ip
        self.user = username
        self.pswd = password

    def show(self, command):
        return f'{command} output from {self.host_ip}'
