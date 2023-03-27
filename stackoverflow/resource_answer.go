package stackoverflow

import (
	"context"
	"fmt"
	"strconv"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAnswer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAnswerCreate,
		ReadContext:   resourceAnswerRead,
		UpdateContext: resourceAnswerUpdate,
		DeleteContext: resourceAnswerDelete,
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAnswerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	answer := &so.Answer{
		BodyMarkdown: d.Get("body_markdown").(string),
		QuestionID:   d.Get("question_id").(int),
		Title:        d.Get("title").(string),
		Tags:         expandTagsToArray(d.Get("tags").([]interface{})),
		Filter:       "omhz)aiL)ei3-sat(rZKVugTgq0f6)", //"!2oF_R8n-Ln(vwVra-FZ1DIV*iJEU3e_yLcG*k1oG5P",
	}

	newAnswer, err := client.CreateAnswer(answer)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(newAnswer.ID))

	return diags
}

func resourceAnswerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)
	var diags diag.Diagnostics
	answerID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	answerIDs := []int{answerID}
	answers, err := client.GetAnswers(&answerIDs)
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

func resourceAnswerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	answerID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	answer := &so.Answer{
		ID:           answerID,
		BodyMarkdown: d.Get("body_markdown").(string),
		QuestionID:   d.Get("question_id").(int),
		Title:        d.Get("title").(string),
		Tags:         expandTagsToArray(d.Get("tags").([]interface{})),
		Filter:       "omhz)aiL)ei3-sat(rZKVugTgq0f6)", //"!2oF_R8n-Ln(vwVra-FZ1DIV*iJEU3e_yLcG*k1oG5P",
	}

	_, err2 := client.UpdateAnswer(answer)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	return diags

}

func resourceAnswerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)
	var diags diag.Diagnostics

	answerID, err := strconv.Atoi(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	err2 := client.DeleteAnswer(answerID)

	if err2 != nil {
		return diag.FromErr(err2)
	}

	return diags
}
