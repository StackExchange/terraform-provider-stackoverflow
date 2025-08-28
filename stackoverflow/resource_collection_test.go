package stackoverflow

import (
	"fmt"
	"strconv"
	so "terraform-provider-stackoverflow/stackoverflow/client"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccStackOverflowCollection(t *testing.T) {
	resourceType := "stackoverflow_collection"
	resourceName := "test"
	resourceIdentifier := fmt.Sprintf("%s.%s", resourceType, resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccStackOverflowCollectionCheckDestroy(resourceIdentifier),
		Steps: []resource.TestStep{
			{
				Config: testAccStackOverflowCollectionConfig(resourceType, resourceName, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccStackOverflowCollectionCheckExists(resourceIdentifier),
				),
			},
			{
				Config: testAccStackOverflowCollectionConfig(resourceType, resourceName, " Updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccStackOverflowCollectionCheckExists(resourceIdentifier),
				),
			},
			{
				ResourceName: resourceIdentifier,
				ImportState:  true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					resourceState, found := s.RootModule().Resources[resourceIdentifier]
					if !found {
						return "", fmt.Errorf("Resource not found: %s", resourceIdentifier)
					}
					return fmt.Sprintf(resourceState.Primary.ID), nil
				},
				ImportStateVerify: true,
			},
		},
	})
}

func testAccStackOverflowCollectionConfig(resourceType, resourceName, update string) string {
	return fmt.Sprintf(`resource "%s" "%s" {
		title = "Unit Test%s"
		description = "Unit testing%s"
	}`, resourceType, resourceName, update, update)
}

func testAccStackOverflowCollectionCheckDestroy(resourceIdentifier string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState, found := s.RootModule().Resources[resourceIdentifier]

		if !found {
			return fmt.Errorf("No resource found with id %s", resourceIdentifier)
		}

		if resourceState.Primary.ID == "" {
			return fmt.Errorf("Resource with ID is empty for %s", resourceIdentifier)
		}

		collectionID, err := strconv.Atoi(resourceState.Primary.ID)

		if err != nil {
			return err
		}

		c := testAccProvider.Meta().(*so.Client)
		collection, err := c.GetCollection(&collectionID)

		if err != nil {
			return err
		}

		if collection != nil && !collection.IsDeleted {
			return fmt.Errorf("Collection with ID %d still exists", collectionID)
		}

		return nil
	}
}

func testAccStackOverflowCollectionCheckExists(n string) resource.TestCheckFunc {
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
