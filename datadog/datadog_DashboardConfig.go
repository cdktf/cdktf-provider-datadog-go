// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The layout type of the dashboard. Valid values are `ordered`, `free`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#layout_type Dashboard#layout_type}
	LayoutType *string `field:"required" json:"layoutType" yaml:"layoutType"`
	// The title of the dashboard.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#title Dashboard#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// A list of dashboard lists this dashboard belongs to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#dashboard_lists Dashboard#dashboard_lists}
	DashboardLists *[]*float64 `field:"optional" json:"dashboardLists" yaml:"dashboardLists"`
	// The description of the dashboard.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#description Dashboard#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#id Dashboard#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether this dashboard is read-only. **Deprecated.** Prefer using `restricted_roles` to define which roles are required to edit the dashboard.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#is_read_only Dashboard#is_read_only}
	IsReadOnly interface{} `field:"optional" json:"isReadOnly" yaml:"isReadOnly"`
	// The list of handles for the users to notify when changes are made to this dashboard.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#notify_list Dashboard#notify_list}
	NotifyList *[]*string `field:"optional" json:"notifyList" yaml:"notifyList"`
	// The reflow type of a new dashboard layout.
	//
	// Set this only when layout type is `ordered`. If set to `fixed`, the dashboard expects all widgets to have a layout, and if it's set to `auto`, widgets should not have layouts. Valid values are `auto`, `fixed`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#reflow_type Dashboard#reflow_type}
	ReflowType *string `field:"optional" json:"reflowType" yaml:"reflowType"`
	// UUIDs of roles whose associated users are authorized to edit the dashboard.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#restricted_roles Dashboard#restricted_roles}
	RestrictedRoles *[]*string `field:"optional" json:"restrictedRoles" yaml:"restrictedRoles"`
	// template_variable block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#template_variable Dashboard#template_variable}
	TemplateVariable interface{} `field:"optional" json:"templateVariable" yaml:"templateVariable"`
	// template_variable_preset block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#template_variable_preset Dashboard#template_variable_preset}
	TemplateVariablePreset interface{} `field:"optional" json:"templateVariablePreset" yaml:"templateVariablePreset"`
	// The URL of the dashboard.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#url Dashboard#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// widget block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#widget Dashboard#widget}
	Widget interface{} `field:"optional" json:"widget" yaml:"widget"`
}

