// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionHttpTokenAuth struct {
	// body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/action_connection#body ActionConnection#body}
	Body *ActionConnectionHttpTokenAuthBody `field:"optional" json:"body" yaml:"body"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/action_connection#header ActionConnection#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/action_connection#token ActionConnection#token}
	Token interface{} `field:"optional" json:"token" yaml:"token"`
	// url_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/action_connection#url_parameter ActionConnection#url_parameter}
	UrlParameter interface{} `field:"optional" json:"urlParameter" yaml:"urlParameter"`
}

