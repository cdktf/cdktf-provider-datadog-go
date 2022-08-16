// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type OrganizationSettingsSettingsSamlIdpInitiatedLogin struct {
	// Whether or not a SAML identity provider metadata file was provided to the Datadog organization.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/organization_settings#enabled OrganizationSettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

