// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/oncallschedule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OnCallScheduleLayerOutputReference interface {
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
	EffectiveDate() *string
	SetEffectiveDate(val *string)
	EffectiveDateInput() *string
	EndDate() *string
	SetEndDate(val *string)
	EndDateInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Interval() OnCallScheduleLayerIntervalOutputReference
	IntervalInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Restriction() OnCallScheduleLayerRestrictionList
	RestrictionInput() interface{}
	RotationStart() *string
	SetRotationStart(val *string)
	RotationStartInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Users() *[]*string
	SetUsers(val *[]*string)
	UsersInput() *[]*string
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
	PutInterval(value *OnCallScheduleLayerInterval)
	PutRestriction(value interface{})
	ResetEndDate()
	ResetInterval()
	ResetRestriction()
	ResetRotationStart()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OnCallScheduleLayerOutputReference
type jsiiProxy_OnCallScheduleLayerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) EffectiveDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) EffectiveDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) EndDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) EndDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) Interval() OnCallScheduleLayerIntervalOutputReference {
	var returns OnCallScheduleLayerIntervalOutputReference
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) IntervalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) Restriction() OnCallScheduleLayerRestrictionList {
	var returns OnCallScheduleLayerRestrictionList
	_jsii_.Get(
		j,
		"restriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) RestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) RotationStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) RotationStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) Users() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference) UsersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}


func NewOnCallScheduleLayerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OnCallScheduleLayerOutputReference {
	_init_.Initialize()

	if err := validateNewOnCallScheduleLayerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OnCallScheduleLayerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.onCallSchedule.OnCallScheduleLayerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOnCallScheduleLayerOutputReference_Override(o OnCallScheduleLayerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.onCallSchedule.OnCallScheduleLayerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference)SetEffectiveDate(val *string) {
	if err := j.validateSetEffectiveDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectiveDate",
		val,
	)
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference)SetEndDate(val *string) {
	if err := j.validateSetEndDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endDate",
		val,
	)
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference)SetRotationStart(val *string) {
	if err := j.validateSetRotationStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationStart",
		val,
	)
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OnCallScheduleLayerOutputReference)SetUsers(val *[]*string) {
	if err := j.validateSetUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"users",
		val,
	)
}

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) PutInterval(value *OnCallScheduleLayerInterval) {
	if err := o.validatePutIntervalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putInterval",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) PutRestriction(value interface{}) {
	if err := o.validatePutRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRestriction",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) ResetEndDate() {
	_jsii_.InvokeVoid(
		o,
		"resetEndDate",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		o,
		"resetInterval",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) ResetRestriction() {
	_jsii_.InvokeVoid(
		o,
		"resetRestriction",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) ResetRotationStart() {
	_jsii_.InvokeVoid(
		o,
		"resetRotationStart",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OnCallScheduleLayerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

