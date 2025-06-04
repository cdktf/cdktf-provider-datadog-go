// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/syntheticstest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestBrowserStepOutputReference interface {
	cdktf.ComplexObject
	AllowFailure() interface{}
	SetAllowFailure(val interface{})
	AllowFailureInput() interface{}
	AlwaysExecute() interface{}
	SetAlwaysExecute(val interface{})
	AlwaysExecuteInput() interface{}
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
	ExitIfSucceed() interface{}
	SetExitIfSucceed(val interface{})
	ExitIfSucceedInput() interface{}
	ForceElementUpdate() interface{}
	SetForceElementUpdate(val interface{})
	ForceElementUpdateInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsCritical() interface{}
	SetIsCritical(val interface{})
	IsCriticalInput() interface{}
	LocalKey() *string
	SetLocalKey(val *string)
	LocalKeyInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NoScreenshot() interface{}
	SetNoScreenshot(val interface{})
	NoScreenshotInput() interface{}
	Params() SyntheticsTestBrowserStepParamsOutputReference
	ParamsInput() *SyntheticsTestBrowserStepParams
	PublicId() *string
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
	PutParams(value *SyntheticsTestBrowserStepParams)
	ResetAllowFailure()
	ResetAlwaysExecute()
	ResetExitIfSucceed()
	ResetForceElementUpdate()
	ResetIsCritical()
	ResetLocalKey()
	ResetNoScreenshot()
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

// The jsii proxy struct for SyntheticsTestBrowserStepOutputReference
type jsiiProxy_SyntheticsTestBrowserStepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) AllowFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) AllowFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) AlwaysExecute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysExecute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) AlwaysExecuteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysExecuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ExitIfSucceed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exitIfSucceed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ExitIfSucceedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exitIfSucceedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ForceElementUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceElementUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ForceElementUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceElementUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) IsCritical() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCritical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) IsCriticalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCriticalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) LocalKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) LocalKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) NoScreenshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noScreenshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) NoScreenshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noScreenshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) Params() SyntheticsTestBrowserStepParamsOutputReference {
	var returns SyntheticsTestBrowserStepParamsOutputReference
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ParamsInput() *SyntheticsTestBrowserStepParams {
	var returns *SyntheticsTestBrowserStepParams
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) PublicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewSyntheticsTestBrowserStepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SyntheticsTestBrowserStepOutputReference {
	_init_.Initialize()

	if err := validateNewSyntheticsTestBrowserStepOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTestBrowserStepOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestBrowserStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSyntheticsTestBrowserStepOutputReference_Override(s SyntheticsTestBrowserStepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestBrowserStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetAllowFailure(val interface{}) {
	if err := j.validateSetAllowFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowFailure",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetAlwaysExecute(val interface{}) {
	if err := j.validateSetAlwaysExecuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysExecute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetExitIfSucceed(val interface{}) {
	if err := j.validateSetExitIfSucceedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exitIfSucceed",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetForceElementUpdate(val interface{}) {
	if err := j.validateSetForceElementUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceElementUpdate",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetIsCritical(val interface{}) {
	if err := j.validateSetIsCriticalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCritical",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetLocalKey(val *string) {
	if err := j.validateSetLocalKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localKey",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetNoScreenshot(val interface{}) {
	if err := j.validateSetNoScreenshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noScreenshot",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) PutParams(value *SyntheticsTestBrowserStepParams) {
	if err := s.validatePutParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putParams",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ResetAllowFailure() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowFailure",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ResetAlwaysExecute() {
	_jsii_.InvokeVoid(
		s,
		"resetAlwaysExecute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ResetExitIfSucceed() {
	_jsii_.InvokeVoid(
		s,
		"resetExitIfSucceed",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ResetForceElementUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetForceElementUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ResetIsCritical() {
	_jsii_.InvokeVoid(
		s,
		"resetIsCritical",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ResetLocalKey() {
	_jsii_.InvokeVoid(
		s,
		"resetLocalKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ResetNoScreenshot() {
	_jsii_.InvokeVoid(
		s,
		"resetNoScreenshot",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SyntheticsTestBrowserStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

