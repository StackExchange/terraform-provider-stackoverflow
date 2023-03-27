package stackoverflow

import (
	"context"
	"fmt"
	"strconv"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceArticle() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceArticleCreate,
		ReadContext:   resourceArticleRead,
		UpdateContext: resourceArticleUpdate,
		DeleteContext: resourceArticleDelete,
		Schema: map[string]*schema.Schema{
			"article_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"body_markdown": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
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

func resourceArticleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	article := &so.Article{
		ArticleType:  d.Get("article_type").(string),
		BodyMarkdown: d.Get("body_markdown").(string),
		Title:        d.Get("title").(string),
		Tags:         expandTagsToArray(d.Get("tags").([]interface{})),
		Filter:       "omhz)aiL)ei3-sat(rZKVugTgq0f6)", //"!2oF_R8n-Ln(vwVra-FZ1DIV*iJEU3e_yLcG*k1oG5P",
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
	articleIDs := []int{articleID}
	articles, err := client.GetArticles(&articleIDs)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(*articles) < 1 {
		return diag.FromErr(fmt.Errorf("no article found matching identifier %d", articleID))
	}

	if len(*articles) > 1 {
		return diag.FromErr(fmt.Errorf("found %d articles matching identifier %d", len(*articles), articleID))
	}

	article := (*articles)[0]

	d.SetId(strconv.Itoa(article.ID))
	d.Set("article_type", article.ArticleType)
	d.Set("body_markdown", article.BodyMarkdown)
	d.Set("title", article.Title)
	d.Set("tags", article.Tags)

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

	article := &so.Article{
		ID:           articleID,
		ArticleType:  d.Get("article_type").(string),
		BodyMarkdown: d.Get("body_markdown").(string),
		Title:        d.Get("title").(string),
		Tags:         expandTagsToArray(d.Get("tags").([]interface{})),
		Filter:       "omhz)aiL)ei3-sat(rZKVugTgq0f6)", //"!2oF_R8n-Ln(vwVra-FZ1DIV*iJEU3e_yLcG*k1oG5P",
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
