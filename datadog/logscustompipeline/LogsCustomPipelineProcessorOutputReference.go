// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/logscustompipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsCustomPipelineProcessorOutputReference interface {
	cdktf.ComplexObject
	ArithmeticProcessor() LogsCustomPipelineProcessorArithmeticProcessorOutputReference
	ArithmeticProcessorInput() *LogsCustomPipelineProcessorArithmeticProcessor
	AttributeRemapper() LogsCustomPipelineProcessorAttributeRemapperOutputReference
	AttributeRemapperInput() *LogsCustomPipelineProcessorAttributeRemapper
	CategoryProcessor() LogsCustomPipelineProcessorCategoryProcessorOutputReference
	CategoryProcessorInput() *LogsCustomPipelineProcessorCategoryProcessor
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
	DateRemapper() LogsCustomPipelineProcessorDateRemapperOutputReference
	DateRemapperInput() *LogsCustomPipelineProcessorDateRemapper
	// Experimental.
	Fqn() *string
	GeoIpParser() LogsCustomPipelineProcessorGeoIpParserOutputReference
	GeoIpParserInput() *LogsCustomPipelineProcessorGeoIpParser
	GrokParser() LogsCustomPipelineProcessorGrokParserOutputReference
	GrokParserInput() *LogsCustomPipelineProcessorGrokParser
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LookupProcessor() LogsCustomPipelineProcessorLookupProcessorOutputReference
	LookupProcessorInput() *LogsCustomPipelineProcessorLookupProcessor
	MessageRemapper() LogsCustomPipelineProcessorMessageRemapperOutputReference
	MessageRemapperInput() *LogsCustomPipelineProcessorMessageRemapper
	Pipeline() LogsCustomPipelineProcessorPipelineOutputReference
	PipelineInput() *LogsCustomPipelineProcessorPipeline
	ReferenceTableLookupProcessor() LogsCustomPipelineProcessorReferenceTableLookupProcessorOutputReference
	ReferenceTableLookupProcessorInput() *LogsCustomPipelineProcessorReferenceTableLookupProcessor
	ServiceRemapper() LogsCustomPipelineProcessorServiceRemapperOutputReference
	ServiceRemapperInput() *LogsCustomPipelineProcessorServiceRemapper
	SpanIdRemapper() LogsCustomPipelineProcessorSpanIdRemapperOutputReference
	SpanIdRemapperInput() *LogsCustomPipelineProcessorSpanIdRemapper
	StatusRemapper() LogsCustomPipelineProcessorStatusRemapperOutputReference
	StatusRemapperInput() *LogsCustomPipelineProcessorStatusRemapper
	StringBuilderProcessor() LogsCustomPipelineProcessorStringBuilderProcessorOutputReference
	StringBuilderProcessorInput() *LogsCustomPipelineProcessorStringBuilderProcessor
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TraceIdRemapper() LogsCustomPipelineProcessorTraceIdRemapperOutputReference
	TraceIdRemapperInput() *LogsCustomPipelineProcessorTraceIdRemapper
	UrlParser() LogsCustomPipelineProcessorUrlParserOutputReference
	UrlParserInput() *LogsCustomPipelineProcessorUrlParser
	UserAgentParser() LogsCustomPipelineProcessorUserAgentParserOutputReference
	UserAgentParserInput() *LogsCustomPipelineProcessorUserAgentParser
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
	PutArithmeticProcessor(value *LogsCustomPipelineProcessorArithmeticProcessor)
	PutAttributeRemapper(value *LogsCustomPipelineProcessorAttributeRemapper)
	PutCategoryProcessor(value *LogsCustomPipelineProcessorCategoryProcessor)
	PutDateRemapper(value *LogsCustomPipelineProcessorDateRemapper)
	PutGeoIpParser(value *LogsCustomPipelineProcessorGeoIpParser)
	PutGrokParser(value *LogsCustomPipelineProcessorGrokParser)
	PutLookupProcessor(value *LogsCustomPipelineProcessorLookupProcessor)
	PutMessageRemapper(value *LogsCustomPipelineProcessorMessageRemapper)
	PutPipeline(value *LogsCustomPipelineProcessorPipeline)
	PutReferenceTableLookupProcessor(value *LogsCustomPipelineProcessorReferenceTableLookupProcessor)
	PutServiceRemapper(value *LogsCustomPipelineProcessorServiceRemapper)
	PutSpanIdRemapper(value *LogsCustomPipelineProcessorSpanIdRemapper)
	PutStatusRemapper(value *LogsCustomPipelineProcessorStatusRemapper)
	PutStringBuilderProcessor(value *LogsCustomPipelineProcessorStringBuilderProcessor)
	PutTraceIdRemapper(value *LogsCustomPipelineProcessorTraceIdRemapper)
	PutUrlParser(value *LogsCustomPipelineProcessorUrlParser)
	PutUserAgentParser(value *LogsCustomPipelineProcessorUserAgentParser)
	ResetArithmeticProcessor()
	ResetAttributeRemapper()
	ResetCategoryProcessor()
	ResetDateRemapper()
	ResetGeoIpParser()
	ResetGrokParser()
	ResetLookupProcessor()
	ResetMessageRemapper()
	ResetPipeline()
	ResetReferenceTableLookupProcessor()
	ResetServiceRemapper()
	ResetSpanIdRemapper()
	ResetStatusRemapper()
	ResetStringBuilderProcessor()
	ResetTraceIdRemapper()
	ResetUrlParser()
	ResetUserAgentParser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogsCustomPipelineProcessorOutputReference
