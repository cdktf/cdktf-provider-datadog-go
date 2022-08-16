// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetGroupDefinitionWidgetLogStreamDefinitionSort struct {
	// The facet path for the column.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#column Dashboard#column}
	Column *string `field:"required" json:"column" yaml:"column"`
	// Widget sorting methods. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#order Dashboard#order}
	Order *string `field:"required" json:"order" yaml:"order"`
}

