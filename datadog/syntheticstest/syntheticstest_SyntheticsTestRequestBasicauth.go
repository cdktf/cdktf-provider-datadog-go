package syntheticstest


type SyntheticsTestRequestBasicauth struct {
	// Access key for `SIGV4` authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#access_key SyntheticsTest#access_key}
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Domain for `ntlm` authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#domain SyntheticsTest#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Password for authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#password SyntheticsTest#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Region for `SIGV4` authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#region SyntheticsTest#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Secret key for `SIGV4` authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#secret_key SyntheticsTest#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// Service name for `SIGV4` authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#service_name SyntheticsTest#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Session token for `SIGV4` authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#session_token SyntheticsTest#session_token}
	SessionToken *string `field:"optional" json:"sessionToken" yaml:"sessionToken"`
	// Type of basic authentication to use when performing the test.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Username for authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#username SyntheticsTest#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Workstation for `ntlm` authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#workstation SyntheticsTest#workstation}
	Workstation *string `field:"optional" json:"workstation" yaml:"workstation"`
}

