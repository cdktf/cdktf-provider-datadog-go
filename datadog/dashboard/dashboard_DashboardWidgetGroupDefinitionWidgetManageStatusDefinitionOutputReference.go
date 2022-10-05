package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v3/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v3/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference interface {
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
	InternalValue() *DashboardWidgetGroupDefinitionWidgetManageStatusDefinition
	SetInternalValue(val *DashboardWidgetGroupDefinitionWidgetManageStatusDefinition)
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	ShowLastTriggered() interface{}
	SetShowLastTriggered(val interface{})
	ShowLastTriggeredInput() interface{}
	ShowPriority() interface{}
	SetShowPriority(val interface{})
	ShowPriorityInput() interface{}
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
	ResetShowPriority()
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

// The jsii proxy struct for DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference
type jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ColorPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ColorPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) DisplayFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) DisplayFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) HideZeroCounts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideZeroCounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) HideZeroCountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideZeroCountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) InternalValue() *DashboardWidgetGroupDefinitionWidgetManageStatusDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetManageStatusDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ShowLastTriggered() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLastTriggered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ShowLastTriggeredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLastTriggeredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ShowPriority() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ShowPriorityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) Sort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) SortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) SummaryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) SummaryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference_Override(d DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetColorPreference(val *string) {
	if err := j.validateSetColorPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colorPreference",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetDisplayFormat(val *string) {
	if err := j.validateSetDisplayFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayFormat",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetHideZeroCounts(val interface{}) {
	if err := j.validateSetHideZeroCountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideZeroCounts",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetInternalValue(val *DashboardWidgetGroupDefinitionWidgetManageStatusDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetShowLastTriggered(val interface{}) {
	if err := j.validateSetShowLastTriggeredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showLastTriggered",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetShowPriority(val interface{}) {
	if err := j.validateSetShowPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showPriority",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetSort(val *string) {
	if err := j.validateSetSortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sort",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetSummaryType(val *string) {
	if err := j.validateSetSummaryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summaryType",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ResetColorPreference() {
	_jsii_.InvokeVoid(
		d,
		"resetColorPreference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ResetDisplayFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ResetHideZeroCounts() {
	_jsii_.InvokeVoid(
		d,
		"resetHideZeroCounts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ResetShowLastTriggered() {
	_jsii_.InvokeVoid(
		d,
		"resetShowLastTriggered",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ResetShowPriority() {
	_jsii_.InvokeVoid(
		d,
		"resetShowPriority",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		d,
		"resetSort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ResetSummaryType() {
	_jsii_.InvokeVoid(
		d,
		"resetSummaryType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

