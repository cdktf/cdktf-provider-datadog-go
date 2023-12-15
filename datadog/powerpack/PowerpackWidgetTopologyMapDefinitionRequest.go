// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetTopologyMapDefinitionRequest struct {
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/powerpack#query Powerpack#query}
	Query interface{} `field:"required" json:"query" yaml:"query"`
	// The request type for the Topology request ('topology'). Valid values are `topology`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/powerpack#request_type Powerpack#request_type}
	RequestType *string `field:"required" json:"requestType" yaml:"requestType"`
}

