// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/syntheticstest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/synthetics_test datadog_synthetics_test}.
type SyntheticsTest interface {
	cdktf.TerraformResource
	ApiStep() SyntheticsTestApiStepList
	ApiStepInput() interface{}
	Assertion() SyntheticsTestAssertionList
	AssertionInput() interface{}
	BrowserStep() SyntheticsTestBrowserStepList
	BrowserStepInput() interface{}
	BrowserVariable() SyntheticsTestBrowserVariableList
	BrowserVariableInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigVariable() SyntheticsTestConfigVariableList
	ConfigVariableInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeviceIds() *[]*string
	SetDeviceIds(val *[]*string)
	DeviceIdsInput() *[]*string
	ForceDeleteDependencies() interface{}
	SetForceDeleteDependencies(val interface{})
	ForceDeleteDependenciesInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locations() *[]*string
	SetLocations(val *[]*string)
	LocationsInput() *[]*string
	Message() *string
	SetMessage(val *string)
	MessageInput() *string
	MonitorId() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OptionsList() SyntheticsTestOptionsListStructOutputReference
	OptionsListInput() *SyntheticsTestOptionsListStruct
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RequestBasicauth() SyntheticsTestRequestBasicauthOutputReference
	RequestBasicauthInput() *SyntheticsTestRequestBasicauth
	RequestClientCertificate() SyntheticsTestRequestClientCertificateOutputReference
	RequestClientCertificateInput() *SyntheticsTestRequestClientCertificate
	RequestDefinition() SyntheticsTestRequestDefinitionOutputReference
	RequestDefinitionInput() *SyntheticsTestRequestDefinition
	RequestFile() SyntheticsTestRequestFileList
	RequestFileInput() interface{}
	RequestHeaders() *map[string]*string
	SetRequestHeaders(val *map[string]*string)
	RequestHeadersInput() *map[string]*string
	RequestMetadata() *map[string]*string
	SetRequestMetadata(val *map[string]*string)
	RequestMetadataInput() *map[string]*string
	RequestProxy() SyntheticsTestRequestProxyOutputReference
	RequestProxyInput() *SyntheticsTestRequestProxy
	RequestQuery() *map[string]*string
	SetRequestQuery(val *map[string]*string)
	RequestQueryInput() *map[string]*string
	SetCookie() *string
	SetSetCookie(val *string)
	SetCookieInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Subtype() *string
	SetSubtype(val *string)
	SubtypeInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VariablesFromScript() *string
	SetVariablesFromScript(val *string)
	VariablesFromScriptInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutApiStep(value interface{})
	PutAssertion(value interface{})
	PutBrowserStep(value interface{})
	PutBrowserVariable(value interface{})
	PutConfigVariable(value interface{})
	PutOptionsList(value *SyntheticsTestOptionsListStruct)
	PutRequestBasicauth(value *SyntheticsTestRequestBasicauth)
	PutRequestClientCertificate(value *SyntheticsTestRequestClientCertificate)
	PutRequestDefinition(value *SyntheticsTestRequestDefinition)
	PutRequestFile(value interface{})
	PutRequestProxy(value *SyntheticsTestRequestProxy)
	ResetApiStep()
	ResetAssertion()
	ResetBrowserStep()
	ResetBrowserVariable()
	ResetConfigVariable()
	ResetDeviceIds()
	ResetForceDeleteDependencies()
	ResetId()
	ResetMessage()
	ResetOptionsList()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestBasicauth()
	ResetRequestClientCertificate()
	ResetRequestDefinition()
	ResetRequestFile()
	ResetRequestHeaders()
	ResetRequestMetadata()
	ResetRequestProxy()
	ResetRequestQuery()
	ResetSetCookie()
	ResetSubtype()
	ResetTags()
	ResetVariablesFromScript()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SyntheticsTest
type jsiiProxy_SyntheticsTest struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SyntheticsTest) ApiStep() SyntheticsTestApiStepList {
	var returns SyntheticsTestApiStepList
	_jsii_.Get(
		j,
		"apiStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) ApiStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Assertion() SyntheticsTestAssertionList {
	var returns SyntheticsTestAssertionList
	_jsii_.Get(
		j,
		"assertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) AssertionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) BrowserStep() SyntheticsTestBrowserStepList {
	var returns SyntheticsTestBrowserStepList
	_jsii_.Get(
		j,
		"browserStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) BrowserStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browserStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) BrowserVariable() SyntheticsTestBrowserVariableList {
	var returns SyntheticsTestBrowserVariableList
	_jsii_.Get(
		j,
		"browserVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) BrowserVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"browserVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) ConfigVariable() SyntheticsTestConfigVariableList {
	var returns SyntheticsTestConfigVariableList
	_jsii_.Get(
		j,
		"configVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) ConfigVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) DeviceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deviceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) DeviceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deviceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) ForceDeleteDependencies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) ForceDeleteDependenciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteDependenciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Locations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) LocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) MessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) MonitorId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) OptionsList() SyntheticsTestOptionsListStructOutputReference {
	var returns SyntheticsTestOptionsListStructOutputReference
	_jsii_.Get(
		j,
		"optionsList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) OptionsListInput() *SyntheticsTestOptionsListStruct {
	var returns *SyntheticsTestOptionsListStruct
	_jsii_.Get(
		j,
		"optionsListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestBasicauth() SyntheticsTestRequestBasicauthOutputReference {
	var returns SyntheticsTestRequestBasicauthOutputReference
	_jsii_.Get(
		j,
		"requestBasicauth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestBasicauthInput() *SyntheticsTestRequestBasicauth {
	var returns *SyntheticsTestRequestBasicauth
	_jsii_.Get(
		j,
		"requestBasicauthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestClientCertificate() SyntheticsTestRequestClientCertificateOutputReference {
	var returns SyntheticsTestRequestClientCertificateOutputReference
	_jsii_.Get(
		j,
		"requestClientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestClientCertificateInput() *SyntheticsTestRequestClientCertificate {
	var returns *SyntheticsTestRequestClientCertificate
	_jsii_.Get(
		j,
		"requestClientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestDefinition() SyntheticsTestRequestDefinitionOutputReference {
	var returns SyntheticsTestRequestDefinitionOutputReference
	_jsii_.Get(
		j,
		"requestDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestDefinitionInput() *SyntheticsTestRequestDefinition {
	var returns *SyntheticsTestRequestDefinition
	_jsii_.Get(
		j,
		"requestDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestFile() SyntheticsTestRequestFileList {
	var returns SyntheticsTestRequestFileList
	_jsii_.Get(
		j,
		"requestFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestProxy() SyntheticsTestRequestProxyOutputReference {
	var returns SyntheticsTestRequestProxyOutputReference
	_jsii_.Get(
		j,
		"requestProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestProxyInput() *SyntheticsTestRequestProxy {
	var returns *SyntheticsTestRequestProxy
	_jsii_.Get(
		j,
		"requestProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestQuery() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) RequestQueryInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) SetCookie() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) SetCookieInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Subtype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) SubtypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) VariablesFromScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variablesFromScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTest) VariablesFromScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variablesFromScriptInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/synthetics_test datadog_synthetics_test} Resource.
func NewSyntheticsTest(scope constructs.Construct, id *string, config *SyntheticsTestConfig) SyntheticsTest {
	_init_.Initialize()

	if err := validateNewSyntheticsTestParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTest{}

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTest",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/synthetics_test datadog_synthetics_test} Resource.
func NewSyntheticsTest_Override(s SyntheticsTest, scope constructs.Construct, id *string, config *SyntheticsTestConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTest",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetDeviceIds(val *[]*string) {
	if err := j.validateSetDeviceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceIds",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetForceDeleteDependencies(val interface{}) {
	if err := j.validateSetForceDeleteDependenciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDeleteDependencies",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetLocations(val *[]*string) {
	if err := j.validateSetLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locations",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetMessage(val *string) {
	if err := j.validateSetMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetRequestHeaders(val *map[string]*string) {
	if err := j.validateSetRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestHeaders",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetRequestMetadata(val *map[string]*string) {
	if err := j.validateSetRequestMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestMetadata",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetRequestQuery(val *map[string]*string) {
	if err := j.validateSetRequestQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestQuery",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetSetCookie(val *string) {
	if err := j.validateSetSetCookieParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setCookie",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetSubtype(val *string) {
	if err := j.validateSetSubtypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subtype",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTest)SetVariablesFromScript(val *string) {
	if err := j.validateSetVariablesFromScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variablesFromScript",
		val,
	)
}

// Generates CDKTF code for importing a SyntheticsTest resource upon running "cdktf plan <stack-name>".
func SyntheticsTest_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSyntheticsTest_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTest",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SyntheticsTest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsTest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsTest_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsTest_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTest",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsTest_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsTest_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTest",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SyntheticsTest_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTest",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SyntheticsTest) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SyntheticsTest) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SyntheticsTest) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SyntheticsTest) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsTest) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SyntheticsTest) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SyntheticsTest) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SyntheticsTest) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SyntheticsTest) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SyntheticsTest) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SyntheticsTest) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SyntheticsTest) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTest) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SyntheticsTest) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTest) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsTest) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SyntheticsTest) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsTest) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SyntheticsTest) PutApiStep(value interface{}) {
	if err := s.validatePutApiStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putApiStep",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTest) PutAssertion(value interface{}) {
	if err := s.validatePutAssertionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAssertion",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTest) PutBrowserStep(value interface{}) {
	if err := s.validatePutBrowserStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBrowserStep",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTest) PutBrowserVariable(value interface{}) {
	if err := s.validatePutBrowserVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBrowserVariable",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTest) PutConfigVariable(value interface{}) {
	if err := s.validatePutConfigVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putConfigVariable",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTest) PutOptionsList(value *SyntheticsTestOptionsListStruct) {
	if err := s.validatePutOptionsListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOptionsList",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTest) PutRequestBasicauth(value *SyntheticsTestRequestBasicauth) {
	if err := s.validatePutRequestBasicauthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRequestBasicauth",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTest) PutRequestClientCertificate(value *SyntheticsTestRequestClientCertificate) {
	if err := s.validatePutRequestClientCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRequestClientCertificate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTest) PutRequestDefinition(value *SyntheticsTestRequestDefinition) {
	if err := s.validatePutRequestDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRequestDefinition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTest) PutRequestFile(value interface{}) {
	if err := s.validatePutRequestFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRequestFile",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTest) PutRequestProxy(value *SyntheticsTestRequestProxy) {
	if err := s.validatePutRequestProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRequestProxy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetApiStep() {
	_jsii_.InvokeVoid(
		s,
		"resetApiStep",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetAssertion() {
	_jsii_.InvokeVoid(
		s,
		"resetAssertion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetBrowserStep() {
	_jsii_.InvokeVoid(
		s,
		"resetBrowserStep",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetBrowserVariable() {
	_jsii_.InvokeVoid(
		s,
		"resetBrowserVariable",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetConfigVariable() {
	_jsii_.InvokeVoid(
		s,
		"resetConfigVariable",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetDeviceIds() {
	_jsii_.InvokeVoid(
		s,
		"resetDeviceIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetForceDeleteDependencies() {
	_jsii_.InvokeVoid(
		s,
		"resetForceDeleteDependencies",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetMessage() {
	_jsii_.InvokeVoid(
		s,
		"resetMessage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetOptionsList() {
	_jsii_.InvokeVoid(
		s,
		"resetOptionsList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetRequestBasicauth() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestBasicauth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetRequestClientCertificate() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestClientCertificate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetRequestDefinition() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestDefinition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetRequestFile() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestFile",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetRequestHeaders() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestHeaders",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetRequestMetadata() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestMetadata",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetRequestProxy() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestProxy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetRequestQuery() {
	_jsii_.InvokeVoid(
		s,
		"resetRequestQuery",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetSetCookie() {
	_jsii_.InvokeVoid(
		s,
		"resetSetCookie",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetSubtype() {
	_jsii_.InvokeVoid(
		s,
		"resetSubtype",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) ResetVariablesFromScript() {
	_jsii_.InvokeVoid(
		s,
		"resetVariablesFromScript",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTest) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTest) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTest) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTest) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTest) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

