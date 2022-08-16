// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type ServiceLevelObjectiveQuery struct {
	// The sum of the `total` events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/service_level_objective#denominator ServiceLevelObjective#denominator}
	Denominator *string `field:"required" json:"denominator" yaml:"denominator"`
	// The sum of all the `good` events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/service_level_objective#numerator ServiceLevelObjective#numerator}
	Numerator *string `field:"required" json:"numerator" yaml:"numerator"`
}

