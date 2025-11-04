// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetScatterplotDefinitionRequestOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackWidgetScatterplotDefinitionRequest
	SetInternalValue(val *PowerpackWidgetScatterplotDefinitionRequest)
	ScatterplotTable() PowerpackWidgetScatterplotDefinitionRequestScatterplotTableList
	ScatterplotTableInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	X() PowerpackWidgetScatterplotDefinitionRequestXList
	XInput() interface{}
	Y() PowerpackWidgetScatterplotDefinitionRequestYList
	YInput() interface{}
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
	PutScatterplotTable(value interface{})
	PutX(value interface{})
	PutY(value interface{})
	ResetScatterplotTable()
	ResetX()
	ResetY()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetScatterplotDefinitionRequestOutputReference
type jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) InternalValue() *PowerpackWidgetScatterplotDefinitionRequest {
	var returns *PowerpackWidgetScatterplotDefinitionRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) ScatterplotTable() PowerpackWidgetScatterplotDefinitionRequestScatterplotTableList {
	var returns PowerpackWidgetScatterplotDefinitionRequestScatterplotTableList
	_jsii_.Get(
		j,
		"scatterplotTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) ScatterplotTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scatterplotTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) X() PowerpackWidgetScatterplotDefinitionRequestXList {
	var returns PowerpackWidgetScatterplotDefinitionRequestXList
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) XInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) Y() PowerpackWidgetScatterplotDefinitionRequestYList {
	var returns PowerpackWidgetScatterplotDefinitionRequestYList
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) YInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"yInput",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetScatterplotDefinitionRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetScatterplotDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetScatterplotDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetScatterplotDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetScatterplotDefinitionRequestOutputReference_Override(p PowerpackWidgetScatterplotDefinitionRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetScatterplotDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference)SetInternalValue(val *PowerpackWidgetScatterplotDefinitionRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) PutScatterplotTable(value interface{}) {
	if err := p.validatePutScatterplotTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putScatterplotTable",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) PutX(value interface{}) {
	if err := p.validatePutXParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putX",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) PutY(value interface{}) {
	if err := p.validatePutYParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putY",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) ResetScatterplotTable() {
	_jsii_.InvokeVoid(
		p,
		"resetScatterplotTable",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) ResetX() {
	_jsii_.InvokeVoid(
		p,
		"resetX",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) ResetY() {
	_jsii_.InvokeVoid(
		p,
		"resetY",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetScatterplotDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

