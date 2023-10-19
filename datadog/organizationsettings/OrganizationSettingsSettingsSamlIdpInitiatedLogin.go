// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationsettings


type OrganizationSettingsSettingsSamlIdpInitiatedLogin struct {
	// Whether or not a SAML identity provider metadata file was provided to the Datadog organization. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.31.0/docs/resources/organization_settings#enabled OrganizationSettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

