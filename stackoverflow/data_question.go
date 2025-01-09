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
			"body_markdown": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The question content in Markdown format",
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The set of tags to be associated with the article",
			},
			"question_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The identifier for the question",
			},
			"title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The title of the article",
			},
		},
	}
}

func dataQuestionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	questionID := d.Get("question_id").(int)

	question, err := c.GetQuestion(&questionID)
	if err != nil {
		return diag.FromErr(err)
	}

	if question == nil {
		return diag.FromErr(fmt.Errorf("no question found matching identifier %d", questionID))
	}

	d.SetId(strconv.Itoa(question.ID))
	d.Set("body_markdown", question.BodyMarkdown)
	d.Set("title", question.Title)
	d.Set("tags", selectTagNamesToArray(question.Tags))

	return diags
}
