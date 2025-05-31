// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustomdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsCustomDestinationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The custom destination name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/logs_custom_destination#name LogsCustomDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// elasticsearch_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/logs_custom_destination#elasticsearch_destination LogsCustomDestination#elasticsearch_destination}
	ElasticsearchDestination interface{} `field:"optional" json:"elasticsearchDestination" yaml:"elasticsearchDestination"`
	// Whether logs matching this custom destination should be forwarded or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/logs_custom_destination#enabled LogsCustomDestination#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Whether tags from the forwarded logs should be forwarded or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/logs_custom_destination#forward_tags LogsCustomDestination#forward_tags}
	ForwardTags interface{} `field:"optional" json:"forwardTags" yaml:"forwardTags"`
	// List of [tag keys](https://docs.datadoghq.com/getting_started/tagging/#define-tags) to be filtered. 				An empty list represents no restriction is in place and either all or no tags will be 				forwarded depending on `forward_tags_restriction_list_type` parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/logs_custom_destination#forward_tags_restriction_list LogsCustomDestination#forward_tags_restriction_list}
	ForwardTagsRestrictionList *[]*string `field:"optional" json:"forwardTagsRestrictionList" yaml:"forwardTagsRestrictionList"`
	// How the `forward_tags_restriction_list` parameter should be interpreted.
	//
	// If `ALLOW_LIST`, then only tags whose keys on the forwarded logs match the ones on the restriction list
	// 				are forwarded.
	// 				`BLOCK_LIST` works the opposite way. It does not forward the tags matching the ones on the list. Valid values are `ALLOW_LIST`, `BLOCK_LIST`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/logs_custom_destination#forward_tags_restriction_list_type LogsCustomDestination#forward_tags_restriction_list_type}
	ForwardTagsRestrictionListType *string `field:"optional" json:"forwardTagsRestrictionListType" yaml:"forwardTagsRestrictionListType"`
	// http_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/logs_custom_destination#http_destination LogsCustomDestination#http_destination}
	HttpDestination interface{} `field:"optional" json:"httpDestination" yaml:"httpDestination"`
	// The custom destination query filter. Logs matching this query are forwarded to the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/logs_custom_destination#query LogsCustomDestination#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// splunk_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/logs_custom_destination#splunk_destination LogsCustomDestination#splunk_destination}
	SplunkDestination interface{} `field:"optional" json:"splunkDestination" yaml:"splunkDestination"`
}

