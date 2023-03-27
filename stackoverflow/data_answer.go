package stackoverflow

import (
	"context"
	"fmt"
	"strconv"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataAnswer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataAnswerRead,
		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"body_markdown": {
				Type:     schema.TypeString,
				Required: true,
			},
			"question_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataAnswerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	answerID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	answerIDs := []int{answerID}

	answers, err := c.GetAnswers(&answerIDs)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(*answers) < 1 {
		return diag.FromErr(fmt.Errorf("no answer found matching identifier %d", answerID))
	}

	if len(*answers) > 1 {
		return diag.FromErr(fmt.Errorf("found %d answers matching identifier %d", len(*answers), answerID))
	}

	answer := (*answers)[0]

	d.SetId(strconv.Itoa(answer.ID))
	d.Set("body_markdown", answer.BodyMarkdown)
	d.Set("question_id", answer.QuestionID)
	d.Set("title", answer.Title)
	d.Set("tags", answer.Tags)

	return diags
}
