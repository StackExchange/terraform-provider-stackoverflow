package stackoverflow

import (
	"fmt"
	"strconv"
	"testing"

	so "terraform-provider-stackoverflow/stackoverflow/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestStackOverflowArticle(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testStackOverflowArticleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testStackOverflowArticleConfig(),
				Check: resource.ComposeTestCheckFunc(
					testStackOverflowArticleExists("stackoverflow_article.test"),
				),
			},
		},
	})
}

func testStackOverflowArticleConfig() string {
	return `resource "stackoverflow_article" "test" {
		article_type = "knowledge-article"
		title = "unit test"
		body_markdown = "unit test"
		tags = ["unit-test"]
		filter = "1234abcd"
	}`
}

func testStackOverflowArticleDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*so.Client)

	for _, resource := range s.RootModule().Resources {
		if resource.Type != "stackoverflow_article" {
			continue
		}

		articleID, err := strconv.Atoi(resource.Primary.ID)
		if err != nil {
			return err
		}

		err = c.DeleteArticle(articleID)
		if err != nil {
			return err
		}
	}

	return nil
}

func testStackOverflowArticleExists(n string) resource.TestCheckFunc {
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
