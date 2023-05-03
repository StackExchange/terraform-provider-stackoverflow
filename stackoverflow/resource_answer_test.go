package stackoverflow

import (
	"fmt"
	"strconv"
	so "terraform-provider-stackoverflow/stackoverflow/client"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestStackOverflowAnswer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testStackOverflowAnswerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testStackOverflowAnswerConfig(),
				Check: resource.ComposeTestCheckFunc(
					testStackOverflowAnswerExists("stackoverflow_answer.test"),
				),
			},
		},
	})
}

func testStackOverflowAnswerConfig() string {
	return `resource "stackoverflow_answer" "test" {
		question_id = 0
		body_markdown = "unit test"
		filter = "1234abcd"
	}`
}

func testStackOverflowAnswerDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*so.Client)

	for _, resource := range s.RootModule().Resources {
		if resource.Type != "stackoverflow_answer" {
			continue
		}

		answerID, err := strconv.Atoi(resource.Primary.ID)
		if err != nil {
			return err
		}

		err = c.DeleteAnswer(answerID)
		if err != nil {
			return err
		}
	}

	return nil
}

func testStackOverflowAnswerExists(n string) resource.TestCheckFunc {
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
