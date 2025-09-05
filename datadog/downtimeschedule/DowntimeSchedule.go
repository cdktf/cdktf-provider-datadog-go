// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package downtimeschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/downtimeschedule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/downtime_schedule datadog_downtime_schedule}.
type DowntimeSchedule interface {
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
	DisplayTimezone() *string
	SetDisplayTimezone(val *string)
	DisplayTimezoneInput() *string
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
	Message() *string
	SetMessage(val *string)
	MessageInput() *string
	MonitorIdentifier() DowntimeScheduleMonitorIdentifierOutputReference
	MonitorIdentifierInput() interface{}
	MuteFirstRecoveryNotification() interface{}
	SetMuteFirstRecoveryNotification(val interface{})
	MuteFirstRecoveryNotificationInput() interface{}
	// The tree node.
	Node() constructs.Node
	NotifyEndStates() *[]*string
	SetNotifyEndStates(val *[]*string)
	NotifyEndStatesInput() *[]*string
	NotifyEndTypes() *[]*string
	SetNotifyEndTypes(val *[]*string)
	NotifyEndTypesInput() *[]*string
	OneTimeSchedule() DowntimeScheduleOneTimeScheduleOutputReference
	OneTimeScheduleInput() interface{}
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
	RecurringSchedule() DowntimeScheduleRecurringScheduleOutputReference
	RecurringScheduleInput() interface{}
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
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
	PutMonitorIdentifier(value *DowntimeScheduleMonitorIdentifier)
	PutOneTimeSchedule(value *DowntimeScheduleOneTimeSchedule)
	PutRecurringSchedule(value *DowntimeScheduleRecurringSchedule)
	ResetDisplayTimezone()
	ResetMessage()
	ResetMonitorIdentifier()
	ResetMuteFirstRecoveryNotification()
	ResetNotifyEndStates()
	ResetNotifyEndTypes()
	ResetOneTimeSchedule()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRecurringSchedule()
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

// The jsii proxy struct for DowntimeSchedule
type jsiiProxy_DowntimeSchedule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DowntimeSchedule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) DisplayTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) DisplayTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) MessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) MonitorIdentifier() DowntimeScheduleMonitorIdentifierOutputReference {
	var returns DowntimeScheduleMonitorIdentifierOutputReference
	_jsii_.Get(
		j,
		"monitorIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) MonitorIdentifierInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitorIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) MuteFirstRecoveryNotification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"muteFirstRecoveryNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) MuteFirstRecoveryNotificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"muteFirstRecoveryNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) NotifyEndStates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifyEndStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) NotifyEndStatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifyEndStatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) NotifyEndTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifyEndTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) NotifyEndTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifyEndTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) OneTimeSchedule() DowntimeScheduleOneTimeScheduleOutputReference {
	var returns DowntimeScheduleOneTimeScheduleOutputReference
	_jsii_.Get(
		j,
		"oneTimeSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) OneTimeScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oneTimeScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) RecurringSchedule() DowntimeScheduleRecurringScheduleOutputReference {
	var returns DowntimeScheduleRecurringScheduleOutputReference
	_jsii_.Get(
		j,
		"recurringSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) RecurringScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recurringScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DowntimeSchedule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/downtime_schedule datadog_downtime_schedule} Resource.
func NewDowntimeSchedule(scope constructs.Construct, id *string, config *DowntimeScheduleConfig) DowntimeSchedule {
	_init_.Initialize()

	if err := validateNewDowntimeScheduleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DowntimeSchedule{}

	_jsii_.Create(
		"@cdktf/provider-datadog.downtimeSchedule.DowntimeSchedule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/downtime_schedule datadog_downtime_schedule} Resource.
func NewDowntimeSchedule_Override(d DowntimeSchedule, scope constructs.Construct, id *string, config *DowntimeScheduleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.downtimeSchedule.DowntimeSchedule",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DowntimeSchedule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DowntimeSchedule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DowntimeSchedule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DowntimeSchedule)SetDisplayTimezone(val *string) {
	if err := j.validateSetDisplayTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayTimezone",
		val,
	)
}

func (j *jsiiProxy_DowntimeSchedule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DowntimeSchedule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DowntimeSchedule)SetMessage(val *string) {
	if err := j.validateSetMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_DowntimeSchedule)SetMuteFirstRecoveryNotification(val interface{}) {
	if err := j.validateSetMuteFirstRecoveryNotificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"muteFirstRecoveryNotification",
		val,
	)
}

func (j *jsiiProxy_DowntimeSchedule)SetNotifyEndStates(val *[]*string) {
	if err := j.validateSetNotifyEndStatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyEndStates",
		val,
	)
}

func (j *jsiiProxy_DowntimeSchedule)SetNotifyEndTypes(val *[]*string) {
	if err := j.validateSetNotifyEndTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyEndTypes",
		val,
	)
}

func (j *jsiiProxy_DowntimeSchedule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DowntimeSchedule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DowntimeSchedule)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

// Generates CDKTF code for importing a DowntimeSchedule resource upon running "cdktf plan <stack-name>".
func DowntimeSchedule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDowntimeSchedule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.downtimeSchedule.DowntimeSchedule",
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
func DowntimeSchedule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDowntimeSchedule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.downtimeSchedule.DowntimeSchedule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DowntimeSchedule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDowntimeSchedule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.downtimeSchedule.DowntimeSchedule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DowntimeSchedule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDowntimeSchedule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.downtimeSchedule.DowntimeSchedule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DowntimeSchedule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.downtimeSchedule.DowntimeSchedule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DowntimeSchedule) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DowntimeSchedule) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DowntimeSchedule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DowntimeSchedule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DowntimeSchedule) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DowntimeSchedule) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DowntimeSchedule) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DowntimeSchedule) PutMonitorIdentifier(value *DowntimeScheduleMonitorIdentifier) {
	if err := d.validatePutMonitorIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMonitorIdentifier",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DowntimeSchedule) PutOneTimeSchedule(value *DowntimeScheduleOneTimeSchedule) {
	if err := d.validatePutOneTimeScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOneTimeSchedule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DowntimeSchedule) PutRecurringSchedule(value *DowntimeScheduleRecurringSchedule) {
	if err := d.validatePutRecurringScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRecurringSchedule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DowntimeSchedule) ResetDisplayTimezone() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayTimezone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeSchedule) ResetMessage() {
	_jsii_.InvokeVoid(
		d,
		"resetMessage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeSchedule) ResetMonitorIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetMonitorIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeSchedule) ResetMuteFirstRecoveryNotification() {
	_jsii_.InvokeVoid(
		d,
		"resetMuteFirstRecoveryNotification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeSchedule) ResetNotifyEndStates() {
	_jsii_.InvokeVoid(
		d,
		"resetNotifyEndStates",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeSchedule) ResetNotifyEndTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetNotifyEndTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeSchedule) ResetOneTimeSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetOneTimeSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeSchedule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeSchedule) ResetRecurringSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetRecurringSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DowntimeSchedule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DowntimeSchedule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

