// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/jsii"

	"github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetManageStatusDefinitionOutputReference interface {
	cdktf.ComplexObject
	ColorPreference() *string
	SetColorPreference(val *string)
	ColorPreferenceInput() *string
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
	DisplayFormat() *string
	SetDisplayFormat(val *string)
	DisplayFormatInput() *string
	// Experimental.
	Fqn() *string
	HideZeroCounts() interface{}
	SetHideZeroCounts(val interface{})
	HideZeroCountsInput() interface{}
	InternalValue() *DashboardWidgetManageStatusDefinition
	SetInternalValue(val *DashboardWidgetManageStatusDefinition)
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	ShowLastTriggered() interface{}
	SetShowLastTriggered(val interface{})
	ShowLastTriggeredInput() interface{}
	Sort() *string
	SetSort(val *string)
	SortInput() *string
	SummaryType() *string
	SetSummaryType(val *string)
	SummaryTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() *string
	SetTitle(val *string)
	TitleAlign() *string
	SetTitleAlign(val *string)
	TitleAlignInput() *string
	TitleInput() *string
	TitleSize() *string
	SetTitleSize(val *string)
	TitleSizeInput() *string
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
	ResetColorPreference()
	ResetDisplayFormat()
	ResetHideZeroCounts()
	ResetShowLastTriggered()
	ResetSort()
	ResetSummaryType()
	ResetTitle()
	ResetTitleAlign()
	ResetTitleSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetManageStatusDefinitionOutputReference
type jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ColorPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ColorPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) DisplayFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) DisplayFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) HideZeroCounts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideZeroCounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) HideZeroCountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideZeroCountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) InternalValue() *DashboardWidgetManageStatusDefinition {
	var returns *DashboardWidgetManageStatusDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ShowLastTriggered() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLastTriggered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ShowLastTriggeredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLastTriggeredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) Sort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SummaryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SummaryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetManageStatusDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetManageStatusDefinitionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.DashboardWidgetManageStatusDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetManageStatusDefinitionOutputReference_Override(d DashboardWidgetManageStatusDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.DashboardWidgetManageStatusDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetColorPreference(val *string) {
	_jsii_.Set(
		j,
		"colorPreference",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetDisplayFormat(val *string) {
	_jsii_.Set(
		j,
		"displayFormat",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetHideZeroCounts(val interface{}) {
	_jsii_.Set(
		j,
		"hideZeroCounts",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetInternalValue(val *DashboardWidgetManageStatusDefinition) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetQuery(val *string) {
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetShowLastTriggered(val interface{}) {
	_jsii_.Set(
		j,
		"showLastTriggered",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetSort(val *string) {
	_jsii_.Set(
		j,
		"sort",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetSummaryType(val *string) {
	_jsii_.Set(
		j,
		"summaryType",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetTitle(val *string) {
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetTitleAlign(val *string) {
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) SetTitleSize(val *string) {
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ResetColorPreference() {
	_jsii_.InvokeVoid(
		d,
		"resetColorPreference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ResetDisplayFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ResetHideZeroCounts() {
	_jsii_.InvokeVoid(
		d,
		"resetHideZeroCounts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ResetShowLastTriggered() {
	_jsii_.InvokeVoid(
		d,
		"resetShowLastTriggered",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		d,
		"resetSort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ResetSummaryType() {
	_jsii_.InvokeVoid(
		d,
		"resetSummaryType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetManageStatusDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

