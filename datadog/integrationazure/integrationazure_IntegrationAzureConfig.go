package integrationazure

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationAzureConfig struct {
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
	// Your Azure web application ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_azure#client_id IntegrationAzure#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// (Required for Initial Creation) Your Azure web application secret key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_azure#client_secret IntegrationAzure#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Your Azure Active Directory ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_azure#tenant_name IntegrationAzure#tenant_name}
	TenantName *string `field:"required" json:"tenantName" yaml:"tenantName"`
	// Silence monitors for expected Azure VM shutdowns.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_azure#automute IntegrationAzure#automute}
	Automute interface{} `field:"optional" json:"automute" yaml:"automute"`
	// String of host tag(s) (in the form `key:value,key:value`) defines a filter that Datadog will use when collecting metrics from Azure.
	//
	// Limit the Azure instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog. e.x. `env:production,deploymentgroup:red`
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_azure#host_filters IntegrationAzure#host_filters}
	HostFilters *string `field:"optional" json:"hostFilters" yaml:"hostFilters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_azure#id IntegrationAzure#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
