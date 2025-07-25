// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestConfig struct {
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
	// Array of locations used to run the test.
	//
	// Refer to [the Datadog Synthetics location data source](https://registry.terraform.io/providers/DataDog/datadog/latest/docs/data-sources/synthetics_locations) to retrieve the list of locations or find the possible values listed in [this API response](https://app.datadoghq.com/api/v1/synthetics/locations?only_public=true).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#locations SyntheticsTest#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
	// Name of Datadog synthetics test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#name SyntheticsTest#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Define whether you want to start (`live`) or pause (`paused`) a Synthetic test. Valid values are `live`, `paused`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#status SyntheticsTest#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// Synthetics test type. Valid values are `api`, `browser`, `mobile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// api_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#api_step SyntheticsTest#api_step}
	ApiStep interface{} `field:"optional" json:"apiStep" yaml:"apiStep"`
	// assertion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#assertion SyntheticsTest#assertion}
	Assertion interface{} `field:"optional" json:"assertion" yaml:"assertion"`
	// browser_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#browser_step SyntheticsTest#browser_step}
	BrowserStep interface{} `field:"optional" json:"browserStep" yaml:"browserStep"`
	// browser_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#browser_variable SyntheticsTest#browser_variable}
	BrowserVariable interface{} `field:"optional" json:"browserVariable" yaml:"browserVariable"`
	// Initial application arguments for the mobile test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#config_initial_application_arguments SyntheticsTest#config_initial_application_arguments}
	ConfigInitialApplicationArguments *map[string]*string `field:"optional" json:"configInitialApplicationArguments" yaml:"configInitialApplicationArguments"`
	// config_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#config_variable SyntheticsTest#config_variable}
	ConfigVariable interface{} `field:"optional" json:"configVariable" yaml:"configVariable"`
	// Required if `type = "browser"`. Array with the different device IDs used to run the test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#device_ids SyntheticsTest#device_ids}
	DeviceIds *[]*string `field:"optional" json:"deviceIds" yaml:"deviceIds"`
	// A boolean indicating whether this synthetics test can be deleted even if it's referenced by other resources (for example, SLOs and composite monitors).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#force_delete_dependencies SyntheticsTest#force_delete_dependencies}
	ForceDeleteDependencies interface{} `field:"optional" json:"forceDeleteDependencies" yaml:"forceDeleteDependencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#id SyntheticsTest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A message to include with notifications for this synthetics test.
	//
	// Email notifications can be sent to specific users by using the same `@username` notation as events. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#message SyntheticsTest#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// mobile_options_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#mobile_options_list SyntheticsTest#mobile_options_list}
	MobileOptionsList *SyntheticsTestMobileOptionsListStruct `field:"optional" json:"mobileOptionsList" yaml:"mobileOptionsList"`
	// mobile_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#mobile_step SyntheticsTest#mobile_step}
	MobileStep interface{} `field:"optional" json:"mobileStep" yaml:"mobileStep"`
	// options_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#options_list SyntheticsTest#options_list}
	OptionsList *SyntheticsTestOptionsListStruct `field:"optional" json:"optionsList" yaml:"optionsList"`
	// request_basicauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#request_basicauth SyntheticsTest#request_basicauth}
	RequestBasicauth *SyntheticsTestRequestBasicauth `field:"optional" json:"requestBasicauth" yaml:"requestBasicauth"`
	// request_client_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#request_client_certificate SyntheticsTest#request_client_certificate}
	RequestClientCertificate *SyntheticsTestRequestClientCertificate `field:"optional" json:"requestClientCertificate" yaml:"requestClientCertificate"`
	// request_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#request_definition SyntheticsTest#request_definition}
	RequestDefinition *SyntheticsTestRequestDefinition `field:"optional" json:"requestDefinition" yaml:"requestDefinition"`
	// request_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#request_file SyntheticsTest#request_file}
	RequestFile interface{} `field:"optional" json:"requestFile" yaml:"requestFile"`
	// Header name and value map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#request_headers SyntheticsTest#request_headers}
	RequestHeaders *map[string]*string `field:"optional" json:"requestHeaders" yaml:"requestHeaders"`
	// Metadata to include when performing the gRPC request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#request_metadata SyntheticsTest#request_metadata}
	RequestMetadata *map[string]*string `field:"optional" json:"requestMetadata" yaml:"requestMetadata"`
	// request_proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#request_proxy SyntheticsTest#request_proxy}
	RequestProxy *SyntheticsTestRequestProxy `field:"optional" json:"requestProxy" yaml:"requestProxy"`
	// Query arguments name and value map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#request_query SyntheticsTest#request_query}
	RequestQuery *map[string]*string `field:"optional" json:"requestQuery" yaml:"requestQuery"`
	// Cookies to be used for a browser test request, using the [Set-Cookie](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie) syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#set_cookie SyntheticsTest#set_cookie}
	SetCookie *string `field:"optional" json:"setCookie" yaml:"setCookie"`
	// The subtype of the Synthetic API test.
	//
	// Defaults to `http`. Valid values are `http`, `ssl`, `tcp`, `dns`, `multi`, `icmp`, `udp`, `websocket`, `grpc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#subtype SyntheticsTest#subtype}
	Subtype *string `field:"optional" json:"subtype" yaml:"subtype"`
	// A list of tags to associate with your synthetics test.
	//
	// This can help you categorize and filter tests in the manage synthetics page of the UI. Default is an empty list (`[]`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#tags SyntheticsTest#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Variables defined from JavaScript code for API HTTP tests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#variables_from_script SyntheticsTest#variables_from_script}
	VariablesFromScript *string `field:"optional" json:"variablesFromScript" yaml:"variablesFromScript"`
}

