// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionHttpTokenAuthBody struct {
	// Serialized body content. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/action_connection#content ActionConnection#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Content type of the body. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/action_connection#content_type ActionConnection#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

