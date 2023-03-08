package logscustompipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v5/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v5/logscustompipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsCustomPipelineProcessorPipelineProcessorOutputReference interface {
	cdktf.ComplexObject
	ArithmeticProcessor() LogsCustomPipelineProcessorPipelineProcessorArithmeticProcessorOutputReference
	ArithmeticProcessorInput() *LogsCustomPipelineProcessorPipelineProcessorArithmeticProcessor
	AttributeRemapper() LogsCustomPipelineProcessorPipelineProcessorAttributeRemapperOutputReference
	AttributeRemapperInput() *LogsCustomPipelineProcessorPipelineProcessorAttributeRemapper
	CategoryProcessor() LogsCustomPipelineProcessorPipelineProcessorCategoryProcessorOutputReference
	CategoryProcessorInput() *LogsCustomPipelineProcessorPipelineProcessorCategoryProcessor
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
	DateRemapper() LogsCustomPipelineProcessorPipelineProcessorDateRemapperOutputReference
	DateRemapperInput() *LogsCustomPipelineProcessorPipelineProcessorDateRemapper
	// Experimental.
	Fqn() *string
	GeoIpParser() LogsCustomPipelineProcessorPipelineProcessorGeoIpParserOutputReference
	GeoIpParserInput() *LogsCustomPipelineProcessorPipelineProcessorGeoIpParser
	GrokParser() LogsCustomPipelineProcessorPipelineProcessorGrokParserOutputReference
	GrokParserInput() *LogsCustomPipelineProcessorPipelineProcessorGrokParser
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LookupProcessor() LogsCustomPipelineProcessorPipelineProcessorLookupProcessorOutputReference
	LookupProcessorInput() *LogsCustomPipelineProcessorPipelineProcessorLookupProcessor
	MessageRemapper() LogsCustomPipelineProcessorPipelineProcessorMessageRemapperOutputReference
	MessageRemapperInput() *LogsCustomPipelineProcessorPipelineProcessorMessageRemapper
	ServiceRemapper() LogsCustomPipelineProcessorPipelineProcessorServiceRemapperOutputReference
	ServiceRemapperInput() *LogsCustomPipelineProcessorPipelineProcessorServiceRemapper
	StatusRemapper() LogsCustomPipelineProcessorPipelineProcessorStatusRemapperOutputReference
	StatusRemapperInput() *LogsCustomPipelineProcessorPipelineProcessorStatusRemapper
	StringBuilderProcessor() LogsCustomPipelineProcessorPipelineProcessorStringBuilderProcessorOutputReference
	StringBuilderProcessorInput() *LogsCustomPipelineProcessorPipelineProcessorStringBuilderProcessor
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TraceIdRemapper() LogsCustomPipelineProcessorPipelineProcessorTraceIdRemapperOutputReference
	TraceIdRemapperInput() *LogsCustomPipelineProcessorPipelineProcessorTraceIdRemapper
	UrlParser() LogsCustomPipelineProcessorPipelineProcessorUrlParserOutputReference
	UrlParserInput() *LogsCustomPipelineProcessorPipelineProcessorUrlParser
	UserAgentParser() LogsCustomPipelineProcessorPipelineProcessorUserAgentParserOutputReference
	UserAgentParserInput() *LogsCustomPipelineProcessorPipelineProcessorUserAgentParser
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
	PutArithmeticProcessor(value *LogsCustomPipelineProcessorPipelineProcessorArithmeticProcessor)
	PutAttributeRemapper(value *LogsCustomPipelineProcessorPipelineProcessorAttributeRemapper)
	PutCategoryProcessor(value *LogsCustomPipelineProcessorPipelineProcessorCategoryProcessor)
	PutDateRemapper(value *LogsCustomPipelineProcessorPipelineProcessorDateRemapper)
	PutGeoIpParser(value *LogsCustomPipelineProcessorPipelineProcessorGeoIpParser)
	PutGrokParser(value *LogsCustomPipelineProcessorPipelineProcessorGrokParser)
	PutLookupProcessor(value *LogsCustomPipelineProcessorPipelineProcessorLookupProcessor)
	PutMessageRemapper(value *LogsCustomPipelineProcessorPipelineProcessorMessageRemapper)
	PutServiceRemapper(value *LogsCustomPipelineProcessorPipelineProcessorServiceRemapper)
	PutStatusRemapper(value *LogsCustomPipelineProcessorPipelineProcessorStatusRemapper)
	PutStringBuilderProcessor(value *LogsCustomPipelineProcessorPipelineProcessorStringBuilderProcessor)
	PutTraceIdRemapper(value *LogsCustomPipelineProcessorPipelineProcessorTraceIdRemapper)
	PutUrlParser(value *LogsCustomPipelineProcessorPipelineProcessorUrlParser)
	PutUserAgentParser(value *LogsCustomPipelineProcessorPipelineProcessorUserAgentParser)
	ResetArithmeticProcessor()
	ResetAttributeRemapper()
	ResetCategoryProcessor()
	ResetDateRemapper()
	ResetGeoIpParser()
	ResetGrokParser()
	ResetLookupProcessor()
	ResetMessageRemapper()
	ResetServiceRemapper()
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

// The jsii proxy struct for LogsCustomPipelineProcessorPipelineProcessorOutputReference
type jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ArithmeticProcessor() LogsCustomPipelineProcessorPipelineProcessorArithmeticProcessorOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorArithmeticProcessorOutputReference
	_jsii_.Get(
		j,
		"arithmeticProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ArithmeticProcessorInput() *LogsCustomPipelineProcessorPipelineProcessorArithmeticProcessor {
	var returns *LogsCustomPipelineProcessorPipelineProcessorArithmeticProcessor
	_jsii_.Get(
		j,
		"arithmeticProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) AttributeRemapper() LogsCustomPipelineProcessorPipelineProcessorAttributeRemapperOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorAttributeRemapperOutputReference
	_jsii_.Get(
		j,
		"attributeRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) AttributeRemapperInput() *LogsCustomPipelineProcessorPipelineProcessorAttributeRemapper {
	var returns *LogsCustomPipelineProcessorPipelineProcessorAttributeRemapper
	_jsii_.Get(
		j,
		"attributeRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) CategoryProcessor() LogsCustomPipelineProcessorPipelineProcessorCategoryProcessorOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorCategoryProcessorOutputReference
	_jsii_.Get(
		j,
		"categoryProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) CategoryProcessorInput() *LogsCustomPipelineProcessorPipelineProcessorCategoryProcessor {
	var returns *LogsCustomPipelineProcessorPipelineProcessorCategoryProcessor
	_jsii_.Get(
		j,
		"categoryProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) DateRemapper() LogsCustomPipelineProcessorPipelineProcessorDateRemapperOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorDateRemapperOutputReference
	_jsii_.Get(
		j,
		"dateRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) DateRemapperInput() *LogsCustomPipelineProcessorPipelineProcessorDateRemapper {
	var returns *LogsCustomPipelineProcessorPipelineProcessorDateRemapper
	_jsii_.Get(
		j,
		"dateRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) GeoIpParser() LogsCustomPipelineProcessorPipelineProcessorGeoIpParserOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorGeoIpParserOutputReference
	_jsii_.Get(
		j,
		"geoIpParser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) GeoIpParserInput() *LogsCustomPipelineProcessorPipelineProcessorGeoIpParser {
	var returns *LogsCustomPipelineProcessorPipelineProcessorGeoIpParser
	_jsii_.Get(
		j,
		"geoIpParserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) GrokParser() LogsCustomPipelineProcessorPipelineProcessorGrokParserOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorGrokParserOutputReference
	_jsii_.Get(
		j,
		"grokParser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) GrokParserInput() *LogsCustomPipelineProcessorPipelineProcessorGrokParser {
	var returns *LogsCustomPipelineProcessorPipelineProcessorGrokParser
	_jsii_.Get(
		j,
		"grokParserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) LookupProcessor() LogsCustomPipelineProcessorPipelineProcessorLookupProcessorOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorLookupProcessorOutputReference
	_jsii_.Get(
		j,
		"lookupProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) LookupProcessorInput() *LogsCustomPipelineProcessorPipelineProcessorLookupProcessor {
	var returns *LogsCustomPipelineProcessorPipelineProcessorLookupProcessor
	_jsii_.Get(
		j,
		"lookupProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) MessageRemapper() LogsCustomPipelineProcessorPipelineProcessorMessageRemapperOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorMessageRemapperOutputReference
	_jsii_.Get(
		j,
		"messageRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) MessageRemapperInput() *LogsCustomPipelineProcessorPipelineProcessorMessageRemapper {
	var returns *LogsCustomPipelineProcessorPipelineProcessorMessageRemapper
	_jsii_.Get(
		j,
		"messageRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ServiceRemapper() LogsCustomPipelineProcessorPipelineProcessorServiceRemapperOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorServiceRemapperOutputReference
	_jsii_.Get(
		j,
		"serviceRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ServiceRemapperInput() *LogsCustomPipelineProcessorPipelineProcessorServiceRemapper {
	var returns *LogsCustomPipelineProcessorPipelineProcessorServiceRemapper
	_jsii_.Get(
		j,
		"serviceRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) StatusRemapper() LogsCustomPipelineProcessorPipelineProcessorStatusRemapperOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorStatusRemapperOutputReference
	_jsii_.Get(
		j,
		"statusRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) StatusRemapperInput() *LogsCustomPipelineProcessorPipelineProcessorStatusRemapper {
	var returns *LogsCustomPipelineProcessorPipelineProcessorStatusRemapper
	_jsii_.Get(
		j,
		"statusRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) StringBuilderProcessor() LogsCustomPipelineProcessorPipelineProcessorStringBuilderProcessorOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorStringBuilderProcessorOutputReference
	_jsii_.Get(
		j,
		"stringBuilderProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) StringBuilderProcessorInput() *LogsCustomPipelineProcessorPipelineProcessorStringBuilderProcessor {
	var returns *LogsCustomPipelineProcessorPipelineProcessorStringBuilderProcessor
	_jsii_.Get(
		j,
		"stringBuilderProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) TraceIdRemapper() LogsCustomPipelineProcessorPipelineProcessorTraceIdRemapperOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorTraceIdRemapperOutputReference
	_jsii_.Get(
		j,
		"traceIdRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) TraceIdRemapperInput() *LogsCustomPipelineProcessorPipelineProcessorTraceIdRemapper {
	var returns *LogsCustomPipelineProcessorPipelineProcessorTraceIdRemapper
	_jsii_.Get(
		j,
		"traceIdRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) UrlParser() LogsCustomPipelineProcessorPipelineProcessorUrlParserOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorUrlParserOutputReference
	_jsii_.Get(
		j,
		"urlParser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) UrlParserInput() *LogsCustomPipelineProcessorPipelineProcessorUrlParser {
	var returns *LogsCustomPipelineProcessorPipelineProcessorUrlParser
	_jsii_.Get(
		j,
		"urlParserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) UserAgentParser() LogsCustomPipelineProcessorPipelineProcessorUserAgentParserOutputReference {
	var returns LogsCustomPipelineProcessorPipelineProcessorUserAgentParserOutputReference
	_jsii_.Get(
		j,
		"userAgentParser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) UserAgentParserInput() *LogsCustomPipelineProcessorPipelineProcessorUserAgentParser {
	var returns *LogsCustomPipelineProcessorPipelineProcessorUserAgentParser
	_jsii_.Get(
		j,
		"userAgentParserInput",
		&returns,
	)
	return returns
}


func NewLogsCustomPipelineProcessorPipelineProcessorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LogsCustomPipelineProcessorPipelineProcessorOutputReference {
	_init_.Initialize()

	if err := validateNewLogsCustomPipelineProcessorPipelineProcessorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorPipelineProcessorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLogsCustomPipelineProcessorPipelineProcessorOutputReference_Override(l LogsCustomPipelineProcessorPipelineProcessorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorPipelineProcessorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutArithmeticProcessor(value *LogsCustomPipelineProcessorPipelineProcessorArithmeticProcessor) {
	if err := l.validatePutArithmeticProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putArithmeticProcessor",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutAttributeRemapper(value *LogsCustomPipelineProcessorPipelineProcessorAttributeRemapper) {
	if err := l.validatePutAttributeRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAttributeRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutCategoryProcessor(value *LogsCustomPipelineProcessorPipelineProcessorCategoryProcessor) {
	if err := l.validatePutCategoryProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCategoryProcessor",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutDateRemapper(value *LogsCustomPipelineProcessorPipelineProcessorDateRemapper) {
	if err := l.validatePutDateRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDateRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutGeoIpParser(value *LogsCustomPipelineProcessorPipelineProcessorGeoIpParser) {
	if err := l.validatePutGeoIpParserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putGeoIpParser",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutGrokParser(value *LogsCustomPipelineProcessorPipelineProcessorGrokParser) {
	if err := l.validatePutGrokParserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putGrokParser",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutLookupProcessor(value *LogsCustomPipelineProcessorPipelineProcessorLookupProcessor) {
	if err := l.validatePutLookupProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLookupProcessor",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutMessageRemapper(value *LogsCustomPipelineProcessorPipelineProcessorMessageRemapper) {
	if err := l.validatePutMessageRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMessageRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutServiceRemapper(value *LogsCustomPipelineProcessorPipelineProcessorServiceRemapper) {
	if err := l.validatePutServiceRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putServiceRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutStatusRemapper(value *LogsCustomPipelineProcessorPipelineProcessorStatusRemapper) {
	if err := l.validatePutStatusRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStatusRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutStringBuilderProcessor(value *LogsCustomPipelineProcessorPipelineProcessorStringBuilderProcessor) {
	if err := l.validatePutStringBuilderProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStringBuilderProcessor",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutTraceIdRemapper(value *LogsCustomPipelineProcessorPipelineProcessorTraceIdRemapper) {
	if err := l.validatePutTraceIdRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTraceIdRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutUrlParser(value *LogsCustomPipelineProcessorPipelineProcessorUrlParser) {
	if err := l.validatePutUrlParserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putUrlParser",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) PutUserAgentParser(value *LogsCustomPipelineProcessorPipelineProcessorUserAgentParser) {
	if err := l.validatePutUserAgentParserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putUserAgentParser",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetArithmeticProcessor() {
	_jsii_.InvokeVoid(
		l,
		"resetArithmeticProcessor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetAttributeRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetAttributeRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetCategoryProcessor() {
	_jsii_.InvokeVoid(
		l,
		"resetCategoryProcessor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetDateRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetDateRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetGeoIpParser() {
	_jsii_.InvokeVoid(
		l,
		"resetGeoIpParser",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetGrokParser() {
	_jsii_.InvokeVoid(
		l,
		"resetGrokParser",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetLookupProcessor() {
	_jsii_.InvokeVoid(
		l,
		"resetLookupProcessor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetMessageRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetMessageRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetServiceRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetServiceRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetStatusRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetStatusRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetStringBuilderProcessor() {
	_jsii_.InvokeVoid(
		l,
		"resetStringBuilderProcessor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetTraceIdRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetTraceIdRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetUrlParser() {
	_jsii_.InvokeVoid(
		l,
		"resetUrlParser",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ResetUserAgentParser() {
	_jsii_.InvokeVoid(
		l,
		"resetUserAgentParser",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineProcessorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

