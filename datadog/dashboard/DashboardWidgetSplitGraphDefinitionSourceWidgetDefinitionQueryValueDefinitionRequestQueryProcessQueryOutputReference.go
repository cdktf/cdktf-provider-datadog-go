// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference interface {
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
	InternalValue() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQuery
	SetInternalValue(val *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQuery)
	IsNormalizedCpu() interface{}
	SetIsNormalizedCpu(val interface{})
	IsNormalizedCpuInput() interface{}
	Limit() *float64
	SetLimit(val *float64)
	LimitInput() *float64
	Metric() *string
	SetMetric(val *string)
	MetricInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Sort() *string
	SetSort(val *string)
	SortInput() *string
	TagFilters() *[]*string
	SetTagFilters(val *[]*string)
	TagFiltersInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextFilter() *string
	SetTextFilter(val *string)
	TextFilterInput() *string
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
	ResetIsNormalizedCpu()
	ResetLimit()
	ResetSort()
	ResetTagFilters()
	ResetTextFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference
type jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) Aggregator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) AggregatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) CrossOrgUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) CrossOrgUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) InternalValue() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQuery {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) IsNormalizedCpu() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNormalizedCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) IsNormalizedCpuInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNormalizedCpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) LimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) Metric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) MetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) Sort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) SortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) TagFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) TagFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) TextFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) TextFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textFilterInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference_Override(d DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetAggregator(val *string) {
	if err := j.validateSetAggregatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregator",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetCrossOrgUuids(val *[]*string) {
	if err := j.validateSetCrossOrgUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossOrgUuids",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetInternalValue(val *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetIsNormalizedCpu(val interface{}) {
	if err := j.validateSetIsNormalizedCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isNormalizedCpu",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetLimit(val *float64) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetMetric(val *string) {
	if err := j.validateSetMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metric",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetSort(val *string) {
	if err := j.validateSetSortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sort",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetTagFilters(val *[]*string) {
	if err := j.validateSetTagFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagFilters",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference)SetTextFilter(val *string) {
	if err := j.validateSetTextFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textFilter",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) ResetAggregator() {
	_jsii_.InvokeVoid(
		d,
		"resetAggregator",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) ResetCrossOrgUuids() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossOrgUuids",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) ResetIsNormalizedCpu() {
	_jsii_.InvokeVoid(
		d,
		"resetIsNormalizedCpu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		d,
		"resetSort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) ResetTagFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetTagFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) ResetTextFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetTextFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryProcessQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

