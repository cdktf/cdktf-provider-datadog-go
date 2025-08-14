// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appbuilderapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppBuilderAppConfig struct {
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
	// The JSON representation of the App. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/app_builder_app#app_json AppBuilderApp#app_json}
	AppJson *string `field:"required" json:"appJson" yaml:"appJson"`
	// If specified, this will override the Action Connection IDs for the specified Action Query Names in the App JSON.
	//
	// Otherwise, a map of the App's Action Query Names to Action Connection IDs will be returned in output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/app_builder_app#action_query_names_to_connection_ids AppBuilderApp#action_query_names_to_connection_ids}
	ActionQueryNamesToConnectionIds *map[string]*string `field:"optional" json:"actionQueryNamesToConnectionIds" yaml:"actionQueryNamesToConnectionIds"`
	// If specified, this will override the human-readable description of the App in the App JSON.
	//
	// String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/app_builder_app#description AppBuilderApp#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If specified, this will override the name of the App in the App JSON.
	//
	// String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/app_builder_app#name AppBuilderApp#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Set the app to published or unpublished.
	//
	// Published apps are available to other users. To ensure the app is accessible to the correct users, you also need to set a [Restriction Policy](https://docs.datadoghq.com/api/latest/restriction-policies/) on the app if a policy does not yet exist. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/app_builder_app#published AppBuilderApp#published}
	Published interface{} `field:"optional" json:"published" yaml:"published"`
	// The name of the root component of the app.
	//
	// This must be a grid component that contains all other components. If specified, this will override the root instance name of the App in the App JSON. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/app_builder_app#root_instance_name AppBuilderApp#root_instance_name}
	RootInstanceName *string `field:"optional" json:"rootInstanceName" yaml:"rootInstanceName"`
}

