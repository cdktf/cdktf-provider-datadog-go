package syntheticstest


type SyntheticsTestOptionsListMonitorOptions struct {
	// Specify a renotification frequency in minutes.
	//
	// Values available by default are `0`, `10`, `20`, `30`, `40`, `50`, `60`, `90`, `120`, `180`, `240`, `300`, `360`, `720`, `1440`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/synthetics_test#renotify_interval SyntheticsTest#renotify_interval}
	RenotifyInterval *float64 `field:"optional" json:"renotifyInterval" yaml:"renotifyInterval"`
}

