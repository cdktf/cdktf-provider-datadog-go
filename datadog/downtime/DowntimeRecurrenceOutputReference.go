package downtime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v7/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v7/downtime/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DowntimeRecurrenceOutputReference interface {
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
	InternalValue() *DowntimeRecurrence
	SetInternalValue(val *DowntimeRecurrence)
	Period() *float64
	SetPeriod(val *float64)
	PeriodInput() *float64
	Rrule() *string
	SetRrule(val *string)
	RruleInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UntilDate() *float64
	SetUntilDate(val *float64)
	UntilDateInput() *float64
	UntilOccurrences() *float64
	SetUntilOccurrences(val *float64)
	UntilOccurrencesInput() *float64
	WeekDays() *[]*string
	SetWeekDays(val *[]*string)
	WeekDaysInput() *[]*string
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
	ResetPeriod()
	ResetRrule()
	ResetUntilDate()
	ResetUntilOccurrences()
	ResetWeekDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DowntimeRecurrenceOutputReference
type jsiiProxy_DowntimeRecurrenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) InternalValue() *DowntimeRecurrence {
	var returns *DowntimeRecurrence
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) Rrule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rrule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) RruleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) UntilDate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"untilDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) UntilDateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"untilDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) UntilOccurrences() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"untilOccurrences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) UntilOccurrencesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"untilOccurrencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) WeekDays() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference) WeekDaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekDaysInput",
		&returns,
	)
	return returns
}


func NewDowntimeRecurrenceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DowntimeRecurrenceOutputReference {
	_init_.Initialize()

	if err := validateNewDowntimeRecurrenceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DowntimeRecurrenceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.downtime.DowntimeRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDowntimeRecurrenceOutputReference_Override(d DowntimeRecurrenceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.downtime.DowntimeRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference)SetInternalValue(val *DowntimeRecurrence) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference)SetPeriod(val *float64) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference)SetRrule(val *string) {
	if err := j.validateSetRruleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rrule",
		val,
	)
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference)SetUntilDate(val *float64) {
	if err := j.validateSetUntilDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"untilDate",
		val,
	)
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference)SetUntilOccurrences(val *float64) {
	if err := j.validateSetUntilOccurrencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"untilOccurrences",
		val,
	)
}

func (j *jsiiProxy_DowntimeRecurrenceOutputReference)SetWeekDays(val *[]*string) {
	if err := j.validateSetWeekDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekDays",
		val,
	)
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) ResetPeriod() {
	_jsii_.InvokeVoid(
		d,
		"resetPeriod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) ResetRrule() {
	_jsii_.InvokeVoid(
		d,
		"resetRrule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) ResetUntilDate() {
	_jsii_.InvokeVoid(
		d,
		"resetUntilDate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) ResetUntilOccurrences() {
	_jsii_.InvokeVoid(
		d,
		"resetUntilOccurrences",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) ResetWeekDays() {
	_jsii_.InvokeVoid(
		d,
		"resetWeekDays",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeRecurrenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

