// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/syntheticstest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestMobileStepOutputReference interface {
	cdktf.ComplexObject
	AllowFailure() interface{}
	SetAllowFailure(val interface{})
	AllowFailureInput() interface{}
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
	HasNewStepElement() interface{}
	SetHasNewStepElement(val interface{})
	HasNewStepElementInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsCritical() interface{}
	SetIsCritical(val interface{})
	IsCriticalInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NoScreenshot() interface{}
	SetNoScreenshot(val interface{})
	NoScreenshotInput() interface{}
	Params() SyntheticsTestMobileStepParamsOutputReference
	ParamsInput() *SyntheticsTestMobileStepParams
	PublicId() *string
	SetPublicId(val *string)
	PublicIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutParams(value *SyntheticsTestMobileStepParams)
	ResetAllowFailure()
	ResetHasNewStepElement()
	ResetIsCritical()
	ResetNoScreenshot()
	ResetPublicId()
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SyntheticsTestMobileStepOutputReference
type jsiiProxy_SyntheticsTestMobileStepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) AllowFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) AllowFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) HasNewStepElement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasNewStepElement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) HasNewStepElementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasNewStepElementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) IsCritical() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCritical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) IsCriticalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCriticalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) NoScreenshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noScreenshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) NoScreenshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noScreenshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) Params() SyntheticsTestMobileStepParamsOutputReference {
	var returns SyntheticsTestMobileStepParamsOutputReference
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) ParamsInput() *SyntheticsTestMobileStepParams {
	var returns *SyntheticsTestMobileStepParams
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) PublicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) PublicIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewSyntheticsTestMobileStepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SyntheticsTestMobileStepOutputReference {
	_init_.Initialize()

	if err := validateNewSyntheticsTestMobileStepOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTestMobileStepOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestMobileStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSyntheticsTestMobileStepOutputReference_Override(s SyntheticsTestMobileStepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestMobileStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference)SetAllowFailure(val interface{}) {
	if err := j.validateSetAllowFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowFailure",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference)SetHasNewStepElement(val interface{}) {
	if err := j.validateSetHasNewStepElementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasNewStepElement",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference)SetIsCritical(val interface{}) {
	if err := j.validateSetIsCriticalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCritical",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference)SetNoScreenshot(val interface{}) {
	if err := j.validateSetNoScreenshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noScreenshot",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference)SetPublicId(val *string) {
	if err := j.validateSetPublicIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicId",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) PutParams(value *SyntheticsTestMobileStepParams) {
	if err := s.validatePutParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putParams",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) ResetAllowFailure() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowFailure",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) ResetHasNewStepElement() {
	_jsii_.InvokeVoid(
		s,
		"resetHasNewStepElement",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) ResetIsCritical() {
	_jsii_.InvokeVoid(
		s,
		"resetIsCritical",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) ResetNoScreenshot() {
	_jsii_.InvokeVoid(
		s,
		"resetNoScreenshot",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) ResetPublicId() {
	_jsii_.InvokeVoid(
		s,
		"resetPublicId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SyntheticsTestMobileStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

