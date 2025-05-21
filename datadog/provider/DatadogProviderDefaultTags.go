// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type DatadogProviderDefaultTags struct {
	// [Experimental - Logs Pipelines, Monitors Security Monitoring Rules, and Service Level Objectives only] Resource tags to be applied by default across all resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs#tags DatadogProvider#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

