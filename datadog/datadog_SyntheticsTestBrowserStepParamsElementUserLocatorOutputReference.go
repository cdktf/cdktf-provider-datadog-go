// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/jsii"

	"github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference interface {
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
	FailTestOnCannotLocate() interface{}
	SetFailTestOnCannotLocate(val interface{})
	FailTestOnCannotLocateInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *SyntheticsTestBrowserStepParamsElementUserLocator
	SetInternalValue(val *SyntheticsTestBrowserStepParamsElementUserLocator)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() SyntheticsTestBrowserStepParamsElementUserLocatorValueOutputReference
	ValueInput() *SyntheticsTestBrowserStepParamsElementUserLocatorValue
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
	PutValue(value *SyntheticsTestBrowserStepParamsElementUserLocatorValue)
	ResetFailTestOnCannotLocate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference
type jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) FailTestOnCannotLocate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTestOnCannotLocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) FailTestOnCannotLocateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failTestOnCannotLocateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) InternalValue() *SyntheticsTestBrowserStepParamsElementUserLocator {
	var returns *SyntheticsTestBrowserStepParamsElementUserLocator
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) Value() SyntheticsTestBrowserStepParamsElementUserLocatorValueOutputReference {
	var returns SyntheticsTestBrowserStepParamsElementUserLocatorValueOutputReference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) ValueInput() *SyntheticsTestBrowserStepParamsElementUserLocatorValue {
	var returns *SyntheticsTestBrowserStepParamsElementUserLocatorValue
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewSyntheticsTestBrowserStepParamsElementUserLocatorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSyntheticsTestBrowserStepParamsElementUserLocatorOutputReference_Override(s SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) SetFailTestOnCannotLocate(val interface{}) {
	_jsii_.Set(
		j,
		"failTestOnCannotLocate",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) SetInternalValue(val *SyntheticsTestBrowserStepParamsElementUserLocator) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) PutValue(value *SyntheticsTestBrowserStepParamsElementUserLocatorValue) {
	_jsii_.InvokeVoid(
		s,
		"putValue",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) ResetFailTestOnCannotLocate() {
	_jsii_.InvokeVoid(
		s,
		"resetFailTestOnCannotLocate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
