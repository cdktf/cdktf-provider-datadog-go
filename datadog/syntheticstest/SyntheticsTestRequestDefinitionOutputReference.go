// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/syntheticstest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestRequestDefinitionOutputReference interface {
	cdktf.ComplexObject
	Body() *string
	SetBody(val *string)
	BodyInput() *string
	BodyType() *string
	SetBodyType(val *string)
	BodyTypeInput() *string
	CallType() *string
	SetCallType(val *string)
	CallTypeInput() *string
	CertificateDomains() *[]*string
	SetCertificateDomains(val *[]*string)
	CertificateDomainsInput() *[]*string
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
	DnsServer() *string
	SetDnsServer(val *string)
	DnsServerInput() *string
	DnsServerPort() *string
	SetDnsServerPort(val *string)
	DnsServerPortInput() *string
	Form() *map[string]*string
	SetForm(val *map[string]*string)
	FormInput() *map[string]*string
	// Experimental.
	Fqn() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	HttpVersion() *string
	SetHttpVersion(val *string)
	HttpVersionInput() *string
	InternalValue() *SyntheticsTestRequestDefinition
	SetInternalValue(val *SyntheticsTestRequestDefinition)
	IsMessageBase64Encoded() interface{}
	SetIsMessageBase64Encoded(val interface{})
	IsMessageBase64EncodedInput() interface{}
	Message() *string
	SetMessage(val *string)
	MessageInput() *string
	Method() *string
	SetMethod(val *string)
	MethodInput() *string
	NoSavingResponseBody() interface{}
	SetNoSavingResponseBody(val interface{})
	NoSavingResponseBodyInput() interface{}
	NumberOfPackets() *float64
	SetNumberOfPackets(val *float64)
	NumberOfPacketsInput() *float64
	PersistCookies() interface{}
	SetPersistCookies(val interface{})
	PersistCookiesInput() interface{}
	PlainProtoFile() *string
	SetPlainProtoFile(val *string)
	PlainProtoFileInput() *string
	Port() *string
	SetPort(val *string)
	PortInput() *string
	ProtoJsonDescriptor() *string
	SetProtoJsonDescriptor(val *string)
	ProtoJsonDescriptorInput() *string
	Servername() *string
	SetServername(val *string)
	ServernameInput() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	ShouldTrackHops() interface{}
	SetShouldTrackHops(val interface{})
	ShouldTrackHopsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	ResetBody()
	ResetBodyType()
	ResetCallType()
	ResetCertificateDomains()
	ResetDnsServer()
	ResetDnsServerPort()
	ResetForm()
	ResetHost()
	ResetHttpVersion()
	ResetIsMessageBase64Encoded()
	ResetMessage()
	ResetMethod()
	ResetNoSavingResponseBody()
	ResetNumberOfPackets()
	ResetPersistCookies()
	ResetPlainProtoFile()
	ResetPort()
	ResetProtoJsonDescriptor()
	ResetServername()
	ResetService()
	ResetShouldTrackHops()
	ResetTimeout()
	ResetUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SyntheticsTestRequestDefinitionOutputReference
