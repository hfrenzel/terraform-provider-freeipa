package freeipa

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	ipa "github.com/tehwalris/go-freeipa/freeipa"
)

func dataFreeIPAGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataFreeIPAGroupRead,
		Schema: map[string]*schema.Schema{
			"group": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Group name",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description",
			},
			//"gidnumber": &schema.Schema{
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Description: "Group ID number",
			//},
			//"no_members": &schema.Schema{
			//	Type:        schema.TypeBool,
			//	Default:     false,
			//	Description: "Suppress processing of membership attributes.",
			//},
		},
	}
}

func dataFreeIPAGroupRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading Group: %s", d.Get("group").(string))
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	group := d.Get("group").(string)

	res, err := client.GroupShow(
		&ipa.GroupShowArgs{
			Cn: group,
		},
		&ipa.GroupShowOptionalArgs{
			All: ipa.Bool(true),
		},
	)
	if err != nil {
		return err
	}

	if res.Result.Description != nil {
		d.Set("description", *res.Result.Description)
	}
	//if res.Result.Gidnumber != nil {
	//	d.Set("gidnumber", *res.Result.Gidnumber)
	//}
	//if res.Result.NoMembers != nil {
	//	d.Set("no_members", *res.Result.NoMembers)
	//}

	d.SetId(group)

	return nil
}
