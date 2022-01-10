terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "3.63.0"
    }
  }
}

provider "aws" {
  # ... potentially other provider configuration ...
  region = "eu-west-2"
  # skip_credentials_validation = true
  # skip_requesting_account_id = true
  # access_key                  = "mock_access_key"
  # secret_key                  = "mock_secret_key"
  
  # endpoints {
  #   iam = "http://localhost:4566"
  #   ec2 = "http://localhost:4566"
  #   sqs = "http://localhost:4566"
  #   sns = "http://localhost:4566"
  #   dynamodb = "http://localhost:4566"
  #   lambda = "http://localhost:4566"
  # }
}
