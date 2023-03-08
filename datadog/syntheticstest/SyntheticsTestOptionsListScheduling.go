package syntheticstest


type SyntheticsTestOptionsListScheduling struct {
	// timeframes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#timeframes SyntheticsTest#timeframes}
	Timeframes interface{} `field:"required" json:"timeframes" yaml:"timeframes"`
	// Timezone in which the timeframe is based.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#timezone SyntheticsTest#timezone}
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

