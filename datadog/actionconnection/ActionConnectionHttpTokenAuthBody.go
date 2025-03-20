// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionHttpTokenAuthBody struct {
	// Serialized body content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/action_connection#content ActionConnection#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Content type of the body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/action_connection#content_type ActionConnection#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

