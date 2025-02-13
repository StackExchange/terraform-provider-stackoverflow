---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stackoverflow_collection Resource - terraform-provider-stackoverflow"
subcategory: ""
description: |-

---

# stackoverflow_collection (Resource)

Manages a collection.

## Example

```
data "stackoverflow_tag" "tag" {
    tag_id = 1
}

resource "stackoverflow_article" "article" {
  article_type = "announcement"
  title = "Terraform Provider for Stack Overflow is available!"
  body_markdown = "Look for the Stack Overflow provider in the Terraform registry"
  tags = [data.stackoverflow_tag.tag.name]
}

resource "stackoverflow_question" "question" {
    title = "Stack Overflow Terraform Provider"
    body_markdown = "What is the Terraform Provider for Stack Overflow?"
    tags = [data.stackoverflow_tag.tag.name]
}

resource "stackoverflow_collection" "collection" {
  title = "Collection Name"
  description = "Example collection"
  content_ids = [
    stackoverflow_article.article.id,
    stackoverflow_question.question.id
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The collection description.
- `title` (String) The title of the collection

### Optional

- `content_ids` (List of Number)

### Read-Only

- `id` (String) The ID of this resource.


