// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/logscustompipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference interface {
	cdktf.ComplexObject
	ClassName() *string
	SetClassName(val *string)
	ClassNameInput() *string
	ClassUid() *float64
	SetClassUid(val *float64)
	ClassUidInput() *float64
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
	Extensions() *[]*string
	SetExtensions(val *[]*string)
	ExtensionsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchema
	SetInternalValue(val *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchema)
	Profiles() *[]*string
	SetProfiles(val *[]*string)
	ProfilesInput() *[]*string
	SchemaType() *string
	SetSchemaType(val *string)
	SchemaTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetExtensions()
	ResetProfiles()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference
type jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) ClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) ClassUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"classUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) ClassUidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"classUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) Extensions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) ExtensionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) InternalValue() *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchema {
	var returns *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchema
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) Profiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"profiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) ProfilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"profilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) SchemaType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) SchemaTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewLogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference {
	_init_.Initialize()

	if err := validateNewLogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference_Override(l LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference)SetClassName(val *string) {
	if err := j.validateSetClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"className",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference)SetClassUid(val *float64) {
	if err := j.validateSetClassUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classUid",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference)SetExtensions(val *[]*string) {
	if err := j.validateSetExtensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensions",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference)SetInternalValue(val *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchema) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference)SetProfiles(val *[]*string) {
	if err := j.validateSetProfilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profiles",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference)SetSchemaType(val *string) {
	if err := j.validateSetSchemaTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaType",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) ResetExtensions() {
	_jsii_.InvokeVoid(
		l,
		"resetExtensions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) ResetProfiles() {
	_jsii_.InvokeVoid(
		l,
		"resetProfiles",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchemaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

