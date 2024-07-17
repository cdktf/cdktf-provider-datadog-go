// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationsettings


type OrganizationSettingsSettingsSamlStrictMode struct {
	// Whether or not the SAML strict mode is enabled.
	//
	// If true, all users must log in with SAML. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/organization_settings#enabled OrganizationSettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

