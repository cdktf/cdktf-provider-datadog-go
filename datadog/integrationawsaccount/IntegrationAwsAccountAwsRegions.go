// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount


type IntegrationAwsAccountAwsRegions struct {
	// Include all regions. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.0/docs/resources/integration_aws_account#include_all IntegrationAwsAccount#include_all}
	IncludeAll interface{} `field:"optional" json:"includeAll" yaml:"includeAll"`
	// Include only these regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.0/docs/resources/integration_aws_account#include_only IntegrationAwsAccount#include_only}
	IncludeOnly *[]*string `field:"optional" json:"includeOnly" yaml:"includeOnly"`
}

