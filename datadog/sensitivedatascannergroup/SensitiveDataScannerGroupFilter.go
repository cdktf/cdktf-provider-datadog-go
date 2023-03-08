package sensitivedatascannergroup


type SensitiveDataScannerGroupFilter struct {
	// Query to filter the events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/sensitive_data_scanner_group#query SensitiveDataScannerGroup#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

