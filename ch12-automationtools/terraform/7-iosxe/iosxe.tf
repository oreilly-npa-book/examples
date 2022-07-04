terraform {
  required_providers {
    iosxe = {
      source  = "CiscoDevNet/iosxe"
      version = "0.1.1"
    }
  }
}

provider "iosxe" {
  host            = "https://sandbox-iosxe-latest-1.cisco.com"
  device_username = "developer"
  device_password = "C1sco12345"
}

data "iosxe_rest" "full_config" {
  path = "/data/Cisco-IOS-XE-native:native"
}

output "response" {
  value = data.iosxe_rest.full_config
}


resource "iosxe_rest" "snmp_example_chassis_id" {
  method = "PUT"
  path   = "/data/Cisco-IOS-XE-native:native/snmp-server/chassis-id"
  payload = jsonencode(
    {
      "Cisco-IOS-XE-snmp:chassis-id" : "a_new_chassis_id"
    }
  )
}
