// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogsoftwarecatalog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogSoftwareCatalogConfig struct {
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
	// Filter entities by excluding snapshotted entities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/data-sources/software_catalog#filter_exclude_snapshot DataDatadogSoftwareCatalog#filter_exclude_snapshot}
	FilterExcludeSnapshot *string `field:"optional" json:"filterExcludeSnapshot" yaml:"filterExcludeSnapshot"`
	// Filter entities by UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/data-sources/software_catalog#filter_id DataDatadogSoftwareCatalog#filter_id}
	FilterId *string `field:"optional" json:"filterId" yaml:"filterId"`
	// Filter entities by kind.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/data-sources/software_catalog#filter_kind DataDatadogSoftwareCatalog#filter_kind}
	FilterKind *string `field:"optional" json:"filterKind" yaml:"filterKind"`
	// Filter entities by name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/data-sources/software_catalog#filter_name DataDatadogSoftwareCatalog#filter_name}
	FilterName *string `field:"optional" json:"filterName" yaml:"filterName"`
	// Filter entities by owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/data-sources/software_catalog#filter_owner DataDatadogSoftwareCatalog#filter_owner}
	FilterOwner *string `field:"optional" json:"filterOwner" yaml:"filterOwner"`
	// Filter entities by reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/data-sources/software_catalog#filter_ref DataDatadogSoftwareCatalog#filter_ref}
	FilterRef *string `field:"optional" json:"filterRef" yaml:"filterRef"`
	// Filter entities by relation type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/data-sources/software_catalog#filter_relation_type DataDatadogSoftwareCatalog#filter_relation_type}
	FilterRelationType *string `field:"optional" json:"filterRelationType" yaml:"filterRelationType"`
}

