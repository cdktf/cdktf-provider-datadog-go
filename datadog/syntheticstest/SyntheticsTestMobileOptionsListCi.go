// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestMobileOptionsListCi struct {
	// Execution rule for a Synthetics test. Valid values are `blocking`, `non_blocking`, `skipped`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/synthetics_test#execution_rule SyntheticsTest#execution_rule}
	ExecutionRule *string `field:"required" json:"executionRule" yaml:"executionRule"`
}

