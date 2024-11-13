// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevelobjective

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceLevelObjectiveConfig struct {
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
	// Name of Datadog service level objective.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#name ServiceLevelObjective#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#thresholds ServiceLevelObjective#thresholds}
	Thresholds interface{} `field:"required" json:"thresholds" yaml:"thresholds"`
	// The type of the service level objective.
	//
	// The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API [documentation page](https://docs.datadoghq.com/api/v1/service-level-objectives/#create-a-slo-object). Valid values are `metric`, `monitor`, `time_slice`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#type ServiceLevelObjective#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description of this service level objective.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#description ServiceLevelObjective#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A boolean indicating whether this monitor can be deleted even if it's referenced by other resources (for example, dashboards).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#force_delete ServiceLevelObjective#force_delete}
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// A static set of groups to filter monitor-based SLOs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#groups ServiceLevelObjective#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#id ServiceLevelObjective#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A static set of monitor IDs to use as part of the SLO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#monitor_ids ServiceLevelObjective#monitor_ids}
	MonitorIds *[]*float64 `field:"optional" json:"monitorIds" yaml:"monitorIds"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#query ServiceLevelObjective#query}
	Query *ServiceLevelObjectiveQuery `field:"optional" json:"query" yaml:"query"`
	// sli_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#sli_specification ServiceLevelObjective#sli_specification}
	SliSpecification *ServiceLevelObjectiveSliSpecification `field:"optional" json:"sliSpecification" yaml:"sliSpecification"`
	// A list of tags to associate with your service level objective.
	//
	// This can help you categorize and filter service level objectives in the service level objectives page of the UI. Note: it's not currently possible to filter by these tags when querying via the API
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#tags ServiceLevelObjective#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// The objective's target in `(0,100)`. This must match the corresponding thresholds of the primary time frame.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#target_threshold ServiceLevelObjective#target_threshold}
	TargetThreshold *float64 `field:"optional" json:"targetThreshold" yaml:"targetThreshold"`
	// The primary time frame for the objective.
	//
	// The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are `7d`, `30d`, `90d`, `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#timeframe ServiceLevelObjective#timeframe}
	Timeframe *string `field:"optional" json:"timeframe" yaml:"timeframe"`
	// Whether or not to validate the SLO. It checks if monitors added to a monitor SLO already exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#validate ServiceLevelObjective#validate}
	Validate interface{} `field:"optional" json:"validate" yaml:"validate"`
	// The objective's warning value in `(0,100)`.
	//
	// This must be greater than the target value and match the corresponding thresholds of the primary time frame.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/service_level_objective#warning_threshold ServiceLevelObjective#warning_threshold}
	WarningThreshold *float64 `field:"optional" json:"warningThreshold" yaml:"warningThreshold"`
}

