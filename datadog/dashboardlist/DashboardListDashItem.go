// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboardlist


type DashboardListDashItem struct {
	// The ID of the dashboard to add.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/dashboard_list#dash_id DashboardList#dash_id}
	DashId *string `field:"required" json:"dashId" yaml:"dashId"`
	// The type of this dashboard. Valid values are `custom_timeboard`, `custom_screenboard`, `integration_screenboard`, `integration_timeboard`, `host_timeboard`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/dashboard_list#type DashboardList#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

