package metricmetadata

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MetricMetadataConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the metric.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/metric_metadata#metric MetricMetadata#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// A description of the metric.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/metric_metadata#description MetricMetadata#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/metric_metadata#id MetricMetadata#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Per unit of the metric such as `second` in `bytes per second`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/metric_metadata#per_unit MetricMetadata#per_unit}
	PerUnit *string `field:"optional" json:"perUnit" yaml:"perUnit"`
	// A short name of the metric.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/metric_metadata#short_name MetricMetadata#short_name}
	ShortName *string `field:"optional" json:"shortName" yaml:"shortName"`
	// If applicable, statsd flush interval in seconds for the metric.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/metric_metadata#statsd_interval MetricMetadata#statsd_interval}
	StatsdInterval *float64 `field:"optional" json:"statsdInterval" yaml:"statsdInterval"`
	// Metric type such as `gauge` or `rate`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/metric_metadata#type MetricMetadata#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Primary unit of the metric such as `byte` or `operation`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/metric_metadata#unit MetricMetadata#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

