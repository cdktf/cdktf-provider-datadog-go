// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type LogsIndexExclusionFilter struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_index#filter LogsIndex#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// A boolean stating if the exclusion is active or not.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_index#is_enabled LogsIndex#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// The name of the exclusion filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_index#name LogsIndex#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}
