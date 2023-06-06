package spansmetric


type SpansMetricFilter struct {
	// The search query - following the span search syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/spans_metric#query SpansMetric#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

