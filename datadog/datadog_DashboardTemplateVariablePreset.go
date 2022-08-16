// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardTemplateVariablePreset struct {
	// The name of the preset.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#name Dashboard#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// template_variable block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#template_variable Dashboard#template_variable}
	TemplateVariable interface{} `field:"optional" json:"templateVariable" yaml:"templateVariable"`
}

