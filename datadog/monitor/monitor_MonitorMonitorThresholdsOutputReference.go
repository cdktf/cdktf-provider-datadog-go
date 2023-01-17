package monitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v5/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v5/monitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorMonitorThresholdsOutputReference interface {
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
	Critical() *string
	SetCritical(val *string)
	CriticalInput() *string
	CriticalRecovery() *string
	SetCriticalRecovery(val *string)
	CriticalRecoveryInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorMonitorThresholds
	SetInternalValue(val *MonitorMonitorThresholds)
	Ok() *string
	SetOk(val *string)
	OkInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unknown() *string
	SetUnknown(val *string)
	UnknownInput() *string
	Warning() *string
	SetWarning(val *string)
	WarningInput() *string
	WarningRecovery() *string
	SetWarningRecovery(val *string)
	WarningRecoveryInput() *string
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
	ResetCritical()
	ResetCriticalRecovery()
	ResetOk()
	ResetUnknown()
	ResetWarning()
	ResetWarningRecovery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorMonitorThresholdsOutputReference
type jsiiProxy_MonitorMonitorThresholdsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) Critical() *string {
	var returns *string
	_jsii_.Get(
		j,
		"critical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) CriticalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"criticalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) CriticalRecovery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"criticalRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) CriticalRecoveryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"criticalRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) InternalValue() *MonitorMonitorThresholds {
	var returns *MonitorMonitorThresholds
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) Ok() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ok",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) OkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"okInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) Unknown() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unknown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) UnknownInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unknownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) Warning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) WarningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) WarningRecovery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference) WarningRecoveryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningRecoveryInput",
		&returns,
	)
	return returns
}


func NewMonitorMonitorThresholdsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorMonitorThresholdsOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorMonitorThresholdsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorMonitorThresholdsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.monitor.MonitorMonitorThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorMonitorThresholdsOutputReference_Override(m MonitorMonitorThresholdsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.monitor.MonitorMonitorThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference)SetCritical(val *string) {
	if err := j.validateSetCriticalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"critical",
		val,
	)
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference)SetCriticalRecovery(val *string) {
	if err := j.validateSetCriticalRecoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"criticalRecovery",
		val,
	)
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference)SetInternalValue(val *MonitorMonitorThresholds) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference)SetOk(val *string) {
	if err := j.validateSetOkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ok",
		val,
	)
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference)SetUnknown(val *string) {
	if err := j.validateSetUnknownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unknown",
		val,
	)
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference)SetWarning(val *string) {
	if err := j.validateSetWarningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warning",
		val,
	)
}

func (j *jsiiProxy_MonitorMonitorThresholdsOutputReference)SetWarningRecovery(val *string) {
	if err := j.validateSetWarningRecoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warningRecovery",
		val,
	)
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) ResetCritical() {
	_jsii_.InvokeVoid(
		m,
		"resetCritical",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) ResetCriticalRecovery() {
	_jsii_.InvokeVoid(
		m,
		"resetCriticalRecovery",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) ResetOk() {
	_jsii_.InvokeVoid(
		m,
		"resetOk",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) ResetUnknown() {
	_jsii_.InvokeVoid(
		m,
		"resetUnknown",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) ResetWarning() {
	_jsii_.InvokeVoid(
		m,
		"resetWarning",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) ResetWarningRecovery() {
	_jsii_.InvokeVoid(
		m,
		"resetWarningRecovery",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMonitorThresholdsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

