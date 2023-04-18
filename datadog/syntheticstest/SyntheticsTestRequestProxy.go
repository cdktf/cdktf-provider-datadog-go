package syntheticstest


type SyntheticsTestRequestProxy struct {
	// URL of the proxy to perform the test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.23.0/docs/resources/synthetics_test#url SyntheticsTest#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Header name and value map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.23.0/docs/resources/synthetics_test#headers SyntheticsTest#headers}
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
}

