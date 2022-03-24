module github.com/camptocamp/terraform-provider-freeipa

go 1.12

replace github.com/tehwalris/go-freeipa => github.com/dickens7/go-freeipa v1.1.0-dickens7

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/tehwalris/go-freeipa v0.0.0-20200322083409-e462fc554b76
)
