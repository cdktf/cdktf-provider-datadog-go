// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/observabilitypipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityPipelineConfigDestinationsOutputReference interface {
	cdktf.ComplexObject
	AmazonOpensearch() ObservabilityPipelineConfigDestinationsAmazonOpensearchList
	AmazonOpensearchInput() interface{}
	AzureStorage() ObservabilityPipelineConfigDestinationsAzureStorageList
	AzureStorageInput() interface{}
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
	DatadogLogs() ObservabilityPipelineConfigDestinationsDatadogLogsList
	DatadogLogsInput() interface{}
	Elasticsearch() ObservabilityPipelineConfigDestinationsElasticsearchList
	ElasticsearchInput() interface{}
	// Experimental.
	Fqn() *string
	GoogleChronicle() ObservabilityPipelineConfigDestinationsGoogleChronicleList
	GoogleChronicleInput() interface{}
	GoogleCloudStorage() ObservabilityPipelineConfigDestinationsGoogleCloudStorageList
	GoogleCloudStorageInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MicrosoftSentinel() ObservabilityPipelineConfigDestinationsMicrosoftSentinelList
	MicrosoftSentinelInput() interface{}
	NewRelic() ObservabilityPipelineConfigDestinationsNewRelicList
	NewRelicInput() interface{}
	Opensearch() ObservabilityPipelineConfigDestinationsOpensearchList
	OpensearchInput() interface{}
	Rsyslog() ObservabilityPipelineConfigDestinationsRsyslogList
	RsyslogInput() interface{}
	SentinelOne() ObservabilityPipelineConfigDestinationsSentinelOneList
	SentinelOneInput() interface{}
	SplunkHec() ObservabilityPipelineConfigDestinationsSplunkHecList
	SplunkHecInput() interface{}
	SumoLogic() ObservabilityPipelineConfigDestinationsSumoLogicList
	SumoLogicInput() interface{}
	SyslogNg() ObservabilityPipelineConfigDestinationsSyslogNgList
	SyslogNgInput() interface{}
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAmazonOpensearch(value interface{})
	PutAzureStorage(value interface{})
	PutDatadogLogs(value interface{})
	PutElasticsearch(value interface{})
	PutGoogleChronicle(value interface{})
	PutGoogleCloudStorage(value interface{})
	PutMicrosoftSentinel(value interface{})
	PutNewRelic(value interface{})
	PutOpensearch(value interface{})
	PutRsyslog(value interface{})
	PutSentinelOne(value interface{})
	PutSplunkHec(value interface{})
	PutSumoLogic(value interface{})
	PutSyslogNg(value interface{})
	ResetAmazonOpensearch()
	ResetAzureStorage()
	ResetDatadogLogs()
	ResetElasticsearch()
	ResetGoogleChronicle()
	ResetGoogleCloudStorage()
	ResetMicrosoftSentinel()
	ResetNewRelic()
	ResetOpensearch()
	ResetRsyslog()
	ResetSentinelOne()
	ResetSplunkHec()
	ResetSumoLogic()
	ResetSyslogNg()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityPipelineConfigDestinationsOutputReference
type jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) AmazonOpensearch() ObservabilityPipelineConfigDestinationsAmazonOpensearchList {
	var returns ObservabilityPipelineConfigDestinationsAmazonOpensearchList
	_jsii_.Get(
		j,
		"amazonOpensearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) AmazonOpensearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonOpensearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) AzureStorage() ObservabilityPipelineConfigDestinationsAzureStorageList {
	var returns ObservabilityPipelineConfigDestinationsAzureStorageList
	_jsii_.Get(
		j,
		"azureStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) AzureStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) DatadogLogs() ObservabilityPipelineConfigDestinationsDatadogLogsList {
	var returns ObservabilityPipelineConfigDestinationsDatadogLogsList
	_jsii_.Get(
		j,
		"datadogLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) DatadogLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datadogLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) Elasticsearch() ObservabilityPipelineConfigDestinationsElasticsearchList {
	var returns ObservabilityPipelineConfigDestinationsElasticsearchList
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ElasticsearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) GoogleChronicle() ObservabilityPipelineConfigDestinationsGoogleChronicleList {
	var returns ObservabilityPipelineConfigDestinationsGoogleChronicleList
	_jsii_.Get(
		j,
		"googleChronicle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) GoogleChronicleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googleChronicleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) GoogleCloudStorage() ObservabilityPipelineConfigDestinationsGoogleCloudStorageList {
	var returns ObservabilityPipelineConfigDestinationsGoogleCloudStorageList
	_jsii_.Get(
		j,
		"googleCloudStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) GoogleCloudStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googleCloudStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) MicrosoftSentinel() ObservabilityPipelineConfigDestinationsMicrosoftSentinelList {
	var returns ObservabilityPipelineConfigDestinationsMicrosoftSentinelList
	_jsii_.Get(
		j,
		"microsoftSentinel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) MicrosoftSentinelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftSentinelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) NewRelic() ObservabilityPipelineConfigDestinationsNewRelicList {
	var returns ObservabilityPipelineConfigDestinationsNewRelicList
	_jsii_.Get(
		j,
		"newRelic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) NewRelicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newRelicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) Opensearch() ObservabilityPipelineConfigDestinationsOpensearchList {
	var returns ObservabilityPipelineConfigDestinationsOpensearchList
	_jsii_.Get(
		j,
		"opensearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) OpensearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opensearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) Rsyslog() ObservabilityPipelineConfigDestinationsRsyslogList {
	var returns ObservabilityPipelineConfigDestinationsRsyslogList
	_jsii_.Get(
		j,
		"rsyslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) RsyslogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rsyslogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) SentinelOne() ObservabilityPipelineConfigDestinationsSentinelOneList {
	var returns ObservabilityPipelineConfigDestinationsSentinelOneList
	_jsii_.Get(
		j,
		"sentinelOne",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) SentinelOneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sentinelOneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) SplunkHec() ObservabilityPipelineConfigDestinationsSplunkHecList {
	var returns ObservabilityPipelineConfigDestinationsSplunkHecList
	_jsii_.Get(
		j,
		"splunkHec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) SplunkHecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkHecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) SumoLogic() ObservabilityPipelineConfigDestinationsSumoLogicList {
	var returns ObservabilityPipelineConfigDestinationsSumoLogicList
	_jsii_.Get(
		j,
		"sumoLogic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) SumoLogicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sumoLogicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) SyslogNg() ObservabilityPipelineConfigDestinationsSyslogNgList {
	var returns ObservabilityPipelineConfigDestinationsSyslogNgList
	_jsii_.Get(
		j,
		"syslogNg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) SyslogNgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syslogNgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigDestinationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ObservabilityPipelineConfigDestinationsOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigDestinationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigDestinationsOutputReference_Override(o ObservabilityPipelineConfigDestinationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutAmazonOpensearch(value interface{}) {
	if err := o.validatePutAmazonOpensearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAmazonOpensearch",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutAzureStorage(value interface{}) {
	if err := o.validatePutAzureStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAzureStorage",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutDatadogLogs(value interface{}) {
	if err := o.validatePutDatadogLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDatadogLogs",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutElasticsearch(value interface{}) {
	if err := o.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutGoogleChronicle(value interface{}) {
	if err := o.validatePutGoogleChronicleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGoogleChronicle",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutGoogleCloudStorage(value interface{}) {
	if err := o.validatePutGoogleCloudStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGoogleCloudStorage",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutMicrosoftSentinel(value interface{}) {
	if err := o.validatePutMicrosoftSentinelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMicrosoftSentinel",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutNewRelic(value interface{}) {
	if err := o.validatePutNewRelicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNewRelic",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutOpensearch(value interface{}) {
	if err := o.validatePutOpensearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOpensearch",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutRsyslog(value interface{}) {
	if err := o.validatePutRsyslogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRsyslog",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutSentinelOne(value interface{}) {
	if err := o.validatePutSentinelOneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSentinelOne",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutSplunkHec(value interface{}) {
	if err := o.validatePutSplunkHecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSplunkHec",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutSumoLogic(value interface{}) {
	if err := o.validatePutSumoLogicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSumoLogic",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) PutSyslogNg(value interface{}) {
	if err := o.validatePutSyslogNgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSyslogNg",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetAmazonOpensearch() {
	_jsii_.InvokeVoid(
		o,
		"resetAmazonOpensearch",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetAzureStorage() {
	_jsii_.InvokeVoid(
		o,
		"resetAzureStorage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetDatadogLogs() {
	_jsii_.InvokeVoid(
		o,
		"resetDatadogLogs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetGoogleChronicle() {
	_jsii_.InvokeVoid(
		o,
		"resetGoogleChronicle",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetGoogleCloudStorage() {
	_jsii_.InvokeVoid(
		o,
		"resetGoogleCloudStorage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetMicrosoftSentinel() {
	_jsii_.InvokeVoid(
		o,
		"resetMicrosoftSentinel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetNewRelic() {
	_jsii_.InvokeVoid(
		o,
		"resetNewRelic",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetOpensearch() {
	_jsii_.InvokeVoid(
		o,
		"resetOpensearch",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetRsyslog() {
	_jsii_.InvokeVoid(
		o,
		"resetRsyslog",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetSentinelOne() {
	_jsii_.InvokeVoid(
		o,
		"resetSentinelOne",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetSplunkHec() {
	_jsii_.InvokeVoid(
		o,
		"resetSplunkHec",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetSumoLogic() {
	_jsii_.InvokeVoid(
		o,
		"resetSumoLogic",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ResetSyslogNg() {
	_jsii_.InvokeVoid(
		o,
		"resetSyslogNg",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

