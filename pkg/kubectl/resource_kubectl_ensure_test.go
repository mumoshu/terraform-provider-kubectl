package kubectl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccHelmfileReleaseSet_basic(t *testing.T) {
	resourceName := "kubectl_ensure.meta"
	releaseID := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckShellScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHelmfileReleaseSetConfig_basic(releaseID),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "labels.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "annotations.%", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func TestAccHelmfileReleaseSet_binaries(t *testing.T) {
	resourceName := "kubectl_ensure.meta"
	releaseID := acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckShellScriptDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHelmfileReleaseSetConfig_binaries(releaseID),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "labels.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "annotations.%", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func testAccCheckShellScriptDestroy(s *terraform.State) error {
	_ = testAccProvider.Meta().(*ProviderInstance)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "kubectl_ensure" {
			continue
		}
		//
		//kubectlYaml := fmt.Sprintf("kubectl-%s.yaml", rs.Primary.ID)
		//
		//cmd := exec.Command("kubectl", "-f", kubectlYaml, "status")
		//if out, err := cmd.CombinedOutput(); err == nil {
		//	return fmt.Errorf("verifying kubectl status: releases still exist for %s", kubectlYaml)
		//} else if !strings.Contains(string(out), "Error: release: not found") {
		//	return fmt.Errorf("verifying kubectl status: unexpected error: %v:\n\nCOMBINED OUTPUT:\n%s", err, string(out))
		//}
	}
	return nil
}

func testAccHelmfileReleaseSetConfig_basic(randVal string) string {
	return fmt.Sprintf(`
resource "kubectl_ensure" "meta" {
  namespace = "kube-system"
  resource = "configmap"
  name = "aws-auth"

  labels = {
    "key1" = "%s"
    "key2" = "two"
  }

  annotations = {
    "key3" = "%s"
    "key4" = "four"
  }

  kubeconfig = pathexpand("~/.kube/config")
}
`, randVal, randVal)
}

func testAccHelmfileReleaseSetConfig_binaries(randVal string) string {
	return fmt.Sprintf(`
resource "kubectl_ensure" "meta" {
  version = ">= 1.18.0, < 1.19.0"

  namespace = "kube-system"
  resource = "configmap"
  name = "aws-auth"

  labels = {
    "key1" = "%s"
    "key2" = "two"
  }

  annotations = {
    "key3" = "%s"
    "key4" = "four"
  }

  kubeconfig = pathexpand("~/.kube/config")
}
`, randVal, randVal)
}
