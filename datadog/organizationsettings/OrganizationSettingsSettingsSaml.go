// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationsettings


type OrganizationSettingsSettingsSaml struct {
	// Whether or not SAML is enabled for this organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/organization_settings#enabled OrganizationSettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

