package stackoverflow

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"stackoverflow": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("STACK_OVERFLOW_ACCESS_TOKEN"); err == "" {
		t.Fatalf("STACK_OVERFLOW_ACCESS_TOKEN must be set for acceptance tests")
	}

	if err := os.Getenv("STACK_OVERFLOW_API_URL"); err == "" {
		t.Fatalf("STACK_OVERFLOW_API_URL must be set for acceptance tests")
	}
}
