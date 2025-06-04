// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/observabilitypipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityPipelineConfigProcessorsOutputReference interface {
	cdktf.ComplexObject
	AddEnvVars() ObservabilityPipelineConfigProcessorsAddEnvVarsList
	AddEnvVarsInput() interface{}
	AddFields() ObservabilityPipelineConfigProcessorsAddFieldsList
	AddFieldsInput() interface{}
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
	Dedupe() ObservabilityPipelineConfigProcessorsDedupeList
	DedupeInput() interface{}
	EnrichmentTable() ObservabilityPipelineConfigProcessorsEnrichmentTableList
	EnrichmentTableInput() interface{}
	Filter() ObservabilityPipelineConfigProcessorsFilterList
	FilterInput() interface{}
	// Experimental.
	Fqn() *string
	GenerateDatadogMetrics() ObservabilityPipelineConfigProcessorsGenerateDatadogMetricsList
	GenerateDatadogMetricsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OcsfMapper() ObservabilityPipelineConfigProcessorsOcsfMapperList
	OcsfMapperInput() interface{}
	ParseGrok() ObservabilityPipelineConfigProcessorsParseGrokList
	ParseGrokInput() interface{}
	ParseJson() ObservabilityPipelineConfigProcessorsParseJsonList
	ParseJsonInput() interface{}
	Quota() ObservabilityPipelineConfigProcessorsQuotaList
	QuotaInput() interface{}
	Reduce() ObservabilityPipelineConfigProcessorsReduceList
	ReduceInput() interface{}
	RemoveFields() ObservabilityPipelineConfigProcessorsRemoveFieldsList
	RemoveFieldsInput() interface{}
	RenameFields() ObservabilityPipelineConfigProcessorsRenameFieldsList
	RenameFieldsInput() interface{}
	Sample() ObservabilityPipelineConfigProcessorsSampleList
	SampleInput() interface{}
	SensitiveDataScanner() ObservabilityPipelineConfigProcessorsSensitiveDataScannerList
	SensitiveDataScannerInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Throttle() ObservabilityPipelineConfigProcessorsThrottleList
	ThrottleInput() interface{}
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
	PutAddEnvVars(value interface{})
	PutAddFields(value interface{})
	PutDedupe(value interface{})
	PutEnrichmentTable(value interface{})
	PutFilter(value interface{})
	PutGenerateDatadogMetrics(value interface{})
	PutOcsfMapper(value interface{})
	PutParseGrok(value interface{})
	PutParseJson(value interface{})
	PutQuota(value interface{})
	PutReduce(value interface{})
	PutRemoveFields(value interface{})
	PutRenameFields(value interface{})
	PutSample(value interface{})
	PutSensitiveDataScanner(value interface{})
	PutThrottle(value interface{})
	ResetAddEnvVars()
	ResetAddFields()
	ResetDedupe()
	ResetEnrichmentTable()
	ResetFilter()
	ResetGenerateDatadogMetrics()
	ResetOcsfMapper()
	ResetParseGrok()
	ResetParseJson()
	ResetQuota()
	ResetReduce()
	ResetRemoveFields()
	ResetRenameFields()
	ResetSample()
	ResetSensitiveDataScanner()
	ResetThrottle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityPipelineConfigProcessorsOutputReference
type jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) AddEnvVars() ObservabilityPipelineConfigProcessorsAddEnvVarsList {
	var returns ObservabilityPipelineConfigProcessorsAddEnvVarsList
	_jsii_.Get(
		j,
		"addEnvVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) AddEnvVarsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addEnvVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) AddFields() ObservabilityPipelineConfigProcessorsAddFieldsList {
	var returns ObservabilityPipelineConfigProcessorsAddFieldsList
	_jsii_.Get(
		j,
		"addFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) AddFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) Dedupe() ObservabilityPipelineConfigProcessorsDedupeList {
	var returns ObservabilityPipelineConfigProcessorsDedupeList
	_jsii_.Get(
		j,
		"dedupe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) DedupeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedupeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) EnrichmentTable() ObservabilityPipelineConfigProcessorsEnrichmentTableList {
	var returns ObservabilityPipelineConfigProcessorsEnrichmentTableList
	_jsii_.Get(
		j,
		"enrichmentTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) EnrichmentTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrichmentTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) Filter() ObservabilityPipelineConfigProcessorsFilterList {
	var returns ObservabilityPipelineConfigProcessorsFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) GenerateDatadogMetrics() ObservabilityPipelineConfigProcessorsGenerateDatadogMetricsList {
	var returns ObservabilityPipelineConfigProcessorsGenerateDatadogMetricsList
	_jsii_.Get(
		j,
		"generateDatadogMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) GenerateDatadogMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateDatadogMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) OcsfMapper() ObservabilityPipelineConfigProcessorsOcsfMapperList {
	var returns ObservabilityPipelineConfigProcessorsOcsfMapperList
	_jsii_.Get(
		j,
		"ocsfMapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) OcsfMapperInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocsfMapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ParseGrok() ObservabilityPipelineConfigProcessorsParseGrokList {
	var returns ObservabilityPipelineConfigProcessorsParseGrokList
	_jsii_.Get(
		j,
		"parseGrok",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ParseGrokInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseGrokInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ParseJson() ObservabilityPipelineConfigProcessorsParseJsonList {
	var returns ObservabilityPipelineConfigProcessorsParseJsonList
	_jsii_.Get(
		j,
		"parseJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ParseJsonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) Quota() ObservabilityPipelineConfigProcessorsQuotaList {
	var returns ObservabilityPipelineConfigProcessorsQuotaList
	_jsii_.Get(
		j,
		"quota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) QuotaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) Reduce() ObservabilityPipelineConfigProcessorsReduceList {
	var returns ObservabilityPipelineConfigProcessorsReduceList
	_jsii_.Get(
		j,
		"reduce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ReduceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reduceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) RemoveFields() ObservabilityPipelineConfigProcessorsRemoveFieldsList {
	var returns ObservabilityPipelineConfigProcessorsRemoveFieldsList
	_jsii_.Get(
		j,
		"removeFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) RemoveFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) RenameFields() ObservabilityPipelineConfigProcessorsRenameFieldsList {
	var returns ObservabilityPipelineConfigProcessorsRenameFieldsList
	_jsii_.Get(
		j,
		"renameFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) RenameFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renameFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) Sample() ObservabilityPipelineConfigProcessorsSampleList {
	var returns ObservabilityPipelineConfigProcessorsSampleList
	_jsii_.Get(
		j,
		"sample",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) SampleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sampleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) SensitiveDataScanner() ObservabilityPipelineConfigProcessorsSensitiveDataScannerList {
	var returns ObservabilityPipelineConfigProcessorsSensitiveDataScannerList
	_jsii_.Get(
		j,
		"sensitiveDataScanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) SensitiveDataScannerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sensitiveDataScannerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) Throttle() ObservabilityPipelineConfigProcessorsThrottleList {
	var returns ObservabilityPipelineConfigProcessorsThrottleList
	_jsii_.Get(
		j,
		"throttle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ThrottleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleInput",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigProcessorsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ObservabilityPipelineConfigProcessorsOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigProcessorsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigProcessorsOutputReference_Override(o ObservabilityPipelineConfigProcessorsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutAddEnvVars(value interface{}) {
	if err := o.validatePutAddEnvVarsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAddEnvVars",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutAddFields(value interface{}) {
	if err := o.validatePutAddFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAddFields",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutDedupe(value interface{}) {
	if err := o.validatePutDedupeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDedupe",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutEnrichmentTable(value interface{}) {
	if err := o.validatePutEnrichmentTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEnrichmentTable",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutFilter(value interface{}) {
	if err := o.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putFilter",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutGenerateDatadogMetrics(value interface{}) {
	if err := o.validatePutGenerateDatadogMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGenerateDatadogMetrics",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutOcsfMapper(value interface{}) {
	if err := o.validatePutOcsfMapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOcsfMapper",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutParseGrok(value interface{}) {
	if err := o.validatePutParseGrokParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putParseGrok",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutParseJson(value interface{}) {
	if err := o.validatePutParseJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putParseJson",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutQuota(value interface{}) {
	if err := o.validatePutQuotaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putQuota",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutReduce(value interface{}) {
	if err := o.validatePutReduceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putReduce",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutRemoveFields(value interface{}) {
	if err := o.validatePutRemoveFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRemoveFields",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutRenameFields(value interface{}) {
	if err := o.validatePutRenameFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRenameFields",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutSample(value interface{}) {
	if err := o.validatePutSampleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSample",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutSensitiveDataScanner(value interface{}) {
	if err := o.validatePutSensitiveDataScannerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSensitiveDataScanner",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) PutThrottle(value interface{}) {
	if err := o.validatePutThrottleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putThrottle",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetAddEnvVars() {
	_jsii_.InvokeVoid(
		o,
		"resetAddEnvVars",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetAddFields() {
	_jsii_.InvokeVoid(
		o,
		"resetAddFields",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetDedupe() {
	_jsii_.InvokeVoid(
		o,
		"resetDedupe",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetEnrichmentTable() {
	_jsii_.InvokeVoid(
		o,
		"resetEnrichmentTable",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		o,
		"resetFilter",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetGenerateDatadogMetrics() {
	_jsii_.InvokeVoid(
		o,
		"resetGenerateDatadogMetrics",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetOcsfMapper() {
	_jsii_.InvokeVoid(
		o,
		"resetOcsfMapper",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetParseGrok() {
	_jsii_.InvokeVoid(
		o,
		"resetParseGrok",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetParseJson() {
	_jsii_.InvokeVoid(
		o,
		"resetParseJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetQuota() {
	_jsii_.InvokeVoid(
		o,
		"resetQuota",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetReduce() {
	_jsii_.InvokeVoid(
		o,
		"resetReduce",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetRemoveFields() {
	_jsii_.InvokeVoid(
		o,
		"resetRemoveFields",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetRenameFields() {
	_jsii_.InvokeVoid(
		o,
		"resetRenameFields",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetSample() {
	_jsii_.InvokeVoid(
		o,
		"resetSample",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetSensitiveDataScanner() {
	_jsii_.InvokeVoid(
		o,
		"resetSensitiveDataScanner",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ResetThrottle() {
	_jsii_.InvokeVoid(
		o,
		"resetThrottle",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

