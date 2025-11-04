// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetImageDefinitionOutputReference interface {
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
	InternalValue() *DashboardWidgetImageDefinition
	SetInternalValue(val *DashboardWidgetImageDefinition)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetHasBackground()
	ResetHasBorder()
	ResetHorizontalAlign()
	ResetMargin()
	ResetSizing()
	ResetUrlDarkTheme()
	ResetVerticalAlign()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetImageDefinitionOutputReference
type jsiiProxy_DashboardWidgetImageDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) HasBackground() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) HasBackgroundInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasBackgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) HasBorder() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasBorder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) HasBorderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasBorderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) HorizontalAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"horizontalAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) HorizontalAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"horizontalAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) InternalValue() *DashboardWidgetImageDefinition {
	var returns *DashboardWidgetImageDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) Margin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"margin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) MarginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) Sizing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) SizingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) UrlDarkTheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlDarkTheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) UrlDarkThemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlDarkThemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) VerticalAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) VerticalAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verticalAlignInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetImageDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetImageDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetImageDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetImageDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetImageDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetImageDefinitionOutputReference_Override(d DashboardWidgetImageDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetImageDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference)SetHasBackground(val interface{}) {
	if err := j.validateSetHasBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasBackground",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference)SetHasBorder(val interface{}) {
	if err := j.validateSetHasBorderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasBorder",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference)SetHorizontalAlign(val *string) {
	if err := j.validateSetHorizontalAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"horizontalAlign",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference)SetInternalValue(val *DashboardWidgetImageDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference)SetMargin(val *string) {
	if err := j.validateSetMarginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"margin",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference)SetSizing(val *string) {
	if err := j.validateSetSizingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizing",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference)SetUrlDarkTheme(val *string) {
	if err := j.validateSetUrlDarkThemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlDarkTheme",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetImageDefinitionOutputReference)SetVerticalAlign(val *string) {
	if err := j.validateSetVerticalAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verticalAlign",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) ResetHasBackground() {
	_jsii_.InvokeVoid(
		d,
		"resetHasBackground",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) ResetHasBorder() {
	_jsii_.InvokeVoid(
		d,
		"resetHasBorder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) ResetHorizontalAlign() {
	_jsii_.InvokeVoid(
		d,
		"resetHorizontalAlign",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) ResetMargin() {
	_jsii_.InvokeVoid(
		d,
		"resetMargin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) ResetSizing() {
	_jsii_.InvokeVoid(
		d,
		"resetSizing",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) ResetUrlDarkTheme() {
	_jsii_.InvokeVoid(
		d,
		"resetUrlDarkTheme",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) ResetVerticalAlign() {
	_jsii_.InvokeVoid(
		d,
		"resetVerticalAlign",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetImageDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

