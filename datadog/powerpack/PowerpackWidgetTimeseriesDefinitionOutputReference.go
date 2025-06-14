// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetTimeseriesDefinitionOutputReference interface {
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
	CustomLink() PowerpackWidgetTimeseriesDefinitionCustomLinkList
	CustomLinkInput() interface{}
	Event() PowerpackWidgetTimeseriesDefinitionEventList
	EventInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackWidgetTimeseriesDefinition
	SetInternalValue(val *PowerpackWidgetTimeseriesDefinition)
	LegendColumns() *[]*string
	SetLegendColumns(val *[]*string)
	LegendColumnsInput() *[]*string
	LegendLayout() *string
	SetLegendLayout(val *string)
	LegendLayoutInput() *string
	LegendSize() *string
	SetLegendSize(val *string)
	LegendSizeInput() *string
	LiveSpan() *string
	SetLiveSpan(val *string)
	LiveSpanInput() *string
	Marker() PowerpackWidgetTimeseriesDefinitionMarkerList
	MarkerInput() interface{}
	Request() PowerpackWidgetTimeseriesDefinitionRequestList
	RequestInput() interface{}
	RightYaxis() PowerpackWidgetTimeseriesDefinitionRightYaxisOutputReference
	RightYaxisInput() *PowerpackWidgetTimeseriesDefinitionRightYaxis
	ShowLegend() interface{}
	SetShowLegend(val interface{})
	ShowLegendInput() interface{}
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
	Yaxis() PowerpackWidgetTimeseriesDefinitionYaxisOutputReference
	YaxisInput() *PowerpackWidgetTimeseriesDefinitionYaxis
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
	PutCustomLink(value interface{})
	PutEvent(value interface{})
	PutMarker(value interface{})
	PutRequest(value interface{})
	PutRightYaxis(value *PowerpackWidgetTimeseriesDefinitionRightYaxis)
	PutYaxis(value *PowerpackWidgetTimeseriesDefinitionYaxis)
	ResetCustomLink()
	ResetEvent()
	ResetLegendColumns()
	ResetLegendLayout()
	ResetLegendSize()
	ResetLiveSpan()
	ResetMarker()
	ResetRequest()
	ResetRightYaxis()
	ResetShowLegend()
	ResetTitle()
	ResetTitleAlign()
	ResetTitleSize()
	ResetYaxis()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetTimeseriesDefinitionOutputReference
type jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) CustomLink() PowerpackWidgetTimeseriesDefinitionCustomLinkList {
	var returns PowerpackWidgetTimeseriesDefinitionCustomLinkList
	_jsii_.Get(
		j,
		"customLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) CustomLinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) Event() PowerpackWidgetTimeseriesDefinitionEventList {
	var returns PowerpackWidgetTimeseriesDefinitionEventList
	_jsii_.Get(
		j,
		"event",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) EventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) InternalValue() *PowerpackWidgetTimeseriesDefinition {
	var returns *PowerpackWidgetTimeseriesDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) LegendColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"legendColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) LegendColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"legendColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) LegendLayout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) LegendLayoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendLayoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) LegendSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) LegendSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) LiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) LiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) Marker() PowerpackWidgetTimeseriesDefinitionMarkerList {
	var returns PowerpackWidgetTimeseriesDefinitionMarkerList
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) MarkerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"markerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) Request() PowerpackWidgetTimeseriesDefinitionRequestList {
	var returns PowerpackWidgetTimeseriesDefinitionRequestList
	_jsii_.Get(
		j,
		"request",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) RequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) RightYaxis() PowerpackWidgetTimeseriesDefinitionRightYaxisOutputReference {
	var returns PowerpackWidgetTimeseriesDefinitionRightYaxisOutputReference
	_jsii_.Get(
		j,
		"rightYaxis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) RightYaxisInput() *PowerpackWidgetTimeseriesDefinitionRightYaxis {
	var returns *PowerpackWidgetTimeseriesDefinitionRightYaxis
	_jsii_.Get(
		j,
		"rightYaxisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ShowLegend() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLegend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ShowLegendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLegendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) Yaxis() PowerpackWidgetTimeseriesDefinitionYaxisOutputReference {
	var returns PowerpackWidgetTimeseriesDefinitionYaxisOutputReference
	_jsii_.Get(
		j,
		"yaxis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) YaxisInput() *PowerpackWidgetTimeseriesDefinitionYaxis {
	var returns *PowerpackWidgetTimeseriesDefinitionYaxis
	_jsii_.Get(
		j,
		"yaxisInput",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetTimeseriesDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetTimeseriesDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetTimeseriesDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetTimeseriesDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetTimeseriesDefinitionOutputReference_Override(p PowerpackWidgetTimeseriesDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetTimeseriesDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference)SetInternalValue(val *PowerpackWidgetTimeseriesDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference)SetLegendColumns(val *[]*string) {
	if err := j.validateSetLegendColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendColumns",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference)SetLegendLayout(val *string) {
	if err := j.validateSetLegendLayoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendLayout",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference)SetLegendSize(val *string) {
	if err := j.validateSetLegendSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendSize",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference)SetLiveSpan(val *string) {
	if err := j.validateSetLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSpan",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference)SetShowLegend(val interface{}) {
	if err := j.validateSetShowLegendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showLegend",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) PutCustomLink(value interface{}) {
	if err := p.validatePutCustomLinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCustomLink",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) PutEvent(value interface{}) {
	if err := p.validatePutEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEvent",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) PutMarker(value interface{}) {
	if err := p.validatePutMarkerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMarker",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) PutRequest(value interface{}) {
	if err := p.validatePutRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRequest",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) PutRightYaxis(value *PowerpackWidgetTimeseriesDefinitionRightYaxis) {
	if err := p.validatePutRightYaxisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRightYaxis",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) PutYaxis(value *PowerpackWidgetTimeseriesDefinitionYaxis) {
	if err := p.validatePutYaxisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putYaxis",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetCustomLink() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomLink",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetEvent() {
	_jsii_.InvokeVoid(
		p,
		"resetEvent",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetLegendColumns() {
	_jsii_.InvokeVoid(
		p,
		"resetLegendColumns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetLegendLayout() {
	_jsii_.InvokeVoid(
		p,
		"resetLegendLayout",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetLegendSize() {
	_jsii_.InvokeVoid(
		p,
		"resetLegendSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetLiveSpan() {
	_jsii_.InvokeVoid(
		p,
		"resetLiveSpan",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetMarker() {
	_jsii_.InvokeVoid(
		p,
		"resetMarker",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetRightYaxis() {
	_jsii_.InvokeVoid(
		p,
		"resetRightYaxis",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetShowLegend() {
	_jsii_.InvokeVoid(
		p,
		"resetShowLegend",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		p,
		"resetTitle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		p,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ResetYaxis() {
	_jsii_.InvokeVoid(
		p,
		"resetYaxis",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

