// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafcustomrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/appsecwafcustomrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsecWafCustomRuleConditionParametersOutputReference interface {
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
	Data() *string
	SetData(val *string)
	DataInput() *string
	// Experimental.
	Fqn() *string
	Input() AppsecWafCustomRuleConditionParametersInputList
	InputInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	List() *[]*string
	SetList(val *[]*string)
	ListInput() *[]*string
	Options() AppsecWafCustomRuleConditionParametersOptionsOutputReference
	OptionsInput() interface{}
	Regex() *string
	SetRegex(val *string)
	RegexInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	PutInput(value interface{})
	PutOptions(value *AppsecWafCustomRuleConditionParametersOptions)
	ResetData()
	ResetInput()
	ResetList()
	ResetOptions()
	ResetRegex()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsecWafCustomRuleConditionParametersOutputReference
type jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) Data() *string {
	var returns *string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) DataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) Input() AppsecWafCustomRuleConditionParametersInputList {
	var returns AppsecWafCustomRuleConditionParametersInputList
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) InputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) List() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"list",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) ListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) Options() AppsecWafCustomRuleConditionParametersOptionsOutputReference {
	var returns AppsecWafCustomRuleConditionParametersOptionsOutputReference
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) OptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) Regex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) RegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewAppsecWafCustomRuleConditionParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppsecWafCustomRuleConditionParametersOutputReference {
	_init_.Initialize()

	if err := validateNewAppsecWafCustomRuleConditionParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.appsecWafCustomRule.AppsecWafCustomRuleConditionParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppsecWafCustomRuleConditionParametersOutputReference_Override(a AppsecWafCustomRuleConditionParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.appsecWafCustomRule.AppsecWafCustomRuleConditionParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference)SetData(val *string) {
	if err := j.validateSetDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"data",
		val,
	)
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference)SetList(val *[]*string) {
	if err := j.validateSetListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"list",
		val,
	)
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference)SetRegex(val *string) {
	if err := j.validateSetRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regex",
		val,
	)
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) PutInput(value interface{}) {
	if err := a.validatePutInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) PutOptions(value *AppsecWafCustomRuleConditionParametersOptions) {
	if err := a.validatePutOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) ResetData() {
	_jsii_.InvokeVoid(
		a,
		"resetData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) ResetInput() {
	_jsii_.InvokeVoid(
		a,
		"resetInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) ResetList() {
	_jsii_.InvokeVoid(
		a,
		"resetList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) ResetOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		a,
		"resetValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafCustomRuleConditionParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

