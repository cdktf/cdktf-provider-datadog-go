package servicelevelobjective


type ServiceLevelObjectiveThresholds struct {
	// The objective's target in `(0,100)`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/service_level_objective#target ServiceLevelObjective#target}
	Target *float64 `field:"required" json:"target" yaml:"target"`
	// The time frame for the objective.
	//
	// The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API documentation page. Valid values are `7d`, `30d`, `90d`, `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/service_level_objective#timeframe ServiceLevelObjective#timeframe}
	Timeframe *string `field:"required" json:"timeframe" yaml:"timeframe"`
	// The objective's warning value in `(0,100)`. This must be greater than the target value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/service_level_objective#warning ServiceLevelObjective#warning}
	Warning *float64 `field:"optional" json:"warning" yaml:"warning"`
}

