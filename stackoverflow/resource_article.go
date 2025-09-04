package stackoverflow

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceArticle() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceArticleCreate,
		ReadContext:   resourceArticleRead,
		UpdateContext: resourceArticleUpdate,
		DeleteContext: resourceArticleDelete,
		Schema: map[string]*schema.Schema{
			"article_type": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The type of article. Must be one of `knowledgeArticle`, `announcement`, `howToGuide`, `policy`",
				ValidateFunc: validation.StringInSlice([]string{"knowledgeArticle", "announcement", "howToGuide", "policy"}, false),
			},
			"body_markdown": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The article content in Markdown format",
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

func resourceArticleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	tags := convertToArray[string](d.Get("tags").([]interface{}))
	sort.Strings(tags)

	article := &so.Article[string]{
		ArticleType: d.Get("article_type").(string),
		Body:        d.Get("body_markdown").(string),
		Title:       d.Get("title").(string),
		Tags:        tags,
	}

	newArticle, err := client.CreateArticle(article)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(newArticle.ID))

	return diags
}

func resourceArticleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)
	var diags diag.Diagnostics
	articleID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	article, err := client.GetArticle(&articleID)
	if err != nil {
		return diag.FromErr(err)
	}

	if article == nil {
		return diag.FromErr(fmt.Errorf("no article found matching identifier %d", articleID))
	}

	tags := selectTagNamesToArray(article.Tags)
	sort.Strings(tags)

	d.SetId(strconv.Itoa(article.ID))
	d.Set("article_type", article.ArticleType)
	d.Set("body_markdown", article.BodyMarkdown)
	d.Set("title", article.Title)
	d.Set("tags", tags)

	return diags
}

func resourceArticleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	articleID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	tags := convertToArray[string](d.Get("tags").([]interface{}))
	sort.Strings(tags)

	article := &so.Article[string]{
		ID:          articleID,
		ArticleType: d.Get("article_type").(string),
		Body:        d.Get("body_markdown").(string),
		Title:       d.Get("title").(string),
		Tags:        tags,
	}

	_, err2 := client.UpdateArticle(article)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	return diags

}

func resourceArticleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)
	var diags diag.Diagnostics

	articleID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	err2 := client.DeleteArticle(articleID)

	if err2 != nil {
		return diag.FromErr(err2)
	}

	return diags
}
