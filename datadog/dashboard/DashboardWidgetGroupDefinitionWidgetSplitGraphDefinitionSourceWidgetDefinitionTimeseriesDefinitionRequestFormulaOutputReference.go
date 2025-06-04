// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference interface {
	cdktf.ComplexObject
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	CellDisplayMode() *string
	SetCellDisplayMode(val *string)
	CellDisplayModeInput() *string
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
	ConditionalFormats() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaConditionalFormatsList
	ConditionalFormatsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FormulaExpression() *string
	SetFormulaExpression(val *string)
	FormulaExpressionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Limit() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaLimitOutputReference
	LimitInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaLimit
	NumberFormat() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaNumberFormatOutputReference
	NumberFormatInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaNumberFormat
	Style() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaStyleOutputReference
	StyleInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaStyle
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
	PutConditionalFormats(value interface{})
	PutLimit(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaLimit)
	PutNumberFormat(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaNumberFormat)
	PutStyle(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaStyle)
	ResetAlias()
	ResetCellDisplayMode()
	ResetConditionalFormats()
	ResetLimit()
	ResetNumberFormat()
	ResetStyle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference
type jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) CellDisplayMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cellDisplayMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) CellDisplayModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cellDisplayModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) ConditionalFormats() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaConditionalFormatsList {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaConditionalFormatsList
	_jsii_.Get(
		j,
		"conditionalFormats",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) ConditionalFormatsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionalFormatsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) FormulaExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formulaExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) FormulaExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formulaExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) Limit() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaLimitOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaLimitOutputReference
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) LimitInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaLimit {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaLimit
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) NumberFormat() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaNumberFormatOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaNumberFormatOutputReference
	_jsii_.Get(
		j,
		"numberFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) NumberFormatInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaNumberFormat {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaNumberFormat
	_jsii_.Get(
		j,
		"numberFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) Style() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaStyleOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) StyleInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaStyle {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference_Override(d DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference)SetAlias(val *string) {
	if err := j.validateSetAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference)SetCellDisplayMode(val *string) {
	if err := j.validateSetCellDisplayModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cellDisplayMode",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference)SetFormulaExpression(val *string) {
	if err := j.validateSetFormulaExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formulaExpression",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) PutConditionalFormats(value interface{}) {
	if err := d.validatePutConditionalFormatsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConditionalFormats",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) PutLimit(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaLimit) {
	if err := d.validatePutLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLimit",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) PutNumberFormat(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaNumberFormat) {
	if err := d.validatePutNumberFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNumberFormat",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) PutStyle(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaStyle) {
	if err := d.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStyle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) ResetAlias() {
	_jsii_.InvokeVoid(
		d,
		"resetAlias",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) ResetCellDisplayMode() {
	_jsii_.InvokeVoid(
		d,
		"resetCellDisplayMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) ResetConditionalFormats() {
	_jsii_.InvokeVoid(
		d,
		"resetConditionalFormats",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) ResetNumberFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		d,
		"resetStyle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

