package main

import (
	"github.com/hashicorp/terraform/plugin"
	//SYMANTEC ODIN CHANGE. Using relative import, to allow local clone
	// and install of odin branch of this plugin:
	// "github.com/terraform-providers/terraform-provider-aws/aws"
	"terraform-provider-aws/aws"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: aws.Provider})
}
