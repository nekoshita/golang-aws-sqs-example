terraform {
  required_version = "= 0.14.8"
  required_providers {
    google = {
      source  = "hashicorp/aws"
      version = "3.33.0"
    }
  }
}
