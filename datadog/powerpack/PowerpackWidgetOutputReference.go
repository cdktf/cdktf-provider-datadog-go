// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetOutputReference interface {
	cdktf.ComplexObject
	AlertGraphDefinition() PowerpackWidgetAlertGraphDefinitionOutputReference
	AlertGraphDefinitionInput() *PowerpackWidgetAlertGraphDefinition
	AlertValueDefinition() PowerpackWidgetAlertValueDefinitionOutputReference
	AlertValueDefinitionInput() *PowerpackWidgetAlertValueDefinition
	ChangeDefinition() PowerpackWidgetChangeDefinitionOutputReference
	ChangeDefinitionInput() *PowerpackWidgetChangeDefinition
	CheckStatusDefinition() PowerpackWidgetCheckStatusDefinitionOutputReference
	CheckStatusDefinitionInput() *PowerpackWidgetCheckStatusDefinition
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
	DistributionDefinition() PowerpackWidgetDistributionDefinitionOutputReference
	DistributionDefinitionInput() *PowerpackWidgetDistributionDefinition
	EventStreamDefinition() PowerpackWidgetEventStreamDefinitionOutputReference
	EventStreamDefinitionInput() *PowerpackWidgetEventStreamDefinition
	EventTimelineDefinition() PowerpackWidgetEventTimelineDefinitionOutputReference
	EventTimelineDefinitionInput() *PowerpackWidgetEventTimelineDefinition
	// Experimental.
	Fqn() *string
	FreeTextDefinition() PowerpackWidgetFreeTextDefinitionOutputReference
	FreeTextDefinitionInput() *PowerpackWidgetFreeTextDefinition
	GeomapDefinition() PowerpackWidgetGeomapDefinitionOutputReference
	GeomapDefinitionInput() *PowerpackWidgetGeomapDefinition
	HeatmapDefinition() PowerpackWidgetHeatmapDefinitionOutputReference
	HeatmapDefinitionInput() *PowerpackWidgetHeatmapDefinition
	HostmapDefinition() PowerpackWidgetHostmapDefinitionOutputReference
	HostmapDefinitionInput() *PowerpackWidgetHostmapDefinition
	Id() *float64
	IframeDefinition() PowerpackWidgetIframeDefinitionOutputReference
	IframeDefinitionInput() *PowerpackWidgetIframeDefinition
	ImageDefinition() PowerpackWidgetImageDefinitionOutputReference
	ImageDefinitionInput() *PowerpackWidgetImageDefinition
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ListStreamDefinition() PowerpackWidgetListStreamDefinitionOutputReference
	ListStreamDefinitionInput() *PowerpackWidgetListStreamDefinition
	LogStreamDefinition() PowerpackWidgetLogStreamDefinitionOutputReference
	LogStreamDefinitionInput() *PowerpackWidgetLogStreamDefinition
	ManageStatusDefinition() PowerpackWidgetManageStatusDefinitionOutputReference
	ManageStatusDefinitionInput() *PowerpackWidgetManageStatusDefinition
	NoteDefinition() PowerpackWidgetNoteDefinitionOutputReference
	NoteDefinitionInput() *PowerpackWidgetNoteDefinition
	QueryTableDefinition() PowerpackWidgetQueryTableDefinitionOutputReference
	QueryTableDefinitionInput() *PowerpackWidgetQueryTableDefinition
	QueryValueDefinition() PowerpackWidgetQueryValueDefinitionOutputReference
	QueryValueDefinitionInput() *PowerpackWidgetQueryValueDefinition
	RunWorkflowDefinition() PowerpackWidgetRunWorkflowDefinitionOutputReference
	RunWorkflowDefinitionInput() *PowerpackWidgetRunWorkflowDefinition
	ScatterplotDefinition() PowerpackWidgetScatterplotDefinitionOutputReference
	ScatterplotDefinitionInput() *PowerpackWidgetScatterplotDefinition
	ServiceLevelObjectiveDefinition() PowerpackWidgetServiceLevelObjectiveDefinitionOutputReference
	ServiceLevelObjectiveDefinitionInput() *PowerpackWidgetServiceLevelObjectiveDefinition
	ServicemapDefinition() PowerpackWidgetServicemapDefinitionOutputReference
	ServicemapDefinitionInput() *PowerpackWidgetServicemapDefinition
	SloListDefinition() PowerpackWidgetSloListDefinitionOutputReference
	SloListDefinitionInput() *PowerpackWidgetSloListDefinition
	SunburstDefinition() PowerpackWidgetSunburstDefinitionOutputReference
	SunburstDefinitionInput() *PowerpackWidgetSunburstDefinition
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeseriesDefinition() PowerpackWidgetTimeseriesDefinitionOutputReference
	TimeseriesDefinitionInput() *PowerpackWidgetTimeseriesDefinition
	ToplistDefinition() PowerpackWidgetToplistDefinitionOutputReference
	ToplistDefinitionInput() *PowerpackWidgetToplistDefinition
	TopologyMapDefinition() PowerpackWidgetTopologyMapDefinitionOutputReference
	TopologyMapDefinitionInput() *PowerpackWidgetTopologyMapDefinition
	TraceServiceDefinition() PowerpackWidgetTraceServiceDefinitionOutputReference
	TraceServiceDefinitionInput() *PowerpackWidgetTraceServiceDefinition
	TreemapDefinition() PowerpackWidgetTreemapDefinitionOutputReference
	TreemapDefinitionInput() *PowerpackWidgetTreemapDefinition
	WidgetLayout() PowerpackWidgetWidgetLayoutOutputReference
	WidgetLayoutInput() *PowerpackWidgetWidgetLayout
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
	PutAlertGraphDefinition(value *PowerpackWidgetAlertGraphDefinition)
	PutAlertValueDefinition(value *PowerpackWidgetAlertValueDefinition)
	PutChangeDefinition(value *PowerpackWidgetChangeDefinition)
	PutCheckStatusDefinition(value *PowerpackWidgetCheckStatusDefinition)
	PutDistributionDefinition(value *PowerpackWidgetDistributionDefinition)
	PutEventStreamDefinition(value *PowerpackWidgetEventStreamDefinition)
	PutEventTimelineDefinition(value *PowerpackWidgetEventTimelineDefinition)
	PutFreeTextDefinition(value *PowerpackWidgetFreeTextDefinition)
	PutGeomapDefinition(value *PowerpackWidgetGeomapDefinition)
	PutHeatmapDefinition(value *PowerpackWidgetHeatmapDefinition)
	PutHostmapDefinition(value *PowerpackWidgetHostmapDefinition)
	PutIframeDefinition(value *PowerpackWidgetIframeDefinition)
	PutImageDefinition(value *PowerpackWidgetImageDefinition)
	PutListStreamDefinition(value *PowerpackWidgetListStreamDefinition)
	PutLogStreamDefinition(value *PowerpackWidgetLogStreamDefinition)
	PutManageStatusDefinition(value *PowerpackWidgetManageStatusDefinition)
	PutNoteDefinition(value *PowerpackWidgetNoteDefinition)
	PutQueryTableDefinition(value *PowerpackWidgetQueryTableDefinition)
	PutQueryValueDefinition(value *PowerpackWidgetQueryValueDefinition)
	PutRunWorkflowDefinition(value *PowerpackWidgetRunWorkflowDefinition)
	PutScatterplotDefinition(value *PowerpackWidgetScatterplotDefinition)
	PutServiceLevelObjectiveDefinition(value *PowerpackWidgetServiceLevelObjectiveDefinition)
	PutServicemapDefinition(value *PowerpackWidgetServicemapDefinition)
	PutSloListDefinition(value *PowerpackWidgetSloListDefinition)
	PutSunburstDefinition(value *PowerpackWidgetSunburstDefinition)
	PutTimeseriesDefinition(value *PowerpackWidgetTimeseriesDefinition)
	PutToplistDefinition(value *PowerpackWidgetToplistDefinition)
	PutTopologyMapDefinition(value *PowerpackWidgetTopologyMapDefinition)
	PutTraceServiceDefinition(value *PowerpackWidgetTraceServiceDefinition)
	PutTreemapDefinition(value *PowerpackWidgetTreemapDefinition)
	PutWidgetLayout(value *PowerpackWidgetWidgetLayout)
	ResetAlertGraphDefinition()
	ResetAlertValueDefinition()
	ResetChangeDefinition()
	ResetCheckStatusDefinition()
	ResetDistributionDefinition()
	ResetEventStreamDefinition()
	ResetEventTimelineDefinition()
	ResetFreeTextDefinition()
	ResetGeomapDefinition()
	ResetHeatmapDefinition()
	ResetHostmapDefinition()
	ResetIframeDefinition()
	ResetImageDefinition()
	ResetListStreamDefinition()
	ResetLogStreamDefinition()
	ResetManageStatusDefinition()
	ResetNoteDefinition()
	ResetQueryTableDefinition()
	ResetQueryValueDefinition()
	ResetRunWorkflowDefinition()
	ResetScatterplotDefinition()
	ResetServiceLevelObjectiveDefinition()
	ResetServicemapDefinition()
	ResetSloListDefinition()
	ResetSunburstDefinition()
	ResetTimeseriesDefinition()
	ResetToplistDefinition()
	ResetTopologyMapDefinition()
	ResetTraceServiceDefinition()
	ResetTreemapDefinition()
	ResetWidgetLayout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetOutputReference
