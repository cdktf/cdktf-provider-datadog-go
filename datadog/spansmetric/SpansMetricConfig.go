package spansmetric

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpansMetricConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// The name of the span-based metric. This field can't be updated after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/spans_metric#name SpansMetric#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/spans_metric#filter SpansMetric#filter}
	Filter *SpansMetricFilter `field:"optional" json:"filter" yaml:"filter"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/spans_metric#group_by SpansMetric#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
}

