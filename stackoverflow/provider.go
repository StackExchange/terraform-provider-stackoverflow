package stackoverflow

import (
	"context"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("STACK_OVERFLOW_ACCESS_TOKEN", nil),
				Description: "The Stack Overflow API access token",
			},
			"team_name": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("STACK_OVERFLOW_TEAM", nil),
				Description: "The Stack Overflow team name",
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("STACK_OVERFLOW_API_URL", nil),
				Description: "The base URL for the Stack Overflow API",
				Default:     "https://api.stackoverflowteams.com/2.3/",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"stackoverflow_article":  resourceArticle(),
			"stackoverflow_question": resourceQuestion(),
			"stackoverflow_answer":   resourceAnswer(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"stackoverflow_article":  dataArticle(),
			"stackoverflow_question": dataQuestion(),
			"stackoverflow_answer":   dataAnswer(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	accessToken := d.Get("access_token").(string)
	teamName := d.Get("team_name").(string)
	baseURL := d.Get("base_url").(string)

	var diags diag.Diagnostics
	client := so.NewClient(&baseURL, &teamName, &accessToken)

	return client, diags
}
