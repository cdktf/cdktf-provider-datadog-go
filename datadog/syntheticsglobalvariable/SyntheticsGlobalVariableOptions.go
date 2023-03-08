package syntheticsglobalvariable


type SyntheticsGlobalVariableOptions struct {
	// totp_parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_global_variable#totp_parameters SyntheticsGlobalVariable#totp_parameters}
	TotpParameters *SyntheticsGlobalVariableOptionsTotpParameters `field:"optional" json:"totpParameters" yaml:"totpParameters"`
}

