// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetTimeseriesDefinitionRequestMetadata struct {
	// The expression name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#expression Dashboard#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The expression alias.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#alias_name Dashboard#alias_name}
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
}

