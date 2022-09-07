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
  region = "eu-west-3"
}

# Create a VPC
resource "aws_vpc" "my_vpc" {
  cidr_block = "192.0.2.0/24"
}
