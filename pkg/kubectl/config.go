package kubectl

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type ProviderInstance struct {
}

func New(_ *schema.ResourceData) *ProviderInstance {
	return &ProviderInstance{}
}
