package dashboard


type DashboardTemplateVariablePreset struct {
	// The name of the preset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.23.0/docs/resources/dashboard#name Dashboard#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// template_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.23.0/docs/resources/dashboard#template_variable Dashboard#template_variable}
	TemplateVariable interface{} `field:"optional" json:"templateVariable" yaml:"templateVariable"`
}

