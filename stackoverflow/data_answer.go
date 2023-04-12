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
				Type:     schema.TypeInt,
				Required: true,
			},
			"filter": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataAnswerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	answerID := d.Get("answer_id").(int)
	filter := d.Get("filter").(string)
	answerIDs := []int{answerID}

	answers, err := c.GetAnswers(&answerIDs, &filter)
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
