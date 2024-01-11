// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logsindex


type LogsIndexFilter struct {
	// Logs filter criteria. Only logs matching this filter criteria are considered for this index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/logs_index#query LogsIndex#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

