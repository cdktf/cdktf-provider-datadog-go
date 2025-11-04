// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/observabilitypipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityPipelineConfigDestinationsSumoLogicOutputReference interface {
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
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	// Experimental.
	Fqn() *string
	HeaderCustomFields() ObservabilityPipelineConfigDestinationsSumoLogicHeaderCustomFieldsList
	HeaderCustomFieldsInput() interface{}
	HeaderHostName() *string
	SetHeaderHostName(val *string)
	HeaderHostNameInput() *string
	HeaderSourceCategory() *string
	SetHeaderSourceCategory(val *string)
	HeaderSourceCategoryInput() *string
	HeaderSourceName() *string
	SetHeaderSourceName(val *string)
	HeaderSourceNameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Inputs() *[]*string
	SetInputs(val *[]*string)
	InputsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutHeaderCustomFields(value interface{})
	ResetEncoding()
	ResetHeaderCustomFields()
	ResetHeaderHostName()
	ResetHeaderSourceCategory()
	ResetHeaderSourceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityPipelineConfigDestinationsSumoLogicOutputReference
type jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) HeaderCustomFields() ObservabilityPipelineConfigDestinationsSumoLogicHeaderCustomFieldsList {
	var returns ObservabilityPipelineConfigDestinationsSumoLogicHeaderCustomFieldsList
	_jsii_.Get(
		j,
		"headerCustomFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) HeaderCustomFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerCustomFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) HeaderHostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerHostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) HeaderHostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerHostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) HeaderSourceCategory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerSourceCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) HeaderSourceCategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerSourceCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) HeaderSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) HeaderSourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerSourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) Inputs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) InputsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigDestinationsSumoLogicOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityPipelineConfigDestinationsSumoLogicOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigDestinationsSumoLogicOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationsSumoLogicOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigDestinationsSumoLogicOutputReference_Override(o ObservabilityPipelineConfigDestinationsSumoLogicOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationsSumoLogicOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference)SetHeaderHostName(val *string) {
	if err := j.validateSetHeaderHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerHostName",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference)SetHeaderSourceCategory(val *string) {
	if err := j.validateSetHeaderSourceCategoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerSourceCategory",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference)SetHeaderSourceName(val *string) {
	if err := j.validateSetHeaderSourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerSourceName",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference)SetInputs(val *[]*string) {
	if err := j.validateSetInputsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputs",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) PutHeaderCustomFields(value interface{}) {
	if err := o.validatePutHeaderCustomFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHeaderCustomFields",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		o,
		"resetEncoding",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) ResetHeaderCustomFields() {
	_jsii_.InvokeVoid(
		o,
		"resetHeaderCustomFields",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) ResetHeaderHostName() {
	_jsii_.InvokeVoid(
		o,
		"resetHeaderHostName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) ResetHeaderSourceCategory() {
	_jsii_.InvokeVoid(
		o,
		"resetHeaderSourceCategory",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) ResetHeaderSourceName() {
	_jsii_.InvokeVoid(
		o,
		"resetHeaderSourceName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsSumoLogicOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

