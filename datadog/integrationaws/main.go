// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationaws

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-datadog.integrationAws.IntegrationAws",
		reflect.TypeOf((*IntegrationAws)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessKeyId", GoGetter: "AccessKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "accessKeyIdInput", GoGetter: "AccessKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "accountSpecificNamespaceRules", GoGetter: "AccountSpecificNamespaceRules"},
			_jsii_.MemberProperty{JsiiProperty: "accountSpecificNamespaceRulesInput", GoGetter: "AccountSpecificNamespaceRulesInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "cspmResourceCollectionEnabled", GoGetter: "CspmResourceCollectionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "cspmResourceCollectionEnabledInput", GoGetter: "CspmResourceCollectionEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "excludedRegions", GoGetter: "ExcludedRegions"},
			_jsii_.MemberProperty{JsiiProperty: "excludedRegionsInput", GoGetter: "ExcludedRegionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "extendedResourceCollectionEnabled", GoGetter: "ExtendedResourceCollectionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "extendedResourceCollectionEnabledInput", GoGetter: "ExtendedResourceCollectionEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalId", GoGetter: "ExternalId"},
			_jsii_.MemberProperty{JsiiProperty: "filterTags", GoGetter: "FilterTags"},
			_jsii_.MemberProperty{JsiiProperty: "filterTagsInput", GoGetter: "FilterTagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "hostTags", GoGetter: "HostTags"},
			_jsii_.MemberProperty{JsiiProperty: "hostTagsInput", GoGetter: "HostTagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "metricsCollectionEnabled", GoGetter: "MetricsCollectionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "metricsCollectionEnabledInput", GoGetter: "MetricsCollectionEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessKeyId", GoMethod: "ResetAccessKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountId", GoMethod: "ResetAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountSpecificNamespaceRules", GoMethod: "ResetAccountSpecificNamespaceRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetCspmResourceCollectionEnabled", GoMethod: "ResetCspmResourceCollectionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedRegions", GoMethod: "ResetExcludedRegions"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtendedResourceCollectionEnabled", GoMethod: "ResetExtendedResourceCollectionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetFilterTags", GoMethod: "ResetFilterTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostTags", GoMethod: "ResetHostTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetricsCollectionEnabled", GoMethod: "ResetMetricsCollectionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceCollectionEnabled", GoMethod: "ResetResourceCollectionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleName", GoMethod: "ResetRoleName"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretAccessKey", GoMethod: "ResetSecretAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "resourceCollectionEnabled", GoGetter: "ResourceCollectionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "resourceCollectionEnabledInput", GoGetter: "ResourceCollectionEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "roleName", GoGetter: "RoleName"},
			_jsii_.MemberProperty{JsiiProperty: "roleNameInput", GoGetter: "RoleNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "secretAccessKey", GoGetter: "SecretAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "secretAccessKeyInput", GoGetter: "SecretAccessKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_IntegrationAws{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-datadog.integrationAws.IntegrationAwsConfig",
		reflect.TypeOf((*IntegrationAwsConfig)(nil)).Elem(),
	)
}
