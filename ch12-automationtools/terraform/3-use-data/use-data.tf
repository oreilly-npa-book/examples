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

variable "vpc_id" {
  type        = string
  description = "The id of the vpc_id."
  sensitive   = true

  validation {
    condition     = length(var.vpc_id) > 4 && substr(var.vpc_id, 0, 4) == "vpc-"
    error_message = "The vpc_id value must be a valid VPC id, starting with \"vpc-\"."
  }
}

data "aws_vpc" "my_vpc" {
  id = var.vpc_id
}

# Create two subnets
resource "aws_subnet" "my_subnets" {
  count      = 2
  vpc_id     = data.aws_vpc.my_vpc.id
  cidr_block = cidrsubnet(data.aws_vpc.my_vpc.cidr_block, 1, count.index)
}
