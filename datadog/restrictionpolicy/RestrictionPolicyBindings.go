// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package restrictionpolicy


type RestrictionPolicyBindings struct {
	// An array of principals.
	//
	// A principal is a subject or group of subjects. Each principal is formatted as `type:id`. Supported types: `role`, `team`, `user`, and `org`. Org ID can be obtained using a `GET /api/v2/current_user` API request. Find it in the `data.relationships.org.data.id` field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/restriction_policy#principals RestrictionPolicy#principals}
	Principals *[]*string `field:"required" json:"principals" yaml:"principals"`
	// The role/level of access. See this page for more details https://docs.datadoghq.com/api/latest/restriction-policies/#supported-relations-for-resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/restriction_policy#relation RestrictionPolicy#relation}
	Relation *string `field:"required" json:"relation" yaml:"relation"`
}

