// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type LogsCustomPipelineProcessorLookupProcessor struct {
	// List of entries of the lookup table using `key,value` format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#lookup_table LogsCustomPipeline#lookup_table}
	LookupTable *[]*string `field:"required" json:"lookupTable" yaml:"lookupTable"`
	// Name of the source attribute used to do the lookup.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#source LogsCustomPipeline#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// Name of the attribute that contains the result of the lookup.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#target LogsCustomPipeline#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// Default lookup value to use if there is no entry in the lookup table for the value of the source attribute.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#default_lookup LogsCustomPipeline#default_lookup}
	DefaultLookup *string `field:"optional" json:"defaultLookup" yaml:"defaultLookup"`
	// If the processor is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#is_enabled LogsCustomPipeline#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Name of the processor.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

