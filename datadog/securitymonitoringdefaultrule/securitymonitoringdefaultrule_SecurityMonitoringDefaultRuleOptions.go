package securitymonitoringdefaultrule


type SecurityMonitoringDefaultRuleOptions struct {
	// If true, signals in non-production environments have a lower severity than what is defined by the rule case, which can reduce noise.
	//
	// The decrement is applied when the environment tag of the signal starts with `staging`, `test`, or `dev`. Only available when the rule type is `log_detection`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_default_rule#decrease_criticality_based_on_env SecurityMonitoringDefaultRule#decrease_criticality_based_on_env}
	DecreaseCriticalityBasedOnEnv interface{} `field:"optional" json:"decreaseCriticalityBasedOnEnv" yaml:"decreaseCriticalityBasedOnEnv"`
}
