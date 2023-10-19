// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference interface {
	cdktf.ComplexObject
	ApmDependencyStatsQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQuery
	ApmResourceStatsQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQuery
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
	EventQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQueryOutputReference
	EventQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQueryOutputReference
	MetricQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQuery
	ProcessQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQueryOutputReference
	ProcessQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQuery
	SloQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQueryOutputReference
	SloQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQuery
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
	PutApmDependencyStatsQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQuery)
	PutApmResourceStatsQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQuery)
	PutEventQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQuery)
	PutMetricQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQuery)
	PutProcessQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQuery)
	PutSloQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQuery)
	ResetApmDependencyStatsQuery()
	ResetApmResourceStatsQuery()
	ResetEventQuery()
	ResetMetricQuery()
	ResetProcessQuery()
	ResetSloQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference
type jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ApmDependencyStatsQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQueryOutputReference {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ApmDependencyStatsQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQuery {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ApmResourceStatsQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQueryOutputReference {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ApmResourceStatsQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQuery {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) EventQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQueryOutputReference {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) EventQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQuery {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) MetricQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQueryOutputReference {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) MetricQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQuery {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ProcessQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQueryOutputReference {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ProcessQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQuery {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) SloQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQueryOutputReference {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) SloQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQuery {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference_Override(d DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) PutApmDependencyStatsQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmDependencyStatsQuery) {
	if err := d.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) PutApmResourceStatsQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryApmResourceStatsQuery) {
	if err := d.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) PutEventQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryEventQuery) {
	if err := d.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) PutMetricQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryMetricQuery) {
	if err := d.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) PutProcessQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) PutSloQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQuerySloQuery) {
	if err := d.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

