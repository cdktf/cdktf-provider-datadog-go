// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/logscustompipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference interface {
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
	InternalValue() *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallback
	SetInternalValue(val *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallback)
	Sources() *map[string]*string
	SetSources(val *map[string]*string)
	SourcesInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *map[string]*string
	SetValues(val *map[string]*string)
	ValuesInput() *map[string]*string
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
	ResetSources()
	ResetValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference
type jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) InternalValue() *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallback {
	var returns *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallback
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) Sources() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) SourcesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) Values() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) ValuesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewLogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference {
	_init_.Initialize()

	if err := validateNewLogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference_Override(l LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference)SetInternalValue(val *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallback) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference)SetSources(val *map[string]*string) {
	if err := j.validateSetSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sources",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference)SetValues(val *map[string]*string) {
	if err := j.validateSetValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) ResetSources() {
	_jsii_.InvokeVoid(
		l,
		"resetSources",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		l,
		"resetValues",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

