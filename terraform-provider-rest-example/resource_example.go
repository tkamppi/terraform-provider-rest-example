package main

import (
	"fmt"
	"httpbin_client/client"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	name := d.Get("name").(string)
	err := apiClient.Post()
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	d.SetId(name)
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	err := apiClient.Get()
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	err := apiClient.Delete()
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	// d.SetId("") it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
