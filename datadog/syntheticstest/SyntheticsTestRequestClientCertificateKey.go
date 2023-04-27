package syntheticstest


type SyntheticsTestRequestClientCertificateKey struct {
	// Content of the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/synthetics_test#content SyntheticsTest#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// File name for the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/synthetics_test#filename SyntheticsTest#filename}
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
}

