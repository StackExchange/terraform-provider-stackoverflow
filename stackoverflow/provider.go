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
				Description: "The Stack Overflow API access token. The `STACK_OVERFLOW_ACCESS_TOKEN` environment variable can be used instead.",
			},
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("STACK_OVERFLOW_API_URL", nil),
				Description: "The base URL for the Stack Overflow API (must end with `/`). For Stack Overflow for Teams this is in the format `https://api.stackoverflowteams.com/v3/teams/{team}/` and for Stack Overflow Enterprise this is in one of the following formats `https://{name}.stackenterprise.co/api/v3/`, `https://{name}.stackenterprise.co/api/v3/teams/{team}/`, `https://{your-custom-domain}/api/v3/`, or `https://{your-custom-domain}/api/v3/teams/{team}/`. The `STACK_OVERFLOW_API_URL` environment variable can be used instead.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"stackoverflow_answer":     resourceAnswer(),
			"stackoverflow_article":    resourceArticle(),
			"stackoverflow_collection": resourceCollection(),
			"stackoverflow_question":   resourceQuestion(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"stackoverflow_answer":     dataAnswer(),
			"stackoverflow_article":    dataArticle(),
			"stackoverflow_collection": dataCollection(),
			"stackoverflow_question":   dataQuestion(),
			"stackoverflow_tag":        dataTag(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	accessToken := d.Get("access_token").(string)
	baseURL := d.Get("base_url").(string)

	var diags diag.Diagnostics
	client := so.NewClient(&baseURL, &accessToken)
	client.DefaultTags = convertToArray[string](d.Get("default_tags").([]interface{}))

	return client, diags
}
