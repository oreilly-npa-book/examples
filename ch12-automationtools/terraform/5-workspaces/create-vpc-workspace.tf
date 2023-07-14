terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  provider = terraform.workspace
}

variable "cidr" {
  type = map(string)

  default = {
    A = "198.51.100.0/24"
    B = "203.0.113.0/24"
  }
}

# Create a VPC
resource "aws_vpc" "my_vpc" {
  cidr_block = var.cidr[terraform.workspace]
}

# Create two subnets
resource "aws_subnet" "my_subnets" {
  count      = 2
  vpc_id     = aws_vpc.my_vpc.id
  cidr_block = cidrsubnet(aws_vpc.my_vpc.cidr_block, 1, count.index)
}
