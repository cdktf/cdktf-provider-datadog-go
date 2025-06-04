// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/syntheticstest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestMobileOptionsListStructOutputReference interface {
	cdktf.ComplexObject
	AllowApplicationCrash() interface{}
	SetAllowApplicationCrash(val interface{})
	AllowApplicationCrashInput() interface{}
	Bindings() SyntheticsTestMobileOptionsListBindingsList
	BindingsInput() interface{}
	Ci() SyntheticsTestMobileOptionsListCiOutputReference
	CiInput() *SyntheticsTestMobileOptionsListCi
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
	DefaultStepTimeout() *float64
	SetDefaultStepTimeout(val *float64)
	DefaultStepTimeoutInput() *float64
	DeviceIds() *[]*string
	SetDeviceIds(val *[]*string)
	DeviceIdsInput() *[]*string
	DisableAutoAcceptAlert() interface{}
	SetDisableAutoAcceptAlert(val interface{})
	DisableAutoAcceptAlertInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *SyntheticsTestMobileOptionsListStruct
	SetInternalValue(val *SyntheticsTestMobileOptionsListStruct)
	MinFailureDuration() *float64
	SetMinFailureDuration(val *float64)
	MinFailureDurationInput() *float64
	MobileApplication() SyntheticsTestMobileOptionsListMobileApplicationOutputReference
	MobileApplicationInput() *SyntheticsTestMobileOptionsListMobileApplication
	MonitorName() *string
	SetMonitorName(val *string)
	MonitorNameInput() *string
	MonitorOptions() SyntheticsTestMobileOptionsListMonitorOptionsOutputReference
	MonitorOptionsInput() *SyntheticsTestMobileOptionsListMonitorOptions
	MonitorPriority() *float64
	SetMonitorPriority(val *float64)
	MonitorPriorityInput() *float64
	NoScreenshot() interface{}
	SetNoScreenshot(val interface{})
	NoScreenshotInput() interface{}
	RestrictedRoles() *[]*string
	SetRestrictedRoles(val *[]*string)
	RestrictedRolesInput() *[]*string
	Retry() SyntheticsTestMobileOptionsListRetryOutputReference
	RetryInput() *SyntheticsTestMobileOptionsListRetry
	Scheduling() SyntheticsTestMobileOptionsListSchedulingOutputReference
	SchedulingInput() *SyntheticsTestMobileOptionsListScheduling
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TickEvery() *float64
	SetTickEvery(val *float64)
	TickEveryInput() *float64
	Verbosity() *float64
	SetVerbosity(val *float64)
	VerbosityInput() *float64
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
	PutBindings(value interface{})
	PutCi(value *SyntheticsTestMobileOptionsListCi)
	PutMobileApplication(value *SyntheticsTestMobileOptionsListMobileApplication)
	PutMonitorOptions(value *SyntheticsTestMobileOptionsListMonitorOptions)
	PutRetry(value *SyntheticsTestMobileOptionsListRetry)
	PutScheduling(value *SyntheticsTestMobileOptionsListScheduling)
	ResetAllowApplicationCrash()
	ResetBindings()
	ResetCi()
	ResetDefaultStepTimeout()
	ResetDisableAutoAcceptAlert()
	ResetMinFailureDuration()
	ResetMonitorName()
	ResetMonitorOptions()
	ResetMonitorPriority()
	ResetNoScreenshot()
	ResetRestrictedRoles()
	ResetRetry()
	ResetScheduling()
	ResetVerbosity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SyntheticsTestMobileOptionsListStructOutputReference
type jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) AllowApplicationCrash() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowApplicationCrash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) AllowApplicationCrashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowApplicationCrashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) Bindings() SyntheticsTestMobileOptionsListBindingsList {
	var returns SyntheticsTestMobileOptionsListBindingsList
	_jsii_.Get(
		j,
		"bindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) BindingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) Ci() SyntheticsTestMobileOptionsListCiOutputReference {
	var returns SyntheticsTestMobileOptionsListCiOutputReference
	_jsii_.Get(
		j,
		"ci",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) CiInput() *SyntheticsTestMobileOptionsListCi {
	var returns *SyntheticsTestMobileOptionsListCi
	_jsii_.Get(
		j,
		"ciInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) DefaultStepTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultStepTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) DefaultStepTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultStepTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) DeviceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deviceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) DeviceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deviceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) DisableAutoAcceptAlert() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutoAcceptAlert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) DisableAutoAcceptAlertInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAutoAcceptAlertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) InternalValue() *SyntheticsTestMobileOptionsListStruct {
	var returns *SyntheticsTestMobileOptionsListStruct
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) MinFailureDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minFailureDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) MinFailureDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minFailureDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) MobileApplication() SyntheticsTestMobileOptionsListMobileApplicationOutputReference {
	var returns SyntheticsTestMobileOptionsListMobileApplicationOutputReference
	_jsii_.Get(
		j,
		"mobileApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) MobileApplicationInput() *SyntheticsTestMobileOptionsListMobileApplication {
	var returns *SyntheticsTestMobileOptionsListMobileApplication
	_jsii_.Get(
		j,
		"mobileApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) MonitorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) MonitorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) MonitorOptions() SyntheticsTestMobileOptionsListMonitorOptionsOutputReference {
	var returns SyntheticsTestMobileOptionsListMonitorOptionsOutputReference
	_jsii_.Get(
		j,
		"monitorOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) MonitorOptionsInput() *SyntheticsTestMobileOptionsListMonitorOptions {
	var returns *SyntheticsTestMobileOptionsListMonitorOptions
	_jsii_.Get(
		j,
		"monitorOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) MonitorPriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitorPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) MonitorPriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitorPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) NoScreenshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noScreenshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) NoScreenshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noScreenshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) RestrictedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) RestrictedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) Retry() SyntheticsTestMobileOptionsListRetryOutputReference {
	var returns SyntheticsTestMobileOptionsListRetryOutputReference
	_jsii_.Get(
		j,
		"retry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) RetryInput() *SyntheticsTestMobileOptionsListRetry {
	var returns *SyntheticsTestMobileOptionsListRetry
	_jsii_.Get(
		j,
		"retryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) Scheduling() SyntheticsTestMobileOptionsListSchedulingOutputReference {
	var returns SyntheticsTestMobileOptionsListSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) SchedulingInput() *SyntheticsTestMobileOptionsListScheduling {
	var returns *SyntheticsTestMobileOptionsListScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) TickEvery() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tickEvery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) TickEveryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tickEveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) Verbosity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"verbosity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) VerbosityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"verbosityInput",
		&returns,
	)
	return returns
}


func NewSyntheticsTestMobileOptionsListStructOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SyntheticsTestMobileOptionsListStructOutputReference {
	_init_.Initialize()

	if err := validateNewSyntheticsTestMobileOptionsListStructOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestMobileOptionsListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSyntheticsTestMobileOptionsListStructOutputReference_Override(s SyntheticsTestMobileOptionsListStructOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestMobileOptionsListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetAllowApplicationCrash(val interface{}) {
	if err := j.validateSetAllowApplicationCrashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowApplicationCrash",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetDefaultStepTimeout(val *float64) {
	if err := j.validateSetDefaultStepTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultStepTimeout",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetDeviceIds(val *[]*string) {
	if err := j.validateSetDeviceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceIds",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetDisableAutoAcceptAlert(val interface{}) {
	if err := j.validateSetDisableAutoAcceptAlertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutoAcceptAlert",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetInternalValue(val *SyntheticsTestMobileOptionsListStruct) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetMinFailureDuration(val *float64) {
	if err := j.validateSetMinFailureDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minFailureDuration",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetMonitorName(val *string) {
	if err := j.validateSetMonitorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitorName",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetMonitorPriority(val *float64) {
	if err := j.validateSetMonitorPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitorPriority",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetNoScreenshot(val interface{}) {
	if err := j.validateSetNoScreenshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noScreenshot",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetRestrictedRoles(val *[]*string) {
	if err := j.validateSetRestrictedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictedRoles",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetTickEvery(val *float64) {
	if err := j.validateSetTickEveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tickEvery",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference)SetVerbosity(val *float64) {
	if err := j.validateSetVerbosityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verbosity",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) PutBindings(value interface{}) {
	if err := s.validatePutBindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBindings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) PutCi(value *SyntheticsTestMobileOptionsListCi) {
	if err := s.validatePutCiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCi",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) PutMobileApplication(value *SyntheticsTestMobileOptionsListMobileApplication) {
	if err := s.validatePutMobileApplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMobileApplication",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) PutMonitorOptions(value *SyntheticsTestMobileOptionsListMonitorOptions) {
	if err := s.validatePutMonitorOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMonitorOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) PutRetry(value *SyntheticsTestMobileOptionsListRetry) {
	if err := s.validatePutRetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRetry",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) PutScheduling(value *SyntheticsTestMobileOptionsListScheduling) {
	if err := s.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putScheduling",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetAllowApplicationCrash() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowApplicationCrash",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetBindings() {
	_jsii_.InvokeVoid(
		s,
		"resetBindings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetCi() {
	_jsii_.InvokeVoid(
		s,
		"resetCi",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetDefaultStepTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultStepTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetDisableAutoAcceptAlert() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableAutoAcceptAlert",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetMinFailureDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetMinFailureDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetMonitorName() {
	_jsii_.InvokeVoid(
		s,
		"resetMonitorName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetMonitorOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetMonitorOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetMonitorPriority() {
	_jsii_.InvokeVoid(
		s,
		"resetMonitorPriority",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetNoScreenshot() {
	_jsii_.InvokeVoid(
		s,
		"resetNoScreenshot",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetRestrictedRoles() {
	_jsii_.InvokeVoid(
		s,
		"resetRestrictedRoles",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetRetry() {
	_jsii_.InvokeVoid(
		s,
		"resetRetry",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetScheduling() {
	_jsii_.InvokeVoid(
		s,
		"resetScheduling",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ResetVerbosity() {
	_jsii_.InvokeVoid(
		s,
		"resetVerbosity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SyntheticsTestMobileOptionsListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

