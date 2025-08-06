// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-datadog.provider.DatadogProvider",
		reflect.TypeOf((*DatadogProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiKey", GoGetter: "ApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyInput", GoGetter: "ApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiUrl", GoGetter: "ApiUrl"},
			_jsii_.MemberProperty{JsiiProperty: "apiUrlInput", GoGetter: "ApiUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "appKey", GoGetter: "AppKey"},
			_jsii_.MemberProperty{JsiiProperty: "appKeyInput", GoGetter: "AppKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccessKeyId", GoGetter: "AwsAccessKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccessKeyIdInput", GoGetter: "AwsAccessKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSecretAccessKey", GoGetter: "AwsSecretAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "awsSecretAccessKeyInput", GoGetter: "AwsSecretAccessKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSessionToken", GoGetter: "AwsSessionToken"},
			_jsii_.MemberProperty{JsiiProperty: "awsSessionTokenInput", GoGetter: "AwsSessionTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cloudProviderRegion", GoGetter: "CloudProviderRegion"},
			_jsii_.MemberProperty{JsiiProperty: "cloudProviderRegionInput", GoGetter: "CloudProviderRegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudProviderType", GoGetter: "CloudProviderType"},
			_jsii_.MemberProperty{JsiiProperty: "cloudProviderTypeInput", GoGetter: "CloudProviderTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTags", GoGetter: "DefaultTags"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTagsInput", GoGetter: "DefaultTagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "httpClientRetryBackoffBase", GoGetter: "HttpClientRetryBackoffBase"},
			_jsii_.MemberProperty{JsiiProperty: "httpClientRetryBackoffBaseInput", GoGetter: "HttpClientRetryBackoffBaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpClientRetryBackoffMultiplier", GoGetter: "HttpClientRetryBackoffMultiplier"},
			_jsii_.MemberProperty{JsiiProperty: "httpClientRetryBackoffMultiplierInput", GoGetter: "HttpClientRetryBackoffMultiplierInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpClientRetryEnabled", GoGetter: "HttpClientRetryEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "httpClientRetryEnabledInput", GoGetter: "HttpClientRetryEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpClientRetryMaxRetries", GoGetter: "HttpClientRetryMaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "httpClientRetryMaxRetriesInput", GoGetter: "HttpClientRetryMaxRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpClientRetryTimeout", GoGetter: "HttpClientRetryTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "httpClientRetryTimeoutInput", GoGetter: "HttpClientRetryTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "orgUuid", GoGetter: "OrgUuid"},
			_jsii_.MemberProperty{JsiiProperty: "orgUuidInput", GoGetter: "OrgUuidInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiKey", GoMethod: "ResetApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiUrl", GoMethod: "ResetApiUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppKey", GoMethod: "ResetAppKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsAccessKeyId", GoMethod: "ResetAwsAccessKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSecretAccessKey", GoMethod: "ResetAwsSecretAccessKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSessionToken", GoMethod: "ResetAwsSessionToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudProviderRegion", GoMethod: "ResetCloudProviderRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudProviderType", GoMethod: "ResetCloudProviderType"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTags", GoMethod: "ResetDefaultTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpClientRetryBackoffBase", GoMethod: "ResetHttpClientRetryBackoffBase"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpClientRetryBackoffMultiplier", GoMethod: "ResetHttpClientRetryBackoffMultiplier"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpClientRetryEnabled", GoMethod: "ResetHttpClientRetryEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpClientRetryMaxRetries", GoMethod: "ResetHttpClientRetryMaxRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpClientRetryTimeout", GoMethod: "ResetHttpClientRetryTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrgUuid", GoMethod: "ResetOrgUuid"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetValidate", GoMethod: "ResetValidate"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "validate", GoGetter: "Validate"},
			_jsii_.MemberProperty{JsiiProperty: "validateInput", GoGetter: "ValidateInput"},
		},
		func() interface{} {
			j := jsiiProxy_DatadogProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-datadog.provider.DatadogProviderConfig",
		reflect.TypeOf((*DatadogProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-datadog.provider.DatadogProviderDefaultTags",
		reflect.TypeOf((*DatadogProviderDefaultTags)(nil)).Elem(),
	)
}
