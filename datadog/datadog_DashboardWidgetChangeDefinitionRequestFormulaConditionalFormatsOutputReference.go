// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/jsii"

	"github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference interface {
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

// The jsii proxy struct for DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference
type jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Comparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ComparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) CustomBgColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customBgColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) CustomBgColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customBgColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) CustomFgColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFgColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) CustomFgColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFgColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) HideValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) HideValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ImageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ImageUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Metric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) MetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Palette() *string {
	var returns *string
	_jsii_.Get(
		j,
		"palette",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) PaletteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paletteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Timeframe() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeframe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) TimeframeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeframeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference_Override(d DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetComparator(val *string) {
	_jsii_.Set(
		j,
		"comparator",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetCustomBgColor(val *string) {
	_jsii_.Set(
		j,
		"customBgColor",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetCustomFgColor(val *string) {
	_jsii_.Set(
		j,
		"customFgColor",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetHideValue(val interface{}) {
	_jsii_.Set(
		j,
		"hideValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetImageUrl(val *string) {
	_jsii_.Set(
		j,
		"imageUrl",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetMetric(val *string) {
	_jsii_.Set(
		j,
		"metric",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetPalette(val *string) {
	_jsii_.Set(
		j,
		"palette",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetTimeframe(val *string) {
	_jsii_.Set(
		j,
		"timeframe",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ResetCustomBgColor() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomBgColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ResetCustomFgColor() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomFgColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ResetHideValue() {
	_jsii_.InvokeVoid(
		d,
		"resetHideValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ResetImageUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetImageUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		d,
		"resetMetric",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ResetTimeframe() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeframe",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestFormulaConditionalFormatsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
