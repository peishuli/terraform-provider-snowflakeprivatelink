package provider

import (
	"context"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePrivatelink() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Use this data source to get privatelink for Snowflake.",

		ReadContext: dataSourcePrivatelinkRead,

		Schema: map[string]*schema.Schema{			
			"privatelink_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of your Snowflake account.",
			},
			
		},
	}
}

func dataSourcePrivatelinkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	
	var diags diag.Diagnostics
	client := meta.(*apiClient).Client
	privatelink_status, err := client.GetPrivatelink()
	if (err != nil) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get privatelink status.",
			Detail:   "Failed to get privatelink status.",
		})
		d.SetId("")
		return diags
	}
	d.Set("privatelink_status", privatelink_status)
	idFromAPI := "my-id"
	d.SetId(idFromAPI)
	
	return diags
}
