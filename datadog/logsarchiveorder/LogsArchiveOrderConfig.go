package logsarchiveorder

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsArchiveOrderConfig struct {
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
	// The archive IDs list.
	//
	// The order of archive IDs in this attribute defines the overall archive order for logs. If `archive_ids` is empty or not specified, it will import the actual archive order, and create the resource. Otherwise, it will try to update the order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/logs_archive_order#archive_ids LogsArchiveOrder#archive_ids}
	ArchiveIds *[]*string `field:"optional" json:"archiveIds" yaml:"archiveIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/logs_archive_order#id LogsArchiveOrder#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

