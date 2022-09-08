// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type MonitorVariablesEventQuerySearch struct {
	// The events search string.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#query Monitor#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

