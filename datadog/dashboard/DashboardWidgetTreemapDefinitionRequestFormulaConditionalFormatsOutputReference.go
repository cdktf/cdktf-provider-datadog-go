// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference interface {
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

// The jsii proxy struct for DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference
type jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) Comparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ComparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) CustomBgColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customBgColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) CustomBgColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customBgColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) CustomFgColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFgColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) CustomFgColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFgColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) HideValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) HideValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ImageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ImageUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) Metric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) MetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) Palette() *string {
	var returns *string
	_jsii_.Get(
		j,
		"palette",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) PaletteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paletteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) Timeframe() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeframe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) TimeframeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeframeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference_Override(d DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetComparator(val *string) {
	if err := j.validateSetComparatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparator",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetCustomBgColor(val *string) {
	if err := j.validateSetCustomBgColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customBgColor",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetCustomFgColor(val *string) {
	if err := j.validateSetCustomFgColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customFgColor",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetHideValue(val interface{}) {
	if err := j.validateSetHideValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetImageUrl(val *string) {
	if err := j.validateSetImageUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUrl",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetMetric(val *string) {
	if err := j.validateSetMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metric",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetPalette(val *string) {
	if err := j.validateSetPaletteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"palette",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetTimeframe(val *string) {
	if err := j.validateSetTimeframeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeframe",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference)SetValue(val *float64) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ResetCustomBgColor() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomBgColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ResetCustomFgColor() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomFgColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ResetHideValue() {
	_jsii_.InvokeVoid(
		d,
		"resetHideValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ResetImageUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetImageUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		d,
		"resetMetric",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ResetTimeframe() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeframe",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaConditionalFormatsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

