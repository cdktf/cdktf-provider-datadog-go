// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/syntheticstest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestApiStepOutputReference interface {
	cdktf.ComplexObject
	AllowFailure() interface{}
	SetAllowFailure(val interface{})
	AllowFailureInput() interface{}
	Assertion() SyntheticsTestApiStepAssertionList
	AssertionInput() interface{}
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
	ExitIfSucceed() interface{}
	SetExitIfSucceed(val interface{})
	ExitIfSucceedInput() interface{}
	ExtractedValue() SyntheticsTestApiStepExtractedValueList
	ExtractedValueInput() interface{}
	ExtractedValuesFromScript() *string
	SetExtractedValuesFromScript(val *string)
	ExtractedValuesFromScriptInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsCritical() interface{}
	SetIsCritical(val interface{})
	IsCriticalInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	RequestBasicauth() SyntheticsTestApiStepRequestBasicauthOutputReference
	RequestBasicauthInput() *SyntheticsTestApiStepRequestBasicauth
	RequestClientCertificate() SyntheticsTestApiStepRequestClientCertificateOutputReference
	RequestClientCertificateInput() *SyntheticsTestApiStepRequestClientCertificate
	RequestDefinition() SyntheticsTestApiStepRequestDefinitionOutputReference
	RequestDefinitionInput() *SyntheticsTestApiStepRequestDefinition
	RequestFile() SyntheticsTestApiStepRequestFileList
	RequestFileInput() interface{}
	RequestHeaders() *map[string]*string
	SetRequestHeaders(val *map[string]*string)
	RequestHeadersInput() *map[string]*string
	RequestMetadata() *map[string]*string
	SetRequestMetadata(val *map[string]*string)
	RequestMetadataInput() *map[string]*string
	RequestProxy() SyntheticsTestApiStepRequestProxyOutputReference
	RequestProxyInput() *SyntheticsTestApiStepRequestProxy
	RequestQuery() *map[string]*string
	SetRequestQuery(val *map[string]*string)
	RequestQueryInput() *map[string]*string
	Retry() SyntheticsTestApiStepRetryOutputReference
	RetryInput() *SyntheticsTestApiStepRetry
	Subtype() *string
	SetSubtype(val *string)
	SubtypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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
	PutAssertion(value interface{})
	PutExtractedValue(value interface{})
	PutRequestBasicauth(value *SyntheticsTestApiStepRequestBasicauth)
	PutRequestClientCertificate(value *SyntheticsTestApiStepRequestClientCertificate)
	PutRequestDefinition(value *SyntheticsTestApiStepRequestDefinition)
	PutRequestFile(value interface{})
	PutRequestProxy(value *SyntheticsTestApiStepRequestProxy)
	PutRetry(value *SyntheticsTestApiStepRetry)
	ResetAllowFailure()
	ResetAssertion()
	ResetExitIfSucceed()
	ResetExtractedValue()
	ResetExtractedValuesFromScript()
	ResetIsCritical()
	ResetRequestBasicauth()
	ResetRequestClientCertificate()
	ResetRequestDefinition()
	ResetRequestFile()
	ResetRequestHeaders()
	ResetRequestMetadata()
	ResetRequestProxy()
	ResetRequestQuery()
	ResetRetry()
	ResetSubtype()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SyntheticsTestApiStepOutputReference
type jsiiProxy_SyntheticsTestApiStepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) AllowFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) AllowFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) Assertion() SyntheticsTestApiStepAssertionList {
	var returns SyntheticsTestApiStepAssertionList
	_jsii_.Get(
		j,
		"assertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) AssertionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) ExitIfSucceed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exitIfSucceed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) ExitIfSucceedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exitIfSucceedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) ExtractedValue() SyntheticsTestApiStepExtractedValueList {
	var returns SyntheticsTestApiStepExtractedValueList
	_jsii_.Get(
		j,
		"extractedValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) ExtractedValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extractedValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) ExtractedValuesFromScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extractedValuesFromScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) ExtractedValuesFromScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extractedValuesFromScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) IsCritical() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCritical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) IsCriticalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCriticalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestBasicauth() SyntheticsTestApiStepRequestBasicauthOutputReference {
	var returns SyntheticsTestApiStepRequestBasicauthOutputReference
	_jsii_.Get(
		j,
		"requestBasicauth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestBasicauthInput() *SyntheticsTestApiStepRequestBasicauth {
	var returns *SyntheticsTestApiStepRequestBasicauth
	_jsii_.Get(
		j,
		"requestBasicauthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestClientCertificate() SyntheticsTestApiStepRequestClientCertificateOutputReference {
	var returns SyntheticsTestApiStepRequestClientCertificateOutputReference
	_jsii_.Get(
		j,
		"requestClientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestClientCertificateInput() *SyntheticsTestApiStepRequestClientCertificate {
	var returns *SyntheticsTestApiStepRequestClientCertificate
	_jsii_.Get(
		j,
		"requestClientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestDefinition() SyntheticsTestApiStepRequestDefinitionOutputReference {
	var returns SyntheticsTestApiStepRequestDefinitionOutputReference
	_jsii_.Get(
		j,
		"requestDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestDefinitionInput() *SyntheticsTestApiStepRequestDefinition {
	var returns *SyntheticsTestApiStepRequestDefinition
	_jsii_.Get(
		j,
		"requestDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestFile() SyntheticsTestApiStepRequestFileList {
	var returns SyntheticsTestApiStepRequestFileList
	_jsii_.Get(
		j,
		"requestFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestProxy() SyntheticsTestApiStepRequestProxyOutputReference {
	var returns SyntheticsTestApiStepRequestProxyOutputReference
	_jsii_.Get(
		j,
		"requestProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestProxyInput() *SyntheticsTestApiStepRequestProxy {
	var returns *SyntheticsTestApiStepRequestProxy
	_jsii_.Get(
		j,
		"requestProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestQuery() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RequestQueryInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) Retry() SyntheticsTestApiStepRetryOutputReference {
	var returns SyntheticsTestApiStepRetryOutputReference
	_jsii_.Get(
		j,
		"retry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) RetryInput() *SyntheticsTestApiStepRetry {
	var returns *SyntheticsTestApiStepRetry
	_jsii_.Get(
		j,
		"retryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) Subtype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) SubtypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewSyntheticsTestApiStepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SyntheticsTestApiStepOutputReference {
	_init_.Initialize()

	if err := validateNewSyntheticsTestApiStepOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTestApiStepOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestApiStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSyntheticsTestApiStepOutputReference_Override(s SyntheticsTestApiStepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestApiStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetAllowFailure(val interface{}) {
	if err := j.validateSetAllowFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowFailure",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetExitIfSucceed(val interface{}) {
	if err := j.validateSetExitIfSucceedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exitIfSucceed",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetExtractedValuesFromScript(val *string) {
	if err := j.validateSetExtractedValuesFromScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extractedValuesFromScript",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetIsCritical(val interface{}) {
	if err := j.validateSetIsCriticalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCritical",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetRequestHeaders(val *map[string]*string) {
	if err := j.validateSetRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestHeaders",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetRequestMetadata(val *map[string]*string) {
	if err := j.validateSetRequestMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestMetadata",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetRequestQuery(val *map[string]*string) {
	if err := j.validateSetRequestQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestQuery",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetSubtype(val *string) {
	if err := j.validateSetSubtypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subtype",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestApiStepOutputReference)SetValue(val *float64) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) PutAssertion(value interface{}) {
	if err := s.validatePutAssertionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAssertion",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) PutExtractedValue(value interface{}) {
	if err := s.validatePutExtractedValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExtractedValue",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) PutRequestBasicauth(value *SyntheticsTestApiStepRequestBasicauth) {
	if err := s.validatePutRequestBasicauthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRequestBasicauth",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) PutRequestClientCertificate(value *SyntheticsTestApiStepRequestClientCertificate) {
	if err := s.validatePutRequestClientCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRequestClientCertificate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) PutRequestDefinition(value *SyntheticsTestApiStepRequestDefinition) {
	if err := s.validatePutRequestDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRequestDefinition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) PutRequestFile(value interface{}) {
	if err := s.validatePutRequestFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRequestFile",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) PutRequestProxy(value *SyntheticsTestApiStepRequestProxy) {
	if err := s.validatePutRequestProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRequestProxy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) PutRetry(value *SyntheticsTestApiStepRetry) {
	if err := s.validatePutRetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRetry",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetAllowFailure() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowFailure",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetAssertion() {
	_jsii_.InvokeVoid(
		s,
		"resetAssertion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetExitIfSucceed() {
	_jsii_.InvokeVoid(
		s,
		"resetExitIfSucceed",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetExtractedValue() {
	_jsii_.InvokeVoid(
		s,
		"resetExtractedValue",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetExtractedValuesFromScript() {
	_jsii_.InvokeVoid(
		s,
		"resetExtractedValuesFromScript",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetIsCritical() {
	_jsii_.InvokeVoid(
		s,
		"resetIsCritical",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetRequestBasicauth() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestBasicauth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetRequestClientCertificate() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestClientCertificate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetRequestDefinition() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestDefinition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetRequestFile() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestFile",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetRequestHeaders() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestHeaders",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetRequestMetadata() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestMetadata",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetRequestProxy() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestProxy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetRequestQuery() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestQuery",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetRetry() {
	_jsii_.InvokeVoid(
		s,
		"resetRetry",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetSubtype() {
	_jsii_.InvokeVoid(
		s,
		"resetSubtype",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		s,
		"resetValue",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SyntheticsTestApiStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

