package freeipa

import (
	"log"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	ipa "github.com/tehwalris/go-freeipa/freeipa"
)

func resourceFreeIPAUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceFreeIPAUserCreate,
		Read:   resourceFreeIPAUserRead,
		Update: resourceFreeIPAUserUpdate,
		Delete: resourceFreeIPAUserDelete,
		Importer: &schema.ResourceImporter{
			State: resourceFreeIPAUserImport,
		},

		Schema: map[string]*schema.Schema{
			// Identity settings
			"jobtitle": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Job title",
			},
			"first_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "First name",
			},
			"last_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Last name",
			},
			"full_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Full name",
			},
			"display_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Display name",
			},
			"initials": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Initials",
			},
			"gecos": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "GECOS",
			},
			"userclass": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "User classes",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			// Account settings
			"uid": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "User login",
			},
			"random": &schema.Schema{
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
				Description: "Generate a random password",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Sensitive:   true,
				Description: "Password",
			},
			//"krb_password_expiration": &schema.Schema{
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//	Description: "User password expiration",
			//},
			//"uidnumber": &schema.Schema{
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//	Description: "User ID number",
			//},
			//"gidnumber": &schema.Schema{
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//	Description: "Group ID number",
			//},
			"krb_principal_name": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "Kerberos principal aliases",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			//"krb_principal_expiration": &schema.Schema{
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//	Description: "Kerberos principal expiration",
			//},
			"login_shell": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Login shell",
			},
			"home_directory": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Home directory",
			},
			"ssh_public_key": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "SSH public keys",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			//"user_certificate": &schema.Schema{
			//	Type:     schema.TypeString,
			//	//Computed: true,
			//	Optional: true,
			//},
			"userauth_type": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "User authentication types",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"radius_config_link": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "RADIUS proxy configuration",
			},
			"radius_username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "RADIUS proxy username",
			},
			"account_disabled": &schema.Schema{
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				Description: "Account disabled",
			},
			//"no_private_group": &schema.Schema{
			//	Type:        schema.TypeBool,
			//	Default:     false,
			//	Optional:    true,
			//	Description: "Don't create user private group",
			//},
			//"no_members": &schema.Schema{
			//	Type:        schema.TypeBool,
			//	Default:     false,
			//	Optional:    true,
			//	Description: "Suppress processing of membership attributes.",
			//},
			"randompassword": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},

			// Contact settings
			"email": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "Email addresses",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"telephone": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Telephone number",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"mobile": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Mobile telephone number",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"pager": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Pager number",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"fax": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Fax number",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			// Mailing address
			"street": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Street address",
			},
			"city": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "City",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "State/Province",
			},
			"zip": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ZIP",
			},

			// Employee information
			"orgunit": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Org. unit",
			},
			"manager": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Manager",
			},
			"department": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Department numbers",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"employee_number": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Employee number",
			},
			"employee_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Employee type",
			},
			"preferred_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Preperred language",
			},

			// Misc. information
			"carlicense": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Car Licenses",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceFreeIPAUserCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Creating User: %s", d.Id())
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	args := ipa.UserAddArgs{}
	optArgs := ipa.UserAddOptionalArgs{}

	// Identity settings
	if v, ok := d.GetOk("jobtitle"); ok {
		_v := v.(string)
		optArgs.Title = &_v
	}
	if v, ok := d.GetOk("first_name"); ok {
		_v := v.(string)
		args.Givenname = _v
	}
	if v, ok := d.GetOk("last_name"); ok {
		_v := v.(string)
		args.Sn = _v
	}
	if v, ok := d.GetOk("full_name"); ok {
		_v := ipa.String(v.(string))
		optArgs.Cn = _v
	}
	if v, ok := d.GetOk("display_name"); ok {
		_v := ipa.String(v.(string))
		optArgs.Displayname = _v
	}
	if v, ok := d.GetOk("initials"); ok {
		_v := ipa.String(v.(string))
		optArgs.Initials = _v
	}
	if v, ok := d.GetOk("gecos"); ok {
		_v := ipa.String(v.(string))
		optArgs.Gecos = _v
	}
	if v, ok := d.GetOk("userclass"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Userclass = &_v
	}

	// Account settings
	if v, ok := d.GetOk("uid"); ok {
		_v := ipa.String(v.(string))
		optArgs.UID = _v
	}
	if v, ok := d.GetOk("random"); ok {
		_v := ipa.Bool(v.(bool))
		optArgs.Random = _v
	}
	//if v, ok := d.GetOk("password"); ok {
	//	_v := ipa.String(v.(string))
	//	optArgs.Userpassword = _v
	//}
	//if v, ok := d.GetOk("krb_password_expiration"); ok {
	//	_v := v.(time.Time)
	//	optArgs.Krbpasswordexpiration = &_v
	//}
	//if v, ok := d.GetOk("uidnumber"); ok {
	//	_v := ipa.Int(v.(int))
	//	optArgs.Uidnumber = _v
	//}
	//if v, ok := d.GetOk("gidnumber"); ok {
	//	_v := ipa.Int(v.(int))
	//	optArgs.Gidnumber = _v
	//}
	if v, ok := d.GetOk("krb_principal_name"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Krbprincipalname = &_v
	}
	//if v, ok := d.GetOk("krb_principal_expiration"); ok {
	//	_v := ipa.String(v.(time.Time))
	//	optArgs.Krbprincipalexpiration = _v
	//}
	if v, ok := d.GetOk("login_shell"); ok {
		_v := ipa.String(v.(string))
		optArgs.Loginshell = _v
	}
	if v, ok := d.GetOk("home_directory"); ok {
		_v := ipa.String(v.(string))
		optArgs.Homedirectory = _v
	}
	if v, ok := d.GetOk("ssh_public_key"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Ipasshpubkey = &_v
	}
	//if v, ok := d.GetOk("user_certificate"); ok {
	//	_v := sliceStrings(v.([]interface{}))
	//	optArgs.Usercertificate = &_v
	//}
	// Certificate mapping data?
	if v, ok := d.GetOk("userauth_type"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Ipauserauthtype = &_v
	}
	if v, ok := d.GetOk("radius_config_link"); ok {
		_v := ipa.String(v.(string))
		optArgs.Ipatokenradiusconfiglink = _v
	}
	if v, ok := d.GetOk("radius_username"); ok {
		_v := ipa.String(v.(string))
		optArgs.Ipatokenradiususername = _v
	}
	if v, ok := d.GetOk("account_disabled"); ok {
		_v := ipa.Bool(v.(bool))
		optArgs.Nsaccountlock = _v
	}
	// TODO: depends on a group to assign to
	//if v, ok := d.GetOk("no_private_group"); ok {
	//	_v := ipa.Bool(v.(bool))
	//	optArgs.Noprivate = _v
	//}
	//if v, ok := d.GetOk("no_members"); ok {
	//	_v := ipa.Bool(v.(bool))
	//	optArgs.NoMembers = _v
	//}

	// Contact settings
	if v, ok := d.GetOk("email"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Mail = &_v
	}
	if v, ok := d.GetOk("telephone"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Telephonenumber = &_v
	}
	if v, ok := d.GetOk("pager"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Pager = &_v
	}
	if v, ok := d.GetOk("mobile"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Mobile = &_v
	}
	if v, ok := d.GetOk("fax"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Facsimiletelephonenumber = &_v
	}

	// Mailing address
	if v, ok := d.GetOk("street"); ok {
		_v := ipa.String(v.(string))
		optArgs.Street = _v
	}
	if v, ok := d.GetOk("city"); ok {
		_v := ipa.String(v.(string))
		optArgs.L = _v
	}
	if v, ok := d.GetOk("state"); ok {
		_v := ipa.String(v.(string))
		optArgs.St = _v
	}
	if v, ok := d.GetOk("zip"); ok {
		_v := ipa.String(v.(string))
		optArgs.Postalcode = _v
	}

	// Employee information
	if v, ok := d.GetOk("orgunit"); ok {
		_v := ipa.String(v.(string))
		optArgs.Ou = _v
	}
	if v, ok := d.GetOk("manager"); ok {
		_v := ipa.String(v.(string))
		optArgs.Manager = _v
	}
	if v, ok := d.GetOk("department"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Departmentnumber = &_v
	}
	if v, ok := d.GetOk("employee_number"); ok {
		_v := ipa.String(v.(string))
		optArgs.Employeenumber = _v
	}
	if v, ok := d.GetOk("employee_type"); ok {
		_v := ipa.String(v.(string))
		optArgs.Employeetype = _v
	}
	if v, ok := d.GetOk("preferred_language"); ok {
		_v := ipa.String(v.(string))
		optArgs.Preferredlanguage = _v
	}

	// Misc. information
	if v, ok := d.GetOk("carlicense"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Carlicense = &_v
	}

	res, err := client.UserAdd(
		&args,
		&optArgs,
	)
	if err != nil {
		return err
	}

	d.SetId(res.Result.UID)

	// randompassword is not returned by UserAdd
	if d.Get("random").(bool) {
		d.Set("randompassword", *res.Result.Randompassword)
	}

	return nil
}

func resourceFreeIPAUserRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Refreshing User: %s", d.Id())
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	uid := d.Id()

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

	return nil
}

func resourceFreeIPAUserUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Updating User: %s", d.Id())
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	optArgs := ipa.UserModOptionalArgs{}

	// Identity settings
	if v, ok := d.GetOk("jobtitle"); ok {
		_v := ipa.String(v.(string))
		optArgs.Title = _v
	}
	if v, ok := d.GetOk("first_name"); ok {
		_v := ipa.String(v.(string))
		optArgs.Givenname = _v
	}
	if v, ok := d.GetOk("last_name"); ok {
		_v := ipa.String(v.(string))
		optArgs.Sn = _v
	}
	if v, ok := d.GetOk("full_name"); ok {
		_v := ipa.String(v.(string))
		optArgs.Cn = _v
	}
	if v, ok := d.GetOk("display_name"); ok {
		_v := ipa.String(v.(string))
		optArgs.Displayname = _v
	}
	if v, ok := d.GetOk("initials"); ok {
		_v := ipa.String(v.(string))
		optArgs.Initials = _v
	}
	if v, ok := d.GetOk("gecos"); ok {
		_v := ipa.String(v.(string))
		optArgs.Gecos = _v
	}
	if v, ok := d.GetOk("userclass"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Userclass = &_v
	} else {
		optArgs.Userclass = &[]string{}
	}

	// Account settings
	if v, ok := d.GetOk("uid"); ok {
		_v := ipa.String(v.(string))
		optArgs.UID = _v
	}
	if v, ok := d.GetOk("random"); ok {
		_v := ipa.Bool(v.(bool))
		optArgs.Random = _v
	}
	//if v, ok := d.GetOk("password"); ok {
	//	_v := ipa.String(v.(string))
	//	optArgs.Userpassword = _v
	//}
	//if v, ok := d.GetOk("krb_password_expiration"); ok {
	//	_v := v.(time.Time)
	//	optArgs.Krbpasswordexpiration = &_v
	//}
	//if v, ok := d.GetOk("uidnumber"); ok {
	//	_v := ipa.Int(v.(int))
	//	optArgs.Uidnumber = _v
	//}
	//if v, ok := d.GetOk("gidnumber"); ok {
	//	_v := ipa.Int(v.(int))
	//	optArgs.Gidnumber = _v
	//}
	if v, ok := d.GetOk("krb_principal_name"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Krbprincipalname = &_v
	} else {
		optArgs.Krbprincipalname = &[]string{}
	}
	//if v, ok := d.GetOk("krb_principal_expiration"); ok {
	//	_v := ipa.String(v.(time.Time))
	//	optArgs.Krbprincipalexpiration = _v
	//}
	if v, ok := d.GetOk("login_shell"); ok {
		_v := ipa.String(v.(string))
		optArgs.Loginshell = _v
	}
	if v, ok := d.GetOk("home_directory"); ok {
		_v := ipa.String(v.(string))
		optArgs.Homedirectory = _v
	}
	if v, ok := d.GetOk("ssh_public_key"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Ipasshpubkey = &_v
	} else {
		optArgs.Ipasshpubkey = &[]string{}
	}
	//if v, ok := d.GetOk("user_certificate"); ok {
	//	_v := sliceStrings(v.([]interface{}))
	//	optArgs.Usercertificate = &_v
	//}
	// Certificate mapping data?
	if v, ok := d.GetOk("userauth_type"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Ipauserauthtype = &_v
	} else {
		optArgs.Ipauserauthtype = &[]string{}
	}
	if v, ok := d.GetOk("radius_config_link"); ok {
		_v := ipa.String(v.(string))
		optArgs.Ipatokenradiusconfiglink = _v
	}
	if v, ok := d.GetOk("radius_username"); ok {
		_v := ipa.String(v.(string))
		optArgs.Ipatokenradiususername = _v
	}
	if v, ok := d.GetOk("account_disabled"); ok {
		_v := ipa.Bool(v.(bool))
		optArgs.Nsaccountlock = _v
	}
	// TODO: depends on a group to assign to
	//if v, ok := d.GetOk("no_private_group"); ok {
	//	_v := ipa.Bool(v.(bool))
	//	optArgs.Noprivate = _v
	//}
	//if v, ok := d.GetOk("no_members"); ok {
	//	_v := ipa.Bool(v.(bool))
	//	optArgs.NoMembers = _v
	//}

	// Contact settings
	if v, ok := d.GetOk("email"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Mail = &_v
	} else {
		optArgs.Mail = &[]string{}
	}
	if v, ok := d.GetOk("telephone"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Telephonenumber = &_v
	} else {
		optArgs.Telephonenumber = &[]string{}
	}
	if v, ok := d.GetOk("pager"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Pager = &_v
	} else {
		optArgs.Pager = &[]string{}
	}
	if v, ok := d.GetOk("mobile"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Mobile = &_v
	} else {
		optArgs.Mobile = &[]string{}
	}
	if v, ok := d.GetOk("fax"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Facsimiletelephonenumber = &_v
	} else {
		optArgs.Facsimiletelephonenumber = &[]string{}
	}

	// Mailing address
	if v, ok := d.GetOk("street"); ok {
		_v := ipa.String(v.(string))
		optArgs.Street = _v
	}
	if v, ok := d.GetOk("city"); ok {
		_v := ipa.String(v.(string))
		optArgs.L = _v
	}
	if v, ok := d.GetOk("state"); ok {
		_v := ipa.String(v.(string))
		optArgs.St = _v
	}
	if v, ok := d.GetOk("zip"); ok {
		_v := ipa.String(v.(string))
		optArgs.Postalcode = _v
	}

	// Employee Information
	if v, ok := d.GetOk("orgunit"); ok {
		_v := ipa.String(v.(string))
		optArgs.Ou = _v
	}
	if v, ok := d.GetOk("manager"); ok {
		_v := ipa.String(v.(string))
		optArgs.Manager = _v
	}
	if v, ok := d.GetOk("department"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Departmentnumber = &_v
	} else {
		optArgs.Departmentnumber = &[]string{}
	}
	if v, ok := d.GetOk("employee_number"); ok {
		_v := ipa.String(v.(string))
		optArgs.Employeenumber = _v
	}
	if v, ok := d.GetOk("employee_type"); ok {
		_v := ipa.String(v.(string))
		optArgs.Employeetype = _v
	}
	if v, ok := d.GetOk("preferred_language"); ok {
		_v := ipa.String(v.(string))
		optArgs.Preferredlanguage = _v
	}

	// Misc. information
	if v, ok := d.GetOk("carlicense"); ok {
		_v := sliceStrings(v.([]interface{}))
		optArgs.Carlicense = &_v
	} else {
		optArgs.Carlicense = &[]string{}
	}

	res, err := client.UserMod(
		&ipa.UserModArgs{},
		&optArgs,
	)
	if err != nil {
		return err
	}

	// randompassword is not returned by UserShow
	if d.Get("random").(bool) {
		d.Set("randompassword", *res.Result.Randompassword)
	}

	return resourceFreeIPAUserRead(d, meta)
}

func resourceFreeIPAUserDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Deleting User: %s", d.Id())
	client, err := meta.(*Config).Client()
	if err != nil {
		return err
	}

	uid := d.Get("uid").(string)
	_, err = client.UserDel(
		&ipa.UserDelArgs{},
		&ipa.UserDelOptionalArgs{
			UID: &[]string{uid},
		},
	)

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceFreeIPAUserImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.SetId(d.Id())
	d.Set("uid", d.Id())

	err := resourceFreeIPAUserRead(d, meta)
	if err != nil {
		return []*schema.ResourceData{}, err
	}

	return []*schema.ResourceData{d}, nil
}
