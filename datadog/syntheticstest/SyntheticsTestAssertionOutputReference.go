// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/syntheticstest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestAssertionOutputReference interface {
	cdktf.ComplexObject
	Code() *string
	SetCode(val *string)
	CodeInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	Property() *string
	SetProperty(val *string)
	PropertyInput() *string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
	Targetjsonpath() SyntheticsTestAssertionTargetjsonpathOutputReference
	TargetjsonpathInput() *SyntheticsTestAssertionTargetjsonpath
	Targetjsonschema() SyntheticsTestAssertionTargetjsonschemaOutputReference
	TargetjsonschemaInput() *SyntheticsTestAssertionTargetjsonschema
	Targetxpath() SyntheticsTestAssertionTargetxpathOutputReference
	TargetxpathInput() *SyntheticsTestAssertionTargetxpath
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimingsScope() *string
	SetTimingsScope(val *string)
	TimingsScopeInput() *string
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
	PutTargetjsonpath(value *SyntheticsTestAssertionTargetjsonpath)
	PutTargetjsonschema(value *SyntheticsTestAssertionTargetjsonschema)
	PutTargetxpath(value *SyntheticsTestAssertionTargetxpath)
	ResetCode()
	ResetOperator()
	ResetProperty()
	ResetTarget()
	ResetTargetjsonpath()
	ResetTargetjsonschema()
	ResetTargetxpath()
	ResetTimingsScope()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SyntheticsTestAssertionOutputReference
type jsiiProxy_SyntheticsTestAssertionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) Code() *string {
	var returns *string
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) CodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) Property() *string {
	var returns *string
	_jsii_.Get(
		j,
		"property",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) PropertyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) Targetjsonpath() SyntheticsTestAssertionTargetjsonpathOutputReference {
	var returns SyntheticsTestAssertionTargetjsonpathOutputReference
	_jsii_.Get(
		j,
		"targetjsonpath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) TargetjsonpathInput() *SyntheticsTestAssertionTargetjsonpath {
	var returns *SyntheticsTestAssertionTargetjsonpath
	_jsii_.Get(
		j,
		"targetjsonpathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) Targetjsonschema() SyntheticsTestAssertionTargetjsonschemaOutputReference {
	var returns SyntheticsTestAssertionTargetjsonschemaOutputReference
	_jsii_.Get(
		j,
		"targetjsonschema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) TargetjsonschemaInput() *SyntheticsTestAssertionTargetjsonschema {
	var returns *SyntheticsTestAssertionTargetjsonschema
	_jsii_.Get(
		j,
		"targetjsonschemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) Targetxpath() SyntheticsTestAssertionTargetxpathOutputReference {
	var returns SyntheticsTestAssertionTargetxpathOutputReference
	_jsii_.Get(
		j,
		"targetxpath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) TargetxpathInput() *SyntheticsTestAssertionTargetxpath {
	var returns *SyntheticsTestAssertionTargetxpath
	_jsii_.Get(
		j,
		"targetxpathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) TimingsScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timingsScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) TimingsScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timingsScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewSyntheticsTestAssertionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SyntheticsTestAssertionOutputReference {
	_init_.Initialize()

	if err := validateNewSyntheticsTestAssertionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTestAssertionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestAssertionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSyntheticsTestAssertionOutputReference_Override(s SyntheticsTestAssertionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestAssertionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference)SetCode(val *string) {
	if err := j.validateSetCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference)SetProperty(val *string) {
	if err := j.validateSetPropertyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"property",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference)SetTimingsScope(val *string) {
	if err := j.validateSetTimingsScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timingsScope",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestAssertionOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) PutTargetjsonpath(value *SyntheticsTestAssertionTargetjsonpath) {
	if err := s.validatePutTargetjsonpathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTargetjsonpath",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) PutTargetjsonschema(value *SyntheticsTestAssertionTargetjsonschema) {
	if err := s.validatePutTargetjsonschemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTargetjsonschema",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) PutTargetxpath(value *SyntheticsTestAssertionTargetxpath) {
	if err := s.validatePutTargetxpathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTargetxpath",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) ResetCode() {
	_jsii_.InvokeVoid(
		s,
		"resetCode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) ResetOperator() {
	_jsii_.InvokeVoid(
		s,
		"resetOperator",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) ResetProperty() {
	_jsii_.InvokeVoid(
		s,
		"resetProperty",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		s,
		"resetTarget",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) ResetTargetjsonpath() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetjsonpath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) ResetTargetjsonschema() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetjsonschema",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) ResetTargetxpath() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetxpath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) ResetTimingsScope() {
	_jsii_.InvokeVoid(
		s,
		"resetTimingsScope",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SyntheticsTestAssertionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

