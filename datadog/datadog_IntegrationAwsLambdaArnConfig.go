// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationAwsLambdaArnConfig struct {
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
	// Your AWS Account ID without dashes. If your account is a GovCloud or China account, specify the `access_key_id` here.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_aws_lambda_arn#account_id IntegrationAwsLambdaArn#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The ARN of the Datadog forwarder Lambda.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_aws_lambda_arn#lambda_arn IntegrationAwsLambdaArn#lambda_arn}
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_aws_lambda_arn#id IntegrationAwsLambdaArn#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

