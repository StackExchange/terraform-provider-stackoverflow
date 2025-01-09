package stackoverflow

import (
	"fmt"
	"strconv"
	"testing"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestStackOverflowCollection(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testStackOverflowCollectionConfig(),
				Check: resource.ComposeTestCheckFunc(
					testStackOverflowArticleExists("stackoverflow_collection.test"),
				),
			},
		},
	})
}

func testStackOverflowCollectionConfig() string {
	return `resource "stackoverflow_collection" "test" {
		title = "unit test"
		description = "unit test"
	}`
}

func testStackOverflowCollectionDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*so.Client)

	for _, resource := range s.RootModule().Resources {
		if resource.Type != "stackoverflow_collection" {
			continue
		}

		collectionID, err := strconv.Atoi(resource.Primary.ID)
		if err != nil {
			return err
		}

		err = c.DeleteCollection(collectionID)
		if err != nil {
			return err
		}
	}

	return nil
}

func testStackOverflowCollectionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resource, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if resource.Primary.ID == "" {
			return fmt.Errorf("No ID set")
		}

		return nil
	}
}
