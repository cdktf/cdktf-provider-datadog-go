// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardTemplateVariable struct {
	// The name of the variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of values that the template variable drop-down is be limited to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#available_values Dashboard#available_values}
	AvailableValues *[]*string `field:"optional" json:"availableValues" yaml:"availableValues"`
	// The default value for the template variable on dashboard load.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#default Dashboard#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
	// The tag prefix associated with the variable. Only tags with this prefix appear in the variable dropdown.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#prefix Dashboard#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}
