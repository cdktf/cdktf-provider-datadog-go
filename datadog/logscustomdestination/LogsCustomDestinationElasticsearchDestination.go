// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustomdestination


type LogsCustomDestinationElasticsearchDestination struct {
	// The destination for which logs will be forwarded to.
	//
	// Must have HTTPS scheme. Forwarding back to Datadog is not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.1/docs/resources/logs_custom_destination#endpoint LogsCustomDestination#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Name of the Elasticsearch index (must follow [Elasticsearch's criteria](https://www.elastic.co/guide/en/elasticsearch/reference/8.11/indices-create-index.html#indices-create-api-path-params)).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.1/docs/resources/logs_custom_destination#index_name LogsCustomDestination#index_name}
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.1/docs/resources/logs_custom_destination#basic_auth LogsCustomDestination#basic_auth}
	BasicAuth interface{} `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// Date pattern with US locale and UTC timezone to be appended to the index name after adding '-' 							(that is, '${index_name}-${indexPattern}').
	//
	// You can customize the index rotation naming pattern by choosing one of these options:
	// 							- Hourly: 'yyyy-MM-dd-HH' (as an example, it would render: '2022-10-19-09')
	// 							- Daily: 'yyyy-MM-dd' (as an example, it would render: '2022-10-19')
	// 							- Weekly: 'yyyy-'W'ww' (as an example, it would render: '2022-W42')
	// 							- Monthly: 'yyyy-MM' (as an example, it would render: '2022-10')
	// 							If this field is missing or is blank, it means that the index name will always be the same
	// 							(that is, no rotation).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.1/docs/resources/logs_custom_destination#index_rotation LogsCustomDestination#index_rotation}
	IndexRotation *string `field:"optional" json:"indexRotation" yaml:"indexRotation"`
}

