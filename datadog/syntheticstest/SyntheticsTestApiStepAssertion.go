// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestApiStepAssertion struct {
	// Type of assertion.
	//
	// **Note:** Only some combinations of `type` and `operator` are valid. Refer to `config.assertions` in the [Datadog API reference](https://docs.datadoghq.com/api/latest/synthetics/#create-an-api-test). Valid values are `body`, `header`, `statusCode`, `certificate`, `responseTime`, `property`, `recordEvery`, `recordSome`, `tlsVersion`, `minTlsVersion`, `latency`, `packetLossPercentage`, `packetsReceived`, `networkHop`, `receivedMessage`, `grpcHealthcheckStatus`, `grpcMetadata`, `grpcProto`, `connection`, `bodyHash`, `javascript`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// If assertion type is `javascript`, this is the JavaScript code that performs the assertions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#code SyntheticsTest#code}
	Code *string `field:"optional" json:"code" yaml:"code"`
	// Assertion operator.
	//
	// **Note:** Only some combinations of `type` and `operator` are valid. Refer to `config.assertions` in the [Datadog API reference](https://docs.datadoghq.com/api/latest/synthetics/#create-an-api-test).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#operator SyntheticsTest#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// If assertion type is `header`, this is the header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#property SyntheticsTest#property}
	Property *string `field:"optional" json:"property" yaml:"property"`
	// Expected value. **Note:** Depends on the assertion type. Refer to `config.assertions` in the [Datadog API reference](https://docs.datadoghq.com/api/latest/synthetics/#create-an-api-test).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#target SyntheticsTest#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// targetjsonpath block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#targetjsonpath SyntheticsTest#targetjsonpath}
	Targetjsonpath *SyntheticsTestApiStepAssertionTargetjsonpath `field:"optional" json:"targetjsonpath" yaml:"targetjsonpath"`
	// targetjsonschema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#targetjsonschema SyntheticsTest#targetjsonschema}
	Targetjsonschema *SyntheticsTestApiStepAssertionTargetjsonschema `field:"optional" json:"targetjsonschema" yaml:"targetjsonschema"`
	// targetxpath block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#targetxpath SyntheticsTest#targetxpath}
	Targetxpath *SyntheticsTestApiStepAssertionTargetxpath `field:"optional" json:"targetxpath" yaml:"targetxpath"`
	// Timings scope for response time assertions. Valid values are `all`, `withoutDNS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#timings_scope SyntheticsTest#timings_scope}
	TimingsScope *string `field:"optional" json:"timingsScope" yaml:"timingsScope"`
}

