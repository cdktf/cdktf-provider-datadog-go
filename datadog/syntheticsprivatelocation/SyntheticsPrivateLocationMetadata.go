// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsprivatelocation


type SyntheticsPrivateLocationMetadata struct {
	// A set of role identifiers pulled from the Roles API to restrict read and write access.
	//
	// **Deprecated.** This field is no longer supported by the Datadog API. Please use `datadog_restriction_policy` instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/synthetics_private_location#restricted_roles SyntheticsPrivateLocation#restricted_roles}
	RestrictedRoles *[]*string `field:"optional" json:"restrictedRoles" yaml:"restrictedRoles"`
}

