// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetImageDefinitionOutputReference interface {
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
	HasBackground() interface{}
	SetHasBackground(val interface{})
	HasBackgroundInput() interface{}
	HasBorder() interface{}
	SetHasBorder(val interface{})
	HasBorderInput() interface{}
	HorizontalAlign() *string
	SetHorizontalAlign(val *string)
	HorizontalAlignInput() *string
	InternalValue() *PowerpackWidgetImageDefinition
	SetInternalValue(val *PowerpackWidgetImageDefinition)
	Margin() *string
	SetMargin(val *string)
	MarginInput() *string
	Sizing() *string
	SetSizing(val *string)
	SizingInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlDarkTheme() *string
	SetUrlDarkTheme(val *string)
	UrlDarkThemeInput() *string
	UrlInput() *string
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
	ResetHasBackground()
	ResetHasBorder()
	ResetHorizontalAlign()
	ResetMargin()
	ResetSizing()
	ResetUrlDarkTheme()
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

// The jsii proxy struct for PowerpackWidgetImageDefinitionOutputReference
type jsiiProxy_PowerpackWidgetImageDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) HasBackground() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) HasBackgroundInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasBackgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) HasBorder() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasBorder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) HasBorderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasBorderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) HorizontalAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"horizontalAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) HorizontalAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"horizontalAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) InternalValue() *PowerpackWidgetImageDefinition {
	var returns *PowerpackWidgetImageDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) Margin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"margin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) MarginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) Sizing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) SizingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) UrlDarkTheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlDarkTheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) UrlDarkThemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlDarkThemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) VerticalAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) VerticalAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlignInput",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetImageDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetImageDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetImageDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetImageDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetImageDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetImageDefinitionOutputReference_Override(p PowerpackWidgetImageDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetImageDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference)SetHasBackground(val interface{}) {
	if err := j.validateSetHasBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasBackground",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference)SetHasBorder(val interface{}) {
	if err := j.validateSetHasBorderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasBorder",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference)SetHorizontalAlign(val *string) {
	if err := j.validateSetHorizontalAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"horizontalAlign",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference)SetInternalValue(val *PowerpackWidgetImageDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference)SetMargin(val *string) {
	if err := j.validateSetMarginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"margin",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference)SetSizing(val *string) {
	if err := j.validateSetSizingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizing",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference)SetUrlDarkTheme(val *string) {
	if err := j.validateSetUrlDarkThemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlDarkTheme",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference)SetVerticalAlign(val *string) {
	if err := j.validateSetVerticalAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verticalAlign",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) ResetHasBackground() {
	_jsii_.InvokeVoid(
		p,
		"resetHasBackground",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) ResetHasBorder() {
	_jsii_.InvokeVoid(
		p,
		"resetHasBorder",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) ResetHorizontalAlign() {
	_jsii_.InvokeVoid(
		p,
		"resetHorizontalAlign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) ResetMargin() {
	_jsii_.InvokeVoid(
		p,
		"resetMargin",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) ResetSizing() {
	_jsii_.InvokeVoid(
		p,
		"resetSizing",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) ResetUrlDarkTheme() {
	_jsii_.InvokeVoid(
		p,
		"resetUrlDarkTheme",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) ResetVerticalAlign() {
	_jsii_.InvokeVoid(
		p,
		"resetVerticalAlign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetImageDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

