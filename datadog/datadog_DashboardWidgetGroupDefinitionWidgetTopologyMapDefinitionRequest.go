// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetGroupDefinitionWidgetTopologyMapDefinitionRequest struct {
	// query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#query Dashboard#query}
	Query interface{} `field:"required" json:"query" yaml:"query"`
	// The request type for the Topology request ('topology'). Valid values are `topology`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#request_type Dashboard#request_type}
	RequestType *string `field:"required" json:"requestType" yaml:"requestType"`
}

