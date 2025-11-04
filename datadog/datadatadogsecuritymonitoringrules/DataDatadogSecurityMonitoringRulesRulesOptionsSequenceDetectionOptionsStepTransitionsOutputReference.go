// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogsecuritymonitoringrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/datadatadogsecuritymonitoringrules/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference interface {
	cdktf.ComplexObject
	Child() *string
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
	EvaluationWindow() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitions
	SetInternalValue(val *DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitions)
	Parent() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference
type jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) Child() *string {
	var returns *string
	_jsii_.Get(
		j,
		"child",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) EvaluationWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) InternalValue() *DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitions {
	var returns *DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) Parent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dataDatadogSecurityMonitoringRules.DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference_Override(d DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dataDatadogSecurityMonitoringRules.DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference)SetInternalValue(val *DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRulesRulesOptionsSequenceDetectionOptionsStepTransitionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

