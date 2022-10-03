package organizationsettings


type OrganizationSettingsSettingsSaml struct {
	// Whether or not SAML is enabled for this organization.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/organization_settings#enabled OrganizationSettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

