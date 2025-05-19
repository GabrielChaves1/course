terraform {
  required_version = ">= 1.2.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "local" {}
}

provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      Environment = "dev"
      Project     = var.project_name
    }
  }
}

module "cognito" {
  source         = "../../modules/cognito"
  user_pool_name = "${var.project_name}-${var.environment}-user-pool"
}
