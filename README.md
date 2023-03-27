# Terraform Provider Stack Overflow

Run the following command to build the provider

```powershell
# Windows
go build -o terraform-provider-stackoverflow.exe
cp terraform-provider-stackoverflow.exe ~/go/bin
```

Create a `terraform.rc` file in the %APPDATA% directory.

```powershell
cd $env:APPDATA
mk terraform.rc
```

Add the following content to the `terraform.rc` file:

```
provider_installation {
  dev_overrides {
    "registry.terraform.io/hashicorp/stackoverflow" = "C:/Users/rbolhofer/go/bin"
  }
  direct {}
}
```

## Test sample configuration

In a new directory, create a file named `main.tf` and add the following content:

```
resource "stackoverflow_article" "test_article" {
  article_type = "knowledge-article"
  title = "Test Article"
  body_markdown = "# Hello World"
  tags = ["test"]
}
```

Then initialize and run Terraform:

```powershell
terraform init
terraform plan -out terraform.tfplan
```