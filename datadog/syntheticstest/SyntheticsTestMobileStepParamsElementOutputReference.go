// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/syntheticstest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestMobileStepParamsElementOutputReference interface {
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
	Context() *string
	SetContext(val *string)
	ContextInput() *string
	ContextType() *string
	SetContextType(val *string)
	ContextTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ElementDescription() *string
	SetElementDescription(val *string)
	ElementDescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SyntheticsTestMobileStepParamsElement
	SetInternalValue(val *SyntheticsTestMobileStepParamsElement)
	MultiLocator() *map[string]*string
	SetMultiLocator(val *map[string]*string)
	MultiLocatorInput() *map[string]*string
	RelativePosition() SyntheticsTestMobileStepParamsElementRelativePositionOutputReference
	RelativePositionInput() *SyntheticsTestMobileStepParamsElementRelativePosition
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextContent() *string
	SetTextContent(val *string)
	TextContentInput() *string
	UserLocator() SyntheticsTestMobileStepParamsElementUserLocatorOutputReference
	UserLocatorInput() *SyntheticsTestMobileStepParamsElementUserLocator
	ViewName() *string
	SetViewName(val *string)
	ViewNameInput() *string
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
	PutRelativePosition(value *SyntheticsTestMobileStepParamsElementRelativePosition)
	PutUserLocator(value *SyntheticsTestMobileStepParamsElementUserLocator)
	ResetContext()
	ResetContextType()
	ResetElementDescription()
	ResetMultiLocator()
	ResetRelativePosition()
	ResetTextContent()
	ResetUserLocator()
	ResetViewName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SyntheticsTestMobileStepParamsElementOutputReference
type jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ContextType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ContextTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ElementDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ElementDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) InternalValue() *SyntheticsTestMobileStepParamsElement {
	var returns *SyntheticsTestMobileStepParamsElement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) MultiLocator() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"multiLocator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) MultiLocatorInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"multiLocatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) RelativePosition() SyntheticsTestMobileStepParamsElementRelativePositionOutputReference {
	var returns SyntheticsTestMobileStepParamsElementRelativePositionOutputReference
	_jsii_.Get(
		j,
		"relativePosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) RelativePositionInput() *SyntheticsTestMobileStepParamsElementRelativePosition {
	var returns *SyntheticsTestMobileStepParamsElementRelativePosition
	_jsii_.Get(
		j,
		"relativePositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) TextContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) TextContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) UserLocator() SyntheticsTestMobileStepParamsElementUserLocatorOutputReference {
	var returns SyntheticsTestMobileStepParamsElementUserLocatorOutputReference
	_jsii_.Get(
		j,
		"userLocator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) UserLocatorInput() *SyntheticsTestMobileStepParamsElementUserLocator {
	var returns *SyntheticsTestMobileStepParamsElementUserLocator
	_jsii_.Get(
		j,
		"userLocatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ViewName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ViewNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewNameInput",
		&returns,
	)
	return returns
}


func NewSyntheticsTestMobileStepParamsElementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SyntheticsTestMobileStepParamsElementOutputReference {
	_init_.Initialize()

	if err := validateNewSyntheticsTestMobileStepParamsElementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestMobileStepParamsElementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSyntheticsTestMobileStepParamsElementOutputReference_Override(s SyntheticsTestMobileStepParamsElementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestMobileStepParamsElementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference)SetContextType(val *string) {
	if err := j.validateSetContextTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contextType",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference)SetElementDescription(val *string) {
	if err := j.validateSetElementDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elementDescription",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference)SetInternalValue(val *SyntheticsTestMobileStepParamsElement) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference)SetMultiLocator(val *map[string]*string) {
	if err := j.validateSetMultiLocatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiLocator",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference)SetTextContent(val *string) {
	if err := j.validateSetTextContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textContent",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference)SetViewName(val *string) {
	if err := j.validateSetViewNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewName",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) PutRelativePosition(value *SyntheticsTestMobileStepParamsElementRelativePosition) {
	if err := s.validatePutRelativePositionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRelativePosition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) PutUserLocator(value *SyntheticsTestMobileStepParamsElementUserLocator) {
	if err := s.validatePutUserLocatorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUserLocator",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ResetContext() {
	_jsii_.InvokeVoid(
		s,
		"resetContext",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ResetContextType() {
	_jsii_.InvokeVoid(
		s,
		"resetContextType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ResetElementDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetElementDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ResetMultiLocator() {
	_jsii_.InvokeVoid(
		s,
		"resetMultiLocator",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ResetRelativePosition() {
	_jsii_.InvokeVoid(
		s,
		"resetRelativePosition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ResetTextContent() {
	_jsii_.InvokeVoid(
		s,
		"resetTextContent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ResetUserLocator() {
	_jsii_.InvokeVoid(
		s,
		"resetUserLocator",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ResetViewName() {
	_jsii_.InvokeVoid(
		s,
		"resetViewName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SyntheticsTestMobileStepParamsElementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