type jsiiProxy_SyntheticsTestRequestDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) BodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) BodyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) BodyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) CallType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) CallTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) CertificateDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) CertificateDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) DnsServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) DnsServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) DnsServerPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) DnsServerPortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServerPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) Form() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"form",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) FormInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"formInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) HttpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) InternalValue() *SyntheticsTestRequestDefinition {
	var returns *SyntheticsTestRequestDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) IsMessageBase64Encoded() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMessageBase64Encoded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) IsMessageBase64EncodedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMessageBase64EncodedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) MessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) MethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) NoSavingResponseBody() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSavingResponseBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) NoSavingResponseBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSavingResponseBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) NumberOfPackets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfPackets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) NumberOfPacketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfPacketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) PersistCookies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) PersistCookiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"persistCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) PlainProtoFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plainProtoFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) PlainProtoFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plainProtoFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) PortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ProtoJsonDescriptor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protoJsonDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ProtoJsonDescriptorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protoJsonDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) Servername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ServernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ShouldTrackHops() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTrackHops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ShouldTrackHopsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTrackHopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewSyntheticsTestRequestDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SyntheticsTestRequestDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewSyntheticsTestRequestDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTestRequestDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestRequestDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSyntheticsTestRequestDefinitionOutputReference_Override(s SyntheticsTestRequestDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestRequestDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetBody(val *string) {
	if err := j.validateSetBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetBodyType(val *string) {
	if err := j.validateSetBodyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bodyType",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetCallType(val *string) {
	if err := j.validateSetCallTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callType",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetCertificateDomains(val *[]*string) {
	if err := j.validateSetCertificateDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateDomains",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetDnsServer(val *string) {
	if err := j.validateSetDnsServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServer",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetDnsServerPort(val *string) {
	if err := j.validateSetDnsServerPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServerPort",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetForm(val *map[string]*string) {
	if err := j.validateSetFormParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"form",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetHttpVersion(val *string) {
	if err := j.validateSetHttpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpVersion",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetInternalValue(val *SyntheticsTestRequestDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetIsMessageBase64Encoded(val interface{}) {
	if err := j.validateSetIsMessageBase64EncodedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMessageBase64Encoded",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetMessage(val *string) {
	if err := j.validateSetMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetNoSavingResponseBody(val interface{}) {
	if err := j.validateSetNoSavingResponseBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSavingResponseBody",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetNumberOfPackets(val *float64) {
	if err := j.validateSetNumberOfPacketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfPackets",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetPersistCookies(val interface{}) {
	if err := j.validateSetPersistCookiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistCookies",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetPlainProtoFile(val *string) {
	if err := j.validateSetPlainProtoFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"plainProtoFile",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetPort(val *string) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetProtoJsonDescriptor(val *string) {
	if err := j.validateSetProtoJsonDescriptorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protoJsonDescriptor",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetServername(val *string) {
	if err := j.validateSetServernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servername",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetShouldTrackHops(val interface{}) {
	if err := j.validateSetShouldTrackHopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldTrackHops",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		s,
		"resetBody",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetBodyType() {
	_jsii_.InvokeVoid(
		s,
		"resetBodyType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetCallType() {
	_jsii_.InvokeVoid(
		s,
		"resetCallType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetCertificateDomains() {
	_jsii_.InvokeVoid(
		s,
		"resetCertificateDomains",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetDnsServer() {
	_jsii_.InvokeVoid(
		s,
		"resetDnsServer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetDnsServerPort() {
	_jsii_.InvokeVoid(
		s,
		"resetDnsServerPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetForm() {
	_jsii_.InvokeVoid(
		s,
		"resetForm",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		s,
		"resetHost",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetHttpVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetIsMessageBase64Encoded() {
	_jsii_.InvokeVoid(
		s,
		"resetIsMessageBase64Encoded",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetMessage() {
	_jsii_.InvokeVoid(
		s,
		"resetMessage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		s,
		"resetMethod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetNoSavingResponseBody() {
	_jsii_.InvokeVoid(
		s,
		"resetNoSavingResponseBody",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetNumberOfPackets() {
	_jsii_.InvokeVoid(
		s,
		"resetNumberOfPackets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetPersistCookies() {
	_jsii_.InvokeVoid(
		s,
		"resetPersistCookies",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetPlainProtoFile() {
	_jsii_.InvokeVoid(
		s,
		"resetPlainProtoFile",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		s,
		"resetPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetProtoJsonDescriptor() {
	_jsii_.InvokeVoid(
		s,
		"resetProtoJsonDescriptor",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetServername() {
	_jsii_.InvokeVoid(
		s,
		"resetServername",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		s,
		"resetService",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetShouldTrackHops() {
	_jsii_.InvokeVoid(
		s,
		"resetShouldTrackHops",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestRequestDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

