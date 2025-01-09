---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "stackoverflow_answer Data Source - terraform-provider-stackoverflow"
subcategory: ""
description: |-

---

# stackoverflow_answer (Data Source)

The `answer` data source allows for referencing an existing answer in Stack Overflow.

## Example

```
data "stackoverflow_answer" "answer" {
    answer_id = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `answer_id` (Number) The identifier for the answer
- `question_id` (Number) The identifier for the question

### Read-Only

- `body_markdown` (String) The answer content in Markdown format
- `id` (String) The ID of this resource.

