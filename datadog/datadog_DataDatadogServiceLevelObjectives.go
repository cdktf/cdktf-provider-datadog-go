// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/datadog/d/service_level_objectives datadog_service_level_objectives}.
type DataDatadogServiceLevelObjectives interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	Ids() *[]*string
	SetIds(val *[]*string)
	IdsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MetricsQuery() *string
	SetMetricsQuery(val *string)
	MetricsQueryInput() *string
	NameQuery() *string
	SetNameQuery(val *string)
	NameQueryInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Slos() DataDatadogServiceLevelObjectivesSlosList
	TagsQuery() *string
	SetTagsQuery(val *string)
	TagsQueryInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetIds()
	ResetMetricsQuery()
	ResetNameQuery()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTagsQuery()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataDatadogServiceLevelObjectives
type jsiiProxy_DataDatadogServiceLevelObjectives struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) Ids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) IdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"idsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) MetricsQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) MetricsQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) NameQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) NameQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) Slos() DataDatadogServiceLevelObjectivesSlosList {
	var returns DataDatadogServiceLevelObjectivesSlosList
	_jsii_.Get(
		j,
		"slos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) TagsQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) TagsQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/datadog/d/service_level_objectives datadog_service_level_objectives} Data Source.
func NewDataDatadogServiceLevelObjectives(scope constructs.Construct, id *string, config *DataDatadogServiceLevelObjectivesConfig) DataDatadogServiceLevelObjectives {
	_init_.Initialize()

	j := jsiiProxy_DataDatadogServiceLevelObjectives{}

	_jsii_.Create(
		"@cdktf/provider-datadog.DataDatadogServiceLevelObjectives",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/datadog/d/service_level_objectives datadog_service_level_objectives} Data Source.
func NewDataDatadogServiceLevelObjectives_Override(d DataDatadogServiceLevelObjectives, scope constructs.Construct, id *string, config *DataDatadogServiceLevelObjectivesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.DataDatadogServiceLevelObjectives",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) SetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"ids",
		val,
	)
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) SetMetricsQuery(val *string) {
	_jsii_.Set(
		j,
		"metricsQuery",
		val,
	)
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) SetNameQuery(val *string) {
	_jsii_.Set(
		j,
		"nameQuery",
		val,
	)
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataDatadogServiceLevelObjectives) SetTagsQuery(val *string) {
	_jsii_.Set(
		j,
		"tagsQuery",
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
func DataDatadogServiceLevelObjectives_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.DataDatadogServiceLevelObjectives",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatadogServiceLevelObjectives_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.DataDatadogServiceLevelObjectives",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) ResetIds() {
	_jsii_.InvokeVoid(
		d,
		"resetIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) ResetMetricsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetMetricsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) ResetNameQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetNameQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) ResetTagsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogServiceLevelObjectives) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
