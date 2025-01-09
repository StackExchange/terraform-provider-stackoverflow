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

func resourceCollection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCollectionCreate,
		ReadContext:   resourceCollectionRead,
		UpdateContext: resourceCollectionUpdate,
		DeleteContext: resourceCollectionDelete,
		Schema: map[string]*schema.Schema{
			"content_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The collection description.",
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The title of the collection",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCollectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	content := convertToArray[int](d.Get("content_ids").([]interface{}))
	sort.Ints(content)

	collection := &so.Collection{
		ContentIds:  content,
		Description: d.Get("description").(string),
		Title:       d.Get("title").(string),
	}

	newCollection, err := client.CreateCollection(collection)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(newCollection.ID))

	return diags
}

func resourceCollectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)
	var diags diag.Diagnostics
	collectionID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	collection, err := client.GetCollection(&collectionID)
	if err != nil {
		return diag.FromErr(err)
	}

	if collection == nil {
		return diag.FromErr(fmt.Errorf("no collection found matching identifier %d", collectionID))
	}

	content := selectCollectionContentIdsToArray(collection.Content)
	sort.Ints(content)

	d.SetId(strconv.Itoa(collection.ID))
	d.Set("description", collection.Description)
	d.Set("title", collection.Title)
	d.Set("content_ids", content)

	return diags
}

func resourceCollectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	collectionID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	content := convertToArray[int](d.Get("content_ids").([]interface{}))
	sort.Ints(content)

	collection := &so.Collection{
		ID:          collectionID,
		Description: d.Get("description").(string),
		Title:       d.Get("title").(string),
		ContentIds:  content,
	}

	_, err2 := client.UpdateCollection(collection)
	if err2 != nil {
		return diag.FromErr(err2)
	}

	return diags

}

func resourceCollectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*so.Client)
	var diags diag.Diagnostics

	collectionID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	err2 := client.DeleteCollection(collectionID)

	if err2 != nil {
		return diag.FromErr(err2)
	}

	return diags
}
