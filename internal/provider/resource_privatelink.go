package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePrivatelink() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider Privatelink.",

		CreateContext: resourcePrivatelinkCreate,
		ReadContext:   resourcePrivatelinkRead,
		UpdateContext: resourcePrivatelinkUpdate,
		DeleteContext: resourcePrivatelinkDelete,

		Schema: map[string]*schema.Schema{
			"dummy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "A dummy field to avoid 'All fields are ForceNew or Computed w/out Optional, Update is superfluous' issue.",
				
			},
			"privatelink_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of your Snowflake account.",
			},
		},
	}
}

func resourcePrivatelinkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).Client
	err := client.AuthorizePrivatelink()
	if (err != nil) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get privatelink status.",
			Detail:   "Failed to get privatelink status.",
		})
		d.SetId("")
		return diags
	}
	idFromAPI := "my-id"
	d.SetId(idFromAPI)
	
	return resourcePrivatelinkRead(ctx, d, meta)
}

func resourcePrivatelinkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).Client
	privatelink_status, err := client.GetPrivatelink()
	if (err != nil) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to authorize privatelink.",
			Detail:   "Failed to authorize privatelink.",
		})
		d.SetId("")
		return diags
	}
	d.Set("privatelink_status", privatelink_status)
	idFromAPI := "my-id"
	d.SetId(idFromAPI)
	
	return diags
}

func resourcePrivatelinkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics	
	
	return diags
}

func resourcePrivatelinkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).Client
	err := client.RevokePrivatelink()
	if (err != nil) {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to revoke privatelink.",
			Detail:   "Failed to revoke privatelink.",
		})
		d.SetId("")
		return diags
	}
	idFromAPI := "my-id"
	d.SetId(idFromAPI)
	
	return resourcePrivatelinkRead(ctx, d, meta)	
}
