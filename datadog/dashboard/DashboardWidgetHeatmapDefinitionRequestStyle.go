// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetHeatmapDefinitionRequestStyle struct {
	// A color palette to apply to the widget. The available options are available at: https://docs.datadoghq.com/dashboards/widgets/timeseries/#appearance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.31.0/docs/resources/dashboard#palette Dashboard#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
}

