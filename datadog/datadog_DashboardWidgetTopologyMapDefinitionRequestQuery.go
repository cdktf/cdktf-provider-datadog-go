// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetTopologyMapDefinitionRequestQuery struct {
	// The data source for the Topology request ('service_map' or 'data_streams'). Valid values are `data_streams`, `service_map`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#data_source Dashboard#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Your environment and primary tag (or `*` if enabled for your account).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#filters Dashboard#filters}
	Filters *[]*string `field:"required" json:"filters" yaml:"filters"`
	// The ID of the service to map.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#service Dashboard#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

