// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ipallowlist


type IpAllowlistEntry struct {
	// IP address or range of addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.0/docs/resources/ip_allowlist#cidr_block IpAllowlist#cidr_block}
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// Note accompanying IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.0/docs/resources/ip_allowlist#note IpAllowlist#note}
	Note *string `field:"optional" json:"note" yaml:"note"`
}

