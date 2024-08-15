// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariables struct {
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.0/docs/resources/monitor#event_query Monitor#event_query}
	EventQuery interface{} `field:"optional" json:"eventQuery" yaml:"eventQuery"`
}

