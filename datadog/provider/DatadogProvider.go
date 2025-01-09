// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs datadog}.
type DatadogProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ApiKey() *string
	SetApiKey(val *string)
	ApiKeyInput() *string
	ApiUrl() *string
	SetApiUrl(val *string)
	ApiUrlInput() *string
	AppKey() *string
	SetAppKey(val *string)
	AppKeyInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	DefaultTags() *DatadogProviderDefaultTags
	SetDefaultTags(val *DatadogProviderDefaultTags)
	DefaultTagsInput() *DatadogProviderDefaultTags
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpClientRetryBackoffBase() *float64
	SetHttpClientRetryBackoffBase(val *float64)
	HttpClientRetryBackoffBaseInput() *float64
	HttpClientRetryBackoffMultiplier() *float64
	SetHttpClientRetryBackoffMultiplier(val *float64)
	HttpClientRetryBackoffMultiplierInput() *float64
	HttpClientRetryEnabled() *string
	SetHttpClientRetryEnabled(val *string)
	HttpClientRetryEnabledInput() *string
	HttpClientRetryMaxRetries() *float64
	SetHttpClientRetryMaxRetries(val *float64)
	HttpClientRetryMaxRetriesInput() *float64
	HttpClientRetryTimeout() *float64
	SetHttpClientRetryTimeout(val *float64)
	HttpClientRetryTimeoutInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Validate() *string
	SetValidate(val *string)
	ValidateInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetApiKey()
	ResetApiUrl()
	ResetAppKey()
	ResetDefaultTags()
	ResetHttpClientRetryBackoffBase()
	ResetHttpClientRetryBackoffMultiplier()
	ResetHttpClientRetryEnabled()
	ResetHttpClientRetryMaxRetries()
	ResetHttpClientRetryTimeout()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetValidate()
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

// The jsii proxy struct for DatadogProvider
type jsiiProxy_DatadogProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_DatadogProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) ApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) ApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) ApiUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) AppKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) AppKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) DefaultTags() *DatadogProviderDefaultTags {
	var returns *DatadogProviderDefaultTags
	_jsii_.Get(
		j,
		"defaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) DefaultTagsInput() *DatadogProviderDefaultTags {
	var returns *DatadogProviderDefaultTags
	_jsii_.Get(
		j,
		"defaultTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) HttpClientRetryBackoffBase() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpClientRetryBackoffBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) HttpClientRetryBackoffBaseInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpClientRetryBackoffBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) HttpClientRetryBackoffMultiplier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpClientRetryBackoffMultiplier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) HttpClientRetryBackoffMultiplierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpClientRetryBackoffMultiplierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) HttpClientRetryEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpClientRetryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) HttpClientRetryEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpClientRetryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) HttpClientRetryMaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpClientRetryMaxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) HttpClientRetryMaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpClientRetryMaxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) HttpClientRetryTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpClientRetryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) HttpClientRetryTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpClientRetryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) Validate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogProvider) ValidateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validateInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs datadog} Resource.
func NewDatadogProvider(scope constructs.Construct, id *string, config *DatadogProviderConfig) DatadogProvider {
	_init_.Initialize()

	if err := validateNewDatadogProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatadogProvider{}

	_jsii_.Create(
		"@cdktf/provider-datadog.provider.DatadogProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs datadog} Resource.
func NewDatadogProvider_Override(d DatadogProvider, scope constructs.Construct, id *string, config *DatadogProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.provider.DatadogProvider",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatadogProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_DatadogProvider)SetApiKey(val *string) {
	_jsii_.Set(
		j,
		"apiKey",
		val,
	)
}

func (j *jsiiProxy_DatadogProvider)SetApiUrl(val *string) {
	_jsii_.Set(
		j,
		"apiUrl",
		val,
	)
}

func (j *jsiiProxy_DatadogProvider)SetAppKey(val *string) {
	_jsii_.Set(
		j,
		"appKey",
		val,
	)
}

func (j *jsiiProxy_DatadogProvider)SetDefaultTags(val *DatadogProviderDefaultTags) {
	if err := j.validateSetDefaultTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTags",
		val,
	)
}

func (j *jsiiProxy_DatadogProvider)SetHttpClientRetryBackoffBase(val *float64) {
	_jsii_.Set(
		j,
		"httpClientRetryBackoffBase",
		val,
	)
}

func (j *jsiiProxy_DatadogProvider)SetHttpClientRetryBackoffMultiplier(val *float64) {
	_jsii_.Set(
		j,
		"httpClientRetryBackoffMultiplier",
		val,
	)
}

func (j *jsiiProxy_DatadogProvider)SetHttpClientRetryEnabled(val *string) {
	_jsii_.Set(
		j,
		"httpClientRetryEnabled",
		val,
	)
}

func (j *jsiiProxy_DatadogProvider)SetHttpClientRetryMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"httpClientRetryMaxRetries",
		val,
	)
}

func (j *jsiiProxy_DatadogProvider)SetHttpClientRetryTimeout(val *float64) {
	_jsii_.Set(
		j,
		"httpClientRetryTimeout",
		val,
	)
}

func (j *jsiiProxy_DatadogProvider)SetValidate(val *string) {
	_jsii_.Set(
		j,
		"validate",
		val,
	)
}

// Generates CDKTF code for importing a DatadogProvider resource upon running "cdktf plan <stack-name>".
func DatadogProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatadogProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.provider.DatadogProvider",
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
func DatadogProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatadogProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.provider.DatadogProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatadogProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatadogProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.provider.DatadogProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatadogProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatadogProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.provider.DatadogProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatadogProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.provider.DatadogProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatadogProvider) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatadogProvider) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatadogProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		d,
		"resetAlias",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatadogProvider) ResetApiKey() {
	_jsii_.InvokeVoid(
		d,
		"resetApiKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatadogProvider) ResetApiUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetApiUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatadogProvider) ResetAppKey() {
	_jsii_.InvokeVoid(
		d,
		"resetAppKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatadogProvider) ResetDefaultTags() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatadogProvider) ResetHttpClientRetryBackoffBase() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpClientRetryBackoffBase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatadogProvider) ResetHttpClientRetryBackoffMultiplier() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpClientRetryBackoffMultiplier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatadogProvider) ResetHttpClientRetryEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpClientRetryEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatadogProvider) ResetHttpClientRetryMaxRetries() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpClientRetryMaxRetries",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatadogProvider) ResetHttpClientRetryTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpClientRetryTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatadogProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatadogProvider) ResetValidate() {
	_jsii_.InvokeVoid(
		d,
		"resetValidate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatadogProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

