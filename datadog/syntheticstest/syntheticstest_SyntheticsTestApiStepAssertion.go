package syntheticstest


type SyntheticsTestApiStepAssertion struct {
	// Assertion operator. **Note** Only some combinations of `type` and `operator` are valid (please refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test)).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#operator SyntheticsTest#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Type of assertion.
	//
	// **Note** Only some combinations of `type` and `operator` are valid (please refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test)). Valid values are `body`, `header`, `statusCode`, `certificate`, `responseTime`, `property`, `recordEvery`, `recordSome`, `tlsVersion`, `minTlsVersion`, `latency`, `packetLossPercentage`, `packetsReceived`, `networkHop`, `receivedMessage`, `grpcHealthcheckStatus`, `connection`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// If assertion type is `header`, this is the header name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#property SyntheticsTest#property}
	Property *string `field:"optional" json:"property" yaml:"property"`
	// Expected value. Depends on the assertion type, refer to [Datadog documentation](https://docs.datadoghq.com/api/latest/synthetics/#create-a-test) for details.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#target SyntheticsTest#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// targetjsonpath block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#targetjsonpath SyntheticsTest#targetjsonpath}
	Targetjsonpath *SyntheticsTestApiStepAssertionTargetjsonpath `field:"optional" json:"targetjsonpath" yaml:"targetjsonpath"`
}

