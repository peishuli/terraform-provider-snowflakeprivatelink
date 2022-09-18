package provider

import (
	"context"
	"fmt"	

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"account": {
					Type:     schema.TypeString,
					Required: true,
					Description: "The name of your Snowflake account.",
				},	
				"username": {
					Type:     schema.TypeString,
					Optional: true,
					DefaultFunc: schema.EnvDefaultFunc("SNOWFLAKE_USER", nil),
					Description: "Your login username.",
				},
				"password": {
					Type:     schema.TypeString,
					Required: true,	
					DefaultFunc: schema.EnvDefaultFunc("SNOWFLAKE_PASSWORD", nil),
					Description: "Your login password.",				
				},
				"region": {
					Type:     schema.TypeString,
					Required: true,			
					Description: "The AWS region of your Snowflake account.",		
				},
				"aws_id": {
					Type:     schema.TypeString,
					Required: true,
					Description: "Your AWS sam Id.",
				},
				"aws_federated_token": {
					Type:     schema.TypeString,
					Required: true,
					Description: "Your federated AWS token.",
				},

			},
			DataSourcesMap: map[string]*schema.Resource{
				"snowflake_privatelink": dataSourcePrivatelink(),
				"snowflake_privatelink_config": dataSourcePrivatelinkConfig(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"snowflake_privatelink": resourcePrivatelink(),
				"snowflake_privatelink_enableinternalstages": resourceInternalStages(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	Client *SnowflakeClient
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	    var diags diag.Diagnostics
		account := d.Get("account").(string)
		username := d.Get("username").(string)
  		password := d.Get("password").(string)
		region := d.Get("region").(string)	
		aws_id := d.Get("aws_id").(string)	
		aws_federated_token := d.Get("aws_federated_token").(string)
		// //TODO: input validations

		c := &apiClient{}
		cred := Credentials{
			Account: account,
			User: username,
			Password: password,
			Region: region,
		}

		fmt.Println(cred.Account)
		sc, err := c.Client.NewSnowflakeClient(cred)
		if (err != nil) {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create new snowflake client",
				Detail:   "Unable to create new snowflake client",
			})
			return nil, diags
		}
		c.Client = sc
		c.Client.AWSId = aws_id
		c.Client.AWSFederatedToken = aws_federated_token
		return c, diags
	}
}
