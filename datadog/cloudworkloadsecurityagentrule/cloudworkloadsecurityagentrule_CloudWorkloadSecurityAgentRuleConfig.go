package cloudworkloadsecurityagentrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudWorkloadSecurityAgentRuleConfig struct {
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
	// The SECL expression of the Agent rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/cloud_workload_security_agent_rule#expression CloudWorkloadSecurityAgentRule#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The name of the Agent rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/cloud_workload_security_agent_rule#name CloudWorkloadSecurityAgentRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the Agent rule.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/cloud_workload_security_agent_rule#description CloudWorkloadSecurityAgentRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the Agent rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/cloud_workload_security_agent_rule#enabled CloudWorkloadSecurityAgentRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/cloud_workload_security_agent_rule#id CloudWorkloadSecurityAgentRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
