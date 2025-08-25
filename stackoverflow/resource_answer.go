package stackoverflow

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

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
			"body_markdown": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The answer content in Markdown format",
			},
			"question_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The question identifier that this answer applies to",
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

	answer := &so.Answer[string]{
		Body:       d.Get("body_markdown").(string),
		QuestionID: d.Get("question_id").(int),
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

	answerID := 0
	questionID := 0
	err := error(nil)

	// Check if the data matches an expression like 10/11 (question id/answer id)
	// this is used to support importing answers where the question id is required
	// but cannot be supplied because the identifier naturally does not contain the question
	// identifier as a logical component
	pattern := regexp.MustCompile(`[0-9]+/[0-9]+`)

	if pattern.MatchString(d.Id()) {
		parts := strings.Split(d.Id(), "/")
		if len(parts) == 2 {
			questionID, err = strconv.Atoi(parts[0])
			if err != nil {
				return diag.FromErr(err)
			}

			answerID, err = strconv.Atoi(parts[1])
			if err != nil {
				return diag.FromErr(err)
			}
		}
	} else {
		answerID, err = strconv.Atoi(d.Id())
		if err != nil {
			return diag.FromErr(err)
		}

		questionID = d.Get("question_id").(int)
	}

	answer, err := client.GetAnswer(&questionID, &answerID)
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

func resourceAnswerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	answerID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	answer := &so.Answer[string]{
		ID:         answerID,
		Body:       d.Get("body_markdown").(string),
		QuestionID: d.Get("question_id").(int),
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

	questionID := d.Get("question_id").(int)

	err2 := client.DeleteAnswer(questionID, answerID)

	if err2 != nil {
		return diag.FromErr(err2)
	}

	return diags
}
