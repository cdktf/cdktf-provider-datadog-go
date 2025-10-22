// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/syntheticstest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestOptionsListStructOutputReference interface {
	cdktf.ComplexObject
	AcceptSelfSigned() interface{}
	SetAcceptSelfSigned(val interface{})
	AcceptSelfSignedInput() interface{}
	AllowInsecure() interface{}
	SetAllowInsecure(val interface{})
	AllowInsecureInput() interface{}
	BlockedRequestPatterns() *[]*string
	SetBlockedRequestPatterns(val *[]*string)
	BlockedRequestPatternsInput() *[]*string
	CheckCertificateRevocation() interface{}
	SetCheckCertificateRevocation(val interface{})
	CheckCertificateRevocationInput() interface{}
	Ci() SyntheticsTestOptionsListCiOutputReference
	CiInput() *SyntheticsTestOptionsListCi
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
	DisableAiaIntermediateFetching() interface{}
	SetDisableAiaIntermediateFetching(val interface{})
	DisableAiaIntermediateFetchingInput() interface{}
	DisableCors() interface{}
	SetDisableCors(val interface{})
	DisableCorsInput() interface{}
	DisableCsp() interface{}
	SetDisableCsp(val interface{})
	DisableCspInput() interface{}
	FollowRedirects() interface{}
	SetFollowRedirects(val interface{})
	FollowRedirectsInput() interface{}
	// Experimental.
	Fqn() *string
	HttpVersion() *string
	SetHttpVersion(val *string)
	HttpVersionInput() *string
	IgnoreServerCertificateError() interface{}
	SetIgnoreServerCertificateError(val interface{})
	IgnoreServerCertificateErrorInput() interface{}
	InitialNavigationTimeout() *float64
	SetInitialNavigationTimeout(val *float64)
	InitialNavigationTimeoutInput() *float64
	InternalValue() *SyntheticsTestOptionsListStruct
	SetInternalValue(val *SyntheticsTestOptionsListStruct)
	MinFailureDuration() *float64
	SetMinFailureDuration(val *float64)
	MinFailureDurationInput() *float64
	MinLocationFailed() *float64
	SetMinLocationFailed(val *float64)
	MinLocationFailedInput() *float64
	MonitorName() *string
	SetMonitorName(val *string)
	MonitorNameInput() *string
	MonitorOptions() SyntheticsTestOptionsListMonitorOptionsOutputReference
	MonitorOptionsInput() *SyntheticsTestOptionsListMonitorOptions
	MonitorPriority() *float64
	SetMonitorPriority(val *float64)
	MonitorPriorityInput() *float64
	NoScreenshot() interface{}
	SetNoScreenshot(val interface{})
	NoScreenshotInput() interface{}
	RestrictedRoles() *[]*string
	SetRestrictedRoles(val *[]*string)
	RestrictedRolesInput() *[]*string
	Retry() SyntheticsTestOptionsListRetryOutputReference
	RetryInput() *SyntheticsTestOptionsListRetry
	RumSettings() SyntheticsTestOptionsListRumSettingsOutputReference
	RumSettingsInput() *SyntheticsTestOptionsListRumSettings
	Scheduling() SyntheticsTestOptionsListSchedulingOutputReference
	SchedulingInput() *SyntheticsTestOptionsListScheduling
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
	PutCi(value *SyntheticsTestOptionsListCi)
	PutMonitorOptions(value *SyntheticsTestOptionsListMonitorOptions)
	PutRetry(value *SyntheticsTestOptionsListRetry)
	PutRumSettings(value *SyntheticsTestOptionsListRumSettings)
	PutScheduling(value *SyntheticsTestOptionsListScheduling)
	ResetAcceptSelfSigned()
	ResetAllowInsecure()
	ResetBlockedRequestPatterns()
	ResetCheckCertificateRevocation()
	ResetCi()
	ResetDisableAiaIntermediateFetching()
	ResetDisableCors()
	ResetDisableCsp()
	ResetFollowRedirects()
	ResetHttpVersion()
	ResetIgnoreServerCertificateError()
	ResetInitialNavigationTimeout()
	ResetMinFailureDuration()
	ResetMinLocationFailed()
	ResetMonitorName()
	ResetMonitorOptions()
	ResetMonitorPriority()
	ResetNoScreenshot()
	ResetRestrictedRoles()
	ResetRetry()
	ResetRumSettings()
	ResetScheduling()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SyntheticsTestOptionsListStructOutputReference
