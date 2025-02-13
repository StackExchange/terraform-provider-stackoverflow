# Terraform Provider for Stack Overflow

The Terraform Provider for Stack Overflow is a Terraform plugin provider that allows you to manage questions, answers, and articles for your Stack Overflow for Teams.

## Using the Provider
---------------------

To use a released version of the Terraform provider in your environment, run `terraform init` and Terraform will automatically install the provider from the Terraform Registry. To specify a particular provider version when installing released providers, see the [Terraform documentation on provider versioning](https://www.terraform.io/docs/configuration/providers.html#version-provider-versions).

## Example
----------

```terraform
provider "stackoverflow" {
  base_url = "https://api.stackoverflowteams.com/v3/teams/{team}/"
  access_token = "xxxx"
}

resource "stackoverflow_article" "article" {
  article_type = "announcement"
  title = "Terraform Provider for Stack Overflow is available!"
  body_markdown = "Look for the Stack Overflow provider in the Terraform registry"
  tags = ["example"]
}

resource "stackoverflow_question" "question" {
    title = "Stack Overflow Terraform Provider"
    body_markdown = "What is the Terraform Provider for Stack Overflow?"
    tags = ["example"]
}

resource "stackoverflow_answer" "answer" {
    question_id = stackoverflow_question.question.id
    body_markdown = "It is a Terraform plugin provider to manage resources in Stack Overflow for Teams"
}
```
