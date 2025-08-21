// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/observabilitypipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObservabilityPipelineConfigSourcesOutputReference interface {
	cdktf.ComplexObject
	AmazonDataFirehose() ObservabilityPipelineConfigSourcesAmazonDataFirehoseList
	AmazonDataFirehoseInput() interface{}
	AmazonS3() ObservabilityPipelineConfigSourcesAmazonS3List
	AmazonS3Input() interface{}
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
	DatadogAgent() ObservabilityPipelineConfigSourcesDatadogAgentList
	DatadogAgentInput() interface{}
	FluentBit() ObservabilityPipelineConfigSourcesFluentBitList
	FluentBitInput() interface{}
	Fluentd() ObservabilityPipelineConfigSourcesFluentdList
	FluentdInput() interface{}
	// Experimental.
	Fqn() *string
	GooglePubsub() ObservabilityPipelineConfigSourcesGooglePubsubList
	GooglePubsubInput() interface{}
	HttpClient() ObservabilityPipelineConfigSourcesHttpClientList
	HttpClientInput() interface{}
	HttpServer() ObservabilityPipelineConfigSourcesHttpServerList
	HttpServerInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Kafka() ObservabilityPipelineConfigSourcesKafkaList
	KafkaInput() interface{}
	Logstash() ObservabilityPipelineConfigSourcesLogstashList
	LogstashInput() interface{}
	Rsyslog() ObservabilityPipelineConfigSourcesRsyslogList
	RsyslogInput() interface{}
	Socket() ObservabilityPipelineConfigSourcesSocketList
	SocketInput() interface{}
	SplunkHec() ObservabilityPipelineConfigSourcesSplunkHecList
	SplunkHecInput() interface{}
	SplunkTcp() ObservabilityPipelineConfigSourcesSplunkTcpList
	SplunkTcpInput() interface{}
	SumoLogic() ObservabilityPipelineConfigSourcesSumoLogicList
	SumoLogicInput() interface{}
	SyslogNg() ObservabilityPipelineConfigSourcesSyslogNgList
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
	PutAmazonDataFirehose(value interface{})
	PutAmazonS3(value interface{})
	PutDatadogAgent(value interface{})
	PutFluentBit(value interface{})
	PutFluentd(value interface{})
	PutGooglePubsub(value interface{})
	PutHttpClient(value interface{})
	PutHttpServer(value interface{})
	PutKafka(value interface{})
	PutLogstash(value interface{})
	PutRsyslog(value interface{})
	PutSocket(value interface{})
	PutSplunkHec(value interface{})
	PutSplunkTcp(value interface{})
	PutSumoLogic(value interface{})
	PutSyslogNg(value interface{})
	ResetAmazonDataFirehose()
	ResetAmazonS3()
	ResetDatadogAgent()
	ResetFluentBit()
	ResetFluentd()
	ResetGooglePubsub()
	ResetHttpClient()
	ResetHttpServer()
	ResetKafka()
	ResetLogstash()
	ResetRsyslog()
	ResetSocket()
	ResetSplunkHec()
	ResetSplunkTcp()
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

