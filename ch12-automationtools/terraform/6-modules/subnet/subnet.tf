variable "vpc_id" {}
variable "vpc_cidr_block" {}

resource "aws_subnet" "my_subnet" {
  vpc_id         = var.vpc_id
  vpc_cidr_block = var.cidr_block
}
