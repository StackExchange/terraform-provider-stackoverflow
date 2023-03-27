package stackoverflow

import (
	"context"
	"fmt"
	"strconv"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataArticle() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataArticleRead,
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
	}
}

func dataArticleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	articleID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	articleIDs := []int{articleID}

	articles, err := c.GetArticles(&articleIDs)
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
