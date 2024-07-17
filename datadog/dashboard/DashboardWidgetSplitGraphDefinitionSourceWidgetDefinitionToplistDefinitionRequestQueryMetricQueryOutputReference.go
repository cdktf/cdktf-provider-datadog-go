// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference interface {
	cdktf.ComplexObject
	Aggregator() *string
	SetAggregator(val *string)
	AggregatorInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQuery
	SetInternalValue(val *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQuery)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
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
	ResetAggregator()
	ResetCrossOrgUuids()
	ResetDataSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference
type jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) Aggregator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) AggregatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) CrossOrgUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) CrossOrgUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) InternalValue() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQuery {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference_Override(d DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference)SetAggregator(val *string) {
	if err := j.validateSetAggregatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregator",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference)SetCrossOrgUuids(val *[]*string) {
	if err := j.validateSetCrossOrgUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossOrgUuids",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference)SetInternalValue(val *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) ResetAggregator() {
	_jsii_.InvokeVoid(
		d,
		"resetAggregator",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) ResetCrossOrgUuids() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossOrgUuids",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) ResetDataSource() {
	_jsii_.InvokeVoid(
		d,
		"resetDataSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestQueryMetricQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

