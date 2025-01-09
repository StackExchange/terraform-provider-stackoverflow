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
			"article_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The identifier for the article",
			},
			"article_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of article. Must be one of `knowledge-article`, `announcement`, `how-to-guide`, `policy`",
			},
			"body_markdown": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The article content in Markdown format",
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The set of tags to be associated with the article",
			},
			"title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The title of the article",
			},
		},
	}
}

func dataArticleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	articleID := d.Get("article_id").(int)

	article, err := c.GetArticle(&articleID)
	if err != nil {
		return diag.FromErr(err)
	}

	if article == nil {
		return diag.FromErr(fmt.Errorf("no article found matching identifier %d", articleID))
	}

	d.SetId(strconv.Itoa(article.ID))
	d.Set("article_type", article.ArticleType)
	d.Set("body_markdown", article.BodyMarkdown)
	d.Set("title", article.Title)
	d.Set("tags", selectTagNamesToArray(article.Tags))

	return diags
}
