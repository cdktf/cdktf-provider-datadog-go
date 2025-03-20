// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustomdestination


type LogsCustomDestinationElasticsearchDestinationBasicAuth struct {
	// The password of the authentication. This field is not returned by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/logs_custom_destination#password LogsCustomDestination#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The username of the authentication. This field is not returned by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/logs_custom_destination#username LogsCustomDestination#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

