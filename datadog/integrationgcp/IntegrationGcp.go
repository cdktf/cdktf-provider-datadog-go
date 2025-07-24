// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationgcp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/integrationgcp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/integration_gcp datadog_integration_gcp}.
type IntegrationGcp interface {
	cdktf.TerraformResource
	Automute() interface{}
	SetAutomute(val interface{})
	AutomuteInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientEmail() *string
	SetClientEmail(val *string)
	ClientEmailInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	CloudRunRevisionFilters() *[]*string
	SetCloudRunRevisionFilters(val *[]*string)
	CloudRunRevisionFiltersInput() *[]*string
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
	CspmResourceCollectionEnabled() interface{}
	SetCspmResourceCollectionEnabled(val interface{})
	CspmResourceCollectionEnabledInput() interface{}
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
	HostFilters() *string
	SetHostFilters(val *string)
	HostFiltersInput() *string
	Id() *string
	IsResourceChangeCollectionEnabled() interface{}
	SetIsResourceChangeCollectionEnabled(val interface{})
	IsResourceChangeCollectionEnabledInput() interface{}
	IsSecurityCommandCenterEnabled() interface{}
	SetIsSecurityCommandCenterEnabled(val interface{})
	IsSecurityCommandCenterEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyId() *string
	SetPrivateKeyId(val *string)
	PrivateKeyIdInput() *string
	PrivateKeyInput() *string
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
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
	ResourceCollectionEnabled() interface{}
	SetResourceCollectionEnabled(val interface{})
	ResourceCollectionEnabledInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAutomute()
	ResetCloudRunRevisionFilters()
	ResetCspmResourceCollectionEnabled()
	ResetHostFilters()
	ResetIsResourceChangeCollectionEnabled()
	ResetIsSecurityCommandCenterEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceCollectionEnabled()
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

// The jsii proxy struct for IntegrationGcp
type jsiiProxy_IntegrationGcp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IntegrationGcp) Automute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) AutomuteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) ClientEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) ClientEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) CloudRunRevisionFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloudRunRevisionFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) CloudRunRevisionFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloudRunRevisionFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) CspmResourceCollectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cspmResourceCollectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) CspmResourceCollectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cspmResourceCollectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) HostFilters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) HostFiltersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) IsResourceChangeCollectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isResourceChangeCollectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) IsResourceChangeCollectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isResourceChangeCollectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) IsSecurityCommandCenterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSecurityCommandCenterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) IsSecurityCommandCenterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSecurityCommandCenterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) PrivateKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) PrivateKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) ResourceCollectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceCollectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) ResourceCollectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceCollectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/integration_gcp datadog_integration_gcp} Resource.
func NewIntegrationGcp(scope constructs.Construct, id *string, config *IntegrationGcpConfig) IntegrationGcp {
	_init_.Initialize()

	if err := validateNewIntegrationGcpParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationGcp{}

	_jsii_.Create(
		"@cdktf/provider-datadog.integrationGcp.IntegrationGcp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/integration_gcp datadog_integration_gcp} Resource.
func NewIntegrationGcp_Override(i IntegrationGcp, scope constructs.Construct, id *string, config *IntegrationGcpConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.integrationGcp.IntegrationGcp",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetAutomute(val interface{}) {
	if err := j.validateSetAutomuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automute",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetClientEmail(val *string) {
	if err := j.validateSetClientEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientEmail",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetCloudRunRevisionFilters(val *[]*string) {
	if err := j.validateSetCloudRunRevisionFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudRunRevisionFilters",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetCspmResourceCollectionEnabled(val interface{}) {
	if err := j.validateSetCspmResourceCollectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cspmResourceCollectionEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetHostFilters(val *string) {
	if err := j.validateSetHostFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostFilters",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetIsResourceChangeCollectionEnabled(val interface{}) {
	if err := j.validateSetIsResourceChangeCollectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isResourceChangeCollectionEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetIsSecurityCommandCenterEnabled(val interface{}) {
	if err := j.validateSetIsSecurityCommandCenterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSecurityCommandCenterEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetPrivateKey(val *string) {
	if err := j.validateSetPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetPrivateKeyId(val *string) {
	if err := j.validateSetPrivateKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeyId",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcp)SetResourceCollectionEnabled(val interface{}) {
	if err := j.validateSetResourceCollectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceCollectionEnabled",
		val,
	)
}

// Generates CDKTF code for importing a IntegrationGcp resource upon running "cdktf plan <stack-name>".
func IntegrationGcp_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIntegrationGcp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.integrationGcp.IntegrationGcp",
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
func IntegrationGcp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationGcp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.integrationGcp.IntegrationGcp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationGcp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationGcp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.integrationGcp.IntegrationGcp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationGcp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationGcp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.integrationGcp.IntegrationGcp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IntegrationGcp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.integrationGcp.IntegrationGcp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IntegrationGcp) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IntegrationGcp) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IntegrationGcp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IntegrationGcp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationGcp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IntegrationGcp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IntegrationGcp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IntegrationGcp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IntegrationGcp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IntegrationGcp) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IntegrationGcp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IntegrationGcp) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationGcp) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IntegrationGcp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationGcp) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationGcp) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IntegrationGcp) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationGcp) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IntegrationGcp) ResetAutomute() {
	_jsii_.InvokeVoid(
		i,
		"resetAutomute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcp) ResetCloudRunRevisionFilters() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudRunRevisionFilters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcp) ResetCspmResourceCollectionEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetCspmResourceCollectionEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcp) ResetHostFilters() {
	_jsii_.InvokeVoid(
		i,
		"resetHostFilters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcp) ResetIsResourceChangeCollectionEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetIsResourceChangeCollectionEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcp) ResetIsSecurityCommandCenterEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetIsSecurityCommandCenterEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcp) ResetResourceCollectionEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetResourceCollectionEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationGcp) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationGcp) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationGcp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationGcp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationGcp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

