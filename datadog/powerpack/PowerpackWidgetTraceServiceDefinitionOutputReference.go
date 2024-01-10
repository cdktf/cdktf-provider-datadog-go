// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetTraceServiceDefinitionOutputReference interface {
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
	InternalValue() *PowerpackWidgetTraceServiceDefinition
	SetInternalValue(val *PowerpackWidgetTraceServiceDefinition)
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

// The jsii proxy struct for PowerpackWidgetTraceServiceDefinitionOutputReference
type jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) DisplayFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) DisplayFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) Env() *string {
	var returns *string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) EnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) InternalValue() *PowerpackWidgetTraceServiceDefinition {
	var returns *PowerpackWidgetTraceServiceDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) LiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) LiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ShowBreakdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showBreakdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ShowBreakdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showBreakdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ShowDistribution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ShowDistributionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showDistributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ShowErrors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ShowErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ShowHits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showHits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ShowHitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showHitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ShowLatency() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ShowLatencyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ShowResourceList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showResourceList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ShowResourceListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showResourceListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) SizeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) SizeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) SpanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spanName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) SpanNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spanNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetTraceServiceDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetTraceServiceDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetTraceServiceDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetTraceServiceDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetTraceServiceDefinitionOutputReference_Override(p PowerpackWidgetTraceServiceDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetTraceServiceDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetDisplayFormat(val *string) {
	if err := j.validateSetDisplayFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayFormat",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetEnv(val *string) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetInternalValue(val *PowerpackWidgetTraceServiceDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetLiveSpan(val *string) {
	if err := j.validateSetLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSpan",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetShowBreakdown(val interface{}) {
	if err := j.validateSetShowBreakdownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showBreakdown",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetShowDistribution(val interface{}) {
	if err := j.validateSetShowDistributionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showDistribution",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetShowErrors(val interface{}) {
	if err := j.validateSetShowErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showErrors",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetShowHits(val interface{}) {
	if err := j.validateSetShowHitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showHits",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetShowLatency(val interface{}) {
	if err := j.validateSetShowLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showLatency",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetShowResourceList(val interface{}) {
	if err := j.validateSetShowResourceListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showResourceList",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetSizeFormat(val *string) {
	if err := j.validateSetSizeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeFormat",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetSpanName(val *string) {
	if err := j.validateSetSpanNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spanName",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ResetDisplayFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetDisplayFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ResetLiveSpan() {
	_jsii_.InvokeVoid(
		p,
		"resetLiveSpan",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ResetShowBreakdown() {
	_jsii_.InvokeVoid(
		p,
		"resetShowBreakdown",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ResetShowDistribution() {
	_jsii_.InvokeVoid(
		p,
		"resetShowDistribution",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ResetShowErrors() {
	_jsii_.InvokeVoid(
		p,
		"resetShowErrors",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ResetShowHits() {
	_jsii_.InvokeVoid(
		p,
		"resetShowHits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ResetShowLatency() {
	_jsii_.InvokeVoid(
		p,
		"resetShowLatency",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ResetShowResourceList() {
	_jsii_.InvokeVoid(
		p,
		"resetShowResourceList",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ResetSizeFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetSizeFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		p,
		"resetTitle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTraceServiceDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

