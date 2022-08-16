// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type SecurityMonitoringFilterExclusionFilter struct {
	// Exclusion filter name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_filter#name SecurityMonitoringFilter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Exclusion filter query. Logs that match this query are excluded from the security filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_filter#query SecurityMonitoringFilter#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

