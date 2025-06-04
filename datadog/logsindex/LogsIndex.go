// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logsindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/logsindex/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/logs_index datadog_logs_index}.
type LogsIndex interface {
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
	DailyLimit() *float64
	SetDailyLimit(val *float64)
	DailyLimitInput() *float64
	DailyLimitReset() LogsIndexDailyLimitResetOutputReference
	DailyLimitResetInput() *LogsIndexDailyLimitReset
	DailyLimitWarningThresholdPercentage() *float64
	SetDailyLimitWarningThresholdPercentage(val *float64)
	DailyLimitWarningThresholdPercentageInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableDailyLimit() interface{}
	SetDisableDailyLimit(val interface{})
	DisableDailyLimitInput() interface{}
	ExclusionFilter() LogsIndexExclusionFilterList
	ExclusionFilterInput() interface{}
	Filter() LogsIndexFilterOutputReference
	FilterInput() *LogsIndexFilter
	FlexRetentionDays() *float64
	SetFlexRetentionDays(val *float64)
	FlexRetentionDaysInput() *float64
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
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	RetentionDays() *float64
	SetRetentionDays(val *float64)
	RetentionDaysInput() *float64
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
	PutDailyLimitReset(value *LogsIndexDailyLimitReset)
	PutExclusionFilter(value interface{})
	PutFilter(value *LogsIndexFilter)
	ResetDailyLimit()
	ResetDailyLimitReset()
	ResetDailyLimitWarningThresholdPercentage()
	ResetDisableDailyLimit()
	ResetExclusionFilter()
	ResetFlexRetentionDays()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetentionDays()
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

// The jsii proxy struct for LogsIndex
type jsiiProxy_LogsIndex struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LogsIndex) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) DailyLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) DailyLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) DailyLimitReset() LogsIndexDailyLimitResetOutputReference {
	var returns LogsIndexDailyLimitResetOutputReference
	_jsii_.Get(
		j,
		"dailyLimitReset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) DailyLimitResetInput() *LogsIndexDailyLimitReset {
	var returns *LogsIndexDailyLimitReset
	_jsii_.Get(
		j,
		"dailyLimitResetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) DailyLimitWarningThresholdPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyLimitWarningThresholdPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) DailyLimitWarningThresholdPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyLimitWarningThresholdPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) DisableDailyLimit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDailyLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) DisableDailyLimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDailyLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) ExclusionFilter() LogsIndexExclusionFilterList {
	var returns LogsIndexExclusionFilterList
	_jsii_.Get(
		j,
		"exclusionFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) ExclusionFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exclusionFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) Filter() LogsIndexFilterOutputReference {
	var returns LogsIndexFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) FilterInput() *LogsIndexFilter {
	var returns *LogsIndexFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) FlexRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flexRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) FlexRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flexRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) RetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) RetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsIndex) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/logs_index datadog_logs_index} Resource.
func NewLogsIndex(scope constructs.Construct, id *string, config *LogsIndexConfig) LogsIndex {
	_init_.Initialize()

	if err := validateNewLogsIndexParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogsIndex{}

	_jsii_.Create(
		"@cdktf/provider-datadog.logsIndex.LogsIndex",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/logs_index datadog_logs_index} Resource.
func NewLogsIndex_Override(l LogsIndex, scope constructs.Construct, id *string, config *LogsIndexConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.logsIndex.LogsIndex",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LogsIndex)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LogsIndex)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LogsIndex)SetDailyLimit(val *float64) {
	if err := j.validateSetDailyLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyLimit",
		val,
	)
}

func (j *jsiiProxy_LogsIndex)SetDailyLimitWarningThresholdPercentage(val *float64) {
	if err := j.validateSetDailyLimitWarningThresholdPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyLimitWarningThresholdPercentage",
		val,
	)
}

func (j *jsiiProxy_LogsIndex)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LogsIndex)SetDisableDailyLimit(val interface{}) {
	if err := j.validateSetDisableDailyLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDailyLimit",
		val,
	)
}

func (j *jsiiProxy_LogsIndex)SetFlexRetentionDays(val *float64) {
	if err := j.validateSetFlexRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flexRetentionDays",
		val,
	)
}

func (j *jsiiProxy_LogsIndex)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LogsIndex)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LogsIndex)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LogsIndex)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LogsIndex)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LogsIndex)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LogsIndex)SetRetentionDays(val *float64) {
	if err := j.validateSetRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionDays",
		val,
	)
}

// Generates CDKTF code for importing a LogsIndex resource upon running "cdktf plan <stack-name>".
func LogsIndex_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLogsIndex_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.logsIndex.LogsIndex",
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
func LogsIndex_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogsIndex_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.logsIndex.LogsIndex",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogsIndex_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogsIndex_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.logsIndex.LogsIndex",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogsIndex_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogsIndex_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.logsIndex.LogsIndex",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LogsIndex_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.logsIndex.LogsIndex",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LogsIndex) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LogsIndex) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LogsIndex) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LogsIndex) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogsIndex) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LogsIndex) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogsIndex) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LogsIndex) PutDailyLimitReset(value *LogsIndexDailyLimitReset) {
	if err := l.validatePutDailyLimitResetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDailyLimitReset",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsIndex) PutExclusionFilter(value interface{}) {
	if err := l.validatePutExclusionFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putExclusionFilter",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsIndex) PutFilter(value *LogsIndexFilter) {
	if err := l.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFilter",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsIndex) ResetDailyLimit() {
	_jsii_.InvokeVoid(
		l,
		"resetDailyLimit",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsIndex) ResetDailyLimitReset() {
	_jsii_.InvokeVoid(
		l,
		"resetDailyLimitReset",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsIndex) ResetDailyLimitWarningThresholdPercentage() {
	_jsii_.InvokeVoid(
		l,
		"resetDailyLimitWarningThresholdPercentage",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsIndex) ResetDisableDailyLimit() {
	_jsii_.InvokeVoid(
		l,
		"resetDisableDailyLimit",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsIndex) ResetExclusionFilter() {
	_jsii_.InvokeVoid(
		l,
		"resetExclusionFilter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsIndex) ResetFlexRetentionDays() {
	_jsii_.InvokeVoid(
		l,
		"resetFlexRetentionDays",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsIndex) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsIndex) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsIndex) ResetRetentionDays() {
	_jsii_.InvokeVoid(
		l,
		"resetRetentionDays",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsIndex) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsIndex) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

