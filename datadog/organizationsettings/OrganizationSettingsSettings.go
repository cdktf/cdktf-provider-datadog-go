// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationsettings


type OrganizationSettingsSettings struct {
	// saml block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.0/docs/resources/organization_settings#saml OrganizationSettings#saml}
	Saml *OrganizationSettingsSettingsSaml `field:"required" json:"saml" yaml:"saml"`
	// saml_autocreate_users_domains block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.0/docs/resources/organization_settings#saml_autocreate_users_domains OrganizationSettings#saml_autocreate_users_domains}
	SamlAutocreateUsersDomains *OrganizationSettingsSettingsSamlAutocreateUsersDomains `field:"required" json:"samlAutocreateUsersDomains" yaml:"samlAutocreateUsersDomains"`
	// saml_idp_initiated_login block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.0/docs/resources/organization_settings#saml_idp_initiated_login OrganizationSettings#saml_idp_initiated_login}
	SamlIdpInitiatedLogin *OrganizationSettingsSettingsSamlIdpInitiatedLogin `field:"required" json:"samlIdpInitiatedLogin" yaml:"samlIdpInitiatedLogin"`
	// saml_strict_mode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.0/docs/resources/organization_settings#saml_strict_mode OrganizationSettings#saml_strict_mode}
	SamlStrictMode *OrganizationSettingsSettingsSamlStrictMode `field:"required" json:"samlStrictMode" yaml:"samlStrictMode"`
	// Whether or not the organization users can share widgets outside of Datadog. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.0/docs/resources/organization_settings#private_widget_share OrganizationSettings#private_widget_share}
	PrivateWidgetShare interface{} `field:"optional" json:"privateWidgetShare" yaml:"privateWidgetShare"`
	// The access role of the user.
	//
	// Options are `st` (standard user), `adm` (admin user), or `ro` (read-only user). Allowed enum values: `st`, `adm` , `ro`, `ERROR` Defaults to `"st"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.0/docs/resources/organization_settings#saml_autocreate_access_role OrganizationSettings#saml_autocreate_access_role}
	SamlAutocreateAccessRole *string `field:"optional" json:"samlAutocreateAccessRole" yaml:"samlAutocreateAccessRole"`
}

