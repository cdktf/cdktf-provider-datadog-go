// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafexclusionfilter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/appsecwafexclusionfilter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/appsec_waf_exclusion_filter datadog_appsec_waf_exclusion_filter}.
type AppsecWafExclusionFilter interface {
	cdktf.TerraformResource
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EventQuery() *string
	SetEventQuery(val *string)
	EventQueryInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IpList() *[]*string
	SetIpList(val *[]*string)
	IpListInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OnMatch() *string
	SetOnMatch(val *string)
	OnMatchInput() *string
	Parameters() *[]*string
	SetParameters(val *[]*string)
	ParametersInput() *[]*string
	PathGlob() *string
	SetPathGlob(val *string)
	PathGlobInput() *string
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
	RulesTarget() AppsecWafExclusionFilterRulesTargetList
	RulesTargetInput() interface{}
	Scope() AppsecWafExclusionFilterScopeList
	ScopeInput() interface{}
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
	PutRulesTarget(value interface{})
	PutScope(value interface{})
	ResetEventQuery()
	ResetIpList()
	ResetOnMatch()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetPathGlob()
	ResetRulesTarget()
	ResetScope()
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

// The jsii proxy struct for AppsecWafExclusionFilter
type jsiiProxy_AppsecWafExclusionFilter struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsecWafExclusionFilter) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) EventQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) EventQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) IpList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) IpListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) OnMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) OnMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) Parameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) ParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) PathGlob() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathGlob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) PathGlobInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathGlobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) RulesTarget() AppsecWafExclusionFilterRulesTargetList {
	var returns AppsecWafExclusionFilterRulesTargetList
	_jsii_.Get(
		j,
		"rulesTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) RulesTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rulesTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) Scope() AppsecWafExclusionFilterScopeList {
	var returns AppsecWafExclusionFilterScopeList
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) ScopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsecWafExclusionFilter) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/appsec_waf_exclusion_filter datadog_appsec_waf_exclusion_filter} Resource.
func NewAppsecWafExclusionFilter(scope constructs.Construct, id *string, config *AppsecWafExclusionFilterConfig) AppsecWafExclusionFilter {
	_init_.Initialize()

	if err := validateNewAppsecWafExclusionFilterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsecWafExclusionFilter{}

	_jsii_.Create(
		"@cdktf/provider-datadog.appsecWafExclusionFilter.AppsecWafExclusionFilter",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/appsec_waf_exclusion_filter datadog_appsec_waf_exclusion_filter} Resource.
func NewAppsecWafExclusionFilter_Override(a AppsecWafExclusionFilter, scope constructs.Construct, id *string, config *AppsecWafExclusionFilterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.appsecWafExclusionFilter.AppsecWafExclusionFilter",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetEventQuery(val *string) {
	if err := j.validateSetEventQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventQuery",
		val,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetIpList(val *[]*string) {
	if err := j.validateSetIpListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipList",
		val,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetOnMatch(val *string) {
	if err := j.validateSetOnMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onMatch",
		val,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetParameters(val *[]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetPathGlob(val *string) {
	if err := j.validateSetPathGlobParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathGlob",
		val,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppsecWafExclusionFilter)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a AppsecWafExclusionFilter resource upon running "cdktf plan <stack-name>".
func AppsecWafExclusionFilter_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAppsecWafExclusionFilter_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.appsecWafExclusionFilter.AppsecWafExclusionFilter",
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
func AppsecWafExclusionFilter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsecWafExclusionFilter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.appsecWafExclusionFilter.AppsecWafExclusionFilter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsecWafExclusionFilter_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsecWafExclusionFilter_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.appsecWafExclusionFilter.AppsecWafExclusionFilter",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppsecWafExclusionFilter_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsecWafExclusionFilter_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.appsecWafExclusionFilter.AppsecWafExclusionFilter",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsecWafExclusionFilter_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.appsecWafExclusionFilter.AppsecWafExclusionFilter",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) PutRulesTarget(value interface{}) {
	if err := a.validatePutRulesTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRulesTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) PutScope(value interface{}) {
	if err := a.validatePutScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScope",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) ResetEventQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) ResetIpList() {
	_jsii_.InvokeVoid(
		a,
		"resetIpList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) ResetOnMatch() {
	_jsii_.InvokeVoid(
		a,
		"resetOnMatch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) ResetParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) ResetPathGlob() {
	_jsii_.InvokeVoid(
		a,
		"resetPathGlob",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) ResetRulesTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetRulesTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) ResetScope() {
	_jsii_.InvokeVoid(
		a,
		"resetScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsecWafExclusionFilter) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsecWafExclusionFilter) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

