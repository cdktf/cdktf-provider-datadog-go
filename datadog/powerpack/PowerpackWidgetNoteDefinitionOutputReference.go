// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetNoteDefinitionOutputReference interface {
	cdktf.ComplexObject
	BackgroundColor() *string
	SetBackgroundColor(val *string)
	BackgroundColorInput() *string
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
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FontSize() *string
	SetFontSize(val *string)
	FontSizeInput() *string
	// Experimental.
	Fqn() *string
	HasPadding() interface{}
	SetHasPadding(val interface{})
	HasPaddingInput() interface{}
	InternalValue() *PowerpackWidgetNoteDefinition
	SetInternalValue(val *PowerpackWidgetNoteDefinition)
	ShowTick() interface{}
	SetShowTick(val interface{})
	ShowTickInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextAlign() *string
	SetTextAlign(val *string)
	TextAlignInput() *string
	TickEdge() *string
	SetTickEdge(val *string)
	TickEdgeInput() *string
	TickPos() *string
	SetTickPos(val *string)
	TickPosInput() *string
	VerticalAlign() *string
	SetVerticalAlign(val *string)
	VerticalAlignInput() *string
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
	ResetBackgroundColor()
	ResetFontSize()
	ResetHasPadding()
	ResetShowTick()
	ResetTextAlign()
	ResetTickEdge()
	ResetTickPos()
	ResetVerticalAlign()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetNoteDefinitionOutputReference
type jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) BackgroundColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) BackgroundColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) FontSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) FontSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) HasPadding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasPadding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) HasPaddingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasPaddingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) InternalValue() *PowerpackWidgetNoteDefinition {
	var returns *PowerpackWidgetNoteDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ShowTick() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showTick",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ShowTickInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showTickInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) TextAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) TextAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) TickEdge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tickEdge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) TickEdgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tickEdgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) TickPos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tickPos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) TickPosInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tickPosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) VerticalAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) VerticalAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlignInput",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetNoteDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetNoteDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetNoteDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetNoteDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetNoteDefinitionOutputReference_Override(p PowerpackWidgetNoteDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetNoteDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetBackgroundColor(val *string) {
	if err := j.validateSetBackgroundColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundColor",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetFontSize(val *string) {
	if err := j.validateSetFontSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontSize",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetHasPadding(val interface{}) {
	if err := j.validateSetHasPaddingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasPadding",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetInternalValue(val *PowerpackWidgetNoteDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetShowTick(val interface{}) {
	if err := j.validateSetShowTickParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showTick",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetTextAlign(val *string) {
	if err := j.validateSetTextAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textAlign",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetTickEdge(val *string) {
	if err := j.validateSetTickEdgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tickEdge",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetTickPos(val *string) {
	if err := j.validateSetTickPosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tickPos",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference)SetVerticalAlign(val *string) {
	if err := j.validateSetVerticalAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verticalAlign",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ResetBackgroundColor() {
	_jsii_.InvokeVoid(
		p,
		"resetBackgroundColor",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ResetFontSize() {
	_jsii_.InvokeVoid(
		p,
		"resetFontSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ResetHasPadding() {
	_jsii_.InvokeVoid(
		p,
		"resetHasPadding",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ResetShowTick() {
	_jsii_.InvokeVoid(
		p,
		"resetShowTick",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ResetTextAlign() {
	_jsii_.InvokeVoid(
		p,
		"resetTextAlign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ResetTickEdge() {
	_jsii_.InvokeVoid(
		p,
		"resetTickEdge",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ResetTickPos() {
	_jsii_.InvokeVoid(
		p,
		"resetTickPos",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ResetVerticalAlign() {
	_jsii_.InvokeVoid(
		p,
		"resetVerticalAlign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetNoteDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

