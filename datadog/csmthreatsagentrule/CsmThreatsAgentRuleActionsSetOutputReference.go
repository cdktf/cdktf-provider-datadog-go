// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csmthreatsagentrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/csmthreatsagentrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CsmThreatsAgentRuleActionsSetOutputReference interface {
	cdktf.ComplexObject
	Append() interface{}
	SetAppend(val interface{})
	AppendInput() interface{}
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
	DefaultValue() *string
	SetDefaultValue(val *string)
	DefaultValueInput() *string
	Expression() *string
	SetExpression(val *string)
	ExpressionInput() *string
	Field() *string
	SetField(val *string)
	FieldInput() *string
	// Experimental.
	Fqn() *string
	Inherited() interface{}
	SetInherited(val interface{})
	InheritedInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAppend()
	ResetDefaultValue()
	ResetExpression()
	ResetField()
	ResetInherited()
	ResetScope()
	ResetSize()
	ResetTtl()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CsmThreatsAgentRuleActionsSetOutputReference
type jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) Append() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"append",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) AppendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) DefaultValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) DefaultValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) Inherited() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inherited",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) InheritedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inheritedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewCsmThreatsAgentRuleActionsSetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CsmThreatsAgentRuleActionsSetOutputReference {
	_init_.Initialize()

	if err := validateNewCsmThreatsAgentRuleActionsSetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.csmThreatsAgentRule.CsmThreatsAgentRuleActionsSetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCsmThreatsAgentRuleActionsSetOutputReference_Override(c CsmThreatsAgentRuleActionsSetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.csmThreatsAgentRule.CsmThreatsAgentRuleActionsSetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetAppend(val interface{}) {
	if err := j.validateSetAppendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"append",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetDefaultValue(val *string) {
	if err := j.validateSetDefaultValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultValue",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetExpression(val *string) {
	if err := j.validateSetExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetInherited(val interface{}) {
	if err := j.validateSetInheritedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inherited",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ResetAppend() {
	_jsii_.InvokeVoid(
		c,
		"resetAppend",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ResetDefaultValue() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ResetExpression() {
	_jsii_.InvokeVoid(
		c,
		"resetExpression",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ResetField() {
	_jsii_.InvokeVoid(
		c,
		"resetField",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ResetInherited() {
	_jsii_.InvokeVoid(
		c,
		"resetInherited",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		c,
		"resetScope",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ResetSize() {
	_jsii_.InvokeVoid(
		c,
		"resetSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		c,
		"resetValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CsmThreatsAgentRuleActionsSetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

