// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type SyntheticsTestRequestClientCertificate struct {
	// cert block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#cert SyntheticsTest#cert}
	Cert *SyntheticsTestRequestClientCertificateCert `field:"required" json:"cert" yaml:"cert"`
	// key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#key SyntheticsTest#key}
	Key *SyntheticsTestRequestClientCertificateKey `field:"required" json:"key" yaml:"key"`
}

