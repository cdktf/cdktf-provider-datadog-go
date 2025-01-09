// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logsarchive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/logsarchive/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/logs_archive datadog_logs_archive}.
type LogsArchive interface {
	cdktf.TerraformResource
	AzureArchive() LogsArchiveAzureArchiveOutputReference
	AzureArchiveInput() *LogsArchiveAzureArchive
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
	GcsArchive() LogsArchiveGcsArchiveOutputReference
	GcsArchiveInput() *LogsArchiveGcsArchive
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
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	// Experimental.
	RawOverrides() interface{}
	RehydrationMaxScanSizeInGb() *float64
	SetRehydrationMaxScanSizeInGb(val *float64)
	RehydrationMaxScanSizeInGbInput() *float64
	RehydrationTags() *[]*string
	SetRehydrationTags(val *[]*string)
	RehydrationTagsInput() *[]*string
	S3Archive() LogsArchiveS3ArchiveOutputReference
	S3ArchiveInput() *LogsArchiveS3Archive
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
	PutAzureArchive(value *LogsArchiveAzureArchive)
	PutGcsArchive(value *LogsArchiveGcsArchive)
	PutS3Archive(value *LogsArchiveS3Archive)
	ResetAzureArchive()
	ResetGcsArchive()
	ResetId()
	ResetIncludeTags()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRehydrationMaxScanSizeInGb()
	ResetRehydrationTags()
	ResetS3Archive()
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

// The jsii proxy struct for LogsArchive
type jsiiProxy_LogsArchive struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LogsArchive) AzureArchive() LogsArchiveAzureArchiveOutputReference {
	var returns LogsArchiveAzureArchiveOutputReference
	_jsii_.Get(
		j,
		"azureArchive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) AzureArchiveInput() *LogsArchiveAzureArchive {
	var returns *LogsArchiveAzureArchive
	_jsii_.Get(
		j,
		"azureArchiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) GcsArchive() LogsArchiveGcsArchiveOutputReference {
	var returns LogsArchiveGcsArchiveOutputReference
	_jsii_.Get(
		j,
		"gcsArchive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) GcsArchiveInput() *LogsArchiveGcsArchive {
	var returns *LogsArchiveGcsArchive
	_jsii_.Get(
		j,
		"gcsArchiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) IncludeTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) IncludeTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) RehydrationMaxScanSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rehydrationMaxScanSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) RehydrationMaxScanSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rehydrationMaxScanSizeInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) RehydrationTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rehydrationTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) RehydrationTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rehydrationTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) S3Archive() LogsArchiveS3ArchiveOutputReference {
	var returns LogsArchiveS3ArchiveOutputReference
	_jsii_.Get(
		j,
		"s3Archive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) S3ArchiveInput() *LogsArchiveS3Archive {
	var returns *LogsArchiveS3Archive
	_jsii_.Get(
		j,
		"s3ArchiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsArchive) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/logs_archive datadog_logs_archive} Resource.
func NewLogsArchive(scope constructs.Construct, id *string, config *LogsArchiveConfig) LogsArchive {
	_init_.Initialize()

	if err := validateNewLogsArchiveParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogsArchive{}

	_jsii_.Create(
		"@cdktf/provider-datadog.logsArchive.LogsArchive",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/logs_archive datadog_logs_archive} Resource.
func NewLogsArchive_Override(l LogsArchive, scope constructs.Construct, id *string, config *LogsArchiveConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.logsArchive.LogsArchive",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LogsArchive)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LogsArchive)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LogsArchive)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LogsArchive)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LogsArchive)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LogsArchive)SetIncludeTags(val interface{}) {
	if err := j.validateSetIncludeTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTags",
		val,
	)
}

func (j *jsiiProxy_LogsArchive)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LogsArchive)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LogsArchive)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LogsArchive)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LogsArchive)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_LogsArchive)SetRehydrationMaxScanSizeInGb(val *float64) {
	if err := j.validateSetRehydrationMaxScanSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rehydrationMaxScanSizeInGb",
		val,
	)
}

func (j *jsiiProxy_LogsArchive)SetRehydrationTags(val *[]*string) {
	if err := j.validateSetRehydrationTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rehydrationTags",
		val,
	)
}

// Generates CDKTF code for importing a LogsArchive resource upon running "cdktf plan <stack-name>".
func LogsArchive_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLogsArchive_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.logsArchive.LogsArchive",
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
func LogsArchive_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogsArchive_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.logsArchive.LogsArchive",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogsArchive_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogsArchive_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.logsArchive.LogsArchive",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogsArchive_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogsArchive_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.logsArchive.LogsArchive",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LogsArchive_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.logsArchive.LogsArchive",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LogsArchive) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LogsArchive) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LogsArchive) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LogsArchive) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogsArchive) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LogsArchive) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LogsArchive) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LogsArchive) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LogsArchive) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LogsArchive) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LogsArchive) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LogsArchive) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsArchive) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LogsArchive) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LogsArchive) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogsArchive) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LogsArchive) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogsArchive) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LogsArchive) PutAzureArchive(value *LogsArchiveAzureArchive) {
	if err := l.validatePutAzureArchiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAzureArchive",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsArchive) PutGcsArchive(value *LogsArchiveGcsArchive) {
	if err := l.validatePutGcsArchiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putGcsArchive",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsArchive) PutS3Archive(value *LogsArchiveS3Archive) {
	if err := l.validatePutS3ArchiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putS3Archive",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsArchive) ResetAzureArchive() {
	_jsii_.InvokeVoid(
		l,
		"resetAzureArchive",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsArchive) ResetGcsArchive() {
	_jsii_.InvokeVoid(
		l,
		"resetGcsArchive",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsArchive) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsArchive) ResetIncludeTags() {
	_jsii_.InvokeVoid(
		l,
		"resetIncludeTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsArchive) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsArchive) ResetRehydrationMaxScanSizeInGb() {
	_jsii_.InvokeVoid(
		l,
		"resetRehydrationMaxScanSizeInGb",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsArchive) ResetRehydrationTags() {
	_jsii_.InvokeVoid(
		l,
		"resetRehydrationTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsArchive) ResetS3Archive() {
	_jsii_.InvokeVoid(
		l,
		"resetS3Archive",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsArchive) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsArchive) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsArchive) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsArchive) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsArchive) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsArchive) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

