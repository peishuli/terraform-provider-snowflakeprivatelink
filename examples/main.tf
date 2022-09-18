###########################################################
# call awssts module to aquire aws id and federated token
# so that we can use it to provision snowflake resources
# via Terraform
###########################################################
module awssts {
  source = "./awssts"  
}

#
# Snowflake Privatelink Provider
#
terraform {
  required_providers {
    snowflakeprivatelink = {
      version = "0.1"
      source  = "peishuli.com/dev/snowflakeprivatelink"
    }
  }
}

provider "snowflakeprivatelink" {
    account = "bc96517"
    region = "us-east-1"
    aws_id = "${module.awssts.user_id}"
    aws_federated_token = "${module.awssts.federation_token}"
}

#
# Create a Snowflake Privatelink Service
#
resource snowflake_privatelink "this" {  
  provider = snowflakeprivatelink
}


#
# The Snowflake Privatelink Configuration data source 
#
data snowflake_privatelink_config "this" {
    provider = snowflakeprivatelink
}

output account {
    value = data.snowflake_privatelink_config.this.account_name
}

output internal_stage {
  value = data.snowflake_privatelink_config.this.internal_stage
}

output aws_vpce_id {
  value = data.snowflake_privatelink_config.this.aws_vpce_id
}

output account_url {
  value = data.snowflake_privatelink_config.this.account_url
}

output regionless_account_url {
  value = data.snowflake_privatelink_config.this.regionless_account_url
}

output ocsp_url {
  value = data.snowflake_privatelink_config.this.ocsp_url
}

#
# The Snowflake data source to view privatelink status
#
data snowflake_privatelink "this" {
  provider = snowflakeprivatelink
}

output privatelink_status {
  value = data.snowflake_privatelink.this.privatelink_status
}

#
# The Snowflake resource to enable/disable internal stages
#
resource snowflake_privatelink_enableinternalstages "this" {
  enable_internal_stages = true
  provider = snowflakeprivatelink
}




