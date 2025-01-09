package stackoverflow

import (
	"context"
	"fmt"
	"strconv"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataCollection() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataCollectionRead,
		Schema: map[string]*schema.Schema{
			"collection_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The identifier for the collection",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The collection description",
			},
			"title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The collection name",
			},
		},
	}
}

func dataCollectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	collectionID := d.Get("collection_id").(int)

	collection, err := c.GetCollection(&collectionID)
	if err != nil {
		return diag.FromErr(err)
	}

	if collection == nil {
		return diag.FromErr(fmt.Errorf("no collection found matching identifier %d", collectionID))
	}

	d.SetId(strconv.Itoa(collection.ID))
	d.Set("title", collection.Title)
	d.Set("description", collection.Description)

	return diags
}
