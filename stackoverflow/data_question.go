package stackoverflow

import (
	"context"
	"fmt"
	"strconv"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataQuestion() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataQuestionRead,
		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"body_markdown": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataQuestionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	questionID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	questionIDs := []int{questionID}

	questions, err := c.GetQuestions(&questionIDs)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(*questions) < 1 {
		return diag.FromErr(fmt.Errorf("no question found matching identifier %d", questionID))
	}

	if len(*questions) > 1 {
		return diag.FromErr(fmt.Errorf("found %d questions matching identifier %d", len(*questions), questionID))
	}

	question := (*questions)[0]

	d.SetId(strconv.Itoa(question.ID))
	d.Set("body_markdown", question.BodyMarkdown)
	d.Set("title", question.Title)
	d.Set("tags", question.Tags)

	return diags
}
