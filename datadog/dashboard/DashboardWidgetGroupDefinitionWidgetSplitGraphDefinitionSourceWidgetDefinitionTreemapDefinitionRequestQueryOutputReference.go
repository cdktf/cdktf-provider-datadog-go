// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference interface {
	cdktf.ComplexObject
	ApmDependencyStatsQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmDependencyStatsQuery
	ApmResourceStatsQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmResourceStatsQuery
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
	EventQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryEventQueryOutputReference
	EventQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryMetricQueryOutputReference
	MetricQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryMetricQuery
	ProcessQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryProcessQueryOutputReference
	ProcessQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryProcessQuery
	SloQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQuerySloQueryOutputReference
	SloQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQuerySloQuery
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
	PutApmDependencyStatsQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmDependencyStatsQuery)
	PutApmResourceStatsQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmResourceStatsQuery)
	PutEventQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryEventQuery)
	PutMetricQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryMetricQuery)
	PutProcessQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryProcessQuery)
	PutSloQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQuerySloQuery)
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

// The jsii proxy struct for DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference
type jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ApmDependencyStatsQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmDependencyStatsQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ApmDependencyStatsQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmDependencyStatsQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ApmResourceStatsQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmResourceStatsQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ApmResourceStatsQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmResourceStatsQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) EventQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryEventQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) EventQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryEventQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) MetricQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryMetricQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) MetricQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryMetricQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ProcessQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryProcessQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ProcessQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryProcessQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) SloQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQuerySloQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) SloQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQuerySloQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference_Override(d DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) PutApmDependencyStatsQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmDependencyStatsQuery) {
	if err := d.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) PutApmResourceStatsQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryApmResourceStatsQuery) {
	if err := d.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) PutEventQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryEventQuery) {
	if err := d.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) PutMetricQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryMetricQuery) {
	if err := d.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) PutProcessQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) PutSloQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQuerySloQuery) {
	if err := d.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