type jsiiProxy_PowerpackWidgetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) AlertGraphDefinition() PowerpackWidgetAlertGraphDefinitionOutputReference {
	var returns PowerpackWidgetAlertGraphDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertGraphDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) AlertGraphDefinitionInput() *PowerpackWidgetAlertGraphDefinition {
	var returns *PowerpackWidgetAlertGraphDefinition
	_jsii_.Get(
		j,
		"alertGraphDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) AlertValueDefinition() PowerpackWidgetAlertValueDefinitionOutputReference {
	var returns PowerpackWidgetAlertValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) AlertValueDefinitionInput() *PowerpackWidgetAlertValueDefinition {
	var returns *PowerpackWidgetAlertValueDefinition
	_jsii_.Get(
		j,
		"alertValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ChangeDefinition() PowerpackWidgetChangeDefinitionOutputReference {
	var returns PowerpackWidgetChangeDefinitionOutputReference
	_jsii_.Get(
		j,
		"changeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ChangeDefinitionInput() *PowerpackWidgetChangeDefinition {
	var returns *PowerpackWidgetChangeDefinition
	_jsii_.Get(
		j,
		"changeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) CheckStatusDefinition() PowerpackWidgetCheckStatusDefinitionOutputReference {
	var returns PowerpackWidgetCheckStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"checkStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) CheckStatusDefinitionInput() *PowerpackWidgetCheckStatusDefinition {
	var returns *PowerpackWidgetCheckStatusDefinition
	_jsii_.Get(
		j,
		"checkStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) DistributionDefinition() PowerpackWidgetDistributionDefinitionOutputReference {
	var returns PowerpackWidgetDistributionDefinitionOutputReference
	_jsii_.Get(
		j,
		"distributionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) DistributionDefinitionInput() *PowerpackWidgetDistributionDefinition {
	var returns *PowerpackWidgetDistributionDefinition
	_jsii_.Get(
		j,
		"distributionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) EventStreamDefinition() PowerpackWidgetEventStreamDefinitionOutputReference {
	var returns PowerpackWidgetEventStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) EventStreamDefinitionInput() *PowerpackWidgetEventStreamDefinition {
	var returns *PowerpackWidgetEventStreamDefinition
	_jsii_.Get(
		j,
		"eventStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) EventTimelineDefinition() PowerpackWidgetEventTimelineDefinitionOutputReference {
	var returns PowerpackWidgetEventTimelineDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventTimelineDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) EventTimelineDefinitionInput() *PowerpackWidgetEventTimelineDefinition {
	var returns *PowerpackWidgetEventTimelineDefinition
	_jsii_.Get(
		j,
		"eventTimelineDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) FreeTextDefinition() PowerpackWidgetFreeTextDefinitionOutputReference {
	var returns PowerpackWidgetFreeTextDefinitionOutputReference
	_jsii_.Get(
		j,
		"freeTextDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) FreeTextDefinitionInput() *PowerpackWidgetFreeTextDefinition {
	var returns *PowerpackWidgetFreeTextDefinition
	_jsii_.Get(
		j,
		"freeTextDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) GeomapDefinition() PowerpackWidgetGeomapDefinitionOutputReference {
	var returns PowerpackWidgetGeomapDefinitionOutputReference
	_jsii_.Get(
		j,
		"geomapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) GeomapDefinitionInput() *PowerpackWidgetGeomapDefinition {
	var returns *PowerpackWidgetGeomapDefinition
	_jsii_.Get(
		j,
		"geomapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) HeatmapDefinition() PowerpackWidgetHeatmapDefinitionOutputReference {
	var returns PowerpackWidgetHeatmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"heatmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) HeatmapDefinitionInput() *PowerpackWidgetHeatmapDefinition {
	var returns *PowerpackWidgetHeatmapDefinition
	_jsii_.Get(
		j,
		"heatmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) HostmapDefinition() PowerpackWidgetHostmapDefinitionOutputReference {
	var returns PowerpackWidgetHostmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"hostmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) HostmapDefinitionInput() *PowerpackWidgetHostmapDefinition {
	var returns *PowerpackWidgetHostmapDefinition
	_jsii_.Get(
		j,
		"hostmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) IframeDefinition() PowerpackWidgetIframeDefinitionOutputReference {
	var returns PowerpackWidgetIframeDefinitionOutputReference
	_jsii_.Get(
		j,
		"iframeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) IframeDefinitionInput() *PowerpackWidgetIframeDefinition {
	var returns *PowerpackWidgetIframeDefinition
	_jsii_.Get(
		j,
		"iframeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ImageDefinition() PowerpackWidgetImageDefinitionOutputReference {
	var returns PowerpackWidgetImageDefinitionOutputReference
	_jsii_.Get(
		j,
		"imageDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ImageDefinitionInput() *PowerpackWidgetImageDefinition {
	var returns *PowerpackWidgetImageDefinition
	_jsii_.Get(
		j,
		"imageDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ListStreamDefinition() PowerpackWidgetListStreamDefinitionOutputReference {
	var returns PowerpackWidgetListStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"listStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ListStreamDefinitionInput() *PowerpackWidgetListStreamDefinition {
	var returns *PowerpackWidgetListStreamDefinition
	_jsii_.Get(
		j,
		"listStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) LogStreamDefinition() PowerpackWidgetLogStreamDefinitionOutputReference {
	var returns PowerpackWidgetLogStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"logStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) LogStreamDefinitionInput() *PowerpackWidgetLogStreamDefinition {
	var returns *PowerpackWidgetLogStreamDefinition
	_jsii_.Get(
		j,
		"logStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ManageStatusDefinition() PowerpackWidgetManageStatusDefinitionOutputReference {
	var returns PowerpackWidgetManageStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"manageStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ManageStatusDefinitionInput() *PowerpackWidgetManageStatusDefinition {
	var returns *PowerpackWidgetManageStatusDefinition
	_jsii_.Get(
		j,
		"manageStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) NoteDefinition() PowerpackWidgetNoteDefinitionOutputReference {
	var returns PowerpackWidgetNoteDefinitionOutputReference
	_jsii_.Get(
		j,
		"noteDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) NoteDefinitionInput() *PowerpackWidgetNoteDefinition {
	var returns *PowerpackWidgetNoteDefinition
	_jsii_.Get(
		j,
		"noteDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) QueryTableDefinition() PowerpackWidgetQueryTableDefinitionOutputReference {
	var returns PowerpackWidgetQueryTableDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) QueryTableDefinitionInput() *PowerpackWidgetQueryTableDefinition {
	var returns *PowerpackWidgetQueryTableDefinition
	_jsii_.Get(
		j,
		"queryTableDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) QueryValueDefinition() PowerpackWidgetQueryValueDefinitionOutputReference {
	var returns PowerpackWidgetQueryValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) QueryValueDefinitionInput() *PowerpackWidgetQueryValueDefinition {
	var returns *PowerpackWidgetQueryValueDefinition
	_jsii_.Get(
		j,
		"queryValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) RunWorkflowDefinition() PowerpackWidgetRunWorkflowDefinitionOutputReference {
	var returns PowerpackWidgetRunWorkflowDefinitionOutputReference
	_jsii_.Get(
		j,
		"runWorkflowDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) RunWorkflowDefinitionInput() *PowerpackWidgetRunWorkflowDefinition {
	var returns *PowerpackWidgetRunWorkflowDefinition
	_jsii_.Get(
		j,
		"runWorkflowDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ScatterplotDefinition() PowerpackWidgetScatterplotDefinitionOutputReference {
	var returns PowerpackWidgetScatterplotDefinitionOutputReference
	_jsii_.Get(
		j,
		"scatterplotDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ScatterplotDefinitionInput() *PowerpackWidgetScatterplotDefinition {
	var returns *PowerpackWidgetScatterplotDefinition
	_jsii_.Get(
		j,
		"scatterplotDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ServiceLevelObjectiveDefinition() PowerpackWidgetServiceLevelObjectiveDefinitionOutputReference {
	var returns PowerpackWidgetServiceLevelObjectiveDefinitionOutputReference
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ServiceLevelObjectiveDefinitionInput() *PowerpackWidgetServiceLevelObjectiveDefinition {
	var returns *PowerpackWidgetServiceLevelObjectiveDefinition
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ServicemapDefinition() PowerpackWidgetServicemapDefinitionOutputReference {
	var returns PowerpackWidgetServicemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"servicemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ServicemapDefinitionInput() *PowerpackWidgetServicemapDefinition {
	var returns *PowerpackWidgetServicemapDefinition
	_jsii_.Get(
		j,
		"servicemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) SloListDefinition() PowerpackWidgetSloListDefinitionOutputReference {
	var returns PowerpackWidgetSloListDefinitionOutputReference
	_jsii_.Get(
		j,
		"sloListDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) SloListDefinitionInput() *PowerpackWidgetSloListDefinition {
	var returns *PowerpackWidgetSloListDefinition
	_jsii_.Get(
		j,
		"sloListDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) SunburstDefinition() PowerpackWidgetSunburstDefinitionOutputReference {
	var returns PowerpackWidgetSunburstDefinitionOutputReference
	_jsii_.Get(
		j,
		"sunburstDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) SunburstDefinitionInput() *PowerpackWidgetSunburstDefinition {
	var returns *PowerpackWidgetSunburstDefinition
	_jsii_.Get(
		j,
		"sunburstDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) TimeseriesDefinition() PowerpackWidgetTimeseriesDefinitionOutputReference {
	var returns PowerpackWidgetTimeseriesDefinitionOutputReference
	_jsii_.Get(
		j,
		"timeseriesDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) TimeseriesDefinitionInput() *PowerpackWidgetTimeseriesDefinition {
	var returns *PowerpackWidgetTimeseriesDefinition
	_jsii_.Get(
		j,
		"timeseriesDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ToplistDefinition() PowerpackWidgetToplistDefinitionOutputReference {
	var returns PowerpackWidgetToplistDefinitionOutputReference
	_jsii_.Get(
		j,
		"toplistDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) ToplistDefinitionInput() *PowerpackWidgetToplistDefinition {
	var returns *PowerpackWidgetToplistDefinition
	_jsii_.Get(
		j,
		"toplistDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) TopologyMapDefinition() PowerpackWidgetTopologyMapDefinitionOutputReference {
	var returns PowerpackWidgetTopologyMapDefinitionOutputReference
	_jsii_.Get(
		j,
		"topologyMapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) TopologyMapDefinitionInput() *PowerpackWidgetTopologyMapDefinition {
	var returns *PowerpackWidgetTopologyMapDefinition
	_jsii_.Get(
		j,
		"topologyMapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) TraceServiceDefinition() PowerpackWidgetTraceServiceDefinitionOutputReference {
	var returns PowerpackWidgetTraceServiceDefinitionOutputReference
	_jsii_.Get(
		j,
		"traceServiceDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) TraceServiceDefinitionInput() *PowerpackWidgetTraceServiceDefinition {
	var returns *PowerpackWidgetTraceServiceDefinition
	_jsii_.Get(
		j,
		"traceServiceDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) TreemapDefinition() PowerpackWidgetTreemapDefinitionOutputReference {
	var returns PowerpackWidgetTreemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"treemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) TreemapDefinitionInput() *PowerpackWidgetTreemapDefinition {
	var returns *PowerpackWidgetTreemapDefinition
	_jsii_.Get(
		j,
		"treemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) WidgetLayout() PowerpackWidgetWidgetLayoutOutputReference {
	var returns PowerpackWidgetWidgetLayoutOutputReference
	_jsii_.Get(
		j,
		"widgetLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetOutputReference) WidgetLayoutInput() *PowerpackWidgetWidgetLayout {
	var returns *PowerpackWidgetWidgetLayout
	_jsii_.Get(
		j,
		"widgetLayoutInput",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackWidgetOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackWidgetOutputReference_Override(p PowerpackWidgetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutAlertGraphDefinition(value *PowerpackWidgetAlertGraphDefinition) {
	if err := p.validatePutAlertGraphDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAlertGraphDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutAlertValueDefinition(value *PowerpackWidgetAlertValueDefinition) {
	if err := p.validatePutAlertValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAlertValueDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutChangeDefinition(value *PowerpackWidgetChangeDefinition) {
	if err := p.validatePutChangeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putChangeDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutCheckStatusDefinition(value *PowerpackWidgetCheckStatusDefinition) {
	if err := p.validatePutCheckStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCheckStatusDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutDistributionDefinition(value *PowerpackWidgetDistributionDefinition) {
	if err := p.validatePutDistributionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDistributionDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutEventStreamDefinition(value *PowerpackWidgetEventStreamDefinition) {
	if err := p.validatePutEventStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventStreamDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutEventTimelineDefinition(value *PowerpackWidgetEventTimelineDefinition) {
	if err := p.validatePutEventTimelineDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventTimelineDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutFreeTextDefinition(value *PowerpackWidgetFreeTextDefinition) {
	if err := p.validatePutFreeTextDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFreeTextDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutGeomapDefinition(value *PowerpackWidgetGeomapDefinition) {
	if err := p.validatePutGeomapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGeomapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutHeatmapDefinition(value *PowerpackWidgetHeatmapDefinition) {
	if err := p.validatePutHeatmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHeatmapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutHostmapDefinition(value *PowerpackWidgetHostmapDefinition) {
	if err := p.validatePutHostmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHostmapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutIframeDefinition(value *PowerpackWidgetIframeDefinition) {
	if err := p.validatePutIframeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIframeDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutImageDefinition(value *PowerpackWidgetImageDefinition) {
	if err := p.validatePutImageDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putImageDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutListStreamDefinition(value *PowerpackWidgetListStreamDefinition) {
	if err := p.validatePutListStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putListStreamDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutLogStreamDefinition(value *PowerpackWidgetLogStreamDefinition) {
	if err := p.validatePutLogStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLogStreamDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutManageStatusDefinition(value *PowerpackWidgetManageStatusDefinition) {
	if err := p.validatePutManageStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putManageStatusDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutNoteDefinition(value *PowerpackWidgetNoteDefinition) {
	if err := p.validatePutNoteDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNoteDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutQueryTableDefinition(value *PowerpackWidgetQueryTableDefinition) {
	if err := p.validatePutQueryTableDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQueryTableDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutQueryValueDefinition(value *PowerpackWidgetQueryValueDefinition) {
	if err := p.validatePutQueryValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQueryValueDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutRunWorkflowDefinition(value *PowerpackWidgetRunWorkflowDefinition) {
	if err := p.validatePutRunWorkflowDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRunWorkflowDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutScatterplotDefinition(value *PowerpackWidgetScatterplotDefinition) {
	if err := p.validatePutScatterplotDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putScatterplotDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutServiceLevelObjectiveDefinition(value *PowerpackWidgetServiceLevelObjectiveDefinition) {
	if err := p.validatePutServiceLevelObjectiveDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putServiceLevelObjectiveDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutServicemapDefinition(value *PowerpackWidgetServicemapDefinition) {
	if err := p.validatePutServicemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putServicemapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutSloListDefinition(value *PowerpackWidgetSloListDefinition) {
	if err := p.validatePutSloListDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSloListDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutSunburstDefinition(value *PowerpackWidgetSunburstDefinition) {
	if err := p.validatePutSunburstDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSunburstDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutTimeseriesDefinition(value *PowerpackWidgetTimeseriesDefinition) {
	if err := p.validatePutTimeseriesDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeseriesDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutToplistDefinition(value *PowerpackWidgetToplistDefinition) {
	if err := p.validatePutToplistDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putToplistDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutTopologyMapDefinition(value *PowerpackWidgetTopologyMapDefinition) {
	if err := p.validatePutTopologyMapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTopologyMapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutTraceServiceDefinition(value *PowerpackWidgetTraceServiceDefinition) {
	if err := p.validatePutTraceServiceDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTraceServiceDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutTreemapDefinition(value *PowerpackWidgetTreemapDefinition) {
	if err := p.validatePutTreemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTreemapDefinition",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) PutWidgetLayout(value *PowerpackWidgetWidgetLayout) {
	if err := p.validatePutWidgetLayoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putWidgetLayout",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetAlertGraphDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetAlertGraphDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetAlertValueDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetAlertValueDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetChangeDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetChangeDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetCheckStatusDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetCheckStatusDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetDistributionDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetDistributionDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetEventStreamDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetEventStreamDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetEventTimelineDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetEventTimelineDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetFreeTextDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetFreeTextDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetGeomapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetGeomapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetHeatmapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetHeatmapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetHostmapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetHostmapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetIframeDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetIframeDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetImageDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetImageDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetListStreamDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetListStreamDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetLogStreamDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetLogStreamDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetManageStatusDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetManageStatusDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetNoteDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetNoteDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetQueryTableDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetQueryTableDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetQueryValueDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetQueryValueDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetRunWorkflowDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetRunWorkflowDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetScatterplotDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetScatterplotDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetServiceLevelObjectiveDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetServiceLevelObjectiveDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetServicemapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetServicemapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetSloListDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetSloListDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetSunburstDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetSunburstDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetTimeseriesDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeseriesDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetToplistDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetToplistDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetTopologyMapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetTopologyMapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetTraceServiceDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetTraceServiceDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetTreemapDefinition() {
	_jsii_.InvokeVoid(
		p,
		"resetTreemapDefinition",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) ResetWidgetLayout() {
	_jsii_.InvokeVoid(
		p,
		"resetWidgetLayout",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

