package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v4/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v4/syntheticstest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestApiStepRequestDefinitionOutputReference interface {
	cdktf.ComplexObject
	AllowInsecure() interface{}
	SetAllowInsecure(val interface{})
	AllowInsecureInput() interface{}
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
	DnsServerPort() *float64
	SetDnsServerPort(val *float64)
	DnsServerPortInput() *float64
	FollowRedirects() interface{}
	SetFollowRedirects(val interface{})
	FollowRedirectsInput() interface{}
	// Experimental.
	Fqn() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() *SyntheticsTestApiStepRequestDefinition
	SetInternalValue(val *SyntheticsTestApiStepRequestDefinition)
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
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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
	ResetAllowInsecure()
	ResetBody()
	ResetBodyType()
	ResetCallType()
	ResetCertificateDomains()
	ResetDnsServer()
	ResetDnsServerPort()
	ResetFollowRedirects()
	ResetHost()
	ResetMessage()
	ResetMethod()
	ResetNoSavingResponseBody()
	ResetNumberOfPackets()
	ResetPort()
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

// The jsii proxy struct for SyntheticsTestApiStepRequestDefinitionOutputReference
type jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) AllowInsecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) AllowInsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) BodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) BodyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) BodyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) CallType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) CallTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) CertificateDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) CertificateDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"certificateDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) DnsServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) DnsServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) DnsServerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dnsServerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) DnsServerPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dnsServerPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) FollowRedirects() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"followRedirects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) FollowRedirectsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"followRedirectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) InternalValue() *SyntheticsTestApiStepRequestDefinition {
	var returns *SyntheticsTestApiStepRequestDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) MessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) MethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) NoSavingResponseBody() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSavingResponseBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) NoSavingResponseBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSavingResponseBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) NumberOfPackets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfPackets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) NumberOfPacketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfPacketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) Servername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ServernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ShouldTrackHops() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTrackHops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ShouldTrackHopsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldTrackHopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewSyntheticsTestApiStepRequestDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SyntheticsTestApiStepRequestDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewSyntheticsTestApiStepRequestDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestApiStepRequestDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSyntheticsTestApiStepRequestDefinitionOutputReference_Override(s SyntheticsTestApiStepRequestDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestApiStepRequestDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetAllowInsecure(val interface{}) {
	if err := j.validateSetAllowInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInsecure",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetBody(val *string) {
	if err := j.validateSetBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetBodyType(val *string) {
	if err := j.validateSetBodyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bodyType",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetCallType(val *string) {
	if err := j.validateSetCallTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callType",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetCertificateDomains(val *[]*string) {
	if err := j.validateSetCertificateDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateDomains",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetDnsServer(val *string) {
	if err := j.validateSetDnsServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServer",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetDnsServerPort(val *float64) {
	if err := j.validateSetDnsServerPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServerPort",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetFollowRedirects(val interface{}) {
	if err := j.validateSetFollowRedirectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"followRedirects",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetInternalValue(val *SyntheticsTestApiStepRequestDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetMessage(val *string) {
	if err := j.validateSetMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetNoSavingResponseBody(val interface{}) {
	if err := j.validateSetNoSavingResponseBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSavingResponseBody",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetNumberOfPackets(val *float64) {
	if err := j.validateSetNumberOfPacketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfPackets",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetServername(val *string) {
	if err := j.validateSetServernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servername",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetShouldTrackHops(val interface{}) {
	if err := j.validateSetShouldTrackHopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldTrackHops",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetAllowInsecure() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowInsecure",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		s,
		"resetBody",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetBodyType() {
	_jsii_.InvokeVoid(
		s,
		"resetBodyType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetCallType() {
	_jsii_.InvokeVoid(
		s,
		"resetCallType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetCertificateDomains() {
	_jsii_.InvokeVoid(
		s,
		"resetCertificateDomains",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetDnsServer() {
	_jsii_.InvokeVoid(
		s,
		"resetDnsServer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetDnsServerPort() {
	_jsii_.InvokeVoid(
		s,
		"resetDnsServerPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetFollowRedirects() {
	_jsii_.InvokeVoid(
		s,
		"resetFollowRedirects",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		s,
		"resetHost",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetMessage() {
	_jsii_.InvokeVoid(
		s,
		"resetMessage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		s,
		"resetMethod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetNoSavingResponseBody() {
	_jsii_.InvokeVoid(
		s,
		"resetNoSavingResponseBody",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetNumberOfPackets() {
	_jsii_.InvokeVoid(
		s,
		"resetNumberOfPackets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		s,
		"resetPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetServername() {
	_jsii_.InvokeVoid(
		s,
		"resetServername",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		s,
		"resetService",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetShouldTrackHops() {
	_jsii_.InvokeVoid(
		s,
		"resetShouldTrackHops",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SyntheticsTestApiStepRequestDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

