// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v9/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v9/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetNoteDefinitionOutputReference interface {
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
	InternalValue() *DashboardWidgetNoteDefinition
	SetInternalValue(val *DashboardWidgetNoteDefinition)
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

// The jsii proxy struct for DashboardWidgetNoteDefinitionOutputReference
type jsiiProxy_DashboardWidgetNoteDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) BackgroundColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) BackgroundColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) FontSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) FontSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) HasPadding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasPadding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) HasPaddingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasPaddingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) InternalValue() *DashboardWidgetNoteDefinition {
	var returns *DashboardWidgetNoteDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ShowTick() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showTick",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ShowTickInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showTickInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) TextAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) TextAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) TickEdge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tickEdge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) TickEdgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tickEdgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) TickPos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tickPos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) TickPosInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tickPosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) VerticalAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) VerticalAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlignInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetNoteDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetNoteDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetNoteDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetNoteDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetNoteDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetNoteDefinitionOutputReference_Override(d DashboardWidgetNoteDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetNoteDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetBackgroundColor(val *string) {
	if err := j.validateSetBackgroundColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundColor",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetFontSize(val *string) {
	if err := j.validateSetFontSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontSize",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetHasPadding(val interface{}) {
	if err := j.validateSetHasPaddingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasPadding",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetInternalValue(val *DashboardWidgetNoteDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetShowTick(val interface{}) {
	if err := j.validateSetShowTickParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showTick",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetTextAlign(val *string) {
	if err := j.validateSetTextAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textAlign",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetTickEdge(val *string) {
	if err := j.validateSetTickEdgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tickEdge",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetTickPos(val *string) {
	if err := j.validateSetTickPosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tickPos",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference)SetVerticalAlign(val *string) {
	if err := j.validateSetVerticalAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verticalAlign",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ResetBackgroundColor() {
	_jsii_.InvokeVoid(
		d,
		"resetBackgroundColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ResetFontSize() {
	_jsii_.InvokeVoid(
		d,
		"resetFontSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ResetHasPadding() {
	_jsii_.InvokeVoid(
		d,
		"resetHasPadding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ResetShowTick() {
	_jsii_.InvokeVoid(
		d,
		"resetShowTick",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ResetTextAlign() {
	_jsii_.InvokeVoid(
		d,
		"resetTextAlign",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ResetTickEdge() {
	_jsii_.InvokeVoid(
		d,
		"resetTickEdge",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ResetTickPos() {
	_jsii_.InvokeVoid(
		d,
		"resetTickPos",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ResetVerticalAlign() {
	_jsii_.InvokeVoid(
		d,
		"resetVerticalAlign",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetNoteDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

