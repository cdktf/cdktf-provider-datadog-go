package datadatadoglogspipelines

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogLogsPipelinesConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/d/logs_pipelines#id DataDatadogLogsPipelines#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Filter parameter for retrieved pipelines.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/d/logs_pipelines#is_read_only DataDatadogLogsPipelines#is_read_only}
	IsReadOnly *string `field:"optional" json:"isReadOnly" yaml:"isReadOnly"`
}
