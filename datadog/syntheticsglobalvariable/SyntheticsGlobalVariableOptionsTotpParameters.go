package syntheticsglobalvariable


type SyntheticsGlobalVariableOptionsTotpParameters struct {
	// Number of digits for the OTP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_global_variable#digits SyntheticsGlobalVariable#digits}
	Digits *float64 `field:"required" json:"digits" yaml:"digits"`
	// Interval for which to refresh the token (in seconds).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_global_variable#refresh_interval SyntheticsGlobalVariable#refresh_interval}
	RefreshInterval *float64 `field:"required" json:"refreshInterval" yaml:"refreshInterval"`
}

