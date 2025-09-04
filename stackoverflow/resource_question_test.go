package stackoverflow

import (
	"fmt"
	"strconv"
	so "terraform-provider-stackoverflow/stackoverflow/client"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccStackOverflowQuestion(t *testing.T) {
	resourceType := "stackoverflow_question"
	resourceName := "test"
	resourceIdentifier := fmt.Sprintf("%s.%s", resourceType, resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccStackOverflowQuestionCheckDestroy(resourceIdentifier),
		Steps: []resource.TestStep{
			{
				Config: testAccStackOverflowQuestionConfig(resourceType, resourceName, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccStackOverflowQuestionCheckExists(resourceIdentifier),
				),
			},
			{
				Config: testAccStackOverflowQuestionConfig(resourceType, resourceName, " Updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccStackOverflowQuestionCheckExists(resourceIdentifier),
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

func testAccStackOverflowQuestionConfig(resourceType, resourceName, update string) string {
	return fmt.Sprintf(`resource "%s" "%s" {
		title = "How do I run the Terraform Provider Acceptance Tests%s?"
		body_markdown = "I want to run the Terraform Provider for Stack Overflow acceptance tests, how do I do that%s?"
		tags = ["test"]
	}`, resourceType, resourceName, update, update)
}

func testAccStackOverflowQuestionCheckDestroy(resourceIdentifier string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resourceState, found := s.RootModule().Resources[resourceIdentifier]

		if !found {
			return fmt.Errorf("No resource found with id %s", resourceIdentifier)
		}

		if resourceState.Primary.ID == "" {
			return fmt.Errorf("Resource with ID is empty for %s", resourceIdentifier)
		}

		questionID, err := strconv.Atoi(resourceState.Primary.ID)

		if err != nil {
			return err
		}

		c := testAccProvider.Meta().(*so.Client)
		question, err := c.GetQuestion(&questionID)

		if err != nil {
			return err
		}

		if question != nil && !question.IsDeleted {
			return fmt.Errorf("Question with ID %d still exists", questionID)
		}

		return nil
	}
}

func testAccStackOverflowQuestionCheckExists(n string) resource.TestCheckFunc {
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
