package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v4/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v4/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetTraceServiceDefinitionOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisplayFormat() *string
	SetDisplayFormat(val *string)
	DisplayFormatInput() *string
	Env() *string
	SetEnv(val *string)
	EnvInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardWidgetTraceServiceDefinition
	SetInternalValue(val *DashboardWidgetTraceServiceDefinition)
	LiveSpan() *string
	SetLiveSpan(val *string)
	LiveSpanInput() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	ShowBreakdown() interface{}
	SetShowBreakdown(val interface{})
	ShowBreakdownInput() interface{}
	ShowDistribution() interface{}
	SetShowDistribution(val interface{})
	ShowDistributionInput() interface{}
	ShowErrors() interface{}
	SetShowErrors(val interface{})
	ShowErrorsInput() interface{}
	ShowHits() interface{}
	SetShowHits(val interface{})
	ShowHitsInput() interface{}
	ShowLatency() interface{}
	SetShowLatency(val interface{})
	ShowLatencyInput() interface{}
	ShowResourceList() interface{}
	SetShowResourceList(val interface{})
	ShowResourceListInput() interface{}
	SizeFormat() *string
	SetSizeFormat(val *string)
	SizeFormatInput() *string
	SpanName() *string
	SetSpanName(val *string)
	SpanNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() *string
	SetTitle(val *string)
	TitleAlign() *string
	SetTitleAlign(val *string)
	TitleAlignInput() *string
	TitleInput() *string
	TitleSize() *string
	SetTitleSize(val *string)
	TitleSizeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDisplayFormat()
	ResetLiveSpan()
	ResetShowBreakdown()
	ResetShowDistribution()
	ResetShowErrors()
	ResetShowHits()
	ResetShowLatency()
	ResetShowResourceList()
	ResetSizeFormat()
	ResetTitle()
	ResetTitleAlign()
	ResetTitleSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetTraceServiceDefinitionOutputReference
type jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) DisplayFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) DisplayFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) Env() *string {
	var returns *string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) EnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) InternalValue() *DashboardWidgetTraceServiceDefinition {
	var returns *DashboardWidgetTraceServiceDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) LiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) LiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ShowBreakdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showBreakdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ShowBreakdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showBreakdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ShowDistribution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ShowDistributionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showDistributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ShowErrors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ShowErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ShowHits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showHits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ShowHitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showHitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ShowLatency() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ShowLatencyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ShowResourceList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showResourceList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ShowResourceListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showResourceListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) SizeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) SizeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) SpanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spanName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) SpanNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spanNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetTraceServiceDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetTraceServiceDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetTraceServiceDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetTraceServiceDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetTraceServiceDefinitionOutputReference_Override(d DashboardWidgetTraceServiceDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetTraceServiceDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetDisplayFormat(val *string) {
	if err := j.validateSetDisplayFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayFormat",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetEnv(val *string) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetInternalValue(val *DashboardWidgetTraceServiceDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetLiveSpan(val *string) {
	if err := j.validateSetLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSpan",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetShowBreakdown(val interface{}) {
	if err := j.validateSetShowBreakdownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showBreakdown",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetShowDistribution(val interface{}) {
	if err := j.validateSetShowDistributionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showDistribution",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetShowErrors(val interface{}) {
	if err := j.validateSetShowErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showErrors",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetShowHits(val interface{}) {
	if err := j.validateSetShowHitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showHits",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetShowLatency(val interface{}) {
	if err := j.validateSetShowLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showLatency",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetShowResourceList(val interface{}) {
	if err := j.validateSetShowResourceListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showResourceList",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetSizeFormat(val *string) {
	if err := j.validateSetSizeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeFormat",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetSpanName(val *string) {
	if err := j.validateSetSpanNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spanName",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ResetDisplayFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ResetLiveSpan() {
	_jsii_.InvokeVoid(
		d,
		"resetLiveSpan",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ResetShowBreakdown() {
	_jsii_.InvokeVoid(
		d,
		"resetShowBreakdown",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ResetShowDistribution() {
	_jsii_.InvokeVoid(
		d,
		"resetShowDistribution",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ResetShowErrors() {
	_jsii_.InvokeVoid(
		d,
		"resetShowErrors",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ResetShowHits() {
	_jsii_.InvokeVoid(
		d,
		"resetShowHits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ResetShowLatency() {
	_jsii_.InvokeVoid(
		d,
		"resetShowLatency",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ResetShowResourceList() {
	_jsii_.InvokeVoid(
		d,
		"resetShowResourceList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ResetSizeFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetSizeFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetTraceServiceDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

