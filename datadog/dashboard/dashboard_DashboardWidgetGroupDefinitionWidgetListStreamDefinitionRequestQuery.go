package dashboard


type DashboardWidgetGroupDefinitionWidgetListStreamDefinitionRequestQuery struct {
	// Source from which to query items to display in the stream. Valid values are `logs_stream`, `audit_stream`, `rum_issue_stream`, `apm_issue_stream`, `logs_pattern_stream`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#data_source Dashboard#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// List of indexes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#indexes Dashboard#indexes}
	Indexes *[]*string `field:"optional" json:"indexes" yaml:"indexes"`
	// Widget query.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#query_string Dashboard#query_string}
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// Storage location (private beta).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#storage Dashboard#storage}
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
}

