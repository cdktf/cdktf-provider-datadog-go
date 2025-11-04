// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/observabilitypipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference interface {
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
	Hash() ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchHashOutputReference
	HashInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PartialRedact() ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchPartialRedactOutputReference
	PartialRedactInput() interface{}
	Redact() ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchRedactOutputReference
	RedactInput() interface{}
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
	PutHash(value *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchHash)
	PutPartialRedact(value *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchPartialRedact)
	PutRedact(value *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchRedact)
	ResetPartialRedact()
	ResetRedact()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference
type jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) Hash() ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchHashOutputReference {
	var returns ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchHashOutputReference
	_jsii_.Get(
		j,
		"hash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) HashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) PartialRedact() ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchPartialRedactOutputReference {
	var returns ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchPartialRedactOutputReference
	_jsii_.Get(
		j,
		"partialRedact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) PartialRedactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partialRedactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) Redact() ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchRedactOutputReference {
	var returns ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchRedactOutputReference
	_jsii_.Get(
		j,
		"redact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) RedactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference_Override(o ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) PutHash(value *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchHash) {
	if err := o.validatePutHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHash",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) PutPartialRedact(value *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchPartialRedact) {
	if err := o.validatePutPartialRedactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPartialRedact",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) PutRedact(value *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchRedact) {
	if err := o.validatePutRedactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRedact",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) ResetPartialRedact() {
	_jsii_.InvokeVoid(
		o,
		"resetPartialRedact",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) ResetRedact() {
	_jsii_.InvokeVoid(
		o,
		"resetRedact",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

