// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/integrationawsaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/integration_aws_account datadog_integration_aws_account}.
type IntegrationAwsAccount interface {
	cdktf.TerraformResource
	AccountTags() *[]*string
	SetAccountTags(val *[]*string)
	AccountTagsInput() *[]*string
	AuthConfig() IntegrationAwsAccountAuthConfigOutputReference
	AuthConfigInput() interface{}
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	AwsAccountIdInput() *string
	AwsPartition() *string
	SetAwsPartition(val *string)
	AwsPartitionInput() *string
	AwsRegions() IntegrationAwsAccountAwsRegionsOutputReference
	AwsRegionsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogsConfig() IntegrationAwsAccountLogsConfigOutputReference
	LogsConfigInput() interface{}
	MetricsConfig() IntegrationAwsAccountMetricsConfigOutputReference
	MetricsConfigInput() interface{}
	// The tree node.
	Node() constructs.Node
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
	ResourcesConfig() IntegrationAwsAccountResourcesConfigOutputReference
	ResourcesConfigInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TracesConfig() IntegrationAwsAccountTracesConfigOutputReference
	TracesConfigInput() interface{}
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
	PutAuthConfig(value *IntegrationAwsAccountAuthConfig)
	PutAwsRegions(value *IntegrationAwsAccountAwsRegions)
	PutLogsConfig(value *IntegrationAwsAccountLogsConfig)
	PutMetricsConfig(value *IntegrationAwsAccountMetricsConfig)
	PutResourcesConfig(value *IntegrationAwsAccountResourcesConfig)
	PutTracesConfig(value *IntegrationAwsAccountTracesConfig)
	ResetAccountTags()
	ResetAuthConfig()
	ResetAwsRegions()
	ResetLogsConfig()
	ResetMetricsConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourcesConfig()
	ResetTracesConfig()
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

// The jsii proxy struct for IntegrationAwsAccount
type jsiiProxy_IntegrationAwsAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IntegrationAwsAccount) AccountTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) AccountTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) AuthConfig() IntegrationAwsAccountAuthConfigOutputReference {
	var returns IntegrationAwsAccountAuthConfigOutputReference
	_jsii_.Get(
		j,
		"authConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) AuthConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) AwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) AwsPartition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsPartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) AwsPartitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsPartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) AwsRegions() IntegrationAwsAccountAwsRegionsOutputReference {
	var returns IntegrationAwsAccountAwsRegionsOutputReference
	_jsii_.Get(
		j,
		"awsRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) AwsRegionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) LogsConfig() IntegrationAwsAccountLogsConfigOutputReference {
	var returns IntegrationAwsAccountLogsConfigOutputReference
	_jsii_.Get(
		j,
		"logsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) LogsConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) MetricsConfig() IntegrationAwsAccountMetricsConfigOutputReference {
	var returns IntegrationAwsAccountMetricsConfigOutputReference
	_jsii_.Get(
		j,
		"metricsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) MetricsConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) ResourcesConfig() IntegrationAwsAccountResourcesConfigOutputReference {
	var returns IntegrationAwsAccountResourcesConfigOutputReference
	_jsii_.Get(
		j,
		"resourcesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) ResourcesConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) TracesConfig() IntegrationAwsAccountTracesConfigOutputReference {
	var returns IntegrationAwsAccountTracesConfigOutputReference
	_jsii_.Get(
		j,
		"tracesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccount) TracesConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracesConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/integration_aws_account datadog_integration_aws_account} Resource.
func NewIntegrationAwsAccount(scope constructs.Construct, id *string, config *IntegrationAwsAccountConfig) IntegrationAwsAccount {
	_init_.Initialize()

	if err := validateNewIntegrationAwsAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationAwsAccount{}

	_jsii_.Create(
		"@cdktf/provider-datadog.integrationAwsAccount.IntegrationAwsAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/integration_aws_account datadog_integration_aws_account} Resource.
func NewIntegrationAwsAccount_Override(i IntegrationAwsAccount, scope constructs.Construct, id *string, config *IntegrationAwsAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.integrationAwsAccount.IntegrationAwsAccount",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IntegrationAwsAccount)SetAccountTags(val *[]*string) {
	if err := j.validateSetAccountTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountTags",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccount)SetAwsAccountId(val *string) {
	if err := j.validateSetAwsAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccount)SetAwsPartition(val *string) {
	if err := j.validateSetAwsPartitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsPartition",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccount)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccount)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccount)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a IntegrationAwsAccount resource upon running "cdktf plan <stack-name>".
func IntegrationAwsAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIntegrationAwsAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.integrationAwsAccount.IntegrationAwsAccount",
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
func IntegrationAwsAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationAwsAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.integrationAwsAccount.IntegrationAwsAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationAwsAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationAwsAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.integrationAwsAccount.IntegrationAwsAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationAwsAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationAwsAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.integrationAwsAccount.IntegrationAwsAccount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IntegrationAwsAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.integrationAwsAccount.IntegrationAwsAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) PutAuthConfig(value *IntegrationAwsAccountAuthConfig) {
	if err := i.validatePutAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAuthConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) PutAwsRegions(value *IntegrationAwsAccountAwsRegions) {
	if err := i.validatePutAwsRegionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAwsRegions",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) PutLogsConfig(value *IntegrationAwsAccountLogsConfig) {
	if err := i.validatePutLogsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLogsConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) PutMetricsConfig(value *IntegrationAwsAccountMetricsConfig) {
	if err := i.validatePutMetricsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMetricsConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) PutResourcesConfig(value *IntegrationAwsAccountResourcesConfig) {
	if err := i.validatePutResourcesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putResourcesConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) PutTracesConfig(value *IntegrationAwsAccountTracesConfig) {
	if err := i.validatePutTracesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTracesConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) ResetAccountTags() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) ResetAuthConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) ResetAwsRegions() {
	_jsii_.InvokeVoid(
		i,
		"resetAwsRegions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) ResetLogsConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetLogsConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) ResetMetricsConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetMetricsConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) ResetResourcesConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetResourcesConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) ResetTracesConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetTracesConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

