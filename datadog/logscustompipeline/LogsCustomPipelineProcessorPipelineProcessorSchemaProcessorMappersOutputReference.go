// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/logscustompipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference interface {
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
	SchemaCategoryMapper() LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperList
	SchemaCategoryMapperInput() interface{}
	SchemaRemapper() LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaRemapperList
	SchemaRemapperInput() interface{}
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
	PutSchemaCategoryMapper(value interface{})
	PutSchemaRemapper(value interface{})
	ResetSchemaCategoryMapper()
	ResetSchemaRemapper()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference
type jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) SchemaCategoryMapper() LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperList {
	var returns LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperList
	_jsii_.Get(
		j,
		"schemaCategoryMapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) SchemaCategoryMapperInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaCategoryMapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) SchemaRemapper() LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaRemapperList {
	var returns LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaRemapperList
	_jsii_.Get(
		j,
		"schemaRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) SchemaRemapperInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference {
	_init_.Initialize()

	if err := validateNewLogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference_Override(l LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) PutSchemaCategoryMapper(value interface{}) {
	if err := l.validatePutSchemaCategoryMapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSchemaCategoryMapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) PutSchemaRemapper(value interface{}) {
	if err := l.validatePutSchemaRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSchemaRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) ResetSchemaCategoryMapper() {
	_jsii_.InvokeVoid(
		l,
		"resetSchemaCategoryMapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) ResetSchemaRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetSchemaRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

