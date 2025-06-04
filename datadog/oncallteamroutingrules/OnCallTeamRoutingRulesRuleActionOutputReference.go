// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallteamroutingrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/oncallteamroutingrules/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OnCallTeamRoutingRulesRuleActionOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SendSlackMessage() OnCallTeamRoutingRulesRuleActionSendSlackMessageOutputReference
	SendSlackMessageInput() interface{}
	SendTeamsMessage() OnCallTeamRoutingRulesRuleActionSendTeamsMessageOutputReference
	SendTeamsMessageInput() interface{}
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
	PutSendSlackMessage(value *OnCallTeamRoutingRulesRuleActionSendSlackMessage)
	PutSendTeamsMessage(value *OnCallTeamRoutingRulesRuleActionSendTeamsMessage)
	ResetSendSlackMessage()
	ResetSendTeamsMessage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OnCallTeamRoutingRulesRuleActionOutputReference
type jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) SendSlackMessage() OnCallTeamRoutingRulesRuleActionSendSlackMessageOutputReference {
	var returns OnCallTeamRoutingRulesRuleActionSendSlackMessageOutputReference
	_jsii_.Get(
		j,
		"sendSlackMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) SendSlackMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendSlackMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) SendTeamsMessage() OnCallTeamRoutingRulesRuleActionSendTeamsMessageOutputReference {
	var returns OnCallTeamRoutingRulesRuleActionSendTeamsMessageOutputReference
	_jsii_.Get(
		j,
		"sendTeamsMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) SendTeamsMessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendTeamsMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOnCallTeamRoutingRulesRuleActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OnCallTeamRoutingRulesRuleActionOutputReference {
	_init_.Initialize()

	if err := validateNewOnCallTeamRoutingRulesRuleActionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.onCallTeamRoutingRules.OnCallTeamRoutingRulesRuleActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOnCallTeamRoutingRulesRuleActionOutputReference_Override(o OnCallTeamRoutingRulesRuleActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.onCallTeamRoutingRules.OnCallTeamRoutingRulesRuleActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) PutSendSlackMessage(value *OnCallTeamRoutingRulesRuleActionSendSlackMessage) {
	if err := o.validatePutSendSlackMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSendSlackMessage",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) PutSendTeamsMessage(value *OnCallTeamRoutingRulesRuleActionSendTeamsMessage) {
	if err := o.validatePutSendTeamsMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSendTeamsMessage",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) ResetSendSlackMessage() {
	_jsii_.InvokeVoid(
		o,
		"resetSendSlackMessage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) ResetSendTeamsMessage() {
	_jsii_.InvokeVoid(
		o,
		"resetSendTeamsMessage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallTeamRoutingRulesRuleActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

