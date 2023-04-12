package stackoverflow

import (
	"context"
	"fmt"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataFilter() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataFilterRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataFilterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	filterID := d.Get("filter").(string)
	filterIDs := []string{filterID}

	filters, err := c.GetFilters(&filterIDs)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(*filters) < 1 {
		return diag.FromErr(fmt.Errorf("no answer found matching identifier %s", filterID))
	}

	if len(*filters) > 1 {
		return diag.FromErr(fmt.Errorf("found %d answers matching identifier %s", len(*filters), filterID))
	}

	filter := (*filters)[0]

	d.SetId(filter.ID)

	return diags
}
