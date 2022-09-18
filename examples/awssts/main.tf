###################################
# Custom AWS STS Provider 
###################################
terraform {
  required_providers {
    awssts = {
      version = "0.1"
      source  = "peishuli.com/dev/awssts"
    }
  }
}

provider "awssts" {
  user_name = "sam" 
}

data "aws_federation_token" "current" { 
  provider = awssts
}

output user_id {
  value = data.aws_federation_token.current.federated_user_id
  sensitive = true
}

output federation_token {
  value = data.aws_federation_token.current.federation_token
  sensitive = true
}
