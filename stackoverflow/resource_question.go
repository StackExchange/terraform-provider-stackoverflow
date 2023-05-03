package stackoverflow

import (
	"context"
	"fmt"
	"strconv"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceQuestion() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceQuestionCreate,
		ReadContext:   resourceQuestionRead,
		UpdateContext: resourceQuestionUpdate,
		DeleteContext: resourceQuestionDelete,
		Schema: map[string]*schema.Schema{
			"body_markdown": {
				Type:     schema.TypeString,
				Required: true,
			},
			"filter": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceQuestionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	question := &so.Question{
		BodyMarkdown: d.Get("body_markdown").(string),
		Title:        d.Get("title").(string),
		Tags:         mergeDefaultTagsWithResourceTags(client.DefaultTags, expandTagsToArray(d.Get("tags").([]interface{}))),
		Filter:       d.Get("filter").(string),
	}

	newQuestion, err := client.CreateQuestion(question)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(newQuestion.ID))

	return diags
}

func resourceQuestionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)
	var diags diag.Diagnostics
	questionID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	questionIDs := []int{questionID}
	filter := d.Get("filter").(string)
	questions, err := client.GetQuestions(&questionIDs, &filter)
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
	d.Set("tags", ignoreDefaultTags(client.DefaultTags, question.Tags, expandTagsToArray(d.Get("tags").([]interface{}))))

	return diags
}

func resourceQuestionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	questionID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	question := &so.Question{
		ID:           questionID,
		BodyMarkdown: d.Get("body_markdown").(string),
		Title:        d.Get("title").(string),
		Tags:         mergeDefaultTagsWithResourceTags(client.DefaultTags, expandTagsToArray(d.Get("tags").([]interface{}))),
		Filter:       d.Get("filter").(string),
	}

	_, err2 := client.UpdateQuestion(question)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	return diags

}

func resourceQuestionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)
	var diags diag.Diagnostics

	questionID, err := strconv.Atoi(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	err2 := client.DeleteQuestion(questionID)

	if err2 != nil {
		return diag.FromErr(err2)
	}

	return diags
}
