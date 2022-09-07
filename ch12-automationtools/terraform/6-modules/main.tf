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

variable "base_cidr" {
  default = "192.0.2.0/24"
}

module "my_vpc_module" {
  source       = "./vpc"
  cidr_for_vpc = var.base_cidr
}

module "my_subnet_module" {
  source     = "./subnet"
  count      = 2
  vpc_id     = module.my_vpc_module.id
  vpc_cidr_block = cidrsubnet(module.my_vpc_module.cidr_block, 1, count.index)
}
