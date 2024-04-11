globals "terraform" {
  version = "1.7.1"
}

globals "terraform" "backend" {
  bucket       = "terramate-atlantis-poc-terraform-state-backend" # Change this to something unique 
  region       = "us-east-1"
  dynamo_table = "terraform_atlantis_state"
}


globals "terraform" "providers" "aws" {
  enabled = true
  source  = "hashicorp/aws"
  version = "~> 5.31"
  config = {
    region = "us-east-1"
  }
}
