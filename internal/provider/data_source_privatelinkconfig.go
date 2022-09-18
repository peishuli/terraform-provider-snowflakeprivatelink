package provider

import (
	"context"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePrivatelinkConfig() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Use this data source to get privatelink configuration for Snowflake.",

		ReadContext: dataSourcePrivatelinkConfigRead,

		Schema: map[string]*schema.Schema{
			"account_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of your Snowflake account.",
			},
		
			"internal_stage": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the internal stage.",
			},

			"aws_vpce_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The AWS VPCE ID for your account.",
			},

			"account_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL used to connect to Snowflake through AWS PrivateLink or Azure Private Link.",
			},
		
			"regionless_account_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The regionless URL used to connect to Snowflake through AWS PrivateLink or Azure Private Link.",
			},

			"ocsp_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The OCSP URL corresponding to your Snowflake account that uses AWS PrivateLink or Azure Private Link.",
			},
		},
	}
}

func dataSourcePrivatelinkConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	
	var diags diag.Diagnostics
	client := meta.(*apiClient).Client
	
	
	privatelink_config, err := client.GetPrivatelinkConfig()
	if (err != nil) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get privatelink config.",
			Detail:   "Failed to get privatelink config.",
		})
		d.SetId("")
		return diags
	}

	
	d.Set("account_name", privatelink_config.Privatelink_account_name)
	d.Set("internal_stage", privatelink_config.Privatelink_internal_stage)
	d.Set("aws_vpce_id", privatelink_config.Privatelink_vpce_id)
	d.Set("account_url", privatelink_config.Privatelink_account_url)
	d.Set("regionless_account_url", privatelink_config.Regionless_privatelink_account_url)
	d.Set("ocsp_url", privatelink_config.Privatelink_ocsp_url)		
	
	idFromAPI := "my-id"
	d.SetId(idFromAPI)
	
	return diags
}
