// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference interface {
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
	InternalValue() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQuery
	SetInternalValue(val *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQuery)
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCrossOrgUuids()
	ResetGroupBy()
	ResetOperationName()
	ResetPrimaryTagName()
	ResetPrimaryTagValue()
	ResetResourceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference
type jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) CrossOrgUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) CrossOrgUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) Env() *string {
	var returns *string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) EnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) GroupBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) GroupByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) InternalValue() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQuery {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) OperationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) PrimaryTagName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) PrimaryTagNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) PrimaryTagValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) PrimaryTagValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) Stat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) StatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference_Override(d DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetCrossOrgUuids(val *[]*string) {
	if err := j.validateSetCrossOrgUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossOrgUuids",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetEnv(val *string) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetGroupBy(val *[]*string) {
	if err := j.validateSetGroupByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupBy",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetInternalValue(val *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetOperationName(val *string) {
	if err := j.validateSetOperationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetPrimaryTagName(val *string) {
	if err := j.validateSetPrimaryTagNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryTagName",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetPrimaryTagValue(val *string) {
	if err := j.validateSetPrimaryTagValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryTagValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetResourceName(val *string) {
	if err := j.validateSetResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceName",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetStat(val *string) {
	if err := j.validateSetStatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stat",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetCrossOrgUuids() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossOrgUuids",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetOperationName() {
	_jsii_.InvokeVoid(
		d,
		"resetOperationName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetPrimaryTagName() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimaryTagName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetPrimaryTagValue() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimaryTagValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetResourceName() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryApmResourceStatsQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

