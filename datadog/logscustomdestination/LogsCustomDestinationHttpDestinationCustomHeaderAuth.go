// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustomdestination


type LogsCustomDestinationHttpDestinationCustomHeaderAuth struct {
	// The header name of the authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/logs_custom_destination#header_name LogsCustomDestination#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The header value of the authentication. This field is not returned by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/logs_custom_destination#header_value LogsCustomDestination#header_value}
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
}

