// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/datadog/r/dashboard datadog_dashboard}.
type Dashboard interface {
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DashboardLists() *[]*float64
	SetDashboardLists(val *[]*float64)
	DashboardListsInput() *[]*float64
	DashboardListsRemoved() *[]*float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	IsReadOnly() interface{}
	SetIsReadOnly(val interface{})
	IsReadOnlyInput() interface{}
	LayoutType() *string
	SetLayoutType(val *string)
	LayoutTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NotifyList() *[]*string
	SetNotifyList(val *[]*string)
	NotifyListInput() *[]*string
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
	ReflowType() *string
	SetReflowType(val *string)
	ReflowTypeInput() *string
	RestrictedRoles() *[]*string
	SetRestrictedRoles(val *[]*string)
	RestrictedRolesInput() *[]*string
	TemplateVariable() DashboardTemplateVariableList
	TemplateVariableInput() interface{}
	TemplateVariablePreset() DashboardTemplateVariablePresetList
	TemplateVariablePresetInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	Widget() DashboardWidgetList
	WidgetInput() interface{}
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
	PutTemplateVariable(value interface{})
	PutTemplateVariablePreset(value interface{})
	PutWidget(value interface{})
	ResetDashboardLists()
	ResetDescription()
	ResetId()
	ResetIsReadOnly()
	ResetNotifyList()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReflowType()
	ResetRestrictedRoles()
	ResetTemplateVariable()
	ResetTemplateVariablePreset()
	ResetUrl()
	ResetWidget()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Dashboard
type jsiiProxy_Dashboard struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Dashboard) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DashboardLists() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"dashboardLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DashboardListsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"dashboardListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DashboardListsRemoved() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"dashboardListsRemoved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) IsReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) IsReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) LayoutType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layoutType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) LayoutTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layoutTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) NotifyList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifyList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) NotifyListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifyListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) ReflowType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reflowType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) ReflowTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reflowTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) RestrictedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) RestrictedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) TemplateVariable() DashboardTemplateVariableList {
	var returns DashboardTemplateVariableList
	_jsii_.Get(
		j,
		"templateVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) TemplateVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) TemplateVariablePreset() DashboardTemplateVariablePresetList {
	var returns DashboardTemplateVariablePresetList
	_jsii_.Get(
		j,
		"templateVariablePreset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) TemplateVariablePresetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateVariablePresetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) Widget() DashboardWidgetList {
	var returns DashboardWidgetList
	_jsii_.Get(
		j,
		"widget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Dashboard) WidgetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/datadog/r/dashboard datadog_dashboard} Resource.
func NewDashboard(scope constructs.Construct, id *string, config *DashboardConfig) Dashboard {
	_init_.Initialize()

	j := jsiiProxy_Dashboard{}

	_jsii_.Create(
		"@cdktf/provider-datadog.Dashboard",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/datadog/r/dashboard datadog_dashboard} Resource.
func NewDashboard_Override(d Dashboard, scope constructs.Construct, id *string, config *DashboardConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.Dashboard",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_Dashboard) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetDashboardLists(val *[]*float64) {
	_jsii_.Set(
		j,
		"dashboardLists",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetIsReadOnly(val interface{}) {
	_jsii_.Set(
		j,
		"isReadOnly",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetLayoutType(val *string) {
	_jsii_.Set(
		j,
		"layoutType",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetNotifyList(val *[]*string) {
	_jsii_.Set(
		j,
		"notifyList",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetReflowType(val *string) {
	_jsii_.Set(
		j,
		"reflowType",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetRestrictedRoles(val *[]*string) {
	_jsii_.Set(
		j,
		"restrictedRoles",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetTitle(val *string) {
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_Dashboard) SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
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
func Dashboard_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-datadog.Dashboard",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Dashboard_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-datadog.Dashboard",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_Dashboard) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_Dashboard) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_Dashboard) PutTemplateVariable(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putTemplateVariable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_Dashboard) PutTemplateVariablePreset(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putTemplateVariablePreset",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_Dashboard) PutWidget(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putWidget",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_Dashboard) ResetDashboardLists() {
	_jsii_.InvokeVoid(
		d,
		"resetDashboardLists",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetIsReadOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetIsReadOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetNotifyList() {
	_jsii_.InvokeVoid(
		d,
		"resetNotifyList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetReflowType() {
	_jsii_.InvokeVoid(
		d,
		"resetReflowType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetRestrictedRoles() {
	_jsii_.InvokeVoid(
		d,
		"resetRestrictedRoles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetTemplateVariable() {
	_jsii_.InvokeVoid(
		d,
		"resetTemplateVariable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetTemplateVariablePreset() {
	_jsii_.InvokeVoid(
		d,
		"resetTemplateVariablePreset",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) ResetWidget() {
	_jsii_.InvokeVoid(
		d,
		"resetWidget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Dashboard) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
