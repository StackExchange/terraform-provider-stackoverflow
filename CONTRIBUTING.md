# Overview

This file will detail instructions on how to develop the Stack Overflow Terraform provider plugin locally.

The original [legacy] documentation used for building custom Terraform providers can be found here:

[https://www.hashicorp.com/en/blog/writing-custom-terraform-providers](https://www.hashicorp.com/en/blog/writing-custom-terraform-providers)

## Getting started

Install Go, ensure that the version you are using meets or exceeds the version listed in the [go.mod](./go.mod) file.

**Build:**

> Test the built executable after building to ensure that it runs correctly.

*Windows:*

```
go build -o terraform-provider-stackoverflow.exe
./terraform-provider-stackoverflow.exe
```

*Linux*

```
go build -o terraform-provider-stackoverflow
./terraform-provider-stackoverflow
```

## Local Development Override

To use the locally built provider, override the `terraform.rc` (Windows) or `.terraformrc` (Linux) file with the following:

`terraform.rc`:
```
provider_installation {
  dev_overrides {
    "StackExchange/stackoverflow" = "~/source/GitHub/StackExchange/terraform-provider-stackoverflow"
  }

  direct {}
}
```

> Note, make sure the provider path matches where your actual repository is located.
> The `terraform.rc` file is located in the user's `%APPDATA%` directory on Windows and `.terraformrc` file is located in the user's home directory (`~`) on Linux

Once you have updated the `terraform.rc` file, you can setup a local testing repository and reference the Stack Overflow Terraform provider.  You _DO NOT_ need to run `terraform init` when using local development provider overrides!  Also, if you encounter any errors, ensure that you remove your `.terraform` folder and `.terraform.lock.hcl` to prevent any version conflicts.

## Acceptance Tests

Acceptence tests (or integration tests) are automatically run from the [test.yml](./github/workflows/test.yml) GitHub Actions Workflow for all pushes and pull requests against a developmentment Stack Overflow Enterprise (SOE) instance.  This instance needs to have the following configurations applied to ensure that the tests do not fail quality control checks that prevent duplicated content.

* Disable content length checking, `QualityChecks.Questions.Rejections.QualityThreshold`: `1`
* Disable duplicate content checking, `Questions.Testing.EnableDuplicateCheckForQuestions`: `False`

SOE site settings can be accessed via the `/developer/site-settings` route.  Additional details can be found on the documentation help site: [https://stackoverflowteams.help/en/articles/9859814-question-quality-and-duplicate-checks](https://stackoverflowteams.help/en/articles/9859814-question-quality-and-duplicate-checks)

> For details on setting up authentication, check the [docs/index.md#authentication-and-configuration](./docs/index.md#authentication-and-configuration) section.
