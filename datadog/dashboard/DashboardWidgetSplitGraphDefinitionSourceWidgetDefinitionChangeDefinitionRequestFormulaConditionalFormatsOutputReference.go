// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference interface {
	cdktf.ComplexObject
	Comparator() *string
	SetComparator(val *string)
	ComparatorInput() *string
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
	CustomBgColor() *string
	SetCustomBgColor(val *string)
	CustomBgColorInput() *string
	CustomFgColor() *string
	SetCustomFgColor(val *string)
	CustomFgColorInput() *string
	// Experimental.
	Fqn() *string
	HideValue() interface{}
	SetHideValue(val interface{})
	HideValueInput() interface{}
	ImageUrl() *string
	SetImageUrl(val *string)
	ImageUrlInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Metric() *string
	SetMetric(val *string)
	MetricInput() *string
	Palette() *string
	SetPalette(val *string)
	PaletteInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeframe() *string
	SetTimeframe(val *string)
	TimeframeInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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
	ResetCustomBgColor()
	ResetCustomFgColor()
	ResetHideValue()
	ResetImageUrl()
	ResetMetric()
	ResetTimeframe()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference
type jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Comparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ComparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) CustomBgColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customBgColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) CustomBgColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customBgColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) CustomFgColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFgColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) CustomFgColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFgColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) HideValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) HideValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ImageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ImageUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Metric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) MetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Palette() *string {
	var returns *string
	_jsii_.Get(
		j,
		"palette",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) PaletteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paletteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Timeframe() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeframe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) TimeframeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeframeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference_Override(d DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetComparator(val *string) {
	if err := j.validateSetComparatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparator",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetCustomBgColor(val *string) {
	if err := j.validateSetCustomBgColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customBgColor",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetCustomFgColor(val *string) {
	if err := j.validateSetCustomFgColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customFgColor",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetHideValue(val interface{}) {
	if err := j.validateSetHideValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetImageUrl(val *string) {
	if err := j.validateSetImageUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUrl",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetMetric(val *string) {
	if err := j.validateSetMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metric",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetPalette(val *string) {
	if err := j.validateSetPaletteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"palette",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetTimeframe(val *string) {
	if err := j.validateSetTimeframeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeframe",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference)SetValue(val *float64) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ResetCustomBgColor() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomBgColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ResetCustomFgColor() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomFgColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ResetHideValue() {
	_jsii_.InvokeVoid(
		d,
		"resetHideValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ResetImageUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetImageUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		d,
		"resetMetric",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ResetTimeframe() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeframe",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

