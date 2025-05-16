// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogsoftwarecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/datadatadogsoftwarecatalog/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/data-sources/software_catalog datadog_software_catalog}.
type DataDatadogSoftwareCatalog interface {
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
	Entities() DataDatadogSoftwareCatalogEntitiesList
	FilterExcludeSnapshot() *string
	SetFilterExcludeSnapshot(val *string)
	FilterExcludeSnapshotInput() *string
	FilterId() *string
	SetFilterId(val *string)
	FilterIdInput() *string
	FilterKind() *string
	SetFilterKind(val *string)
	FilterKindInput() *string
	FilterName() *string
	SetFilterName(val *string)
	FilterNameInput() *string
	FilterOwner() *string
	SetFilterOwner(val *string)
	FilterOwnerInput() *string
	FilterRef() *string
	SetFilterRef(val *string)
	FilterRefInput() *string
	FilterRelationType() *string
	SetFilterRelationType(val *string)
	FilterRelationTypeInput() *string
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
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
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
	ResetFilterExcludeSnapshot()
	ResetFilterId()
	ResetFilterKind()
	ResetFilterName()
	ResetFilterOwner()
	ResetFilterRef()
	ResetFilterRelationType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataDatadogSoftwareCatalog
type jsiiProxy_DataDatadogSoftwareCatalog struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) Entities() DataDatadogSoftwareCatalogEntitiesList {
	var returns DataDatadogSoftwareCatalogEntitiesList
	_jsii_.Get(
		j,
		"entities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterExcludeSnapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterExcludeSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterExcludeSnapshotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterExcludeSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterKindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterRelationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterRelationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FilterRelationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterRelationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/data-sources/software_catalog datadog_software_catalog} Data Source.
func NewDataDatadogSoftwareCatalog(scope constructs.Construct, id *string, config *DataDatadogSoftwareCatalogConfig) DataDatadogSoftwareCatalog {
	_init_.Initialize()

	if err := validateNewDataDatadogSoftwareCatalogParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatadogSoftwareCatalog{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dataDatadogSoftwareCatalog.DataDatadogSoftwareCatalog",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/data-sources/software_catalog datadog_software_catalog} Data Source.
func NewDataDatadogSoftwareCatalog_Override(d DataDatadogSoftwareCatalog, scope constructs.Construct, id *string, config *DataDatadogSoftwareCatalogConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dataDatadogSoftwareCatalog.DataDatadogSoftwareCatalog",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog)SetFilterExcludeSnapshot(val *string) {
	if err := j.validateSetFilterExcludeSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterExcludeSnapshot",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog)SetFilterId(val *string) {
	if err := j.validateSetFilterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterId",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog)SetFilterKind(val *string) {
	if err := j.validateSetFilterKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterKind",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog)SetFilterName(val *string) {
	if err := j.validateSetFilterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterName",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog)SetFilterOwner(val *string) {
	if err := j.validateSetFilterOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterOwner",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog)SetFilterRef(val *string) {
	if err := j.validateSetFilterRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterRef",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog)SetFilterRelationType(val *string) {
	if err := j.validateSetFilterRelationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterRelationType",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataDatadogSoftwareCatalog)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataDatadogSoftwareCatalog resource upon running "cdktf plan <stack-name>".
func DataDatadogSoftwareCatalog_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataDatadogSoftwareCatalog_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.dataDatadogSoftwareCatalog.DataDatadogSoftwareCatalog",
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
func DataDatadogSoftwareCatalog_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatadogSoftwareCatalog_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.dataDatadogSoftwareCatalog.DataDatadogSoftwareCatalog",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatadogSoftwareCatalog_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatadogSoftwareCatalog_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.dataDatadogSoftwareCatalog.DataDatadogSoftwareCatalog",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataDatadogSoftwareCatalog_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataDatadogSoftwareCatalog_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.dataDatadogSoftwareCatalog.DataDatadogSoftwareCatalog",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataDatadogSoftwareCatalog_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.dataDatadogSoftwareCatalog.DataDatadogSoftwareCatalog",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatadogSoftwareCatalog) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatadogSoftwareCatalog) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatadogSoftwareCatalog) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatadogSoftwareCatalog) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatadogSoftwareCatalog) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatadogSoftwareCatalog) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatadogSoftwareCatalog) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatadogSoftwareCatalog) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatadogSoftwareCatalog) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatadogSoftwareCatalog) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) ResetFilterExcludeSnapshot() {
	_jsii_.InvokeVoid(
		d,
		"resetFilterExcludeSnapshot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) ResetFilterId() {
	_jsii_.InvokeVoid(
		d,
		"resetFilterId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) ResetFilterKind() {
	_jsii_.InvokeVoid(
		d,
		"resetFilterKind",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) ResetFilterName() {
	_jsii_.InvokeVoid(
		d,
		"resetFilterName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) ResetFilterOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetFilterOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) ResetFilterRef() {
	_jsii_.InvokeVoid(
		d,
		"resetFilterRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) ResetFilterRelationType() {
	_jsii_.InvokeVoid(
		d,
		"resetFilterRelationType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogSoftwareCatalog) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

