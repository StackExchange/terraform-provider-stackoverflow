---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stackoverflow_question Data Source - terraform-provider-stackoverflow"
subcategory: ""
description: |-

---

# stackoverflow_question (Data Source)

The `question` data source allows for referencing an existing question in Stack Overflow.

## Example

```
data "stackoverflow_question" "question" {
    question_id = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `question_id` (Number) The identifier for the question

### Read-Only

- `body_markdown` (String) The question content in Markdown format
- `id` (String) The ID of this resource.
- `tags` (List of String) The set of tags to be associated with the article
- `title` (String) The title of the article


