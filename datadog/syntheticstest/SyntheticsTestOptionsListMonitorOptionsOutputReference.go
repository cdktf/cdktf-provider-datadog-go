// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/syntheticstest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestOptionsListMonitorOptionsOutputReference interface {
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
	EscalationMessage() *string
	SetEscalationMessage(val *string)
	EscalationMessageInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SyntheticsTestOptionsListMonitorOptions
	SetInternalValue(val *SyntheticsTestOptionsListMonitorOptions)
	NotificationPresetName() *string
	SetNotificationPresetName(val *string)
	NotificationPresetNameInput() *string
	RenotifyInterval() *float64
	SetRenotifyInterval(val *float64)
	RenotifyIntervalInput() *float64
	RenotifyOccurrences() *float64
	SetRenotifyOccurrences(val *float64)
	RenotifyOccurrencesInput() *float64
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
	ResetEscalationMessage()
	ResetNotificationPresetName()
	ResetRenotifyInterval()
	ResetRenotifyOccurrences()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SyntheticsTestOptionsListMonitorOptionsOutputReference
type jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) EscalationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escalationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) EscalationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escalationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) InternalValue() *SyntheticsTestOptionsListMonitorOptions {
	var returns *SyntheticsTestOptionsListMonitorOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) NotificationPresetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPresetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) NotificationPresetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPresetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) RenotifyInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renotifyInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) RenotifyIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renotifyIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) RenotifyOccurrences() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renotifyOccurrences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) RenotifyOccurrencesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renotifyOccurrencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSyntheticsTestOptionsListMonitorOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SyntheticsTestOptionsListMonitorOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewSyntheticsTestOptionsListMonitorOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestOptionsListMonitorOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSyntheticsTestOptionsListMonitorOptionsOutputReference_Override(s SyntheticsTestOptionsListMonitorOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestOptionsListMonitorOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference)SetEscalationMessage(val *string) {
	if err := j.validateSetEscalationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"escalationMessage",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference)SetInternalValue(val *SyntheticsTestOptionsListMonitorOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference)SetNotificationPresetName(val *string) {
	if err := j.validateSetNotificationPresetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationPresetName",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference)SetRenotifyInterval(val *float64) {
	if err := j.validateSetRenotifyIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renotifyInterval",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference)SetRenotifyOccurrences(val *float64) {
	if err := j.validateSetRenotifyOccurrencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renotifyOccurrences",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) ResetEscalationMessage() {
	_jsii_.InvokeVoid(
		s,
		"resetEscalationMessage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) ResetNotificationPresetName() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationPresetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) ResetRenotifyInterval() {
	_jsii_.InvokeVoid(
		s,
		"resetRenotifyInterval",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) ResetRenotifyOccurrences() {
	_jsii_.InvokeVoid(
		s,
		"resetRenotifyOccurrences",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SyntheticsTestOptionsListMonitorOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

