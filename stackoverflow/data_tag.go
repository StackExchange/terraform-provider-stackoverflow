package stackoverflow

import (
	"context"
	"fmt"
	"strconv"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataTag() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataTagRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The tag description",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The tag name",
			},
			"tag_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The identifier for the tag",
			},
		},
	}
}

func dataTagRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	tagID := d.Get("tag_id").(int)

	tag, err := c.GetTag(&tagID)
	if err != nil {
		return diag.FromErr(err)
	}

	if tag == nil {
		return diag.FromErr(fmt.Errorf("no tag found matching identifier %d", tagID))
	}

	d.SetId(strconv.Itoa(tag.ID))
	d.Set("name", tag.Name)
	d.Set("description", tag.Description)

	return diags
}
