// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/securitymonitoringrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference interface {
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
	DefaultNotifications() *[]*string
	SetDefaultNotifications(val *[]*string)
	DefaultNotificationsInput() *[]*string
	DefaultStatus() *string
	SetDefaultStatus(val *string)
	DefaultStatusInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SecurityMonitoringRuleOptionsThirdPartyRuleOptions
	SetInternalValue(val *SecurityMonitoringRuleOptionsThirdPartyRuleOptions)
	RootQuery() SecurityMonitoringRuleOptionsThirdPartyRuleOptionsRootQueryList
	RootQueryInput() interface{}
	SignalTitleTemplate() *string
	SetSignalTitleTemplate(val *string)
	SignalTitleTemplateInput() *string
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
	PutRootQuery(value interface{})
	ResetDefaultNotifications()
	ResetSignalTitleTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference
type jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) DefaultNotifications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultNotifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) DefaultNotificationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultNotificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) DefaultStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) DefaultStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) InternalValue() *SecurityMonitoringRuleOptionsThirdPartyRuleOptions {
	var returns *SecurityMonitoringRuleOptionsThirdPartyRuleOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) RootQuery() SecurityMonitoringRuleOptionsThirdPartyRuleOptionsRootQueryList {
	var returns SecurityMonitoringRuleOptionsThirdPartyRuleOptionsRootQueryList
	_jsii_.Get(
		j,
		"rootQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) RootQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) SignalTitleTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signalTitleTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) SignalTitleTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signalTitleTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.securityMonitoringRule.SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference_Override(s SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.securityMonitoringRule.SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference)SetDefaultNotifications(val *[]*string) {
	if err := j.validateSetDefaultNotificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNotifications",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference)SetDefaultStatus(val *string) {
	if err := j.validateSetDefaultStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultStatus",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference)SetInternalValue(val *SecurityMonitoringRuleOptionsThirdPartyRuleOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference)SetSignalTitleTemplate(val *string) {
	if err := j.validateSetSignalTitleTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signalTitleTemplate",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) PutRootQuery(value interface{}) {
	if err := s.validatePutRootQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRootQuery",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) ResetDefaultNotifications() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultNotifications",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) ResetSignalTitleTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetSignalTitleTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsThirdPartyRuleOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

