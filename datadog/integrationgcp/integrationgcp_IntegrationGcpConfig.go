package integrationgcp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationGcpConfig struct {
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
	// Your email found in your JSON service account key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_gcp#client_email IntegrationGcp#client_email}
	ClientEmail *string `field:"required" json:"clientEmail" yaml:"clientEmail"`
	// Your ID found in your JSON service account key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_gcp#client_id IntegrationGcp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Your private key name found in your JSON service account key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_gcp#private_key IntegrationGcp#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Your private key ID found in your JSON service account key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_gcp#private_key_id IntegrationGcp#private_key_id}
	PrivateKeyId *string `field:"required" json:"privateKeyId" yaml:"privateKeyId"`
	// Your Google Cloud project ID found in your JSON service account key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_gcp#project_id IntegrationGcp#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Silence monitors for expected GCE instance shutdowns.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_gcp#automute IntegrationGcp#automute}
	Automute interface{} `field:"optional" json:"automute" yaml:"automute"`
	// Limit the GCE instances that are pulled into Datadog by using tags.
	//
	// Only hosts that match one of the defined tags are imported into Datadog.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_gcp#host_filters IntegrationGcp#host_filters}
	HostFilters *string `field:"optional" json:"hostFilters" yaml:"hostFilters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/integration_gcp#id IntegrationGcp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

