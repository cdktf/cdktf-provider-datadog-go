// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/datadog/d/security_monitoring_rules datadog_security_monitoring_rules}.
type DataDatadogSecurityMonitoringRules interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DefaultOnlyFilter() interface{}
	SetDefaultOnlyFilter(val interface{})
	DefaultOnlyFilterInput() interface{}
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
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NameFilter() *string
	SetNameFilter(val *string)
	NameFilterInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RuleIds() *[]*string
	Rules() DataDatadogSecurityMonitoringRulesRulesList
	TagsFilter() *[]*string
	SetTagsFilter(val *[]*string)
	TagsFilterInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserOnlyFilter() interface{}
	SetUserOnlyFilter(val interface{})
	UserOnlyFilterInput() interface{}
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
	ResetDefaultOnlyFilter()
	ResetId()
	ResetNameFilter()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTagsFilter()
	ResetUserOnlyFilter()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataDatadogSecurityMonitoringRules
type jsiiProxy_DataDatadogSecurityMonitoringRules struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) DefaultOnlyFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultOnlyFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) DefaultOnlyFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultOnlyFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) NameFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) NameFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) RuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ruleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) Rules() DataDatadogSecurityMonitoringRulesRulesList {
	var returns DataDatadogSecurityMonitoringRulesRulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) TagsFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) TagsFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) UserOnlyFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userOnlyFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) UserOnlyFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userOnlyFilterInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/datadog/d/security_monitoring_rules datadog_security_monitoring_rules} Data Source.
func NewDataDatadogSecurityMonitoringRules(scope constructs.Construct, id *string, config *DataDatadogSecurityMonitoringRulesConfig) DataDatadogSecurityMonitoringRules {
	_init_.Initialize()

	j := jsiiProxy_DataDatadogSecurityMonitoringRules{}

	_jsii_.Create(
		"@cdktf/provider-datadog.DataDatadogSecurityMonitoringRules",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/datadog/d/security_monitoring_rules datadog_security_monitoring_rules} Data Source.
func NewDataDatadogSecurityMonitoringRules_Override(d DataDatadogSecurityMonitoringRules, scope constructs.Construct, id *string, config *DataDatadogSecurityMonitoringRulesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.DataDatadogSecurityMonitoringRules",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) SetDefaultOnlyFilter(val interface{}) {
	_jsii_.Set(
		j,
		"defaultOnlyFilter",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) SetNameFilter(val *string) {
	_jsii_.Set(
		j,
		"nameFilter",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) SetTagsFilter(val *[]*string) {
	_jsii_.Set(
		j,
		"tagsFilter",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSecurityMonitoringRules) SetUserOnlyFilter(val interface{}) {
	_jsii_.Set(
		j,
		"userOnlyFilter",
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
func DataDatadogSecurityMonitoringRules_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.DataDatadogSecurityMonitoringRules",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatadogSecurityMonitoringRules_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.DataDatadogSecurityMonitoringRules",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) ResetDefaultOnlyFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultOnlyFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) ResetNameFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetNameFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) ResetTagsFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) ResetUserOnlyFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetUserOnlyFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSecurityMonitoringRules) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
