package dashboard


type DashboardWidgetGroupDefinitionWidgetRunWorkflowDefinitionInput struct {
	// Name of the workflow input.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Dashboard template variable. Can be suffixed with `.value` or `.key`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#value Dashboard#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

