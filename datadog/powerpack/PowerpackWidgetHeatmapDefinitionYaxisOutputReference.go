// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetHeatmapDefinitionYaxisOutputReference interface {
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
	IncludeZero() interface{}
	SetIncludeZero(val interface{})
	IncludeZeroInput() interface{}
	InternalValue() *PowerpackWidgetHeatmapDefinitionYaxis
	SetInternalValue(val *PowerpackWidgetHeatmapDefinitionYaxis)
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Max() *string
	SetMax(val *string)
	MaxInput() *string
	Min() *string
	SetMin(val *string)
	MinInput() *string
	Scale() *string
	SetScale(val *string)
	ScaleInput() *string
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
	ResetIncludeZero()
	ResetLabel()
	ResetMax()
	ResetMin()
	ResetScale()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetHeatmapDefinitionYaxisOutputReference
type jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) IncludeZero() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeZero",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) IncludeZeroInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeZeroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) InternalValue() *PowerpackWidgetHeatmapDefinitionYaxis {
	var returns *PowerpackWidgetHeatmapDefinitionYaxis
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) Max() *string {
	var returns *string
	_jsii_.Get(
		j,
		"max",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) MaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) Min() *string {
	var returns *string
	_jsii_.Get(
		j,
		"min",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) MinInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) Scale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) ScaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetHeatmapDefinitionYaxisOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetHeatmapDefinitionYaxisOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetHeatmapDefinitionYaxisOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetHeatmapDefinitionYaxisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetHeatmapDefinitionYaxisOutputReference_Override(p PowerpackWidgetHeatmapDefinitionYaxisOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetHeatmapDefinitionYaxisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference)SetIncludeZero(val interface{}) {
	if err := j.validateSetIncludeZeroParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeZero",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference)SetInternalValue(val *PowerpackWidgetHeatmapDefinitionYaxis) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference)SetMax(val *string) {
	if err := j.validateSetMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"max",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference)SetMin(val *string) {
	if err := j.validateSetMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"min",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference)SetScale(val *string) {
	if err := j.validateSetScaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scale",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) ResetIncludeZero() {
	_jsii_.InvokeVoid(
		p,
		"resetIncludeZero",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		p,
		"resetLabel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) ResetMax() {
	_jsii_.InvokeVoid(
		p,
		"resetMax",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) ResetMin() {
	_jsii_.InvokeVoid(
		p,
		"resetMin",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) ResetScale() {
	_jsii_.InvokeVoid(
		p,
		"resetScale",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionYaxisOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

