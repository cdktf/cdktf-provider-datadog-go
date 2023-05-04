package syntheticsglobalvariable


type SyntheticsGlobalVariableOptions struct {
	// totp_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/synthetics_global_variable#totp_parameters SyntheticsGlobalVariable#totp_parameters}
	TotpParameters *SyntheticsGlobalVariableOptionsTotpParameters `field:"optional" json:"totpParameters" yaml:"totpParameters"`
}

