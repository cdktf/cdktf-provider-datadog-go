// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustomdestination


type LogsCustomDestinationHttpDestination struct {
	// The destination for which logs will be forwarded to.
	//
	// Must have HTTPS scheme. Forwarding back to Datadog is not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/logs_custom_destination#endpoint LogsCustomDestination#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/logs_custom_destination#basic_auth LogsCustomDestination#basic_auth}
	BasicAuth interface{} `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// custom_header_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/logs_custom_destination#custom_header_auth LogsCustomDestination#custom_header_auth}
	CustomHeaderAuth interface{} `field:"optional" json:"customHeaderAuth" yaml:"customHeaderAuth"`
}

