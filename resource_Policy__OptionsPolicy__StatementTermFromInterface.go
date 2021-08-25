
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlPolicy__OptionsPolicy__StatementTermFromInterface struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy__statement  struct {
			XMLName xml.Name `xml:"policy-statement"`
			V_name  string  `xml:"name"`
			V_term  struct {
				XMLName xml.Name `xml:"term"`
				V_name__1  string  `xml:"name"`
				V_from  struct {
					XMLName xml.Name `xml:"from"`
					V_interface  string  `xml:"interface"`
				} `xml:"from"`
			} `xml:"term"`
		} `xml:"policy-options>policy-statement"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosPolicy__OptionsPolicy__StatementTermFromInterfaceCreate(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_interface := d.Get("interface").(string)
	commit := true

	config := xmlPolicy__OptionsPolicy__StatementTermFromInterface{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = V_name
	config.Groups.V_policy__statement.V_term.V_name__1 = V_name__1
	config.Groups.V_policy__statement.V_term.V_from.V_interface = V_interface

    err = client.SendTransaction("", config, commit)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosPolicy__OptionsPolicy__StatementTermFromInterfaceRead(d,m)
}

func junosPolicy__OptionsPolicy__StatementTermFromInterfaceRead(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlPolicy__OptionsPolicy__StatementTermFromInterface{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_policy__statement.V_name)
	d.Set("name__1", config.Groups.V_policy__statement.V_term.V_name__1)
	d.Set("interface", config.Groups.V_policy__statement.V_term.V_from.V_interface)

	return nil
}

func junosPolicy__OptionsPolicy__StatementTermFromInterfaceUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_interface := d.Get("interface").(string)
	commit := true

	config := xmlPolicy__OptionsPolicy__StatementTermFromInterface{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = V_name
	config.Groups.V_policy__statement.V_term.V_name__1 = V_name__1
	config.Groups.V_policy__statement.V_term.V_from.V_interface = V_interface

    err = client.SendTransaction(id, config, commit)
    check(err)
    
	return junosPolicy__OptionsPolicy__StatementTermFromInterfaceRead(d,m)
}

func junosPolicy__OptionsPolicy__StatementTermFromInterfaceDelete(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfig(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosPolicy__OptionsPolicy__StatementTermFromInterface() *schema.Resource {
	return &schema.Resource{
		Create: junosPolicy__OptionsPolicy__StatementTermFromInterfaceCreate,
		Read: junosPolicy__OptionsPolicy__StatementTermFromInterfaceRead,
		Update: junosPolicy__OptionsPolicy__StatementTermFromInterfaceUpdate,
		Delete: junosPolicy__OptionsPolicy__StatementTermFromInterfaceDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement",
			},
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement.V_term",
			},
			"interface": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement.V_term.V_from. Interface name or address",
			},
		},
	}
}