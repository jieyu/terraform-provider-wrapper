package main

import (
	goplugin "github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jieyu/terraform-provider-wrapper/pkg/wrapper"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

func main() {
	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: plugin.Handshake,
		Plugins: map[string]goplugin.Plugin{
			"provider": &plugin.ResourceProviderPlugin{
				F: func() terraform.ResourceProvider {
					return wrapper.NewProvider(aws.Provider())
				},
			},
			"provisioner": &plugin.ResourceProvisionerPlugin{
				F: nil,
			},
		},
	})
}
