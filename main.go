package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/mumoshu/terraform-provider-kubectl/pkg/kubectl"
	"github.com/mumoshu/terraform-provider-kubectl/pkg/profile"
)

func main() {
	defer profile.Start().Stop()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kubectl.Provider})
}
