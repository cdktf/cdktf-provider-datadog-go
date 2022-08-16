// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type LogsArchiveS3Archive struct {
	// Your AWS account id.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_archive#account_id LogsArchive#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of your s3 bucket.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_archive#bucket LogsArchive#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Path where the archive will be stored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_archive#path LogsArchive#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Your AWS role name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_archive#role_name LogsArchive#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
}