type jsiiProxy_SyntheticsTestOptionsListStructOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) AcceptSelfSigned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptSelfSigned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) AcceptSelfSignedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptSelfSignedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) AllowInsecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) AllowInsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) BlockedRequestPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedRequestPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) BlockedRequestPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedRequestPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) CheckCertificateRevocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkCertificateRevocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) CheckCertificateRevocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkCertificateRevocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) Ci() SyntheticsTestOptionsListCiOutputReference {
	var returns SyntheticsTestOptionsListCiOutputReference
	_jsii_.Get(
		j,
		"ci",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) CiInput() *SyntheticsTestOptionsListCi {
	var returns *SyntheticsTestOptionsListCi
	_jsii_.Get(
		j,
		"ciInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) DisableAiaIntermediateFetching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAiaIntermediateFetching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) DisableAiaIntermediateFetchingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableAiaIntermediateFetchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) DisableCors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) DisableCorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) DisableCsp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCsp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) DisableCspInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCspInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) FollowRedirects() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"followRedirects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) FollowRedirectsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"followRedirectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) HttpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) IgnoreServerCertificateError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreServerCertificateError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) IgnoreServerCertificateErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreServerCertificateErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) InitialNavigationTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNavigationTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) InitialNavigationTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNavigationTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) InternalValue() *SyntheticsTestOptionsListStruct {
	var returns *SyntheticsTestOptionsListStruct
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) MinFailureDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minFailureDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) MinFailureDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minFailureDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) MinLocationFailed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLocationFailed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) MinLocationFailedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLocationFailedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) MonitorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) MonitorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) MonitorOptions() SyntheticsTestOptionsListMonitorOptionsOutputReference {
	var returns SyntheticsTestOptionsListMonitorOptionsOutputReference
	_jsii_.Get(
		j,
		"monitorOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) MonitorOptionsInput() *SyntheticsTestOptionsListMonitorOptions {
	var returns *SyntheticsTestOptionsListMonitorOptions
	_jsii_.Get(
		j,
		"monitorOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) MonitorPriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitorPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) MonitorPriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitorPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) NoScreenshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noScreenshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) NoScreenshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noScreenshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) RestrictedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) RestrictedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) Retry() SyntheticsTestOptionsListRetryOutputReference {
	var returns SyntheticsTestOptionsListRetryOutputReference
	_jsii_.Get(
		j,
		"retry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) RetryInput() *SyntheticsTestOptionsListRetry {
	var returns *SyntheticsTestOptionsListRetry
	_jsii_.Get(
		j,
		"retryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) RumSettings() SyntheticsTestOptionsListRumSettingsOutputReference {
	var returns SyntheticsTestOptionsListRumSettingsOutputReference
	_jsii_.Get(
		j,
		"rumSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) RumSettingsInput() *SyntheticsTestOptionsListRumSettings {
	var returns *SyntheticsTestOptionsListRumSettings
	_jsii_.Get(
		j,
		"rumSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) Scheduling() SyntheticsTestOptionsListSchedulingOutputReference {
	var returns SyntheticsTestOptionsListSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) SchedulingInput() *SyntheticsTestOptionsListScheduling {
	var returns *SyntheticsTestOptionsListScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) TickEvery() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tickEvery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) TickEveryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tickEveryInput",
		&returns,
	)
	return returns
}


func NewSyntheticsTestOptionsListStructOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SyntheticsTestOptionsListStructOutputReference {
	_init_.Initialize()

	if err := validateNewSyntheticsTestOptionsListStructOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTestOptionsListStructOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestOptionsListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSyntheticsTestOptionsListStructOutputReference_Override(s SyntheticsTestOptionsListStructOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestOptionsListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetAcceptSelfSigned(val interface{}) {
	if err := j.validateSetAcceptSelfSignedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptSelfSigned",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetAllowInsecure(val interface{}) {
	if err := j.validateSetAllowInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInsecure",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetBlockedRequestPatterns(val *[]*string) {
	if err := j.validateSetBlockedRequestPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockedRequestPatterns",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetCheckCertificateRevocation(val interface{}) {
	if err := j.validateSetCheckCertificateRevocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkCertificateRevocation",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetDisableAiaIntermediateFetching(val interface{}) {
	if err := j.validateSetDisableAiaIntermediateFetchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAiaIntermediateFetching",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetDisableCors(val interface{}) {
	if err := j.validateSetDisableCorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableCors",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetDisableCsp(val interface{}) {
	if err := j.validateSetDisableCspParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableCsp",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetFollowRedirects(val interface{}) {
	if err := j.validateSetFollowRedirectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"followRedirects",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetHttpVersion(val *string) {
	if err := j.validateSetHttpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpVersion",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetIgnoreServerCertificateError(val interface{}) {
	if err := j.validateSetIgnoreServerCertificateErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreServerCertificateError",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetInitialNavigationTimeout(val *float64) {
	if err := j.validateSetInitialNavigationTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialNavigationTimeout",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetInternalValue(val *SyntheticsTestOptionsListStruct) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetMinFailureDuration(val *float64) {
	if err := j.validateSetMinFailureDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minFailureDuration",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetMinLocationFailed(val *float64) {
	if err := j.validateSetMinLocationFailedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLocationFailed",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetMonitorName(val *string) {
	if err := j.validateSetMonitorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitorName",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetMonitorPriority(val *float64) {
	if err := j.validateSetMonitorPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitorPriority",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetNoScreenshot(val interface{}) {
	if err := j.validateSetNoScreenshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noScreenshot",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetRestrictedRoles(val *[]*string) {
	if err := j.validateSetRestrictedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictedRoles",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestOptionsListStructOutputReference)SetTickEvery(val *float64) {
	if err := j.validateSetTickEveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tickEvery",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) PutCi(value *SyntheticsTestOptionsListCi) {
	if err := s.validatePutCiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCi",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) PutMonitorOptions(value *SyntheticsTestOptionsListMonitorOptions) {
	if err := s.validatePutMonitorOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMonitorOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) PutRetry(value *SyntheticsTestOptionsListRetry) {
	if err := s.validatePutRetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRetry",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) PutRumSettings(value *SyntheticsTestOptionsListRumSettings) {
	if err := s.validatePutRumSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRumSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) PutScheduling(value *SyntheticsTestOptionsListScheduling) {
	if err := s.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putScheduling",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetAcceptSelfSigned() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceptSelfSigned",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetAllowInsecure() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowInsecure",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetBlockedRequestPatterns() {
	_jsii_.InvokeVoid(
		s,
		"resetBlockedRequestPatterns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetCheckCertificateRevocation() {
	_jsii_.InvokeVoid(
		s,
		"resetCheckCertificateRevocation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetCi() {
	_jsii_.InvokeVoid(
		s,
		"resetCi",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetDisableAiaIntermediateFetching() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableAiaIntermediateFetching",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetDisableCors() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableCors",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetDisableCsp() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableCsp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetFollowRedirects() {
	_jsii_.InvokeVoid(
		s,
		"resetFollowRedirects",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetHttpVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetIgnoreServerCertificateError() {
	_jsii_.InvokeVoid(
		s,
		"resetIgnoreServerCertificateError",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetInitialNavigationTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetInitialNavigationTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetMinFailureDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetMinFailureDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetMinLocationFailed() {
	_jsii_.InvokeVoid(
		s,
		"resetMinLocationFailed",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetMonitorName() {
	_jsii_.InvokeVoid(
		s,
		"resetMonitorName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetMonitorOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetMonitorOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetMonitorPriority() {
	_jsii_.InvokeVoid(
		s,
		"resetMonitorPriority",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetNoScreenshot() {
	_jsii_.InvokeVoid(
		s,
		"resetNoScreenshot",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetRestrictedRoles() {
	_jsii_.InvokeVoid(
		s,
		"resetRestrictedRoles",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetRetry() {
	_jsii_.InvokeVoid(
		s,
		"resetRetry",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetRumSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetRumSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ResetScheduling() {
	_jsii_.InvokeVoid(
		s,
		"resetScheduling",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SyntheticsTestOptionsListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

