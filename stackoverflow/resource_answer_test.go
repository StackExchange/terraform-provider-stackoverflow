package stackoverflow

import (
	"fmt"
	"strconv"
	so "terraform-provider-stackoverflow/stackoverflow/client"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccStackOverflowAnswer(t *testing.T) {
	resourceType := "stackoverflow_answer"
	resourceName := "test"
	resourceIdentifier := fmt.Sprintf("%s.%s", resourceType, resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccStackOverflowAnswerCheckDestroy(resourceIdentifier),
		Steps: []resource.TestStep{
			{
				Config: testAccStackOverflowAnswerConfig(resourceType, resourceName, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccStackOverflowAnswerCheckExists(resourceIdentifier),
				),
			},
			{
				Config: testAccStackOverflowAnswerConfig(resourceType, resourceName, " Updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccStackOverflowAnswerCheckExists(resourceIdentifier),
				),
			},
		},
	})
}

func testAccStackOverflowAnswerConfig(resourceType, resourceName, update string) string {
	return fmt.Sprintf(`resource "stackoverflow_question" "test" {
		title = "How do I run the Terraform Provider Acceptance Tests?"
		body_markdown = "I want to run the Terraform Provider for Stack Overflow acceptance tests, how do I do that?"
		tags = ["test"]
	}
	resource "%s" "%s" {
		question_id = stackoverflow_question.test.id
		body_markdown = "Unit test%s."
	}`, resourceType, resourceName, update)
}

func testAccStackOverflowAnswerCheckDestroy(resourceIdentifier string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState, found := s.RootModule().Resources[resourceIdentifier]

		if !found {
			return fmt.Errorf("No resource found with id %s", resourceIdentifier)
		}

		if resourceState.Primary.ID == "" {
			return fmt.Errorf("Resource with ID is empty for %s", resourceIdentifier)
		}

		answerID, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		questionID, err := strconv.Atoi(resourceState.Primary.Attributes["question_id"])
		if err != nil {
			return err
		}

		c := testAccProvider.Meta().(*so.Client)
		answer, err := c.GetAnswer(&questionID, &answerID)

		if err != nil {
			return err
		}

		if answer != nil && !answer.IsDeleted {
			return fmt.Errorf("Resource still exists for %s", resourceIdentifier)
		}

		return nil
	}
}

func testAccStackOverflowAnswerCheckExists(n string) resource.TestCheckFunc {
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
