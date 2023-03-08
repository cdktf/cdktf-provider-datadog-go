package syntheticstest


type SyntheticsTestApiStepRequestClientCertificate struct {
	// cert block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#cert SyntheticsTest#cert}
	Cert *SyntheticsTestApiStepRequestClientCertificateCert `field:"required" json:"cert" yaml:"cert"`
	// key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#key SyntheticsTest#key}
	Key *SyntheticsTestApiStepRequestClientCertificateKey `field:"required" json:"key" yaml:"key"`
}