// The jsii proxy struct for ObservabilityPipelineConfigSourcesOutputReference
type jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) AmazonDataFirehose() ObservabilityPipelineConfigSourcesAmazonDataFirehoseList {
	var returns ObservabilityPipelineConfigSourcesAmazonDataFirehoseList
	_jsii_.Get(
		j,
		"amazonDataFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) AmazonDataFirehoseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonDataFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) AmazonS3() ObservabilityPipelineConfigSourcesAmazonS3List {
	var returns ObservabilityPipelineConfigSourcesAmazonS3List
	_jsii_.Get(
		j,
		"amazonS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) AmazonS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonS3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) DatadogAgent() ObservabilityPipelineConfigSourcesDatadogAgentList {
	var returns ObservabilityPipelineConfigSourcesDatadogAgentList
	_jsii_.Get(
		j,
		"datadogAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) DatadogAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datadogAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) FluentBit() ObservabilityPipelineConfigSourcesFluentBitList {
	var returns ObservabilityPipelineConfigSourcesFluentBitList
	_jsii_.Get(
		j,
		"fluentBit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) FluentBitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fluentBitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) Fluentd() ObservabilityPipelineConfigSourcesFluentdList {
	var returns ObservabilityPipelineConfigSourcesFluentdList
	_jsii_.Get(
		j,
		"fluentd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) FluentdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fluentdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) GooglePubsub() ObservabilityPipelineConfigSourcesGooglePubsubList {
	var returns ObservabilityPipelineConfigSourcesGooglePubsubList
	_jsii_.Get(
		j,
		"googlePubsub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) GooglePubsubInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googlePubsubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) HttpClient() ObservabilityPipelineConfigSourcesHttpClientList {
	var returns ObservabilityPipelineConfigSourcesHttpClientList
	_jsii_.Get(
		j,
		"httpClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) HttpClientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) HttpServer() ObservabilityPipelineConfigSourcesHttpServerList {
	var returns ObservabilityPipelineConfigSourcesHttpServerList
	_jsii_.Get(
		j,
		"httpServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) HttpServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) Kafka() ObservabilityPipelineConfigSourcesKafkaList {
	var returns ObservabilityPipelineConfigSourcesKafkaList
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) KafkaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) Logstash() ObservabilityPipelineConfigSourcesLogstashList {
	var returns ObservabilityPipelineConfigSourcesLogstashList
	_jsii_.Get(
		j,
		"logstash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) LogstashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logstashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) Rsyslog() ObservabilityPipelineConfigSourcesRsyslogList {
	var returns ObservabilityPipelineConfigSourcesRsyslogList
	_jsii_.Get(
		j,
		"rsyslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) RsyslogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rsyslogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) Socket() ObservabilityPipelineConfigSourcesSocketList {
	var returns ObservabilityPipelineConfigSourcesSocketList
	_jsii_.Get(
		j,
		"socket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) SocketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"socketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) SplunkHec() ObservabilityPipelineConfigSourcesSplunkHecList {
	var returns ObservabilityPipelineConfigSourcesSplunkHecList
	_jsii_.Get(
		j,
		"splunkHec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) SplunkHecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkHecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) SplunkTcp() ObservabilityPipelineConfigSourcesSplunkTcpList {
	var returns ObservabilityPipelineConfigSourcesSplunkTcpList
	_jsii_.Get(
		j,
		"splunkTcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) SplunkTcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkTcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) SumoLogic() ObservabilityPipelineConfigSourcesSumoLogicList {
	var returns ObservabilityPipelineConfigSourcesSumoLogicList
	_jsii_.Get(
		j,
		"sumoLogic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) SumoLogicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sumoLogicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) SyslogNg() ObservabilityPipelineConfigSourcesSyslogNgList {
	var returns ObservabilityPipelineConfigSourcesSyslogNgList
	_jsii_.Get(
		j,
		"syslogNg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) SyslogNgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syslogNgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigSourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ObservabilityPipelineConfigSourcesOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigSourcesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigSourcesOutputReference_Override(o ObservabilityPipelineConfigSourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutAmazonDataFirehose(value interface{}) {
	if err := o.validatePutAmazonDataFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAmazonDataFirehose",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutAmazonS3(value interface{}) {
	if err := o.validatePutAmazonS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAmazonS3",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutDatadogAgent(value interface{}) {
	if err := o.validatePutDatadogAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDatadogAgent",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutFluentBit(value interface{}) {
	if err := o.validatePutFluentBitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putFluentBit",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutFluentd(value interface{}) {
	if err := o.validatePutFluentdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putFluentd",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutGooglePubsub(value interface{}) {
	if err := o.validatePutGooglePubsubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGooglePubsub",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutHttpClient(value interface{}) {
	if err := o.validatePutHttpClientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHttpClient",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutHttpServer(value interface{}) {
	if err := o.validatePutHttpServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHttpServer",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutKafka(value interface{}) {
	if err := o.validatePutKafkaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putKafka",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutLogstash(value interface{}) {
	if err := o.validatePutLogstashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLogstash",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutRsyslog(value interface{}) {
	if err := o.validatePutRsyslogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRsyslog",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutSocket(value interface{}) {
	if err := o.validatePutSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSocket",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutSplunkHec(value interface{}) {
	if err := o.validatePutSplunkHecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSplunkHec",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutSplunkTcp(value interface{}) {
	if err := o.validatePutSplunkTcpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSplunkTcp",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutSumoLogic(value interface{}) {
	if err := o.validatePutSumoLogicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSumoLogic",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) PutSyslogNg(value interface{}) {
	if err := o.validatePutSyslogNgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSyslogNg",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetAmazonDataFirehose() {
	_jsii_.InvokeVoid(
		o,
		"resetAmazonDataFirehose",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetAmazonS3() {
	_jsii_.InvokeVoid(
		o,
		"resetAmazonS3",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetDatadogAgent() {
	_jsii_.InvokeVoid(
		o,
		"resetDatadogAgent",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetFluentBit() {
	_jsii_.InvokeVoid(
		o,
		"resetFluentBit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetFluentd() {
	_jsii_.InvokeVoid(
		o,
		"resetFluentd",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetGooglePubsub() {
	_jsii_.InvokeVoid(
		o,
		"resetGooglePubsub",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetHttpClient() {
	_jsii_.InvokeVoid(
		o,
		"resetHttpClient",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetHttpServer() {
	_jsii_.InvokeVoid(
		o,
		"resetHttpServer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetKafka() {
	_jsii_.InvokeVoid(
		o,
		"resetKafka",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetLogstash() {
	_jsii_.InvokeVoid(
		o,
		"resetLogstash",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetRsyslog() {
	_jsii_.InvokeVoid(
		o,
		"resetRsyslog",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetSocket() {
	_jsii_.InvokeVoid(
		o,
		"resetSocket",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetSplunkHec() {
	_jsii_.InvokeVoid(
		o,
		"resetSplunkHec",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetSplunkTcp() {
	_jsii_.InvokeVoid(
		o,
		"resetSplunkTcp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetSumoLogic() {
	_jsii_.InvokeVoid(
		o,
		"resetSumoLogic",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ResetSyslogNg() {
	_jsii_.InvokeVoid(
		o,
		"resetSyslogNg",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

