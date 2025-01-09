package stackoverflow

import (
	"context"
	"fmt"
	"sort"
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
				Type:        schema.TypeString,
				Required:    true,
				Description: "The question content in Markdown format",
			},
			"tags": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The set of tags to be associated with the article",
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The title of the article",
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

	tags := convertToArray[string](d.Get("tags").([]interface{}))
	sort.Strings(tags)

	question := &so.Question[string]{
		Body:  d.Get("body_markdown").(string),
		Title: d.Get("title").(string),
		Tags:  tags,
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
	question, err := client.GetQuestion(&questionID)
	if err != nil {
		return diag.FromErr(err)
	}

	if question == nil {
		return diag.FromErr(fmt.Errorf("no question found matching identifier %d", questionID))
	}

	tags := selectTagNamesToArray(question.Tags)
	sort.Strings(tags)

	d.SetId(strconv.Itoa(question.ID))
	d.Set("body_markdown", question.BodyMarkdown)
	d.Set("title", question.Title)
	d.Set("tags", tags)

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

	tags := convertToArray[string](d.Get("tags").([]interface{}))
	sort.Strings(tags)

	question := &so.Question[string]{
		ID:    questionID,
		Body:  d.Get("body_markdown").(string),
		Title: d.Get("title").(string),
		Tags:  tags,
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
