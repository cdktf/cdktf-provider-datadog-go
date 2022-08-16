// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type OrganizationSettingsSettings struct {
	// saml block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/organization_settings#saml OrganizationSettings#saml}
	Saml *OrganizationSettingsSettingsSaml `field:"required" json:"saml" yaml:"saml"`
	// saml_autocreate_users_domains block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/organization_settings#saml_autocreate_users_domains OrganizationSettings#saml_autocreate_users_domains}
	SamlAutocreateUsersDomains *OrganizationSettingsSettingsSamlAutocreateUsersDomains `field:"required" json:"samlAutocreateUsersDomains" yaml:"samlAutocreateUsersDomains"`
	// saml_idp_initiated_login block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/organization_settings#saml_idp_initiated_login OrganizationSettings#saml_idp_initiated_login}
	SamlIdpInitiatedLogin *OrganizationSettingsSettingsSamlIdpInitiatedLogin `field:"required" json:"samlIdpInitiatedLogin" yaml:"samlIdpInitiatedLogin"`
	// saml_strict_mode block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/organization_settings#saml_strict_mode OrganizationSettings#saml_strict_mode}
	SamlStrictMode *OrganizationSettingsSettingsSamlStrictMode `field:"required" json:"samlStrictMode" yaml:"samlStrictMode"`
	// Whether or not the organization users can share widgets outside of Datadog.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/organization_settings#private_widget_share OrganizationSettings#private_widget_share}
	PrivateWidgetShare interface{} `field:"optional" json:"privateWidgetShare" yaml:"privateWidgetShare"`
	// The access role of the user.
	//
	// Options are `st` (standard user), `adm` (admin user), or `ro` (read-only user). Allowed enum values: `st`, `adm` , `ro`, `ERROR`
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/organization_settings#saml_autocreate_access_role OrganizationSettings#saml_autocreate_access_role}
	SamlAutocreateAccessRole *string `field:"optional" json:"samlAutocreateAccessRole" yaml:"samlAutocreateAccessRole"`
}

