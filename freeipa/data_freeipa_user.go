package freeipa

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	ipa "github.com/tehwalris/go-freeipa/freeipa"
)

func dataFreeIPAUser() *schema.Resource {
	return &schema.Resource{
		Read: dataFreeIPAUserRead,
		Schema: map[string]*schema.Schema{
			// Identity settings
			"jobtitle": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job title",
			},
			"first_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "First name",
			},
			"last_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last name",
			},
			"full_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Full name",
			},
			"display_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Display name",
			},
			"initials": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Initials",
			},
			"gecos": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GECOS",
			},
			"userclass": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User classes",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			// Account settings
			"uid": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "User login",
			},
			"random": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Generate a random password",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Password",
			},
			//"krb_password_expiration": &schema.Schema{
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Description: "User password expiration",
			//},
			//"uidnumber": &schema.Schema{
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Description: "User ID number",
			//},
			//"gidnumber": &schema.Schema{
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Description: "Group ID number",
			//},
			"krb_principal_name": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Kerberos principal aliases",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			//"krb_principal_expiration": &schema.Schema{
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Description: "Kerberos principal expiration",
			//},
			"login_shell": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Login shell",
			},
			"home_directory": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Home directory",
			},
			"ssh_public_key": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SSH public keys",
			},
			//"user_certificate": &schema.Schema{
			//	Type:     schema.TypeString,
			//	Computed: true,
			//},
			"userauth_type": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User authentication types",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"radius_config_link": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "RADIUS proxy configuration",
			},
			"radius_username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "RADIUS proxy username",
			},
			"account_disabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Account disabled",
			},
			//"no_private_group": &schema.Schema{
			//	Type:        schema.TypeBool,
			//	Default:     false,
			//	Description: "Don't create user private group",
			//},
			//"no_members": &schema.Schema{
			//	Type:        schema.TypeBool,
			//	Default:     false,
			//	Description: "Suppress processing of membership attributes.",
			//},
			"randompassword": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			// Contact settings
			"email": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Email addresses",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"telephone": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Telephone number",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"mobile": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Mobile telephone number",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"pager": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Pager number",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"fax": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Fax number",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			// Mailing settings
			"street": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Street address",
			},
			"city": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "City",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State/Province",
			},
			"zip": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ZIP",
			},

			// Employee information
			"orgunit": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Org. unit",
			},
			"manager": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Manager",
			},
			"department": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Department numbers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"employee_number": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Employee number",
			},
			"employee_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Employee type",
			},
			"preferred_language": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Preperred language",
			},

			// Misc. information
			"carlicense": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Car Licenses",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataFreeIPAUserRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Reading User: %s", d.Get("uid").(string))
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	uid := d.Get("uid").(string)

	res, err := client.UserShow(
		&ipa.UserShowArgs{},
		&ipa.UserShowOptionalArgs{
			UID: ipa.String(uid),
			All: ipa.Bool(true),
		},
	)
	if err != nil {
		return err
	}

	// Identity settings
	if res.Result.Title != nil {
		d.Set("jobtitle", *res.Result.Title)
	}
	if res.Result.Givenname != nil {
		d.Set("first_name", *res.Result.Givenname)
	}
	if res.Result.Sn != "" {
		d.Set("last_name", res.Result.Sn)
	}
	if res.Result.Cn != nil {
		d.Set("full_name", *res.Result.Cn)
	}
	if res.Result.Displayname != nil {
		d.Set("display_name", *res.Result.Displayname)
	}
	if res.Result.Initials != nil {
		d.Set("initials", *res.Result.Initials)
	}
	if res.Result.Gecos != nil {
		d.Set("gecos", *res.Result.Gecos)
	}
	if res.Result.Userclass != nil {
		d.Set("userclass", *res.Result.Userclass)
	}

	// Account settings
	if res.Result.Random != nil {
		d.Set("random", *res.Result.Random)
	}
	//if res.Result.Userpassword != nil {
	//	d.Set("password", *res.Result.Userpassword)
	//}
	//if res.Result.Krbpasswordexpiration != nil {
	//	d.Set("krb_password_expiration", *res.Result.Krbpasswordexpiration)
	//}
	//if res.Result.Uidnumber != nil {
	//	d.Set("uidnumber", *res.Result.Uidnumber)
	//}
	//if res.Result.Gidnumber != nil {
	//	d.Set("gidnumber", *res.Result.Gidnumber)
	//}
	if res.Result.Krbprincipalname != nil {
		d.Set("krb_principal_name", *res.Result.Krbprincipalname)
	}
	//if res.Result.Krbprincipalexpiration != nil {
	//	d.Set("krb_principal_expiration", *res.Result.Krbprincipalexpiration)
	//}
	if res.Result.Loginshell != nil {
		d.Set("login_shell", *res.Result.Loginshell)
	}
	if res.Result.Homedirectory != nil {
		d.Set("home_directory", *res.Result.Homedirectory)
	}
	if res.Result.Ipasshpubkey != nil {
		d.Set("ssh_public_key", *res.Result.Ipasshpubkey)
	}
	//if res.Result.Usercertificate != nil {
	//	d.Set("user_certificate", *res.Result.Usercertificate)
	//}
	if res.Result.Ipauserauthtype != nil {
		d.Set("userauth_type", *res.Result.Ipauserauthtype)
	}
	if res.Result.Ipatokenradiusconfiglink != nil {
		d.Set("radius_config_link", *res.Result.Ipatokenradiusconfiglink)
	}
	if res.Result.Ipatokenradiususername != nil {
		d.Set("radius_username", *res.Result.Ipatokenradiususername)
	}
	if res.Result.Nsaccountlock != nil {
		d.Set("account_disabled", *res.Result.Nsaccountlock)
	}
	//if res.Result.Noprivate != nil {
	//	d.Set("no_private_group", *res.Result.Noprivate)
	//}
	//if res.Result.NoMembers != nil {
	//	d.Set("no_members", *res.Result.NoMembers)
	//}

	// Contact settings
	if res.Result.Mail != nil {
		d.Set("email", *res.Result.Mail)
	}
	if res.Result.Telephonenumber != nil {
		d.Set("telephone", *res.Result.Telephonenumber)
	}
	if res.Result.Mobile != nil {
		d.Set("mobile", *res.Result.Mobile)
	}
	if res.Result.Pager != nil {
		d.Set("pager", *res.Result.Pager)
	}
	if res.Result.Facsimiletelephonenumber != nil {
		d.Set("fax", *res.Result.Facsimiletelephonenumber)
	}

	// Mailing address
	if res.Result.Street != nil {
		d.Set("street", *res.Result.Street)
	}
	if res.Result.L != nil {
		d.Set("city", *res.Result.L)
	}
	if res.Result.St != nil {
		d.Set("state", *res.Result.St)
	}
	if res.Result.Postalcode != nil {
		d.Set("zip", *res.Result.Postalcode)
	}

	// Employee information
	if res.Result.Ou != nil {
		d.Set("orgunit", *res.Result.Ou)
	}
	if res.Result.Manager != nil {
		d.Set("manager", *res.Result.Manager)
	}
	if res.Result.Departmentnumber != nil {
		d.Set("department", *res.Result.Departmentnumber)
	}
	if res.Result.Employeenumber != nil {
		d.Set("employee_number", *res.Result.Employeenumber)
	}
	if res.Result.Employeetype != nil {
		d.Set("employee_type", *res.Result.Employeetype)
	}
	if res.Result.Preferredlanguage != nil {
		d.Set("preferred_language", *res.Result.Preferredlanguage)
	}

	// Misc. information
	if res.Result.Carlicense != nil {
		d.Set("carlicense", *res.Result.Carlicense)
	}

	d.SetId(uid)

	return nil
}
