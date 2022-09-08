// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type MonitorVariablesEventQueryGroupBy struct {
	// The event facet.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#facet Monitor#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
	// The number of groups to return.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#limit Monitor#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#sort Monitor#sort}
	Sort *MonitorVariablesEventQueryGroupBySort `field:"optional" json:"sort" yaml:"sort"`
}

