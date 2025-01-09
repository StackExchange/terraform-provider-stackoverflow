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
			"answer_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The identifier for the answer",
			},
			"body_markdown": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The answer content in Markdown format",
			},
			"question_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The identifier for the question",
			},
		},
	}
}

func dataAnswerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	answerID := d.Get("answer_id").(int)
	questionID := d.Get("question_id").(int)

	answer, err := c.GetAnswer(&questionID, &answerID)
	if err != nil {
		return diag.FromErr(err)
	}

	if answer == nil {
		return diag.FromErr(fmt.Errorf("no answer found matching identifier %d", answerID))
	}

	d.SetId(strconv.Itoa(answer.ID))
	d.Set("body_markdown", answer.BodyMarkdown)
	d.Set("question_id", answer.QuestionID)

	return diags
}
