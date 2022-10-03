package dashboard


type DashboardTemplateVariablePresetTemplateVariable struct {
	// The name of the template variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#name Dashboard#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value that should be assumed by the template variable in this preset.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#value Dashboard#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

