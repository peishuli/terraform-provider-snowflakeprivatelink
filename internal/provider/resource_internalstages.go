package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInternalStages() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider InternalStages.",

		CreateContext: resourceInternalStagesCreate,
		ReadContext:   resourceInternalStagesRead,
		UpdateContext: resourceInternalStagesUpdate,
		DeleteContext: resourceInternalStagesDelete,

		Schema: map[string]*schema.Schema{
			"enable_internal_stages": {
				// This description is used by the documentation generator and the language server.
				Description: "A flag that indicate whether to enable or disable internal stages for the privatelink.",
				Type:        schema.TypeBool,
				Required:    true,
			},
		},
	}
}

func resourceInternalStagesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).Client
	err := client.EnableInternalStagesForPrivatelink(true)
	if (err != nil)	{
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
	return diags
}

func resourceInternalStagesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// return diag.Errorf("not implemented")
	return diags
}

func resourceInternalStagesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).Client
	enable_internal_stages := d.Get("enable_internal_stages").(bool)
	err := client.EnableInternalStagesForPrivatelink(enable_internal_stages)
	if (err != nil)	{
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
	return diags

}

func resourceInternalStagesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient).Client
	err := client.EnableInternalStagesForPrivatelink(false)
	if (err != nil)	{
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
	return diags
}
