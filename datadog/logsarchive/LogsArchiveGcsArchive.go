// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logsarchive


type LogsArchiveGcsArchive struct {
	// Name of your GCS bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/logs_archive#bucket LogsArchive#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Your client email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/logs_archive#client_email LogsArchive#client_email}
	ClientEmail *string `field:"required" json:"clientEmail" yaml:"clientEmail"`
	// Path where the archive is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/logs_archive#path LogsArchive#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Your project id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/logs_archive#project_id LogsArchive#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

