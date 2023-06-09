---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stackoverflow_article Resource - terraform-provider-stackoverflow"
subcategory: ""
description: |-

---

# stackoverflow_article (Resource)

Manages a long form article. At least one tag is required which can be either be set as a default tag on the provider or set at the resource level.

## Example

```
resource "stackoverflow_article" "article" {
  article_type = "announcement"
  title = "Terraform Provider for Stack Overflow is available!"
  body_markdown = "Look for the Stack Overflow provider in the Terraform registry"
  tags = ["example"]
  filter = "XXXX"
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `article_type` (String) The type of article. Must be one of `knowledge-article`, `announcement`, `how-to-guide`, `policy`
- `filter` (String) The API filter to use
- `title` (String) The title of the article

### Optional

- `body_markdown` (String) The article content in Markdown format
- `tags` (List of String) The set of tags to be associated with the article

### Read-Only

- `id` (String) The ID of this resource.


