package organizationsettings


type OrganizationSettingsSettingsSamlIdpInitiatedLogin struct {
	// Whether or not a SAML identity provider metadata file was provided to the Datadog organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.23.0/docs/resources/organization_settings#enabled OrganizationSettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

