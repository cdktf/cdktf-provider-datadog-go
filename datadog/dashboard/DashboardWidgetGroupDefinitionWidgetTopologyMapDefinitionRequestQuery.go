// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetTopologyMapDefinitionRequestQuery struct {
	// The data source for the Topology request ('service_map' or 'data_streams'). Valid values are `data_streams`, `service_map`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#data_source Dashboard#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Your environment and primary tag (or `*` if enabled for your account).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#filters Dashboard#filters}
	Filters *[]*string `field:"required" json:"filters" yaml:"filters"`
	// The ID of the service to map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#service Dashboard#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

