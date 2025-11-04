// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference interface {
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
	CustomLink() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionCustomLinkList
	CustomLinkInput() interface{}
	Event() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionEventList
	EventInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition
	SetInternalValue(val *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition)
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
	Marker() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionMarkerList
	MarkerInput() interface{}
	Request() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestList
	RequestInput() interface{}
	RightYaxis() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxisOutputReference
	RightYaxisInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxis
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
	Yaxis() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxisOutputReference
	YaxisInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxis
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutCustomLink(value interface{})
	PutEvent(value interface{})
	PutMarker(value interface{})
	PutRequest(value interface{})
	PutRightYaxis(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxis)
	PutYaxis(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxis)
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
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference
type jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) CustomLink() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionCustomLinkList {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionCustomLinkList
	_jsii_.Get(
		j,
		"customLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) CustomLinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Event() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionEventList {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionEventList
	_jsii_.Get(
		j,
		"event",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) EventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) InternalValue() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LegendColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"legendColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LegendColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"legendColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LegendLayout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LegendLayoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendLayoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LegendSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LegendSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legendSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) LiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Marker() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionMarkerList {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionMarkerList
	_jsii_.Get(
		j,
		"marker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) MarkerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"markerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Request() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestList {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestList
	_jsii_.Get(
		j,
		"request",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) RequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) RightYaxis() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxisOutputReference {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxisOutputReference
	_jsii_.Get(
		j,
		"rightYaxis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) RightYaxisInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxis {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxis
	_jsii_.Get(
		j,
		"rightYaxisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ShowLegend() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLegend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ShowLegendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLegendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Yaxis() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxisOutputReference {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxisOutputReference
	_jsii_.Get(
		j,
		"yaxis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) YaxisInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxis {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxis
	_jsii_.Get(
		j,
		"yaxisInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference_Override(d DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetInternalValue(val *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetLegendColumns(val *[]*string) {
	if err := j.validateSetLegendColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendColumns",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetLegendLayout(val *string) {
	if err := j.validateSetLegendLayoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendLayout",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetLegendSize(val *string) {
	if err := j.validateSetLegendSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendSize",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetLiveSpan(val *string) {
	if err := j.validateSetLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSpan",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetShowLegend(val interface{}) {
	if err := j.validateSetShowLegendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showLegend",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) PutCustomLink(value interface{}) {
	if err := d.validatePutCustomLinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomLink",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) PutEvent(value interface{}) {
	if err := d.validatePutEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEvent",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) PutMarker(value interface{}) {
	if err := d.validatePutMarkerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMarker",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) PutRequest(value interface{}) {
	if err := d.validatePutRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRequest",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) PutRightYaxis(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRightYaxis) {
	if err := d.validatePutRightYaxisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRightYaxis",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) PutYaxis(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionYaxis) {
	if err := d.validatePutYaxisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putYaxis",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetCustomLink() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomLink",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetEvent() {
	_jsii_.InvokeVoid(
		d,
		"resetEvent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetLegendColumns() {
	_jsii_.InvokeVoid(
		d,
		"resetLegendColumns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetLegendLayout() {
	_jsii_.InvokeVoid(
		d,
		"resetLegendLayout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetLegendSize() {
	_jsii_.InvokeVoid(
		d,
		"resetLegendSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetLiveSpan() {
	_jsii_.InvokeVoid(
		d,
		"resetLiveSpan",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetMarker() {
	_jsii_.InvokeVoid(
		d,
		"resetMarker",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetRequest() {
	_jsii_.InvokeVoid(
		d,
		"resetRequest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetRightYaxis() {
	_jsii_.InvokeVoid(
		d,
		"resetRightYaxis",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetShowLegend() {
	_jsii_.InvokeVoid(
		d,
		"resetShowLegend",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ResetYaxis() {
	_jsii_.InvokeVoid(
		d,
		"resetYaxis",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