type jsiiProxy_LogsCustomPipelineProcessorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ArithmeticProcessor() LogsCustomPipelineProcessorArithmeticProcessorOutputReference {
	var returns LogsCustomPipelineProcessorArithmeticProcessorOutputReference
	_jsii_.Get(
		j,
		"arithmeticProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ArithmeticProcessorInput() *LogsCustomPipelineProcessorArithmeticProcessor {
	var returns *LogsCustomPipelineProcessorArithmeticProcessor
	_jsii_.Get(
		j,
		"arithmeticProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) AttributeRemapper() LogsCustomPipelineProcessorAttributeRemapperOutputReference {
	var returns LogsCustomPipelineProcessorAttributeRemapperOutputReference
	_jsii_.Get(
		j,
		"attributeRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) AttributeRemapperInput() *LogsCustomPipelineProcessorAttributeRemapper {
	var returns *LogsCustomPipelineProcessorAttributeRemapper
	_jsii_.Get(
		j,
		"attributeRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) CategoryProcessor() LogsCustomPipelineProcessorCategoryProcessorOutputReference {
	var returns LogsCustomPipelineProcessorCategoryProcessorOutputReference
	_jsii_.Get(
		j,
		"categoryProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) CategoryProcessorInput() *LogsCustomPipelineProcessorCategoryProcessor {
	var returns *LogsCustomPipelineProcessorCategoryProcessor
	_jsii_.Get(
		j,
		"categoryProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) DateRemapper() LogsCustomPipelineProcessorDateRemapperOutputReference {
	var returns LogsCustomPipelineProcessorDateRemapperOutputReference
	_jsii_.Get(
		j,
		"dateRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) DateRemapperInput() *LogsCustomPipelineProcessorDateRemapper {
	var returns *LogsCustomPipelineProcessorDateRemapper
	_jsii_.Get(
		j,
		"dateRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) GeoIpParser() LogsCustomPipelineProcessorGeoIpParserOutputReference {
	var returns LogsCustomPipelineProcessorGeoIpParserOutputReference
	_jsii_.Get(
		j,
		"geoIpParser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) GeoIpParserInput() *LogsCustomPipelineProcessorGeoIpParser {
	var returns *LogsCustomPipelineProcessorGeoIpParser
	_jsii_.Get(
		j,
		"geoIpParserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) GrokParser() LogsCustomPipelineProcessorGrokParserOutputReference {
	var returns LogsCustomPipelineProcessorGrokParserOutputReference
	_jsii_.Get(
		j,
		"grokParser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) GrokParserInput() *LogsCustomPipelineProcessorGrokParser {
	var returns *LogsCustomPipelineProcessorGrokParser
	_jsii_.Get(
		j,
		"grokParserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) LookupProcessor() LogsCustomPipelineProcessorLookupProcessorOutputReference {
	var returns LogsCustomPipelineProcessorLookupProcessorOutputReference
	_jsii_.Get(
		j,
		"lookupProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) LookupProcessorInput() *LogsCustomPipelineProcessorLookupProcessor {
	var returns *LogsCustomPipelineProcessorLookupProcessor
	_jsii_.Get(
		j,
		"lookupProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) MessageRemapper() LogsCustomPipelineProcessorMessageRemapperOutputReference {
	var returns LogsCustomPipelineProcessorMessageRemapperOutputReference
	_jsii_.Get(
		j,
		"messageRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) MessageRemapperInput() *LogsCustomPipelineProcessorMessageRemapper {
	var returns *LogsCustomPipelineProcessorMessageRemapper
	_jsii_.Get(
		j,
		"messageRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) Pipeline() LogsCustomPipelineProcessorPipelineOutputReference {
	var returns LogsCustomPipelineProcessorPipelineOutputReference
	_jsii_.Get(
		j,
		"pipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PipelineInput() *LogsCustomPipelineProcessorPipeline {
	var returns *LogsCustomPipelineProcessorPipeline
	_jsii_.Get(
		j,
		"pipelineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ReferenceTableLookupProcessor() LogsCustomPipelineProcessorReferenceTableLookupProcessorOutputReference {
	var returns LogsCustomPipelineProcessorReferenceTableLookupProcessorOutputReference
	_jsii_.Get(
		j,
		"referenceTableLookupProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ReferenceTableLookupProcessorInput() *LogsCustomPipelineProcessorReferenceTableLookupProcessor {
	var returns *LogsCustomPipelineProcessorReferenceTableLookupProcessor
	_jsii_.Get(
		j,
		"referenceTableLookupProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ServiceRemapper() LogsCustomPipelineProcessorServiceRemapperOutputReference {
	var returns LogsCustomPipelineProcessorServiceRemapperOutputReference
	_jsii_.Get(
		j,
		"serviceRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ServiceRemapperInput() *LogsCustomPipelineProcessorServiceRemapper {
	var returns *LogsCustomPipelineProcessorServiceRemapper
	_jsii_.Get(
		j,
		"serviceRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) SpanIdRemapper() LogsCustomPipelineProcessorSpanIdRemapperOutputReference {
	var returns LogsCustomPipelineProcessorSpanIdRemapperOutputReference
	_jsii_.Get(
		j,
		"spanIdRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) SpanIdRemapperInput() *LogsCustomPipelineProcessorSpanIdRemapper {
	var returns *LogsCustomPipelineProcessorSpanIdRemapper
	_jsii_.Get(
		j,
		"spanIdRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) StatusRemapper() LogsCustomPipelineProcessorStatusRemapperOutputReference {
	var returns LogsCustomPipelineProcessorStatusRemapperOutputReference
	_jsii_.Get(
		j,
		"statusRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) StatusRemapperInput() *LogsCustomPipelineProcessorStatusRemapper {
	var returns *LogsCustomPipelineProcessorStatusRemapper
	_jsii_.Get(
		j,
		"statusRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) StringBuilderProcessor() LogsCustomPipelineProcessorStringBuilderProcessorOutputReference {
	var returns LogsCustomPipelineProcessorStringBuilderProcessorOutputReference
	_jsii_.Get(
		j,
		"stringBuilderProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) StringBuilderProcessorInput() *LogsCustomPipelineProcessorStringBuilderProcessor {
	var returns *LogsCustomPipelineProcessorStringBuilderProcessor
	_jsii_.Get(
		j,
		"stringBuilderProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) TraceIdRemapper() LogsCustomPipelineProcessorTraceIdRemapperOutputReference {
	var returns LogsCustomPipelineProcessorTraceIdRemapperOutputReference
	_jsii_.Get(
		j,
		"traceIdRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) TraceIdRemapperInput() *LogsCustomPipelineProcessorTraceIdRemapper {
	var returns *LogsCustomPipelineProcessorTraceIdRemapper
	_jsii_.Get(
		j,
		"traceIdRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) UrlParser() LogsCustomPipelineProcessorUrlParserOutputReference {
	var returns LogsCustomPipelineProcessorUrlParserOutputReference
	_jsii_.Get(
		j,
		"urlParser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) UrlParserInput() *LogsCustomPipelineProcessorUrlParser {
	var returns *LogsCustomPipelineProcessorUrlParser
	_jsii_.Get(
		j,
		"urlParserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) UserAgentParser() LogsCustomPipelineProcessorUserAgentParserOutputReference {
	var returns LogsCustomPipelineProcessorUserAgentParserOutputReference
	_jsii_.Get(
		j,
		"userAgentParser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference) UserAgentParserInput() *LogsCustomPipelineProcessorUserAgentParser {
	var returns *LogsCustomPipelineProcessorUserAgentParser
	_jsii_.Get(
		j,
		"userAgentParserInput",
		&returns,
	)
	return returns
}


func NewLogsCustomPipelineProcessorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LogsCustomPipelineProcessorOutputReference {
	_init_.Initialize()

	if err := validateNewLogsCustomPipelineProcessorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogsCustomPipelineProcessorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLogsCustomPipelineProcessorOutputReference_Override(l LogsCustomPipelineProcessorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutArithmeticProcessor(value *LogsCustomPipelineProcessorArithmeticProcessor) {
	if err := l.validatePutArithmeticProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putArithmeticProcessor",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutAttributeRemapper(value *LogsCustomPipelineProcessorAttributeRemapper) {
	if err := l.validatePutAttributeRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAttributeRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutCategoryProcessor(value *LogsCustomPipelineProcessorCategoryProcessor) {
	if err := l.validatePutCategoryProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCategoryProcessor",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutDateRemapper(value *LogsCustomPipelineProcessorDateRemapper) {
	if err := l.validatePutDateRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDateRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutGeoIpParser(value *LogsCustomPipelineProcessorGeoIpParser) {
	if err := l.validatePutGeoIpParserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putGeoIpParser",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutGrokParser(value *LogsCustomPipelineProcessorGrokParser) {
	if err := l.validatePutGrokParserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putGrokParser",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutLookupProcessor(value *LogsCustomPipelineProcessorLookupProcessor) {
	if err := l.validatePutLookupProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLookupProcessor",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutMessageRemapper(value *LogsCustomPipelineProcessorMessageRemapper) {
	if err := l.validatePutMessageRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMessageRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutPipeline(value *LogsCustomPipelineProcessorPipeline) {
	if err := l.validatePutPipelineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPipeline",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutReferenceTableLookupProcessor(value *LogsCustomPipelineProcessorReferenceTableLookupProcessor) {
	if err := l.validatePutReferenceTableLookupProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putReferenceTableLookupProcessor",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutServiceRemapper(value *LogsCustomPipelineProcessorServiceRemapper) {
	if err := l.validatePutServiceRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putServiceRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutSpanIdRemapper(value *LogsCustomPipelineProcessorSpanIdRemapper) {
	if err := l.validatePutSpanIdRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSpanIdRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutStatusRemapper(value *LogsCustomPipelineProcessorStatusRemapper) {
	if err := l.validatePutStatusRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStatusRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutStringBuilderProcessor(value *LogsCustomPipelineProcessorStringBuilderProcessor) {
	if err := l.validatePutStringBuilderProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStringBuilderProcessor",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutTraceIdRemapper(value *LogsCustomPipelineProcessorTraceIdRemapper) {
	if err := l.validatePutTraceIdRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTraceIdRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutUrlParser(value *LogsCustomPipelineProcessorUrlParser) {
	if err := l.validatePutUrlParserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putUrlParser",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) PutUserAgentParser(value *LogsCustomPipelineProcessorUserAgentParser) {
	if err := l.validatePutUserAgentParserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putUserAgentParser",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetArithmeticProcessor() {
	_jsii_.InvokeVoid(
		l,
		"resetArithmeticProcessor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetAttributeRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetAttributeRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetCategoryProcessor() {
	_jsii_.InvokeVoid(
		l,
		"resetCategoryProcessor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetDateRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetDateRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetGeoIpParser() {
	_jsii_.InvokeVoid(
		l,
		"resetGeoIpParser",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetGrokParser() {
	_jsii_.InvokeVoid(
		l,
		"resetGrokParser",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetLookupProcessor() {
	_jsii_.InvokeVoid(
		l,
		"resetLookupProcessor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetMessageRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetMessageRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetPipeline() {
	_jsii_.InvokeVoid(
		l,
		"resetPipeline",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetReferenceTableLookupProcessor() {
	_jsii_.InvokeVoid(
		l,
		"resetReferenceTableLookupProcessor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetServiceRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetServiceRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetSpanIdRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetSpanIdRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetStatusRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetStatusRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetStringBuilderProcessor() {
	_jsii_.InvokeVoid(
		l,
		"resetStringBuilderProcessor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetTraceIdRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetTraceIdRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetUrlParser() {
	_jsii_.InvokeVoid(
		l,
		"resetUrlParser",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ResetUserAgentParser() {
	_jsii_.InvokeVoid(
		l,
		"resetUserAgentParser",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

