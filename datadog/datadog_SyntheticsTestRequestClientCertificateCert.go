// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type SyntheticsTestRequestClientCertificateCert struct {
	// Content of the certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#content SyntheticsTest#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// File name for the certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#filename SyntheticsTest#filename}
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
}

