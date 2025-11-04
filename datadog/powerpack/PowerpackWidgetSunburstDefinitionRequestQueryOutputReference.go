// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetSunburstDefinitionRequestQueryOutputReference interface {
	cdktf.ComplexObject
	ApmDependencyStatsQuery() PowerpackWidgetSunburstDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	ApmDependencyStatsQueryInput() *PowerpackWidgetSunburstDefinitionRequestQueryApmDependencyStatsQuery
	ApmResourceStatsQuery() PowerpackWidgetSunburstDefinitionRequestQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *PowerpackWidgetSunburstDefinitionRequestQueryApmResourceStatsQuery
	CloudCostQuery() PowerpackWidgetSunburstDefinitionRequestQueryCloudCostQueryOutputReference
	CloudCostQueryInput() *PowerpackWidgetSunburstDefinitionRequestQueryCloudCostQuery
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
	EventQuery() PowerpackWidgetSunburstDefinitionRequestQueryEventQueryOutputReference
	EventQueryInput() *PowerpackWidgetSunburstDefinitionRequestQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricQuery() PowerpackWidgetSunburstDefinitionRequestQueryMetricQueryOutputReference
	MetricQueryInput() *PowerpackWidgetSunburstDefinitionRequestQueryMetricQuery
	ProcessQuery() PowerpackWidgetSunburstDefinitionRequestQueryProcessQueryOutputReference
	ProcessQueryInput() *PowerpackWidgetSunburstDefinitionRequestQueryProcessQuery
	SloQuery() PowerpackWidgetSunburstDefinitionRequestQuerySloQueryOutputReference
	SloQueryInput() *PowerpackWidgetSunburstDefinitionRequestQuerySloQuery
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
	PutApmDependencyStatsQuery(value *PowerpackWidgetSunburstDefinitionRequestQueryApmDependencyStatsQuery)
	PutApmResourceStatsQuery(value *PowerpackWidgetSunburstDefinitionRequestQueryApmResourceStatsQuery)
	PutCloudCostQuery(value *PowerpackWidgetSunburstDefinitionRequestQueryCloudCostQuery)
	PutEventQuery(value *PowerpackWidgetSunburstDefinitionRequestQueryEventQuery)
	PutMetricQuery(value *PowerpackWidgetSunburstDefinitionRequestQueryMetricQuery)
	PutProcessQuery(value *PowerpackWidgetSunburstDefinitionRequestQueryProcessQuery)
	PutSloQuery(value *PowerpackWidgetSunburstDefinitionRequestQuerySloQuery)
	ResetApmDependencyStatsQuery()
	ResetApmResourceStatsQuery()
	ResetCloudCostQuery()
	ResetEventQuery()
	ResetMetricQuery()
	ResetProcessQuery()
	ResetSloQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetSunburstDefinitionRequestQueryOutputReference
type jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ApmDependencyStatsQuery() PowerpackWidgetSunburstDefinitionRequestQueryApmDependencyStatsQueryOutputReference {
	var returns PowerpackWidgetSunburstDefinitionRequestQueryApmDependencyStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmDependencyStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ApmDependencyStatsQueryInput() *PowerpackWidgetSunburstDefinitionRequestQueryApmDependencyStatsQuery {
	var returns *PowerpackWidgetSunburstDefinitionRequestQueryApmDependencyStatsQuery
	_jsii_.Get(
		j,
		"apmDependencyStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ApmResourceStatsQuery() PowerpackWidgetSunburstDefinitionRequestQueryApmResourceStatsQueryOutputReference {
	var returns PowerpackWidgetSunburstDefinitionRequestQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ApmResourceStatsQueryInput() *PowerpackWidgetSunburstDefinitionRequestQueryApmResourceStatsQuery {
	var returns *PowerpackWidgetSunburstDefinitionRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) CloudCostQuery() PowerpackWidgetSunburstDefinitionRequestQueryCloudCostQueryOutputReference {
	var returns PowerpackWidgetSunburstDefinitionRequestQueryCloudCostQueryOutputReference
	_jsii_.Get(
		j,
		"cloudCostQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) CloudCostQueryInput() *PowerpackWidgetSunburstDefinitionRequestQueryCloudCostQuery {
	var returns *PowerpackWidgetSunburstDefinitionRequestQueryCloudCostQuery
	_jsii_.Get(
		j,
		"cloudCostQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) EventQuery() PowerpackWidgetSunburstDefinitionRequestQueryEventQueryOutputReference {
	var returns PowerpackWidgetSunburstDefinitionRequestQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) EventQueryInput() *PowerpackWidgetSunburstDefinitionRequestQueryEventQuery {
	var returns *PowerpackWidgetSunburstDefinitionRequestQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) MetricQuery() PowerpackWidgetSunburstDefinitionRequestQueryMetricQueryOutputReference {
	var returns PowerpackWidgetSunburstDefinitionRequestQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) MetricQueryInput() *PowerpackWidgetSunburstDefinitionRequestQueryMetricQuery {
	var returns *PowerpackWidgetSunburstDefinitionRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ProcessQuery() PowerpackWidgetSunburstDefinitionRequestQueryProcessQueryOutputReference {
	var returns PowerpackWidgetSunburstDefinitionRequestQueryProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ProcessQueryInput() *PowerpackWidgetSunburstDefinitionRequestQueryProcessQuery {
	var returns *PowerpackWidgetSunburstDefinitionRequestQueryProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) SloQuery() PowerpackWidgetSunburstDefinitionRequestQuerySloQueryOutputReference {
	var returns PowerpackWidgetSunburstDefinitionRequestQuerySloQueryOutputReference
	_jsii_.Get(
		j,
		"sloQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) SloQueryInput() *PowerpackWidgetSunburstDefinitionRequestQuerySloQuery {
	var returns *PowerpackWidgetSunburstDefinitionRequestQuerySloQuery
	_jsii_.Get(
		j,
		"sloQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetSunburstDefinitionRequestQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackWidgetSunburstDefinitionRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetSunburstDefinitionRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetSunburstDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackWidgetSunburstDefinitionRequestQueryOutputReference_Override(p PowerpackWidgetSunburstDefinitionRequestQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetSunburstDefinitionRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) PutApmDependencyStatsQuery(value *PowerpackWidgetSunburstDefinitionRequestQueryApmDependencyStatsQuery) {
	if err := p.validatePutApmDependencyStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmDependencyStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) PutApmResourceStatsQuery(value *PowerpackWidgetSunburstDefinitionRequestQueryApmResourceStatsQuery) {
	if err := p.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) PutCloudCostQuery(value *PowerpackWidgetSunburstDefinitionRequestQueryCloudCostQuery) {
	if err := p.validatePutCloudCostQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudCostQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) PutEventQuery(value *PowerpackWidgetSunburstDefinitionRequestQueryEventQuery) {
	if err := p.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) PutMetricQuery(value *PowerpackWidgetSunburstDefinitionRequestQueryMetricQuery) {
	if err := p.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) PutProcessQuery(value *PowerpackWidgetSunburstDefinitionRequestQueryProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) PutSloQuery(value *PowerpackWidgetSunburstDefinitionRequestQuerySloQuery) {
	if err := p.validatePutSloQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSloQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ResetApmDependencyStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmDependencyStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ResetCloudCostQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudCostQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ResetSloQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSloQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetSunburstDefinitionRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

