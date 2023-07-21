package integrationconfluentaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationConfluentAccountConfig struct {
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
	// The API key associated with your Confluent account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/integration_confluent_account#api_key IntegrationConfluentAccount#api_key}
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The API secret associated with your Confluent account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/integration_confluent_account#api_secret IntegrationConfluentAccount#api_secret}
	ApiSecret *string `field:"required" json:"apiSecret" yaml:"apiSecret"`
	// A list of strings representing tags. Can be a single key, or key-value pairs separated by a colon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/integration_confluent_account#tags IntegrationConfluentAccount#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

