// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/syntheticstest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestMobileStepParamsOutputReference interface {
	cdktf.ComplexObject
	Check() *string
	SetCheck(val *string)
	CheckInput() *string
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
	Delay() *float64
	SetDelay(val *float64)
	DelayInput() *float64
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	Element() SyntheticsTestMobileStepParamsElementOutputReference
	ElementInput() *SyntheticsTestMobileStepParamsElement
	Enable() interface{}
	SetEnable(val interface{})
	EnableInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *SyntheticsTestMobileStepParams
	SetInternalValue(val *SyntheticsTestMobileStepParams)
	MaxScrolls() *float64
	SetMaxScrolls(val *float64)
	MaxScrollsInput() *float64
	Positions() SyntheticsTestMobileStepParamsPositionsList
	PositionsInput() interface{}
	SubtestPublicId() *string
	SetSubtestPublicId(val *string)
	SubtestPublicIdInput() *string
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
	Variable() SyntheticsTestMobileStepParamsVariableOutputReference
	VariableInput() *SyntheticsTestMobileStepParamsVariable
	WithEnter() interface{}
	SetWithEnter(val interface{})
	WithEnterInput() interface{}
	X() *float64
	SetX(val *float64)
	XInput() *float64
	Y() *float64
	SetY(val *float64)
	YInput() *float64
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
	PutElement(value *SyntheticsTestMobileStepParamsElement)
	PutPositions(value interface{})
	PutVariable(value *SyntheticsTestMobileStepParamsVariable)
	ResetCheck()
	ResetDelay()
	ResetDirection()
	ResetElement()
	ResetEnable()
	ResetMaxScrolls()
	ResetPositions()
	ResetSubtestPublicId()
	ResetValue()
	ResetVariable()
	ResetWithEnter()
	ResetX()
	ResetY()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SyntheticsTestMobileStepParamsOutputReference
type jsiiProxy_SyntheticsTestMobileStepParamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) Check() *string {
	var returns *string
	_jsii_.Get(
		j,
		"check",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) CheckInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) Delay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) DelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) Element() SyntheticsTestMobileStepParamsElementOutputReference {
	var returns SyntheticsTestMobileStepParamsElementOutputReference
	_jsii_.Get(
		j,
		"element",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ElementInput() *SyntheticsTestMobileStepParamsElement {
	var returns *SyntheticsTestMobileStepParamsElement
	_jsii_.Get(
		j,
		"elementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) Enable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) EnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) InternalValue() *SyntheticsTestMobileStepParams {
	var returns *SyntheticsTestMobileStepParams
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) MaxScrolls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxScrolls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) MaxScrollsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxScrollsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) Positions() SyntheticsTestMobileStepParamsPositionsList {
	var returns SyntheticsTestMobileStepParamsPositionsList
	_jsii_.Get(
		j,
		"positions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) PositionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"positionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) SubtestPublicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtestPublicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) SubtestPublicIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtestPublicIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) Variable() SyntheticsTestMobileStepParamsVariableOutputReference {
	var returns SyntheticsTestMobileStepParamsVariableOutputReference
	_jsii_.Get(
		j,
		"variable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) VariableInput() *SyntheticsTestMobileStepParamsVariable {
	var returns *SyntheticsTestMobileStepParamsVariable
	_jsii_.Get(
		j,
		"variableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) WithEnter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withEnter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) WithEnterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withEnterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) XInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"xInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) YInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yInput",
		&returns,
	)
	return returns
}


func NewSyntheticsTestMobileStepParamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SyntheticsTestMobileStepParamsOutputReference {
	_init_.Initialize()

	if err := validateNewSyntheticsTestMobileStepParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTestMobileStepParamsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestMobileStepParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSyntheticsTestMobileStepParamsOutputReference_Override(s SyntheticsTestMobileStepParamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestMobileStepParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetCheck(val *string) {
	if err := j.validateSetCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"check",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetDelay(val *float64) {
	if err := j.validateSetDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delay",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetEnable(val interface{}) {
	if err := j.validateSetEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enable",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetInternalValue(val *SyntheticsTestMobileStepParams) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetMaxScrolls(val *float64) {
	if err := j.validateSetMaxScrollsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxScrolls",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetSubtestPublicId(val *string) {
	if err := j.validateSetSubtestPublicIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subtestPublicId",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetWithEnter(val interface{}) {
	if err := j.validateSetWithEnterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withEnter",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetX(val *float64) {
	if err := j.validateSetXParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference)SetY(val *float64) {
	if err := j.validateSetYParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) PutElement(value *SyntheticsTestMobileStepParamsElement) {
	if err := s.validatePutElementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putElement",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) PutPositions(value interface{}) {
	if err := s.validatePutPositionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPositions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) PutVariable(value *SyntheticsTestMobileStepParamsVariable) {
	if err := s.validatePutVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVariable",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ResetCheck() {
	_jsii_.InvokeVoid(
		s,
		"resetCheck",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ResetDelay() {
	_jsii_.InvokeVoid(
		s,
		"resetDelay",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ResetDirection() {
	_jsii_.InvokeVoid(
		s,
		"resetDirection",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ResetElement() {
	_jsii_.InvokeVoid(
		s,
		"resetElement",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ResetEnable() {
	_jsii_.InvokeVoid(
		s,
		"resetEnable",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ResetMaxScrolls() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxScrolls",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ResetPositions() {
	_jsii_.InvokeVoid(
		s,
		"resetPositions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ResetSubtestPublicId() {
	_jsii_.InvokeVoid(
		s,
		"resetSubtestPublicId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		s,
		"resetValue",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ResetVariable() {
	_jsii_.InvokeVoid(
		s,
		"resetVariable",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ResetWithEnter() {
	_jsii_.InvokeVoid(
		s,
		"resetWithEnter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ResetX() {
	_jsii_.InvokeVoid(
		s,
		"resetX",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ResetY() {
	_jsii_.InvokeVoid(
		s,
		"resetY",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

