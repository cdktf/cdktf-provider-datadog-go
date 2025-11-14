// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogactionconnection


type DataDatadogActionConnectionHttpTokenAuth struct {
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/data-sources/action_connection#header DataDatadogActionConnection#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/data-sources/action_connection#token DataDatadogActionConnection#token}
	Token interface{} `field:"optional" json:"token" yaml:"token"`
	// url_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/data-sources/action_connection#url_parameter DataDatadogActionConnection#url_parameter}
	UrlParameter interface{} `field:"optional" json:"urlParameter" yaml:"urlParameter"`
}

