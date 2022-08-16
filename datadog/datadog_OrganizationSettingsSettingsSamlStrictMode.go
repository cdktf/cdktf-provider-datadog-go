// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type OrganizationSettingsSettingsSamlStrictMode struct {
	// Whether or not the SAML strict mode is enabled. If true, all users must log in with SAML.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/organization_settings#enabled OrganizationSettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

