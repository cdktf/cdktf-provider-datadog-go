// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevelobjective

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/servicelevelobjective/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/service_level_objective datadog_service_level_objective}.
type ServiceLevelObjective interface {
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
	Groups() *[]*string
	SetGroups(val *[]*string)
	GroupsInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MonitorIds() *[]*float64
	SetMonitorIds(val *[]*float64)
	MonitorIdsInput() *[]*float64
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
	Query() ServiceLevelObjectiveQueryOutputReference
	QueryInput() *ServiceLevelObjectiveQuery
	// Experimental.
	RawOverrides() interface{}
	SliSpecification() ServiceLevelObjectiveSliSpecificationOutputReference
	SliSpecificationInput() *ServiceLevelObjectiveSliSpecification
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	TargetThreshold() *float64
	SetTargetThreshold(val *float64)
	TargetThresholdInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Thresholds() ServiceLevelObjectiveThresholdsList
	ThresholdsInput() interface{}
	Timeframe() *string
	SetTimeframe(val *string)
	TimeframeInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Validate() interface{}
	SetValidate(val interface{})
	ValidateInput() interface{}
	WarningThreshold() *float64
	SetWarningThreshold(val *float64)
	WarningThresholdInput() *float64
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
	PutQuery(value *ServiceLevelObjectiveQuery)
	PutSliSpecification(value *ServiceLevelObjectiveSliSpecification)
	PutThresholds(value interface{})
	ResetDescription()
	ResetForceDelete()
	ResetGroups()
	ResetId()
	ResetMonitorIds()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQuery()
	ResetSliSpecification()
	ResetTags()
	ResetTargetThreshold()
	ResetTimeframe()
	ResetValidate()
	ResetWarningThreshold()
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

// The jsii proxy struct for ServiceLevelObjective
type jsiiProxy_ServiceLevelObjective struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServiceLevelObjective) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) ForceDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) ForceDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) GroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) MonitorIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"monitorIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) MonitorIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"monitorIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Query() ServiceLevelObjectiveQueryOutputReference {
	var returns ServiceLevelObjectiveQueryOutputReference
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) QueryInput() *ServiceLevelObjectiveQuery {
	var returns *ServiceLevelObjectiveQuery
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) SliSpecification() ServiceLevelObjectiveSliSpecificationOutputReference {
	var returns ServiceLevelObjectiveSliSpecificationOutputReference
	_jsii_.Get(
		j,
		"sliSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) SliSpecificationInput() *ServiceLevelObjectiveSliSpecification {
	var returns *ServiceLevelObjectiveSliSpecification
	_jsii_.Get(
		j,
		"sliSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) TargetThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) TargetThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Thresholds() ServiceLevelObjectiveThresholdsList {
	var returns ServiceLevelObjectiveThresholdsList
	_jsii_.Get(
		j,
		"thresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) ThresholdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Timeframe() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeframe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) TimeframeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeframeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) Validate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) ValidateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) WarningThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warningThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjective) WarningThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warningThresholdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/service_level_objective datadog_service_level_objective} Resource.
func NewServiceLevelObjective(scope constructs.Construct, id *string, config *ServiceLevelObjectiveConfig) ServiceLevelObjective {
	_init_.Initialize()

	if err := validateNewServiceLevelObjectiveParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceLevelObjective{}

	_jsii_.Create(
		"@cdktf/provider-datadog.serviceLevelObjective.ServiceLevelObjective",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/service_level_objective datadog_service_level_objective} Resource.
func NewServiceLevelObjective_Override(s ServiceLevelObjective, scope constructs.Construct, id *string, config *ServiceLevelObjectiveConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.serviceLevelObjective.ServiceLevelObjective",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetForceDelete(val interface{}) {
	if err := j.validateSetForceDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDelete",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetGroups(val *[]*string) {
	if err := j.validateSetGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetMonitorIds(val *[]*float64) {
	if err := j.validateSetMonitorIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitorIds",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetTargetThreshold(val *float64) {
	if err := j.validateSetTargetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetThreshold",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetTimeframe(val *string) {
	if err := j.validateSetTimeframeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeframe",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetValidate(val interface{}) {
	if err := j.validateSetValidateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validate",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjective)SetWarningThreshold(val *float64) {
	if err := j.validateSetWarningThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warningThreshold",
		val,
	)
}

// Generates CDKTF code for importing a ServiceLevelObjective resource upon running "cdktf plan <stack-name>".
func ServiceLevelObjective_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServiceLevelObjective_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.serviceLevelObjective.ServiceLevelObjective",
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
func ServiceLevelObjective_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceLevelObjective_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.serviceLevelObjective.ServiceLevelObjective",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceLevelObjective_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceLevelObjective_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.serviceLevelObjective.ServiceLevelObjective",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceLevelObjective_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceLevelObjective_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.serviceLevelObjective.ServiceLevelObjective",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServiceLevelObjective_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.serviceLevelObjective.ServiceLevelObjective",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_ServiceLevelObjective) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServiceLevelObjective) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_ServiceLevelObjective) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServiceLevelObjective) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_ServiceLevelObjective) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServiceLevelObjective) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServiceLevelObjective) PutQuery(value *ServiceLevelObjectiveQuery) {
	if err := s.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putQuery",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceLevelObjective) PutSliSpecification(value *ServiceLevelObjectiveSliSpecification) {
	if err := s.validatePutSliSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSliSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceLevelObjective) PutThresholds(value interface{}) {
	if err := s.validatePutThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putThresholds",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceLevelObjective) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelObjective) ResetForceDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetForceDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelObjective) ResetGroups() {
	_jsii_.InvokeVoid(
		s,
		"resetGroups",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelObjective) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelObjective) ResetMonitorIds() {
	_jsii_.InvokeVoid(
		s,
		"resetMonitorIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelObjective) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelObjective) ResetQuery() {
	_jsii_.InvokeVoid(
		s,
		"resetQuery",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelObjective) ResetSliSpecification() {
	_jsii_.InvokeVoid(
		s,
		"resetSliSpecification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelObjective) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelObjective) ResetTargetThreshold() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetThreshold",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelObjective) ResetTimeframe() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeframe",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelObjective) ResetValidate() {
	_jsii_.InvokeVoid(
		s,
		"resetValidate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelObjective) ResetWarningThreshold() {
	_jsii_.InvokeVoid(
		s,
		"resetWarningThreshold",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelObjective) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjective) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

