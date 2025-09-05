// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataset


type DatasetProductFilters struct {
	// A list of tag-based filters used to restrict access to the product type. Each filter is formatted as `@tag.key:value`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dataset#filters Dataset#filters}
	Filters *[]*string `field:"required" json:"filters" yaml:"filters"`
	// The product type of the dataset. Supported types: `apm`, `rum`, `synthetics`, `metrics`, `logs`, `sd_repoinfo`, `error_tracking`, `cloud_cost`, and `ml_obs`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dataset#product Dataset#product}
	Product *string `field:"required" json:"product" yaml:"product"`
}

