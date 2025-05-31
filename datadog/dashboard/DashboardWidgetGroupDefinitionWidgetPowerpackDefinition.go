// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetPowerpackDefinition struct {
	// UUID of the associated powerpack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/dashboard#powerpack_id Dashboard#powerpack_id}
	PowerpackId *string `field:"required" json:"powerpackId" yaml:"powerpackId"`
	// The background color of the powerpack title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/dashboard#background_color Dashboard#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// URL of image to display as a banner for the powerpack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/dashboard#banner_img Dashboard#banner_img}
	BannerImg *string `field:"optional" json:"bannerImg" yaml:"bannerImg"`
	// Whether to show the title of the powerpack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/dashboard#show_title Dashboard#show_title}
	ShowTitle interface{} `field:"optional" json:"showTitle" yaml:"showTitle"`
	// template_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/dashboard#template_variables Dashboard#template_variables}
	TemplateVariables *DashboardWidgetGroupDefinitionWidgetPowerpackDefinitionTemplateVariables `field:"optional" json:"templateVariables" yaml:"templateVariables"`
	// Title of the powerpack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/dashboard#title Dashboard#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

