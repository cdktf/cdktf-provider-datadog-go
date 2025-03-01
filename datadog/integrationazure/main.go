// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationazure

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-datadog.integrationAzure.IntegrationAzure",
		reflect.TypeOf((*IntegrationAzure)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "appServicePlanFilters", GoGetter: "AppServicePlanFilters"},
			_jsii_.MemberProperty{JsiiProperty: "appServicePlanFiltersInput", GoGetter: "AppServicePlanFiltersInput"},
			_jsii_.MemberProperty{JsiiProperty: "automute", GoGetter: "Automute"},
			_jsii_.MemberProperty{JsiiProperty: "automuteInput", GoGetter: "AutomuteInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretInput", GoGetter: "ClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "containerAppFilters", GoGetter: "ContainerAppFilters"},
			_jsii_.MemberProperty{JsiiProperty: "containerAppFiltersInput", GoGetter: "ContainerAppFiltersInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "cspmEnabled", GoGetter: "CspmEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "cspmEnabledInput", GoGetter: "CspmEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "customMetricsEnabled", GoGetter: "CustomMetricsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "customMetricsEnabledInput", GoGetter: "CustomMetricsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
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
			_jsii_.MemberProperty{JsiiProperty: "hostFilters", GoGetter: "HostFilters"},
			_jsii_.MemberProperty{JsiiProperty: "hostFiltersInput", GoGetter: "HostFiltersInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "metricsEnabled", GoGetter: "MetricsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "metricsEnabledDefault", GoGetter: "MetricsEnabledDefault"},
			_jsii_.MemberProperty{JsiiProperty: "metricsEnabledDefaultInput", GoGetter: "MetricsEnabledDefaultInput"},
			_jsii_.MemberProperty{JsiiProperty: "metricsEnabledInput", GoGetter: "MetricsEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putResourceProviderConfigs", GoMethod: "PutResourceProviderConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppServicePlanFilters", GoMethod: "ResetAppServicePlanFilters"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutomute", GoMethod: "ResetAutomute"},
			_jsii_.MemberMethod{JsiiMethod: "resetContainerAppFilters", GoMethod: "ResetContainerAppFilters"},
			_jsii_.MemberMethod{JsiiMethod: "resetCspmEnabled", GoMethod: "ResetCspmEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomMetricsEnabled", GoMethod: "ResetCustomMetricsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostFilters", GoMethod: "ResetHostFilters"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetricsEnabled", GoMethod: "ResetMetricsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetricsEnabledDefault", GoMethod: "ResetMetricsEnabledDefault"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceCollectionEnabled", GoMethod: "ResetResourceCollectionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceProviderConfigs", GoMethod: "ResetResourceProviderConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsageMetricsEnabled", GoMethod: "ResetUsageMetricsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "resourceCollectionEnabled", GoGetter: "ResourceCollectionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "resourceCollectionEnabledInput", GoGetter: "ResourceCollectionEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceProviderConfigs", GoGetter: "ResourceProviderConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "resourceProviderConfigsInput", GoGetter: "ResourceProviderConfigsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tenantName", GoGetter: "TenantName"},
			_jsii_.MemberProperty{JsiiProperty: "tenantNameInput", GoGetter: "TenantNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "usageMetricsEnabled", GoGetter: "UsageMetricsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "usageMetricsEnabledInput", GoGetter: "UsageMetricsEnabledInput"},
		},
		func() interface{} {
			j := jsiiProxy_IntegrationAzure{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-datadog.integrationAzure.IntegrationAzureConfig",
		reflect.TypeOf((*IntegrationAzureConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-datadog.integrationAzure.IntegrationAzureResourceProviderConfigs",
		reflect.TypeOf((*IntegrationAzureResourceProviderConfigs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-datadog.integrationAzure.IntegrationAzureResourceProviderConfigsList",
		reflect.TypeOf((*IntegrationAzureResourceProviderConfigsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_IntegrationAzureResourceProviderConfigsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-datadog.integrationAzure.IntegrationAzureResourceProviderConfigsOutputReference",
		reflect.TypeOf((*IntegrationAzureResourceProviderConfigsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "metricsEnabled", GoGetter: "MetricsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "metricsEnabledInput", GoGetter: "MetricsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetricsEnabled", GoMethod: "ResetMetricsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IntegrationAzureResourceProviderConfigsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
