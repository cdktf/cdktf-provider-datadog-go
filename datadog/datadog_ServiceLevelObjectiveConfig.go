// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceLevelObjectiveConfig struct {
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
	// Name of Datadog service level objective.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/service_level_objective#name ServiceLevelObjective#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// thresholds block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/service_level_objective#thresholds ServiceLevelObjective#thresholds}
	Thresholds interface{} `field:"required" json:"thresholds" yaml:"thresholds"`
	// The type of the service level objective.
	//
	// The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API [documentation page](https://docs.datadoghq.com/api/v1/service-level-objectives/#create-a-slo-object). Valid values are `metric`, `monitor`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/service_level_objective#type ServiceLevelObjective#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description of this service level objective.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/service_level_objective#description ServiceLevelObjective#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A boolean indicating whether this monitor can be deleted even if it’s referenced by other resources (e.g. dashboards).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/service_level_objective#force_delete ServiceLevelObjective#force_delete}
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// A static set of groups to filter monitor-based SLOs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/service_level_objective#groups ServiceLevelObjective#groups}
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/service_level_objective#id ServiceLevelObjective#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A static set of monitor IDs to use as part of the SLO.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/service_level_objective#monitor_ids ServiceLevelObjective#monitor_ids}
	MonitorIds *[]*float64 `field:"optional" json:"monitorIds" yaml:"monitorIds"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/service_level_objective#query ServiceLevelObjective#query}
	Query *ServiceLevelObjectiveQuery `field:"optional" json:"query" yaml:"query"`
	// A list of tags to associate with your service level objective.
	//
	// This can help you categorize and filter service level objectives in the service level objectives page of the UI. Note: it's not currently possible to filter by these tags when querying via the API
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/service_level_objective#tags ServiceLevelObjective#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Whether or not to validate the SLO.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/service_level_objective#validate ServiceLevelObjective#validate}
	Validate interface{} `field:"optional" json:"validate" yaml:"validate"`
}
