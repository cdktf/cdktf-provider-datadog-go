// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationsettings


type OrganizationSettingsSettingsSaml struct {
	// Whether or not SAML is enabled for this organization. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/organization_settings#enabled OrganizationSettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

