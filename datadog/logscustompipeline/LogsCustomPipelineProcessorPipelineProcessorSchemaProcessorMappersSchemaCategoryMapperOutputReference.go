// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/logscustompipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference interface {
	cdktf.ComplexObject
	Categories() LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperCategoriesList
	CategoriesInput() interface{}
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
	Fallback() LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference
	FallbackInput() *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallback
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Targets() LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperTargetsOutputReference
	TargetsInput() *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperTargets
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
	PutCategories(value interface{})
	PutFallback(value *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallback)
	PutTargets(value *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperTargets)
	ResetFallback()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference
type jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) Categories() LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperCategoriesList {
	var returns LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperCategoriesList
	_jsii_.Get(
		j,
		"categories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) CategoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"categoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) Fallback() LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallbackOutputReference
	_jsii_.Get(
		j,
		"fallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) FallbackInput() *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallback {
	var returns *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallback
	_jsii_.Get(
		j,
		"fallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) Targets() LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperTargetsOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperTargetsOutputReference
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) TargetsInput() *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperTargets {
	var returns *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperTargets
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference {
	_init_.Initialize()

	if err := validateNewLogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference_Override(l LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) PutCategories(value interface{}) {
	if err := l.validatePutCategoriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCategories",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) PutFallback(value *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallback) {
	if err := l.validatePutFallbackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFallback",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) PutTargets(value *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperTargets) {
	if err := l.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTargets",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) ResetFallback() {
	_jsii_.InvokeVoid(
		l,
		"resetFallback",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

