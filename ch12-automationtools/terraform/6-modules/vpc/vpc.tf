variable "my_cidr" {}

resource "aws_vpc" "my_vpc" {
  cidr_block = var.my_cidr
}

output "cidr_block" {
  value = aws_vpc.my_vpc.cidr_block
}

output "id" {
  value = aws_vpc.my_vpc.id
}
