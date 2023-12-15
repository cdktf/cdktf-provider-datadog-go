// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetHostmapDefinitionRequestFillOutputReference interface {
	cdktf.ComplexObject
	ApmQuery() PowerpackWidgetHostmapDefinitionRequestFillApmQueryOutputReference
	ApmQueryInput() *PowerpackWidgetHostmapDefinitionRequestFillApmQuery
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() PowerpackWidgetHostmapDefinitionRequestFillLogQueryOutputReference
	LogQueryInput() *PowerpackWidgetHostmapDefinitionRequestFillLogQuery
	ProcessQuery() PowerpackWidgetHostmapDefinitionRequestFillProcessQueryOutputReference
	ProcessQueryInput() *PowerpackWidgetHostmapDefinitionRequestFillProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	RumQuery() PowerpackWidgetHostmapDefinitionRequestFillRumQueryOutputReference
	RumQueryInput() *PowerpackWidgetHostmapDefinitionRequestFillRumQuery
	SecurityQuery() PowerpackWidgetHostmapDefinitionRequestFillSecurityQueryOutputReference
	SecurityQueryInput() *PowerpackWidgetHostmapDefinitionRequestFillSecurityQuery
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
	PutApmQuery(value *PowerpackWidgetHostmapDefinitionRequestFillApmQuery)
	PutLogQuery(value *PowerpackWidgetHostmapDefinitionRequestFillLogQuery)
	PutProcessQuery(value *PowerpackWidgetHostmapDefinitionRequestFillProcessQuery)
	PutRumQuery(value *PowerpackWidgetHostmapDefinitionRequestFillRumQuery)
	PutSecurityQuery(value *PowerpackWidgetHostmapDefinitionRequestFillSecurityQuery)
	ResetApmQuery()
	ResetLogQuery()
	ResetProcessQuery()
	ResetQ()
	ResetRumQuery()
	ResetSecurityQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetHostmapDefinitionRequestFillOutputReference
type jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ApmQuery() PowerpackWidgetHostmapDefinitionRequestFillApmQueryOutputReference {
	var returns PowerpackWidgetHostmapDefinitionRequestFillApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ApmQueryInput() *PowerpackWidgetHostmapDefinitionRequestFillApmQuery {
	var returns *PowerpackWidgetHostmapDefinitionRequestFillApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) LogQuery() PowerpackWidgetHostmapDefinitionRequestFillLogQueryOutputReference {
	var returns PowerpackWidgetHostmapDefinitionRequestFillLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) LogQueryInput() *PowerpackWidgetHostmapDefinitionRequestFillLogQuery {
	var returns *PowerpackWidgetHostmapDefinitionRequestFillLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ProcessQuery() PowerpackWidgetHostmapDefinitionRequestFillProcessQueryOutputReference {
	var returns PowerpackWidgetHostmapDefinitionRequestFillProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ProcessQueryInput() *PowerpackWidgetHostmapDefinitionRequestFillProcessQuery {
	var returns *PowerpackWidgetHostmapDefinitionRequestFillProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) RumQuery() PowerpackWidgetHostmapDefinitionRequestFillRumQueryOutputReference {
	var returns PowerpackWidgetHostmapDefinitionRequestFillRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) RumQueryInput() *PowerpackWidgetHostmapDefinitionRequestFillRumQuery {
	var returns *PowerpackWidgetHostmapDefinitionRequestFillRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) SecurityQuery() PowerpackWidgetHostmapDefinitionRequestFillSecurityQueryOutputReference {
	var returns PowerpackWidgetHostmapDefinitionRequestFillSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) SecurityQueryInput() *PowerpackWidgetHostmapDefinitionRequestFillSecurityQuery {
	var returns *PowerpackWidgetHostmapDefinitionRequestFillSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetHostmapDefinitionRequestFillOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackWidgetHostmapDefinitionRequestFillOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetHostmapDefinitionRequestFillOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetHostmapDefinitionRequestFillOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackWidgetHostmapDefinitionRequestFillOutputReference_Override(p PowerpackWidgetHostmapDefinitionRequestFillOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetHostmapDefinitionRequestFillOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) PutApmQuery(value *PowerpackWidgetHostmapDefinitionRequestFillApmQuery) {
	if err := p.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) PutLogQuery(value *PowerpackWidgetHostmapDefinitionRequestFillLogQuery) {
	if err := p.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) PutProcessQuery(value *PowerpackWidgetHostmapDefinitionRequestFillProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) PutRumQuery(value *PowerpackWidgetHostmapDefinitionRequestFillRumQuery) {
	if err := p.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) PutSecurityQuery(value *PowerpackWidgetHostmapDefinitionRequestFillSecurityQuery) {
	if err := p.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		p,
		"resetQ",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetHostmapDefinitionRequestFillOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

