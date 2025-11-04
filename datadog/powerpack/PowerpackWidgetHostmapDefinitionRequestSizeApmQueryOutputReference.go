// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference interface {
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
	ComputeQuery() PowerpackWidgetHostmapDefinitionRequestSizeApmQueryComputeQueryOutputReference
	ComputeQueryInput() *PowerpackWidgetHostmapDefinitionRequestSizeApmQueryComputeQuery
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GroupBy() PowerpackWidgetHostmapDefinitionRequestSizeApmQueryGroupByList
	GroupByInput() interface{}
	Index() *string
	SetIndex(val *string)
	IndexInput() *string
	InternalValue() *PowerpackWidgetHostmapDefinitionRequestSizeApmQuery
	SetInternalValue(val *PowerpackWidgetHostmapDefinitionRequestSizeApmQuery)
	MultiCompute() PowerpackWidgetHostmapDefinitionRequestSizeApmQueryMultiComputeList
	MultiComputeInput() interface{}
	SearchQuery() *string
	SetSearchQuery(val *string)
	SearchQueryInput() *string
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
	PutComputeQuery(value *PowerpackWidgetHostmapDefinitionRequestSizeApmQueryComputeQuery)
	PutGroupBy(value interface{})
	PutMultiCompute(value interface{})
	ResetComputeQuery()
	ResetGroupBy()
	ResetMultiCompute()
	ResetSearchQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference
type jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) ComputeQuery() PowerpackWidgetHostmapDefinitionRequestSizeApmQueryComputeQueryOutputReference {
	var returns PowerpackWidgetHostmapDefinitionRequestSizeApmQueryComputeQueryOutputReference
	_jsii_.Get(
		j,
		"computeQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) ComputeQueryInput() *PowerpackWidgetHostmapDefinitionRequestSizeApmQueryComputeQuery {
	var returns *PowerpackWidgetHostmapDefinitionRequestSizeApmQueryComputeQuery
	_jsii_.Get(
		j,
		"computeQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) GroupBy() PowerpackWidgetHostmapDefinitionRequestSizeApmQueryGroupByList {
	var returns PowerpackWidgetHostmapDefinitionRequestSizeApmQueryGroupByList
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) GroupByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) IndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) InternalValue() *PowerpackWidgetHostmapDefinitionRequestSizeApmQuery {
	var returns *PowerpackWidgetHostmapDefinitionRequestSizeApmQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) MultiCompute() PowerpackWidgetHostmapDefinitionRequestSizeApmQueryMultiComputeList {
	var returns PowerpackWidgetHostmapDefinitionRequestSizeApmQueryMultiComputeList
	_jsii_.Get(
		j,
		"multiCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) MultiComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) SearchQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) SearchQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference_Override(p PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference)SetIndex(val *string) {
	if err := j.validateSetIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference)SetInternalValue(val *PowerpackWidgetHostmapDefinitionRequestSizeApmQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference)SetSearchQuery(val *string) {
	if err := j.validateSetSearchQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchQuery",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) PutComputeQuery(value *PowerpackWidgetHostmapDefinitionRequestSizeApmQueryComputeQuery) {
	if err := p.validatePutComputeQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putComputeQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) PutGroupBy(value interface{}) {
	if err := p.validatePutGroupByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGroupBy",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) PutMultiCompute(value interface{}) {
	if err := p.validatePutMultiComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMultiCompute",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) ResetComputeQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetComputeQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) ResetMultiCompute() {
	_jsii_.InvokeVoid(
		p,
		"resetMultiCompute",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) ResetSearchQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSearchQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestSizeApmQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

