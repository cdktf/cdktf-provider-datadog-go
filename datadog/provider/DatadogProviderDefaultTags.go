// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type DatadogProviderDefaultTags struct {
	// [Experimental - Monitors only] Resource tags to be applied by default across all resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs#tags DatadogProvider#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

