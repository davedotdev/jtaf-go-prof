
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlProtocolsBgpGroupNeighborExport struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_group  struct {
			XMLName xml.Name `xml:"group"`
			V_name  string  `xml:"name"`
			V_neighbor  struct {
				XMLName xml.Name `xml:"neighbor"`
				V_name__1  string  `xml:"name"`
				V_export  string  `xml:"export"`
			} `xml:"neighbor"`
		} `xml:"protocols>bgp>group"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsBgpGroupNeighborExportCreate(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_export := d.Get("export").(string)
	commit := true

	config := xmlProtocolsBgpGroupNeighborExport{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_group.V_name = V_name
	config.Groups.V_group.V_neighbor.V_name__1 = V_name__1
	config.Groups.V_group.V_neighbor.V_export = V_export

    err = client.SendTransaction("", config, commit)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsBgpGroupNeighborExportRead(d,m)
}

func junosProtocolsBgpGroupNeighborExportRead(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsBgpGroupNeighborExport{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_group.V_name)
	d.Set("name__1", config.Groups.V_group.V_neighbor.V_name__1)
	d.Set("export", config.Groups.V_group.V_neighbor.V_export)

	return nil
}

func junosProtocolsBgpGroupNeighborExportUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_export := d.Get("export").(string)
	commit := true

	config := xmlProtocolsBgpGroupNeighborExport{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_group.V_name = V_name
	config.Groups.V_group.V_neighbor.V_name__1 = V_name__1
	config.Groups.V_group.V_neighbor.V_export = V_export

    err = client.SendTransaction(id, config, commit)
    check(err)
    
	return junosProtocolsBgpGroupNeighborExportRead(d,m)
}

func junosProtocolsBgpGroupNeighborExportDelete(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfig(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosProtocolsBgpGroupNeighborExport() *schema.Resource {
	return &schema.Resource{
		Create: junosProtocolsBgpGroupNeighborExportCreate,
		Read: junosProtocolsBgpGroupNeighborExportRead,
		Update: junosProtocolsBgpGroupNeighborExportUpdate,
		Delete: junosProtocolsBgpGroupNeighborExportDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group",
			},
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_neighbor",
			},
			"export": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_neighbor. Export policy",
			},
		},
	}
}