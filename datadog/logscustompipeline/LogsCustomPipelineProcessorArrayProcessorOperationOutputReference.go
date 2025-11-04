// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/logscustompipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsCustomPipelineProcessorArrayProcessorOperationOutputReference interface {
	cdktf.ComplexObject
	Append() LogsCustomPipelineProcessorArrayProcessorOperationAppendOutputReference
	AppendInput() *LogsCustomPipelineProcessorArrayProcessorOperationAppend
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
	InternalValue() *LogsCustomPipelineProcessorArrayProcessorOperation
	SetInternalValue(val *LogsCustomPipelineProcessorArrayProcessorOperation)
	Length() LogsCustomPipelineProcessorArrayProcessorOperationLengthOutputReference
	LengthInput() *LogsCustomPipelineProcessorArrayProcessorOperationLength
	Select() LogsCustomPipelineProcessorArrayProcessorOperationSelectOutputReference
	SelectInput() *LogsCustomPipelineProcessorArrayProcessorOperationSelect
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
	PutAppend(value *LogsCustomPipelineProcessorArrayProcessorOperationAppend)
	PutLength(value *LogsCustomPipelineProcessorArrayProcessorOperationLength)
	PutSelect(value *LogsCustomPipelineProcessorArrayProcessorOperationSelect)
	ResetAppend()
	ResetLength()
	ResetSelect()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogsCustomPipelineProcessorArrayProcessorOperationOutputReference
type jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) Append() LogsCustomPipelineProcessorArrayProcessorOperationAppendOutputReference {
	var returns LogsCustomPipelineProcessorArrayProcessorOperationAppendOutputReference
	_jsii_.Get(
		j,
		"append",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) AppendInput() *LogsCustomPipelineProcessorArrayProcessorOperationAppend {
	var returns *LogsCustomPipelineProcessorArrayProcessorOperationAppend
	_jsii_.Get(
		j,
		"appendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) InternalValue() *LogsCustomPipelineProcessorArrayProcessorOperation {
	var returns *LogsCustomPipelineProcessorArrayProcessorOperation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) Length() LogsCustomPipelineProcessorArrayProcessorOperationLengthOutputReference {
	var returns LogsCustomPipelineProcessorArrayProcessorOperationLengthOutputReference
	_jsii_.Get(
		j,
		"length",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) LengthInput() *LogsCustomPipelineProcessorArrayProcessorOperationLength {
	var returns *LogsCustomPipelineProcessorArrayProcessorOperationLength
	_jsii_.Get(
		j,
		"lengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) Select() LogsCustomPipelineProcessorArrayProcessorOperationSelectOutputReference {
	var returns LogsCustomPipelineProcessorArrayProcessorOperationSelectOutputReference
	_jsii_.Get(
		j,
		"select",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) SelectInput() *LogsCustomPipelineProcessorArrayProcessorOperationSelect {
	var returns *LogsCustomPipelineProcessorArrayProcessorOperationSelect
	_jsii_.Get(
		j,
		"selectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLogsCustomPipelineProcessorArrayProcessorOperationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LogsCustomPipelineProcessorArrayProcessorOperationOutputReference {
	_init_.Initialize()

	if err := validateNewLogsCustomPipelineProcessorArrayProcessorOperationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorArrayProcessorOperationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLogsCustomPipelineProcessorArrayProcessorOperationOutputReference_Override(l LogsCustomPipelineProcessorArrayProcessorOperationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorArrayProcessorOperationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference)SetInternalValue(val *LogsCustomPipelineProcessorArrayProcessorOperation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) PutAppend(value *LogsCustomPipelineProcessorArrayProcessorOperationAppend) {
	if err := l.validatePutAppendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAppend",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) PutLength(value *LogsCustomPipelineProcessorArrayProcessorOperationLength) {
	if err := l.validatePutLengthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLength",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) PutSelect(value *LogsCustomPipelineProcessorArrayProcessorOperationSelect) {
	if err := l.validatePutSelectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSelect",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) ResetAppend() {
	_jsii_.InvokeVoid(
		l,
		"resetAppend",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) ResetLength() {
	_jsii_.InvokeVoid(
		l,
		"resetLength",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) ResetSelect() {
	_jsii_.InvokeVoid(
		l,
		"resetSelect",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

