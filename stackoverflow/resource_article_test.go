package stackoverflow

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccStackOverflowArticle(t *testing.T) {
	resourceType := "stackoverflow_article"
	resourceName := "test"
	resourceIdentifier := fmt.Sprintf("%s.%s", resourceType, resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccStackOverflowArticleCheckDestroy(resourceIdentifier),
		Steps: []resource.TestStep{
			{
				Config: testAccStackOverflowArticleConfig(resourceType, resourceName, "knowledgeArticle", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccStackOverflowArticleCheckExists(resourceIdentifier),
				),
			},
			{
				Config:      testAccStackOverflowArticleConfig(resourceType, resourceName, "error", ""),
				ExpectError: regexp.MustCompile("expected article_type to be one of"),
			},
			{
				Config: testAccStackOverflowArticleConfig(resourceType, resourceName, "knowledgeArticle", " Updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccStackOverflowArticleCheckExists(resourceIdentifier),
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

func testAccStackOverflowArticleConfig(resourceType, resourceName, articleType, update string) string {
	return fmt.Sprintf(`resource "%s" "%s" {
		article_type = "%s"
		title = "unit test%s"
		body_markdown = "unit test%s"
		tags = ["unit-test"]
	}`, resourceType, resourceName, articleType, update, update)
}

func testAccStackOverflowArticleCheckDestroy(resourceIdentifier string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState, found := s.RootModule().Resources[resourceIdentifier]

		if !found {
			return fmt.Errorf("No resource found with id %s", resourceIdentifier)
		}

		if resourceState.Primary.ID == "" {
			return fmt.Errorf("Resource with ID is empty for %s", resourceIdentifier)
		}

		articleID, err := strconv.Atoi(resourceState.Primary.ID)

		if err != nil {
			return err
		}

		c := testAccProvider.Meta().(*so.Client)
		article, err := c.GetArticle(&articleID)

		if err != nil {
			return err
		}

		if article != nil && !article.IsDeleted {
			return fmt.Errorf("Article with ID %d still exists", articleID)
		}

		return nil
	}
}

func testAccStackOverflowArticleCheckExists(n string) resource.TestCheckFunc {
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
