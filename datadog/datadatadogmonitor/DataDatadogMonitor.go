package datadatadogmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v6/datadatadogmonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.23.0/docs/data-sources/monitor datadog_monitor}.
type DataDatadogMonitor interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	EnableLogsSample() cdktf.IResolvable
	EnableSamples() cdktf.IResolvable
	EscalationMessage() *string
	EvaluationDelay() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupbySimpleMonitor() cdktf.IResolvable
	GroupRetentionDuration() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IncludeTags() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locked() cdktf.IResolvable
	Message() *string
	MonitorTagsFilter() *[]*string
	SetMonitorTagsFilter(val *[]*string)
	MonitorTagsFilterInput() *[]*string
	MonitorThresholds() DataDatadogMonitorMonitorThresholdsList
	MonitorThresholdWindows() DataDatadogMonitorMonitorThresholdWindowsList
	Name() *string
	NameFilter() *string
	SetNameFilter(val *string)
	NameFilterInput() *string
	NewGroupDelay() *float64
	NewHostDelay() *float64
	NoDataTimeframe() *float64
	// The tree node.
	Node() constructs.Node
	NotificationPresetName() *string
	NotifyAudit() cdktf.IResolvable
	NotifyBy() *[]*string
	NotifyNoData() cdktf.IResolvable
	OnMissingData() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	Query() *string
	// Experimental.
	RawOverrides() interface{}
	RenotifyInterval() *float64
	RenotifyOccurrences() *float64
	RenotifyStatuses() *[]*string
	RequireFullWindow() cdktf.IResolvable
	RestrictedRoles() *[]*string
	SchedulingOptions() DataDatadogMonitorSchedulingOptionsList
	Tags() *[]*string
	TagsFilter() *[]*string
	SetTagsFilter(val *[]*string)
	TagsFilterInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeoutH() *float64
	Type() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
	ResetMonitorTagsFilter()
	ResetNameFilter()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTagsFilter()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataDatadogMonitor
type jsiiProxy_DataDatadogMonitor struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatadogMonitor) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) EnableLogsSample() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableLogsSample",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) EnableSamples() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableSamples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) EscalationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escalationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) EvaluationDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) GroupbySimpleMonitor() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"groupbySimpleMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) GroupRetentionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupRetentionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) IncludeTags() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"includeTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) Locked() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"locked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) MonitorTagsFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monitorTagsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) MonitorTagsFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monitorTagsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) MonitorThresholds() DataDatadogMonitorMonitorThresholdsList {
	var returns DataDatadogMonitorMonitorThresholdsList
	_jsii_.Get(
		j,
		"monitorThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) MonitorThresholdWindows() DataDatadogMonitorMonitorThresholdWindowsList {
	var returns DataDatadogMonitorMonitorThresholdWindowsList
	_jsii_.Get(
		j,
		"monitorThresholdWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) NameFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) NameFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) NewGroupDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newGroupDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) NewHostDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newHostDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) NoDataTimeframe() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"noDataTimeframe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) NotificationPresetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPresetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) NotifyAudit() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"notifyAudit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) NotifyBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifyBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) NotifyNoData() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"notifyNoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) OnMissingData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onMissingData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) RenotifyInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renotifyInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) RenotifyOccurrences() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"renotifyOccurrences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) RenotifyStatuses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"renotifyStatuses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) RequireFullWindow() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"requireFullWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) RestrictedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) SchedulingOptions() DataDatadogMonitorSchedulingOptionsList {
	var returns DataDatadogMonitorSchedulingOptionsList
	_jsii_.Get(
		j,
		"schedulingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) TagsFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) TagsFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) TimeoutH() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutH",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitor) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.23.0/docs/data-sources/monitor datadog_monitor} Data Source.
func NewDataDatadogMonitor(scope constructs.Construct, id *string, config *DataDatadogMonitorConfig) DataDatadogMonitor {
	_init_.Initialize()

	if err := validateNewDataDatadogMonitorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatadogMonitor{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dataDatadogMonitor.DataDatadogMonitor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.23.0/docs/data-sources/monitor datadog_monitor} Data Source.
func NewDataDatadogMonitor_Override(d DataDatadogMonitor, scope constructs.Construct, id *string, config *DataDatadogMonitorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dataDatadogMonitor.DataDatadogMonitor",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatadogMonitor)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatadogMonitor)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatadogMonitor)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatadogMonitor)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDatadogMonitor)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatadogMonitor)SetMonitorTagsFilter(val *[]*string) {
	if err := j.validateSetMonitorTagsFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitorTagsFilter",
		val,
	)
}

func (j *jsiiProxy_DataDatadogMonitor)SetNameFilter(val *string) {
	if err := j.validateSetNameFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameFilter",
		val,
	)
}

func (j *jsiiProxy_DataDatadogMonitor)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDatadogMonitor)SetTagsFilter(val *[]*string) {
	if err := j.validateSetTagsFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsFilter",
		val,
	)
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
func DataDatadogMonitor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatadogMonitor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.dataDatadogMonitor.DataDatadogMonitor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatadogMonitor_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatadogMonitor_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.dataDatadogMonitor.DataDatadogMonitor",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatadogMonitor_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatadogMonitor_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.dataDatadogMonitor.DataDatadogMonitor",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatadogMonitor_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.dataDatadogMonitor.DataDatadogMonitor",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatadogMonitor) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatadogMonitor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatadogMonitor) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatadogMonitor) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatadogMonitor) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatadogMonitor) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatadogMonitor) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatadogMonitor) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatadogMonitor) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatadogMonitor) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatadogMonitor) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatadogMonitor) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatadogMonitor) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogMonitor) ResetMonitorTagsFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetMonitorTagsFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogMonitor) ResetNameFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetNameFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogMonitor) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogMonitor) ResetTagsFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogMonitor) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogMonitor) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogMonitor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogMonitor) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

