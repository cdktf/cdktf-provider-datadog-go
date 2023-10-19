// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetOutputReference interface {
	cdktf.ComplexObject
	AlertGraphDefinition() DashboardWidgetAlertGraphDefinitionOutputReference
	AlertGraphDefinitionInput() *DashboardWidgetAlertGraphDefinition
	AlertValueDefinition() DashboardWidgetAlertValueDefinitionOutputReference
	AlertValueDefinitionInput() *DashboardWidgetAlertValueDefinition
	ChangeDefinition() DashboardWidgetChangeDefinitionOutputReference
	ChangeDefinitionInput() *DashboardWidgetChangeDefinition
	CheckStatusDefinition() DashboardWidgetCheckStatusDefinitionOutputReference
	CheckStatusDefinitionInput() *DashboardWidgetCheckStatusDefinition
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
	DistributionDefinition() DashboardWidgetDistributionDefinitionOutputReference
	DistributionDefinitionInput() *DashboardWidgetDistributionDefinition
	EventStreamDefinition() DashboardWidgetEventStreamDefinitionOutputReference
	EventStreamDefinitionInput() *DashboardWidgetEventStreamDefinition
	EventTimelineDefinition() DashboardWidgetEventTimelineDefinitionOutputReference
	EventTimelineDefinitionInput() *DashboardWidgetEventTimelineDefinition
	// Experimental.
	Fqn() *string
	FreeTextDefinition() DashboardWidgetFreeTextDefinitionOutputReference
	FreeTextDefinitionInput() *DashboardWidgetFreeTextDefinition
	GeomapDefinition() DashboardWidgetGeomapDefinitionOutputReference
	GeomapDefinitionInput() *DashboardWidgetGeomapDefinition
	GroupDefinition() DashboardWidgetGroupDefinitionOutputReference
	GroupDefinitionInput() *DashboardWidgetGroupDefinition
	HeatmapDefinition() DashboardWidgetHeatmapDefinitionOutputReference
	HeatmapDefinitionInput() *DashboardWidgetHeatmapDefinition
	HostmapDefinition() DashboardWidgetHostmapDefinitionOutputReference
	HostmapDefinitionInput() *DashboardWidgetHostmapDefinition
	Id() *float64
	IframeDefinition() DashboardWidgetIframeDefinitionOutputReference
	IframeDefinitionInput() *DashboardWidgetIframeDefinition
	ImageDefinition() DashboardWidgetImageDefinitionOutputReference
	ImageDefinitionInput() *DashboardWidgetImageDefinition
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ListStreamDefinition() DashboardWidgetListStreamDefinitionOutputReference
	ListStreamDefinitionInput() *DashboardWidgetListStreamDefinition
	LogStreamDefinition() DashboardWidgetLogStreamDefinitionOutputReference
	LogStreamDefinitionInput() *DashboardWidgetLogStreamDefinition
	ManageStatusDefinition() DashboardWidgetManageStatusDefinitionOutputReference
	ManageStatusDefinitionInput() *DashboardWidgetManageStatusDefinition
	NoteDefinition() DashboardWidgetNoteDefinitionOutputReference
	NoteDefinitionInput() *DashboardWidgetNoteDefinition
	QueryTableDefinition() DashboardWidgetQueryTableDefinitionOutputReference
	QueryTableDefinitionInput() *DashboardWidgetQueryTableDefinition
	QueryValueDefinition() DashboardWidgetQueryValueDefinitionOutputReference
	QueryValueDefinitionInput() *DashboardWidgetQueryValueDefinition
	RunWorkflowDefinition() DashboardWidgetRunWorkflowDefinitionOutputReference
	RunWorkflowDefinitionInput() *DashboardWidgetRunWorkflowDefinition
	ScatterplotDefinition() DashboardWidgetScatterplotDefinitionOutputReference
	ScatterplotDefinitionInput() *DashboardWidgetScatterplotDefinition
	ServiceLevelObjectiveDefinition() DashboardWidgetServiceLevelObjectiveDefinitionOutputReference
	ServiceLevelObjectiveDefinitionInput() *DashboardWidgetServiceLevelObjectiveDefinition
	ServicemapDefinition() DashboardWidgetServicemapDefinitionOutputReference
	ServicemapDefinitionInput() *DashboardWidgetServicemapDefinition
	SloListDefinition() DashboardWidgetSloListDefinitionOutputReference
	SloListDefinitionInput() *DashboardWidgetSloListDefinition
	SplitGraphDefinition() DashboardWidgetSplitGraphDefinitionOutputReference
	SplitGraphDefinitionInput() *DashboardWidgetSplitGraphDefinition
	SunburstDefinition() DashboardWidgetSunburstDefinitionOutputReference
	SunburstDefinitionInput() *DashboardWidgetSunburstDefinition
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeseriesDefinition() DashboardWidgetTimeseriesDefinitionOutputReference
	TimeseriesDefinitionInput() *DashboardWidgetTimeseriesDefinition
	ToplistDefinition() DashboardWidgetToplistDefinitionOutputReference
	ToplistDefinitionInput() *DashboardWidgetToplistDefinition
	TopologyMapDefinition() DashboardWidgetTopologyMapDefinitionOutputReference
	TopologyMapDefinitionInput() *DashboardWidgetTopologyMapDefinition
	TraceServiceDefinition() DashboardWidgetTraceServiceDefinitionOutputReference
	TraceServiceDefinitionInput() *DashboardWidgetTraceServiceDefinition
	TreemapDefinition() DashboardWidgetTreemapDefinitionOutputReference
	TreemapDefinitionInput() *DashboardWidgetTreemapDefinition
	WidgetLayout() DashboardWidgetWidgetLayoutOutputReference
	WidgetLayoutInput() *DashboardWidgetWidgetLayout
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
	PutAlertGraphDefinition(value *DashboardWidgetAlertGraphDefinition)
	PutAlertValueDefinition(value *DashboardWidgetAlertValueDefinition)
	PutChangeDefinition(value *DashboardWidgetChangeDefinition)
	PutCheckStatusDefinition(value *DashboardWidgetCheckStatusDefinition)
	PutDistributionDefinition(value *DashboardWidgetDistributionDefinition)
	PutEventStreamDefinition(value *DashboardWidgetEventStreamDefinition)
	PutEventTimelineDefinition(value *DashboardWidgetEventTimelineDefinition)
	PutFreeTextDefinition(value *DashboardWidgetFreeTextDefinition)
	PutGeomapDefinition(value *DashboardWidgetGeomapDefinition)
	PutGroupDefinition(value *DashboardWidgetGroupDefinition)
	PutHeatmapDefinition(value *DashboardWidgetHeatmapDefinition)
	PutHostmapDefinition(value *DashboardWidgetHostmapDefinition)
	PutIframeDefinition(value *DashboardWidgetIframeDefinition)
	PutImageDefinition(value *DashboardWidgetImageDefinition)
	PutListStreamDefinition(value *DashboardWidgetListStreamDefinition)
	PutLogStreamDefinition(value *DashboardWidgetLogStreamDefinition)
	PutManageStatusDefinition(value *DashboardWidgetManageStatusDefinition)
	PutNoteDefinition(value *DashboardWidgetNoteDefinition)
	PutQueryTableDefinition(value *DashboardWidgetQueryTableDefinition)
	PutQueryValueDefinition(value *DashboardWidgetQueryValueDefinition)
	PutRunWorkflowDefinition(value *DashboardWidgetRunWorkflowDefinition)
	PutScatterplotDefinition(value *DashboardWidgetScatterplotDefinition)
	PutServiceLevelObjectiveDefinition(value *DashboardWidgetServiceLevelObjectiveDefinition)
	PutServicemapDefinition(value *DashboardWidgetServicemapDefinition)
	PutSloListDefinition(value *DashboardWidgetSloListDefinition)
	PutSplitGraphDefinition(value *DashboardWidgetSplitGraphDefinition)
	PutSunburstDefinition(value *DashboardWidgetSunburstDefinition)
	PutTimeseriesDefinition(value *DashboardWidgetTimeseriesDefinition)
	PutToplistDefinition(value *DashboardWidgetToplistDefinition)
	PutTopologyMapDefinition(value *DashboardWidgetTopologyMapDefinition)
	PutTraceServiceDefinition(value *DashboardWidgetTraceServiceDefinition)
	PutTreemapDefinition(value *DashboardWidgetTreemapDefinition)
	PutWidgetLayout(value *DashboardWidgetWidgetLayout)
	ResetAlertGraphDefinition()
	ResetAlertValueDefinition()
	ResetChangeDefinition()
	ResetCheckStatusDefinition()
	ResetDistributionDefinition()
	ResetEventStreamDefinition()
	ResetEventTimelineDefinition()
	ResetFreeTextDefinition()
	ResetGeomapDefinition()
	ResetGroupDefinition()
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
	ResetSplitGraphDefinition()
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

