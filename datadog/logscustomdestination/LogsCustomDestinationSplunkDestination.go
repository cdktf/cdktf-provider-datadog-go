// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustomdestination


type LogsCustomDestinationSplunkDestination struct {
	// Access token of the Splunk HTTP Event Collector. This field is not returned by the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/logs_custom_destination#access_token LogsCustomDestination#access_token}
	AccessToken *string `field:"required" json:"accessToken" yaml:"accessToken"`
	// The destination for which logs will be forwarded to.
	//
	// Must have HTTPS scheme. Forwarding back to Datadog is not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/logs_custom_destination#endpoint LogsCustomDestination#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
}

