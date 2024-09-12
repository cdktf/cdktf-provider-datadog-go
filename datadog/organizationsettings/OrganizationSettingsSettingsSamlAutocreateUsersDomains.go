// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationsettings


type OrganizationSettingsSettingsSamlAutocreateUsersDomains struct {
	// List of domains where the SAML automated user creation is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/organization_settings#domains OrganizationSettings#domains}
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// Whether or not the automated user creation based on SAML domain is enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/organization_settings#enabled OrganizationSettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

