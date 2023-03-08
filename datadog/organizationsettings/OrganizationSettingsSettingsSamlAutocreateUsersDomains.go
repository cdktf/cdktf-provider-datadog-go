package organizationsettings


type OrganizationSettingsSettingsSamlAutocreateUsersDomains struct {
	// List of domains where the SAML automated user creation is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/organization_settings#domains OrganizationSettings#domains}
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// Whether or not the automated user creation based on SAML domain is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/organization_settings#enabled OrganizationSettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

