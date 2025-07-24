// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionHttp struct {
	// Base HTTP url for the integration. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/action_connection#base_url ActionConnection#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// token_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/action_connection#token_auth ActionConnection#token_auth}
	TokenAuth *ActionConnectionHttpTokenAuth `field:"optional" json:"tokenAuth" yaml:"tokenAuth"`
}