// The jsii proxy struct for DashboardWidgetOutputReference
type jsiiProxy_DashboardWidgetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetOutputReference) AlertGraphDefinition() DashboardWidgetAlertGraphDefinitionOutputReference {
	var returns DashboardWidgetAlertGraphDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertGraphDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) AlertGraphDefinitionInput() *DashboardWidgetAlertGraphDefinition {
	var returns *DashboardWidgetAlertGraphDefinition
	_jsii_.Get(
		j,
		"alertGraphDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) AlertValueDefinition() DashboardWidgetAlertValueDefinitionOutputReference {
	var returns DashboardWidgetAlertValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) AlertValueDefinitionInput() *DashboardWidgetAlertValueDefinition {
	var returns *DashboardWidgetAlertValueDefinition
	_jsii_.Get(
		j,
		"alertValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ChangeDefinition() DashboardWidgetChangeDefinitionOutputReference {
	var returns DashboardWidgetChangeDefinitionOutputReference
	_jsii_.Get(
		j,
		"changeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ChangeDefinitionInput() *DashboardWidgetChangeDefinition {
	var returns *DashboardWidgetChangeDefinition
	_jsii_.Get(
		j,
		"changeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) CheckStatusDefinition() DashboardWidgetCheckStatusDefinitionOutputReference {
	var returns DashboardWidgetCheckStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"checkStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) CheckStatusDefinitionInput() *DashboardWidgetCheckStatusDefinition {
	var returns *DashboardWidgetCheckStatusDefinition
	_jsii_.Get(
		j,
		"checkStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) DistributionDefinition() DashboardWidgetDistributionDefinitionOutputReference {
	var returns DashboardWidgetDistributionDefinitionOutputReference
	_jsii_.Get(
		j,
		"distributionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) DistributionDefinitionInput() *DashboardWidgetDistributionDefinition {
	var returns *DashboardWidgetDistributionDefinition
	_jsii_.Get(
		j,
		"distributionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) EventStreamDefinition() DashboardWidgetEventStreamDefinitionOutputReference {
	var returns DashboardWidgetEventStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) EventStreamDefinitionInput() *DashboardWidgetEventStreamDefinition {
	var returns *DashboardWidgetEventStreamDefinition
	_jsii_.Get(
		j,
		"eventStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) EventTimelineDefinition() DashboardWidgetEventTimelineDefinitionOutputReference {
	var returns DashboardWidgetEventTimelineDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventTimelineDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) EventTimelineDefinitionInput() *DashboardWidgetEventTimelineDefinition {
	var returns *DashboardWidgetEventTimelineDefinition
	_jsii_.Get(
		j,
		"eventTimelineDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) FreeTextDefinition() DashboardWidgetFreeTextDefinitionOutputReference {
	var returns DashboardWidgetFreeTextDefinitionOutputReference
	_jsii_.Get(
		j,
		"freeTextDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) FreeTextDefinitionInput() *DashboardWidgetFreeTextDefinition {
	var returns *DashboardWidgetFreeTextDefinition
	_jsii_.Get(
		j,
		"freeTextDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) GeomapDefinition() DashboardWidgetGeomapDefinitionOutputReference {
	var returns DashboardWidgetGeomapDefinitionOutputReference
	_jsii_.Get(
		j,
		"geomapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) GeomapDefinitionInput() *DashboardWidgetGeomapDefinition {
	var returns *DashboardWidgetGeomapDefinition
	_jsii_.Get(
		j,
		"geomapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) GroupDefinition() DashboardWidgetGroupDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionOutputReference
	_jsii_.Get(
		j,
		"groupDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) GroupDefinitionInput() *DashboardWidgetGroupDefinition {
	var returns *DashboardWidgetGroupDefinition
	_jsii_.Get(
		j,
		"groupDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) HeatmapDefinition() DashboardWidgetHeatmapDefinitionOutputReference {
	var returns DashboardWidgetHeatmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"heatmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) HeatmapDefinitionInput() *DashboardWidgetHeatmapDefinition {
	var returns *DashboardWidgetHeatmapDefinition
	_jsii_.Get(
		j,
		"heatmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) HostmapDefinition() DashboardWidgetHostmapDefinitionOutputReference {
	var returns DashboardWidgetHostmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"hostmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) HostmapDefinitionInput() *DashboardWidgetHostmapDefinition {
	var returns *DashboardWidgetHostmapDefinition
	_jsii_.Get(
		j,
		"hostmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) IframeDefinition() DashboardWidgetIframeDefinitionOutputReference {
	var returns DashboardWidgetIframeDefinitionOutputReference
	_jsii_.Get(
		j,
		"iframeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) IframeDefinitionInput() *DashboardWidgetIframeDefinition {
	var returns *DashboardWidgetIframeDefinition
	_jsii_.Get(
		j,
		"iframeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ImageDefinition() DashboardWidgetImageDefinitionOutputReference {
	var returns DashboardWidgetImageDefinitionOutputReference
	_jsii_.Get(
		j,
		"imageDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ImageDefinitionInput() *DashboardWidgetImageDefinition {
	var returns *DashboardWidgetImageDefinition
	_jsii_.Get(
		j,
		"imageDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ListStreamDefinition() DashboardWidgetListStreamDefinitionOutputReference {
	var returns DashboardWidgetListStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"listStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ListStreamDefinitionInput() *DashboardWidgetListStreamDefinition {
	var returns *DashboardWidgetListStreamDefinition
	_jsii_.Get(
		j,
		"listStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) LogStreamDefinition() DashboardWidgetLogStreamDefinitionOutputReference {
	var returns DashboardWidgetLogStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"logStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) LogStreamDefinitionInput() *DashboardWidgetLogStreamDefinition {
	var returns *DashboardWidgetLogStreamDefinition
	_jsii_.Get(
		j,
		"logStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ManageStatusDefinition() DashboardWidgetManageStatusDefinitionOutputReference {
	var returns DashboardWidgetManageStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"manageStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ManageStatusDefinitionInput() *DashboardWidgetManageStatusDefinition {
	var returns *DashboardWidgetManageStatusDefinition
	_jsii_.Get(
		j,
		"manageStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) NoteDefinition() DashboardWidgetNoteDefinitionOutputReference {
	var returns DashboardWidgetNoteDefinitionOutputReference
	_jsii_.Get(
		j,
		"noteDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) NoteDefinitionInput() *DashboardWidgetNoteDefinition {
	var returns *DashboardWidgetNoteDefinition
	_jsii_.Get(
		j,
		"noteDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) QueryTableDefinition() DashboardWidgetQueryTableDefinitionOutputReference {
	var returns DashboardWidgetQueryTableDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) QueryTableDefinitionInput() *DashboardWidgetQueryTableDefinition {
	var returns *DashboardWidgetQueryTableDefinition
	_jsii_.Get(
		j,
		"queryTableDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) QueryValueDefinition() DashboardWidgetQueryValueDefinitionOutputReference {
	var returns DashboardWidgetQueryValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) QueryValueDefinitionInput() *DashboardWidgetQueryValueDefinition {
	var returns *DashboardWidgetQueryValueDefinition
	_jsii_.Get(
		j,
		"queryValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) RunWorkflowDefinition() DashboardWidgetRunWorkflowDefinitionOutputReference {
	var returns DashboardWidgetRunWorkflowDefinitionOutputReference
	_jsii_.Get(
		j,
		"runWorkflowDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) RunWorkflowDefinitionInput() *DashboardWidgetRunWorkflowDefinition {
	var returns *DashboardWidgetRunWorkflowDefinition
	_jsii_.Get(
		j,
		"runWorkflowDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ScatterplotDefinition() DashboardWidgetScatterplotDefinitionOutputReference {
	var returns DashboardWidgetScatterplotDefinitionOutputReference
	_jsii_.Get(
		j,
		"scatterplotDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ScatterplotDefinitionInput() *DashboardWidgetScatterplotDefinition {
	var returns *DashboardWidgetScatterplotDefinition
	_jsii_.Get(
		j,
		"scatterplotDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ServiceLevelObjectiveDefinition() DashboardWidgetServiceLevelObjectiveDefinitionOutputReference {
	var returns DashboardWidgetServiceLevelObjectiveDefinitionOutputReference
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ServiceLevelObjectiveDefinitionInput() *DashboardWidgetServiceLevelObjectiveDefinition {
	var returns *DashboardWidgetServiceLevelObjectiveDefinition
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ServicemapDefinition() DashboardWidgetServicemapDefinitionOutputReference {
	var returns DashboardWidgetServicemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"servicemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ServicemapDefinitionInput() *DashboardWidgetServicemapDefinition {
	var returns *DashboardWidgetServicemapDefinition
	_jsii_.Get(
		j,
		"servicemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) SloListDefinition() DashboardWidgetSloListDefinitionOutputReference {
	var returns DashboardWidgetSloListDefinitionOutputReference
	_jsii_.Get(
		j,
		"sloListDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) SloListDefinitionInput() *DashboardWidgetSloListDefinition {
	var returns *DashboardWidgetSloListDefinition
	_jsii_.Get(
		j,
		"sloListDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) SplitGraphDefinition() DashboardWidgetSplitGraphDefinitionOutputReference {
	var returns DashboardWidgetSplitGraphDefinitionOutputReference
	_jsii_.Get(
		j,
		"splitGraphDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) SplitGraphDefinitionInput() *DashboardWidgetSplitGraphDefinition {
	var returns *DashboardWidgetSplitGraphDefinition
	_jsii_.Get(
		j,
		"splitGraphDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) SunburstDefinition() DashboardWidgetSunburstDefinitionOutputReference {
	var returns DashboardWidgetSunburstDefinitionOutputReference
	_jsii_.Get(
		j,
		"sunburstDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) SunburstDefinitionInput() *DashboardWidgetSunburstDefinition {
	var returns *DashboardWidgetSunburstDefinition
	_jsii_.Get(
		j,
		"sunburstDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) TimeseriesDefinition() DashboardWidgetTimeseriesDefinitionOutputReference {
	var returns DashboardWidgetTimeseriesDefinitionOutputReference
	_jsii_.Get(
		j,
		"timeseriesDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) TimeseriesDefinitionInput() *DashboardWidgetTimeseriesDefinition {
	var returns *DashboardWidgetTimeseriesDefinition
	_jsii_.Get(
		j,
		"timeseriesDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ToplistDefinition() DashboardWidgetToplistDefinitionOutputReference {
	var returns DashboardWidgetToplistDefinitionOutputReference
	_jsii_.Get(
		j,
		"toplistDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) ToplistDefinitionInput() *DashboardWidgetToplistDefinition {
	var returns *DashboardWidgetToplistDefinition
	_jsii_.Get(
		j,
		"toplistDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) TopologyMapDefinition() DashboardWidgetTopologyMapDefinitionOutputReference {
	var returns DashboardWidgetTopologyMapDefinitionOutputReference
	_jsii_.Get(
		j,
		"topologyMapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) TopologyMapDefinitionInput() *DashboardWidgetTopologyMapDefinition {
	var returns *DashboardWidgetTopologyMapDefinition
	_jsii_.Get(
		j,
		"topologyMapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) TraceServiceDefinition() DashboardWidgetTraceServiceDefinitionOutputReference {
	var returns DashboardWidgetTraceServiceDefinitionOutputReference
	_jsii_.Get(
		j,
		"traceServiceDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) TraceServiceDefinitionInput() *DashboardWidgetTraceServiceDefinition {
	var returns *DashboardWidgetTraceServiceDefinition
	_jsii_.Get(
		j,
		"traceServiceDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) TreemapDefinition() DashboardWidgetTreemapDefinitionOutputReference {
	var returns DashboardWidgetTreemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"treemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) TreemapDefinitionInput() *DashboardWidgetTreemapDefinition {
	var returns *DashboardWidgetTreemapDefinition
	_jsii_.Get(
		j,
		"treemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) WidgetLayout() DashboardWidgetWidgetLayoutOutputReference {
	var returns DashboardWidgetWidgetLayoutOutputReference
	_jsii_.Get(
		j,
		"widgetLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetOutputReference) WidgetLayoutInput() *DashboardWidgetWidgetLayout {
	var returns *DashboardWidgetWidgetLayout
	_jsii_.Get(
		j,
		"widgetLayoutInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardWidgetOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardWidgetOutputReference_Override(d DashboardWidgetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetOutputReference) PutAlertGraphDefinition(value *DashboardWidgetAlertGraphDefinition) {
	if err := d.validatePutAlertGraphDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAlertGraphDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutAlertValueDefinition(value *DashboardWidgetAlertValueDefinition) {
	if err := d.validatePutAlertValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAlertValueDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutChangeDefinition(value *DashboardWidgetChangeDefinition) {
	if err := d.validatePutChangeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putChangeDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutCheckStatusDefinition(value *DashboardWidgetCheckStatusDefinition) {
	if err := d.validatePutCheckStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCheckStatusDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutDistributionDefinition(value *DashboardWidgetDistributionDefinition) {
	if err := d.validatePutDistributionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDistributionDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutEventStreamDefinition(value *DashboardWidgetEventStreamDefinition) {
	if err := d.validatePutEventStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventStreamDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutEventTimelineDefinition(value *DashboardWidgetEventTimelineDefinition) {
	if err := d.validatePutEventTimelineDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventTimelineDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutFreeTextDefinition(value *DashboardWidgetFreeTextDefinition) {
	if err := d.validatePutFreeTextDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFreeTextDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutGeomapDefinition(value *DashboardWidgetGeomapDefinition) {
	if err := d.validatePutGeomapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGeomapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutGroupDefinition(value *DashboardWidgetGroupDefinition) {
	if err := d.validatePutGroupDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGroupDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutHeatmapDefinition(value *DashboardWidgetHeatmapDefinition) {
	if err := d.validatePutHeatmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHeatmapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutHostmapDefinition(value *DashboardWidgetHostmapDefinition) {
	if err := d.validatePutHostmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHostmapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutIframeDefinition(value *DashboardWidgetIframeDefinition) {
	if err := d.validatePutIframeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIframeDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutImageDefinition(value *DashboardWidgetImageDefinition) {
	if err := d.validatePutImageDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putImageDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutListStreamDefinition(value *DashboardWidgetListStreamDefinition) {
	if err := d.validatePutListStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putListStreamDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutLogStreamDefinition(value *DashboardWidgetLogStreamDefinition) {
	if err := d.validatePutLogStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLogStreamDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutManageStatusDefinition(value *DashboardWidgetManageStatusDefinition) {
	if err := d.validatePutManageStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putManageStatusDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutNoteDefinition(value *DashboardWidgetNoteDefinition) {
	if err := d.validatePutNoteDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNoteDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutQueryTableDefinition(value *DashboardWidgetQueryTableDefinition) {
	if err := d.validatePutQueryTableDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryTableDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutQueryValueDefinition(value *DashboardWidgetQueryValueDefinition) {
	if err := d.validatePutQueryValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryValueDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutRunWorkflowDefinition(value *DashboardWidgetRunWorkflowDefinition) {
	if err := d.validatePutRunWorkflowDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRunWorkflowDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutScatterplotDefinition(value *DashboardWidgetScatterplotDefinition) {
	if err := d.validatePutScatterplotDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putScatterplotDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutServiceLevelObjectiveDefinition(value *DashboardWidgetServiceLevelObjectiveDefinition) {
	if err := d.validatePutServiceLevelObjectiveDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceLevelObjectiveDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutServicemapDefinition(value *DashboardWidgetServicemapDefinition) {
	if err := d.validatePutServicemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServicemapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutSloListDefinition(value *DashboardWidgetSloListDefinition) {
	if err := d.validatePutSloListDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloListDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutSplitGraphDefinition(value *DashboardWidgetSplitGraphDefinition) {
	if err := d.validatePutSplitGraphDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSplitGraphDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutSunburstDefinition(value *DashboardWidgetSunburstDefinition) {
	if err := d.validatePutSunburstDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSunburstDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutTimeseriesDefinition(value *DashboardWidgetTimeseriesDefinition) {
	if err := d.validatePutTimeseriesDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeseriesDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutToplistDefinition(value *DashboardWidgetToplistDefinition) {
	if err := d.validatePutToplistDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putToplistDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutTopologyMapDefinition(value *DashboardWidgetTopologyMapDefinition) {
	if err := d.validatePutTopologyMapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTopologyMapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutTraceServiceDefinition(value *DashboardWidgetTraceServiceDefinition) {
	if err := d.validatePutTraceServiceDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTraceServiceDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutTreemapDefinition(value *DashboardWidgetTreemapDefinition) {
	if err := d.validatePutTreemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTreemapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) PutWidgetLayout(value *DashboardWidgetWidgetLayout) {
	if err := d.validatePutWidgetLayoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWidgetLayout",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetAlertGraphDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetAlertGraphDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetAlertValueDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetAlertValueDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetChangeDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetChangeDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetCheckStatusDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetCheckStatusDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetDistributionDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetDistributionDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetEventStreamDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetEventStreamDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetEventTimelineDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetEventTimelineDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetFreeTextDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetFreeTextDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetGeomapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetGeomapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetGroupDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetHeatmapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetHeatmapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetHostmapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetHostmapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetIframeDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetIframeDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetImageDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetImageDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetListStreamDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetListStreamDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetLogStreamDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetLogStreamDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetManageStatusDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetManageStatusDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetNoteDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetNoteDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetQueryTableDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryTableDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetQueryValueDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryValueDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetRunWorkflowDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetRunWorkflowDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetScatterplotDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetScatterplotDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetServiceLevelObjectiveDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceLevelObjectiveDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetServicemapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetServicemapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetSloListDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSloListDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetSplitGraphDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSplitGraphDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetSunburstDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSunburstDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetTimeseriesDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeseriesDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetToplistDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetToplistDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetTopologyMapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTopologyMapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetTraceServiceDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTraceServiceDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetTreemapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTreemapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) ResetWidgetLayout() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetLayout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

