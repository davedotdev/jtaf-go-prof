
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlProtocolsBgpGroupPeer__As struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_group  struct {
			XMLName xml.Name `xml:"group"`
			V_name  string  `xml:"name"`
			V_peer__as  string  `xml:"peer-as"`
		} `xml:"protocols>bgp>group"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsBgpGroupPeer__AsCreate(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_peer__as := d.Get("peer__as").(string)
	commit := true

	config := xmlProtocolsBgpGroupPeer__As{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_group.V_name = V_name
	config.Groups.V_group.V_peer__as = V_peer__as

    err = client.SendTransaction("", config, commit)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsBgpGroupPeer__AsRead(d,m)
}

func junosProtocolsBgpGroupPeer__AsRead(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsBgpGroupPeer__As{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_group.V_name)
	d.Set("peer__as", config.Groups.V_group.V_peer__as)

	return nil
}

func junosProtocolsBgpGroupPeer__AsUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_peer__as := d.Get("peer__as").(string)
	commit := true

	config := xmlProtocolsBgpGroupPeer__As{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_group.V_name = V_name
	config.Groups.V_group.V_peer__as = V_peer__as

    err = client.SendTransaction(id, config, commit)
    check(err)
    
	return junosProtocolsBgpGroupPeer__AsRead(d,m)
}

func junosProtocolsBgpGroupPeer__AsDelete(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfig(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosProtocolsBgpGroupPeer__As() *schema.Resource {
	return &schema.Resource{
		Create: junosProtocolsBgpGroupPeer__AsCreate,
		Read: junosProtocolsBgpGroupPeer__AsRead,
		Update: junosProtocolsBgpGroupPeer__AsUpdate,
		Delete: junosProtocolsBgpGroupPeer__AsDelete,

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
			"peer__as": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.  Autonomous system number in plain number or 'higher 16bits'.'Lower 16 bits' (asdot notation) format",
			},
		},
	}
}