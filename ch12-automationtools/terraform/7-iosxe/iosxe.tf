terraform {
  required_providers {
    iosxe = {
      source = "CiscoDevNet/iosxe"
      version = "0.1.1"
    }
  }
}

provider "iosxe" {
  host = "https://sandbox-iosxe-latest-1.cisco.com"
  device_username = "developer"
  device_password = "C1sco12345"
}

data "iosxe_rest" "vlan_data" {
   path = "/data/Cisco-IOS-XE-native:native"
}

output "response" {
  value = data.iosxe_rest.vlan_data
}

variable "vlan_id" {
  default = 456
}

# resource "iosxe_rest" "vlan_example_patch" {
#   method = "PATCH"
#   path = "/data/Cisco-IOS-XE-native:native/vlan"
#   payload = jsonencode(
#     {
#       "Cisco-IOS-XE-native:vlan": {
#         "Cisco-IOS-XE-vlan:vlan-list": [
#             {
#                 "id": var.vlan_id,
#                 "name": "VLAN-${var.vlan_id}"
#             }
#         ]
#       }
#     }
#   )
# }





# resource "iosxe_rest" "feature_put" {
#     method = "PUT"
#     path   = <RESTCONF_XPATH>
#     payload = jsonencode(
#         {
#             <JSON_RESPONSE>
#         }
#     )
# }
