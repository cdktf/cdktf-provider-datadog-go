package syntheticstest


type SyntheticsTestOptionsListMonitorOptions struct {
	// Specify a renotification frequency.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#renotify_interval SyntheticsTest#renotify_interval}
	RenotifyInterval *float64 `field:"optional" json:"renotifyInterval" yaml:"renotifyInterval"`
}

