package stackoverflow

import (
	"fmt"
	"strconv"
	so "terraform-provider-stackoverflow/stackoverflow/client"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestStackOverflowQuestion(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testStackOverflowQuestionConfig(),
				Check: resource.ComposeTestCheckFunc(
					testStackOverflowQuestionExists("stackoverflow_question.test"),
				),
			},
		},
	})
}

func testStackOverflowQuestionConfig() string {
	return `resource "stackoverflow_question" "test" {
		title = "Unit testing"
		body_markdown = "How do I unit test?"
		tags = ["test"]
	}`
}

func testStackOverflowQuestionDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*so.Client)

	for _, resource := range s.RootModule().Resources {
		if resource.Type != "stackoverflow_question" {
			continue
		}

		questionID, err := strconv.Atoi(resource.Primary.ID)
		if err != nil {
			return err
		}

		err = c.DeleteQuestion(questionID)
		if err != nil {
			return err
		}
	}

	return nil
}

func testStackOverflowQuestionExists(n string) resource.TestCheckFunc {
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
