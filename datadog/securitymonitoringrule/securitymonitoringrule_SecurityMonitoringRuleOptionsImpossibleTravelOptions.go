package securitymonitoringrule


type SecurityMonitoringRuleOptionsImpossibleTravelOptions struct {
	// If true, signals are suppressed for the first 24 hours.
	//
	// During that time, Datadog learns the user's regular access locations. This can be helpful to reduce noise and infer VPN usage or credentialed API access.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#baseline_user_locations SecurityMonitoringRule#baseline_user_locations}
	BaselineUserLocations interface{} `field:"optional" json:"baselineUserLocations" yaml:"baselineUserLocations"`
}
