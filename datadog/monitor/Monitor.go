// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/monitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/monitor datadog_monitor}.
type Monitor interface {
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
	EnableLogsSample() interface{}
	SetEnableLogsSample(val interface{})
	EnableLogsSampleInput() interface{}
	EnableSamples() interface{}
	SetEnableSamples(val interface{})
	EnableSamplesInput() interface{}
	EscalationMessage() *string
	SetEscalationMessage(val *string)
	EscalationMessageInput() *string
	EvaluationDelay() *float64
	SetEvaluationDelay(val *float64)
	EvaluationDelayInput() *float64
	ForceDelete() interface{}
	SetForceDelete(val interface{})
	ForceDeleteInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupbySimpleMonitor() interface{}
	SetGroupbySimpleMonitor(val interface{})
	GroupbySimpleMonitorInput() interface{}
	GroupRetentionDuration() *string
	SetGroupRetentionDuration(val *string)
	GroupRetentionDurationInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IncludeTags() interface{}
	SetIncludeTags(val interface{})
	IncludeTagsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locked() interface{}
	SetLocked(val interface{})
	LockedInput() interface{}
	Message() *string
	SetMessage(val *string)
	MessageInput() *string
	MonitorThresholds() MonitorMonitorThresholdsOutputReference
	MonitorThresholdsInput() *MonitorMonitorThresholds
	MonitorThresholdWindows() MonitorMonitorThresholdWindowsOutputReference
	MonitorThresholdWindowsInput() *MonitorMonitorThresholdWindows
	Name() *string
	SetName(val *string)
	NameInput() *string
	NewGroupDelay() *float64
	SetNewGroupDelay(val *float64)
	NewGroupDelayInput() *float64
	NewHostDelay() *float64
	SetNewHostDelay(val *float64)
	NewHostDelayInput() *float64
	NoDataTimeframe() *float64
	SetNoDataTimeframe(val *float64)
	NoDataTimeframeInput() *float64
	// The tree node.
	Node() constructs.Node
	NotificationPresetName() *string
	SetNotificationPresetName(val *string)
	NotificationPresetNameInput() *string
	NotifyAudit() interface{}
	SetNotifyAudit(val interface{})
	NotifyAuditInput() interface{}
	NotifyBy() *[]*string
	SetNotifyBy(val *[]*string)
	NotifyByInput() *[]*string
	NotifyNoData() interface{}
	SetNotifyNoData(val interface{})
	NotifyNoDataInput() interface{}
	OnMissingData() *string
	SetOnMissingData(val *string)
	OnMissingDataInput() *string
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	// Experimental.
	RawOverrides() interface{}
	RenotifyInterval() *float64
	SetRenotifyInterval(val *float64)
	RenotifyIntervalInput() *float64
	RenotifyOccurrences() *float64
	SetRenotifyOccurrences(val *float64)
	RenotifyOccurrencesInput() *float64
	RenotifyStatuses() *[]*string
	SetRenotifyStatuses(val *[]*string)
	RenotifyStatusesInput() *[]*string
	RequireFullWindow() interface{}
	SetRequireFullWindow(val interface{})
	RequireFullWindowInput() interface{}
	RestrictedRoles() *[]*string
	SetRestrictedRoles(val *[]*string)
	RestrictedRolesInput() *[]*string
	SchedulingOptions() MonitorSchedulingOptionsOutputReference
	SchedulingOptionsInput() *MonitorSchedulingOptions
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeoutH() *float64
	SetTimeoutH(val *float64)
	TimeoutHInput() *float64
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Validate() interface{}
	SetValidate(val interface{})
	ValidateInput() interface{}
	Variables() MonitorVariablesOutputReference
	VariablesInput() *MonitorVariables
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
	PutMonitorThresholds(value *MonitorMonitorThresholds)
	PutMonitorThresholdWindows(value *MonitorMonitorThresholdWindows)
	PutSchedulingOptions(value *MonitorSchedulingOptions)
	PutVariables(value *MonitorVariables)
	ResetEnableLogsSample()
	ResetEnableSamples()
	ResetEscalationMessage()
	ResetEvaluationDelay()
	ResetForceDelete()
	ResetGroupbySimpleMonitor()
	ResetGroupRetentionDuration()
	ResetId()
	ResetIncludeTags()
	ResetLocked()
	ResetMonitorThresholds()
	ResetMonitorThresholdWindows()
	ResetNewGroupDelay()
	ResetNewHostDelay()
	ResetNoDataTimeframe()
	ResetNotificationPresetName()
	ResetNotifyAudit()
	ResetNotifyBy()
	ResetNotifyNoData()
	ResetOnMissingData()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPriority()
	ResetRenotifyInterval()
	ResetRenotifyOccurrences()
	ResetRenotifyStatuses()
	ResetRequireFullWindow()
	ResetRestrictedRoles()
	ResetSchedulingOptions()
	ResetTags()
	ResetTimeoutH()
	ResetValidate()
	ResetVariables()
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

// The jsii proxy struct for Monitor
type jsiiProxy_Monitor struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Monitor) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) EnableLogsSample() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogsSample",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) EnableLogsSampleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogsSampleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) EnableSamples() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSamples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) EnableSamplesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSamplesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) EscalationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escalationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) EscalationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escalationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) EvaluationDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) EvaluationDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) ForceDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) ForceDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) GroupbySimpleMonitor() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupbySimpleMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) GroupbySimpleMonitorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupbySimpleMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) GroupRetentionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupRetentionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) GroupRetentionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupRetentionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) IncludeTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) IncludeTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Locked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) LockedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) MessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) MonitorThresholds() MonitorMonitorThresholdsOutputReference {
	var returns MonitorMonitorThresholdsOutputReference
	_jsii_.Get(
		j,
		"monitorThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) MonitorThresholdsInput() *MonitorMonitorThresholds {
	var returns *MonitorMonitorThresholds
	_jsii_.Get(
		j,
		"monitorThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) MonitorThresholdWindows() MonitorMonitorThresholdWindowsOutputReference {
	var returns MonitorMonitorThresholdWindowsOutputReference
	_jsii_.Get(
		j,
		"monitorThresholdWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) MonitorThresholdWindowsInput() *MonitorMonitorThresholdWindows {
	var returns *MonitorMonitorThresholdWindows
	_jsii_.Get(
		j,
		"monitorThresholdWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NewGroupDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newGroupDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NewGroupDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newGroupDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NewHostDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newHostDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NewHostDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newHostDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NoDataTimeframe() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"noDataTimeframe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NoDataTimeframeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"noDataTimeframeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NotificationPresetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPresetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NotificationPresetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPresetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NotifyAudit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyAudit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NotifyAuditInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyAuditInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NotifyBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifyBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NotifyByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifyByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NotifyNoData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyNoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) NotifyNoDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyNoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) OnMissingData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onMissingData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) OnMissingDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onMissingDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) RenotifyInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renotifyInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) RenotifyIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renotifyIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) RenotifyOccurrences() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renotifyOccurrences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) RenotifyOccurrencesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renotifyOccurrencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) RenotifyStatuses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"renotifyStatuses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) RenotifyStatusesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"renotifyStatusesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) RequireFullWindow() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireFullWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) RequireFullWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireFullWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) RestrictedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) RestrictedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) SchedulingOptions() MonitorSchedulingOptionsOutputReference {
	var returns MonitorSchedulingOptionsOutputReference
	_jsii_.Get(
		j,
		"schedulingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) SchedulingOptionsInput() *MonitorSchedulingOptions {
	var returns *MonitorSchedulingOptions
	_jsii_.Get(
		j,
		"schedulingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) TimeoutH() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutH",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) TimeoutHInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutHInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Validate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) ValidateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) Variables() MonitorVariablesOutputReference {
	var returns MonitorVariablesOutputReference
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitor) VariablesInput() *MonitorVariables {
	var returns *MonitorVariables
	_jsii_.Get(
		j,
		"variablesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/monitor datadog_monitor} Resource.
func NewMonitor(scope constructs.Construct, id *string, config *MonitorConfig) Monitor {
	_init_.Initialize()

	if err := validateNewMonitorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Monitor{}

	_jsii_.Create(
		"@cdktf/provider-datadog.monitor.Monitor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/monitor datadog_monitor} Resource.
func NewMonitor_Override(m Monitor, scope constructs.Construct, id *string, config *MonitorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.monitor.Monitor",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_Monitor)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetEnableLogsSample(val interface{}) {
	if err := j.validateSetEnableLogsSampleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogsSample",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetEnableSamples(val interface{}) {
	if err := j.validateSetEnableSamplesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSamples",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetEscalationMessage(val *string) {
	if err := j.validateSetEscalationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"escalationMessage",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetEvaluationDelay(val *float64) {
	if err := j.validateSetEvaluationDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationDelay",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetForceDelete(val interface{}) {
	if err := j.validateSetForceDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDelete",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetGroupbySimpleMonitor(val interface{}) {
	if err := j.validateSetGroupbySimpleMonitorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupbySimpleMonitor",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetGroupRetentionDuration(val *string) {
	if err := j.validateSetGroupRetentionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupRetentionDuration",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetIncludeTags(val interface{}) {
	if err := j.validateSetIncludeTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTags",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetLocked(val interface{}) {
	if err := j.validateSetLockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locked",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetMessage(val *string) {
	if err := j.validateSetMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetNewGroupDelay(val *float64) {
	if err := j.validateSetNewGroupDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newGroupDelay",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetNewHostDelay(val *float64) {
	if err := j.validateSetNewHostDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newHostDelay",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetNoDataTimeframe(val *float64) {
	if err := j.validateSetNoDataTimeframeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDataTimeframe",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetNotificationPresetName(val *string) {
	if err := j.validateSetNotificationPresetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationPresetName",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetNotifyAudit(val interface{}) {
	if err := j.validateSetNotifyAuditParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyAudit",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetNotifyBy(val *[]*string) {
	if err := j.validateSetNotifyByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyBy",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetNotifyNoData(val interface{}) {
	if err := j.validateSetNotifyNoDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyNoData",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetOnMissingData(val *string) {
	if err := j.validateSetOnMissingDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onMissingData",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetRenotifyInterval(val *float64) {
	if err := j.validateSetRenotifyIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renotifyInterval",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetRenotifyOccurrences(val *float64) {
	if err := j.validateSetRenotifyOccurrencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renotifyOccurrences",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetRenotifyStatuses(val *[]*string) {
	if err := j.validateSetRenotifyStatusesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renotifyStatuses",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetRequireFullWindow(val interface{}) {
	if err := j.validateSetRequireFullWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireFullWindow",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetRestrictedRoles(val *[]*string) {
	if err := j.validateSetRestrictedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictedRoles",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetTimeoutH(val *float64) {
	if err := j.validateSetTimeoutHParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutH",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_Monitor)SetValidate(val interface{}) {
	if err := j.validateSetValidateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validate",
		val,
	)
}

// Generates CDKTF code for importing a Monitor resource upon running "cdktf plan <stack-name>".
func Monitor_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMonitor_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.monitor.Monitor",
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
func Monitor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.monitor.Monitor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Monitor_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitor_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.monitor.Monitor",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Monitor_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitor_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.monitor.Monitor",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Monitor_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.monitor.Monitor",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_Monitor) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_Monitor) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_Monitor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_Monitor) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_Monitor) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_Monitor) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_Monitor) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_Monitor) PutMonitorThresholds(value *MonitorMonitorThresholds) {
	if err := m.validatePutMonitorThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMonitorThresholds",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Monitor) PutMonitorThresholdWindows(value *MonitorMonitorThresholdWindows) {
	if err := m.validatePutMonitorThresholdWindowsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMonitorThresholdWindows",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Monitor) PutSchedulingOptions(value *MonitorSchedulingOptions) {
	if err := m.validatePutSchedulingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSchedulingOptions",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Monitor) PutVariables(value *MonitorVariables) {
	if err := m.validatePutVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putVariables",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_Monitor) ResetEnableLogsSample() {
	_jsii_.InvokeVoid(
		m,
		"resetEnableLogsSample",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetEnableSamples() {
	_jsii_.InvokeVoid(
		m,
		"resetEnableSamples",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetEscalationMessage() {
	_jsii_.InvokeVoid(
		m,
		"resetEscalationMessage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetEvaluationDelay() {
	_jsii_.InvokeVoid(
		m,
		"resetEvaluationDelay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetForceDelete() {
	_jsii_.InvokeVoid(
		m,
		"resetForceDelete",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetGroupbySimpleMonitor() {
	_jsii_.InvokeVoid(
		m,
		"resetGroupbySimpleMonitor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetGroupRetentionDuration() {
	_jsii_.InvokeVoid(
		m,
		"resetGroupRetentionDuration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetIncludeTags() {
	_jsii_.InvokeVoid(
		m,
		"resetIncludeTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetLocked() {
	_jsii_.InvokeVoid(
		m,
		"resetLocked",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetMonitorThresholds() {
	_jsii_.InvokeVoid(
		m,
		"resetMonitorThresholds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetMonitorThresholdWindows() {
	_jsii_.InvokeVoid(
		m,
		"resetMonitorThresholdWindows",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetNewGroupDelay() {
	_jsii_.InvokeVoid(
		m,
		"resetNewGroupDelay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetNewHostDelay() {
	_jsii_.InvokeVoid(
		m,
		"resetNewHostDelay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetNoDataTimeframe() {
	_jsii_.InvokeVoid(
		m,
		"resetNoDataTimeframe",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetNotificationPresetName() {
	_jsii_.InvokeVoid(
		m,
		"resetNotificationPresetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetNotifyAudit() {
	_jsii_.InvokeVoid(
		m,
		"resetNotifyAudit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetNotifyBy() {
	_jsii_.InvokeVoid(
		m,
		"resetNotifyBy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetNotifyNoData() {
	_jsii_.InvokeVoid(
		m,
		"resetNotifyNoData",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetOnMissingData() {
	_jsii_.InvokeVoid(
		m,
		"resetOnMissingData",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetPriority() {
	_jsii_.InvokeVoid(
		m,
		"resetPriority",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetRenotifyInterval() {
	_jsii_.InvokeVoid(
		m,
		"resetRenotifyInterval",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetRenotifyOccurrences() {
	_jsii_.InvokeVoid(
		m,
		"resetRenotifyOccurrences",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetRenotifyStatuses() {
	_jsii_.InvokeVoid(
		m,
		"resetRenotifyStatuses",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetRequireFullWindow() {
	_jsii_.InvokeVoid(
		m,
		"resetRequireFullWindow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetRestrictedRoles() {
	_jsii_.InvokeVoid(
		m,
		"resetRestrictedRoles",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetSchedulingOptions() {
	_jsii_.InvokeVoid(
		m,
		"resetSchedulingOptions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetTimeoutH() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeoutH",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetValidate() {
	_jsii_.InvokeVoid(
		m,
		"resetValidate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) ResetVariables() {
	_jsii_.InvokeVoid(
		m,
		"resetVariables",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Monitor) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitor) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

