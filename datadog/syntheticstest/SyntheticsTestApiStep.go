package syntheticstest


type SyntheticsTestApiStep struct {
	// The name of the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/synthetics_test#name SyntheticsTest#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Determines whether or not to continue with test if this step fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/synthetics_test#allow_failure SyntheticsTest#allow_failure}
	AllowFailure interface{} `field:"optional" json:"allowFailure" yaml:"allowFailure"`
	// assertion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/synthetics_test#assertion SyntheticsTest#assertion}
	Assertion interface{} `field:"optional" json:"assertion" yaml:"assertion"`
	// extracted_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/synthetics_test#extracted_value SyntheticsTest#extracted_value}
	ExtractedValue interface{} `field:"optional" json:"extractedValue" yaml:"extractedValue"`
	// Determines whether or not to consider the entire test as failed if this step fails.
	//
	// Can be used only if `allow_failure` is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/synthetics_test#is_critical SyntheticsTest#is_critical}
	IsCritical interface{} `field:"optional" json:"isCritical" yaml:"isCritical"`
	// request_basicauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/synthetics_test#request_basicauth SyntheticsTest#request_basicauth}
	RequestBasicauth *SyntheticsTestApiStepRequestBasicauth `field:"optional" json:"requestBasicauth" yaml:"requestBasicauth"`
	// request_client_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/synthetics_test#request_client_certificate SyntheticsTest#request_client_certificate}
	RequestClientCertificate *SyntheticsTestApiStepRequestClientCertificate `field:"optional" json:"requestClientCertificate" yaml:"requestClientCertificate"`
	// request_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/synthetics_test#request_definition SyntheticsTest#request_definition}
	RequestDefinition *SyntheticsTestApiStepRequestDefinition `field:"optional" json:"requestDefinition" yaml:"requestDefinition"`
	// Header name and value map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/synthetics_test#request_headers SyntheticsTest#request_headers}
	RequestHeaders *map[string]*string `field:"optional" json:"requestHeaders" yaml:"requestHeaders"`
	// request_proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/synthetics_test#request_proxy SyntheticsTest#request_proxy}
	RequestProxy *SyntheticsTestApiStepRequestProxy `field:"optional" json:"requestProxy" yaml:"requestProxy"`
	// Query arguments name and value map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/synthetics_test#request_query SyntheticsTest#request_query}
	RequestQuery *map[string]*string `field:"optional" json:"requestQuery" yaml:"requestQuery"`
	// retry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/synthetics_test#retry SyntheticsTest#retry}
	Retry *SyntheticsTestApiStepRetry `field:"optional" json:"retry" yaml:"retry"`
	// The subtype of the Synthetic multistep API test step. Valid values are `http`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/synthetics_test#subtype SyntheticsTest#subtype}
	Subtype *string `field:"optional" json:"subtype" yaml:"subtype"`
}

