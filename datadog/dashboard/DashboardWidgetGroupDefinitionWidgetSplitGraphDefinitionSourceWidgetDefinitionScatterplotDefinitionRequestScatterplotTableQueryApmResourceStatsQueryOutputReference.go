// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CrossOrgUuids() *[]*string
	SetCrossOrgUuids(val *[]*string)
	CrossOrgUuidsInput() *[]*string
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	Env() *string
	SetEnv(val *string)
	EnvInput() *string
	// Experimental.
	Fqn() *string
	GroupBy() *[]*string
	SetGroupBy(val *[]*string)
	GroupByInput() *[]*string
	InternalValue() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery
	SetInternalValue(val *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery)
	Name() *string
	SetName(val *string)
	NameInput() *string
	OperationName() *string
	SetOperationName(val *string)
	OperationNameInput() *string
	PrimaryTagName() *string
	SetPrimaryTagName(val *string)
	PrimaryTagNameInput() *string
	PrimaryTagValue() *string
	SetPrimaryTagValue(val *string)
	PrimaryTagValueInput() *string
	ResourceName() *string
	SetResourceName(val *string)
	ResourceNameInput() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	Stat() *string
	SetStat(val *string)
	StatInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetCrossOrgUuids()
	ResetGroupBy()
	ResetOperationName()
	ResetPrimaryTagName()
	ResetPrimaryTagValue()
	ResetResourceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference
type jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) CrossOrgUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) CrossOrgUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) Env() *string {
	var returns *string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) EnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) GroupBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) GroupByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) InternalValue() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) OperationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) PrimaryTagName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) PrimaryTagNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) PrimaryTagValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) PrimaryTagValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) ResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) Stat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) StatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference_Override(d DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetCrossOrgUuids(val *[]*string) {
	if err := j.validateSetCrossOrgUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossOrgUuids",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetEnv(val *string) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetGroupBy(val *[]*string) {
	if err := j.validateSetGroupByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupBy",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetInternalValue(val *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetOperationName(val *string) {
	if err := j.validateSetOperationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetPrimaryTagName(val *string) {
	if err := j.validateSetPrimaryTagNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryTagName",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetPrimaryTagValue(val *string) {
	if err := j.validateSetPrimaryTagValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryTagValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetResourceName(val *string) {
	if err := j.validateSetResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceName",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetStat(val *string) {
	if err := j.validateSetStatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stat",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) ResetCrossOrgUuids() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossOrgUuids",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) ResetOperationName() {
	_jsii_.InvokeVoid(
		d,
		"resetOperationName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) ResetPrimaryTagName() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimaryTagName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) ResetPrimaryTagValue() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimaryTagValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) ResetResourceName() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

