// TERRAMATE: GENERATED AUTOMATICALLY DO NOT EDIT

terraform {
  backend "s3" {
    bucket         = "terramate-atlantis-poc-terraform-state-backend"
    dynamodb_table = "terraform_atlantis_state"
    encrypt        = true
    key            = "terraform/stacks/by-id/98bd7abe-1ae0-4f8f-9a27-af4a171f40af/terraform.tfstate"
    region         = "us-east-1"
  }
}
