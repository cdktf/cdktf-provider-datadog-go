// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetToplistDefinitionRequestQueryEventQuerySearch struct {
	// The events search string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/powerpack#query Powerpack#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

