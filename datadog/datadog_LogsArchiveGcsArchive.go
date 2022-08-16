// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type LogsArchiveGcsArchive struct {
	// Name of your GCS bucket.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_archive#bucket LogsArchive#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Your client email.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_archive#client_email LogsArchive#client_email}
	ClientEmail *string `field:"required" json:"clientEmail" yaml:"clientEmail"`
	// Path where the archive will be stored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_archive#path LogsArchive#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Your project id.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/logs_archive#project_id LogsArchive#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

