package securitymonitoringrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v7/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v7/securitymonitoringrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityMonitoringRuleOptionsOutputReference interface {
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
	DecreaseCriticalityBasedOnEnv() interface{}
	SetDecreaseCriticalityBasedOnEnv(val interface{})
	DecreaseCriticalityBasedOnEnvInput() interface{}
	DetectionMethod() *string
	SetDetectionMethod(val *string)
	DetectionMethodInput() *string
	EvaluationWindow() *float64
	SetEvaluationWindow(val *float64)
	EvaluationWindowInput() *float64
	// Experimental.
	Fqn() *string
	ImpossibleTravelOptions() SecurityMonitoringRuleOptionsImpossibleTravelOptionsOutputReference
	ImpossibleTravelOptionsInput() *SecurityMonitoringRuleOptionsImpossibleTravelOptions
	InternalValue() *SecurityMonitoringRuleOptions
	SetInternalValue(val *SecurityMonitoringRuleOptions)
	KeepAlive() *float64
	SetKeepAlive(val *float64)
	KeepAliveInput() *float64
	MaxSignalDuration() *float64
	SetMaxSignalDuration(val *float64)
	MaxSignalDurationInput() *float64
	NewValueOptions() SecurityMonitoringRuleOptionsNewValueOptionsOutputReference
	NewValueOptionsInput() *SecurityMonitoringRuleOptionsNewValueOptions
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
	PutImpossibleTravelOptions(value *SecurityMonitoringRuleOptionsImpossibleTravelOptions)
	PutNewValueOptions(value *SecurityMonitoringRuleOptionsNewValueOptions)
	ResetDecreaseCriticalityBasedOnEnv()
	ResetDetectionMethod()
	ResetEvaluationWindow()
	ResetImpossibleTravelOptions()
	ResetNewValueOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityMonitoringRuleOptionsOutputReference
type jsiiProxy_SecurityMonitoringRuleOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) DecreaseCriticalityBasedOnEnv() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decreaseCriticalityBasedOnEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) DecreaseCriticalityBasedOnEnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decreaseCriticalityBasedOnEnvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) DetectionMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectionMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) DetectionMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectionMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) EvaluationWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) EvaluationWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) ImpossibleTravelOptions() SecurityMonitoringRuleOptionsImpossibleTravelOptionsOutputReference {
	var returns SecurityMonitoringRuleOptionsImpossibleTravelOptionsOutputReference
	_jsii_.Get(
		j,
		"impossibleTravelOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) ImpossibleTravelOptionsInput() *SecurityMonitoringRuleOptionsImpossibleTravelOptions {
	var returns *SecurityMonitoringRuleOptionsImpossibleTravelOptions
	_jsii_.Get(
		j,
		"impossibleTravelOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) InternalValue() *SecurityMonitoringRuleOptions {
	var returns *SecurityMonitoringRuleOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) KeepAlive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) KeepAliveInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) MaxSignalDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSignalDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) MaxSignalDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSignalDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) NewValueOptions() SecurityMonitoringRuleOptionsNewValueOptionsOutputReference {
	var returns SecurityMonitoringRuleOptionsNewValueOptionsOutputReference
	_jsii_.Get(
		j,
		"newValueOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) NewValueOptionsInput() *SecurityMonitoringRuleOptionsNewValueOptions {
	var returns *SecurityMonitoringRuleOptionsNewValueOptions
	_jsii_.Get(
		j,
		"newValueOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecurityMonitoringRuleOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SecurityMonitoringRuleOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityMonitoringRuleOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityMonitoringRuleOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.securityMonitoringRule.SecurityMonitoringRuleOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityMonitoringRuleOptionsOutputReference_Override(s SecurityMonitoringRuleOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.securityMonitoringRule.SecurityMonitoringRuleOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference)SetDecreaseCriticalityBasedOnEnv(val interface{}) {
	if err := j.validateSetDecreaseCriticalityBasedOnEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decreaseCriticalityBasedOnEnv",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference)SetDetectionMethod(val *string) {
	if err := j.validateSetDetectionMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectionMethod",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference)SetEvaluationWindow(val *float64) {
	if err := j.validateSetEvaluationWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationWindow",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference)SetInternalValue(val *SecurityMonitoringRuleOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference)SetKeepAlive(val *float64) {
	if err := j.validateSetKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepAlive",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference)SetMaxSignalDuration(val *float64) {
	if err := j.validateSetMaxSignalDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSignalDuration",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) PutImpossibleTravelOptions(value *SecurityMonitoringRuleOptionsImpossibleTravelOptions) {
	if err := s.validatePutImpossibleTravelOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putImpossibleTravelOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) PutNewValueOptions(value *SecurityMonitoringRuleOptionsNewValueOptions) {
	if err := s.validatePutNewValueOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNewValueOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) ResetDecreaseCriticalityBasedOnEnv() {
	_jsii_.InvokeVoid(
		s,
		"resetDecreaseCriticalityBasedOnEnv",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) ResetDetectionMethod() {
	_jsii_.InvokeVoid(
		s,
		"resetDetectionMethod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) ResetEvaluationWindow() {
	_jsii_.InvokeVoid(
		s,
		"resetEvaluationWindow",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) ResetImpossibleTravelOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetImpossibleTravelOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) ResetNewValueOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetNewValueOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

