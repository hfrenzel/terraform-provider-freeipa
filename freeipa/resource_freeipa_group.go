package freeipa

import (
	"log"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	ipa "github.com/tehwalris/go-freeipa/freeipa"
)

func resourceFreeIPAGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFreeIPAGroupCreate,
		Read:   resourceFreeIPAGroupRead,
		Update: resourceFreeIPAGroupUpdate,
		Delete: resourceFreeIPAGroupDelete,
		Importer: &schema.ResourceImporter{
			State: resourceFreeIPAGroupImport,
		},

		Schema: map[string]*schema.Schema{
			"group": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Group name",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description",
			},
			//"gidnumber": &schema.Schema{
			//	Type:        schema.TypeInt,
			//	Optional:    true,
			//	Computed:    true,
			//	Description: "Group ID number",
			//},
			"non_posix": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Is a non-POSIX group",
			},
			//"no_members": &schema.Schema{
			//	Type:        schema.TypeBool,
			//	Optional:    true,
			//	Default:     false,
			//	Description: "Suppress processing of membership attributes.",
			//},
		},
	}
}

func resourceFreeIPAGroupCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Creating Group: %s", d.Get("group").(string))
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	args := ipa.GroupAddArgs{}
	optArgs := ipa.GroupAddOptionalArgs{}

	args.Cn = d.Get("group").(string)

	if v, ok := d.GetOk("description"); ok {
		_v := ipa.String(v.(string))
		optArgs.Description = _v
	}
	//if v, ok := d.GetOk("gidnumber"); ok {
	//	_v := ipa.Int(v.(int))
	//	optArgs.Gidnumber = _v
	//}
	if v, ok := d.GetOk("non_posix"); ok {
		_v := ipa.Bool(v.(bool))
		optArgs.Nonposix = _v
	}
	//if v, ok := d.GetOk("no_members"); ok {
	//	_v := ipa.Bool(v.(bool))
	//	optArgs.NoMembers = _v
	//}

	res, err := client.GroupAdd(
		&args,
		&optArgs,
	)
	if err != nil {
		return err
	}

	d.SetId(res.Result.Cn)

	return nil
}

func resourceFreeIPAGroupRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Refreshing Group: %s", d.Id())
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

func resourceFreeIPAGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Updating Group: %s", d.Get("group").(string))
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	args := ipa.GroupModArgs{}
	optArgs := ipa.GroupModOptionalArgs{}

	args.Cn = d.Get("group").(string)

	if v, ok := d.GetOk("description"); ok {
		_v := ipa.String(v.(string))
		optArgs.Description = _v
	}
	//if v, ok := d.GetOk("gidnumber"); ok {
	//	_v := ipa.Int(v.(int))
	//	optArgs.Gidnumber = _v
	//}
	//if v, ok := d.GetOk("no_members"); ok {
	//	_v := ipa.Bool(v.(bool))
	//	optArgs.NoMembers = _v
	//}

	_, err = client.GroupMod(
		&args,
		&optArgs,
	)
	if err != nil {
		return err
	}

	return resourceFreeIPAGroupRead(d, meta)
}

func resourceFreeIPAGroupDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Deleting Group: %s", d.Get("group").(string))
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	group := d.Get("group").([]string)

	_, err = client.GroupDel(
		&ipa.GroupDelArgs{
			Cn: group,
		},
		&ipa.GroupDelOptionalArgs{},
	)

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceFreeIPAGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.SetId(d.Id())
	d.Set("group", d.Id())

	err := resourceFreeIPAGroupRead(d, meta)
	if err != nil {
		return []*schema.ResourceData{}, err
	}

	return []*schema.ResourceData{d}, nil
}
