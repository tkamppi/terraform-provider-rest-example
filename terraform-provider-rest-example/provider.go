package main

import (
	"httpbin_client/client"

	"github.com/hashicorp/terraform/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HTTPBIN_HOSTNAME", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"restexample_server": resourceServer(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	hostname := d.Get("hostname").(string)
	return client.NewClient(hostname), nil
}
