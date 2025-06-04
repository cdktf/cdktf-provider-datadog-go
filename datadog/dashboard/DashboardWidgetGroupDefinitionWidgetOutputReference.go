// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetGroupDefinitionWidgetOutputReference interface {
	cdktf.ComplexObject
	AlertGraphDefinition() DashboardWidgetGroupDefinitionWidgetAlertGraphDefinitionOutputReference
	AlertGraphDefinitionInput() *DashboardWidgetGroupDefinitionWidgetAlertGraphDefinition
	AlertValueDefinition() DashboardWidgetGroupDefinitionWidgetAlertValueDefinitionOutputReference
	AlertValueDefinitionInput() *DashboardWidgetGroupDefinitionWidgetAlertValueDefinition
	ChangeDefinition() DashboardWidgetGroupDefinitionWidgetChangeDefinitionOutputReference
	ChangeDefinitionInput() *DashboardWidgetGroupDefinitionWidgetChangeDefinition
	CheckStatusDefinition() DashboardWidgetGroupDefinitionWidgetCheckStatusDefinitionOutputReference
	CheckStatusDefinitionInput() *DashboardWidgetGroupDefinitionWidgetCheckStatusDefinition
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
	DistributionDefinition() DashboardWidgetGroupDefinitionWidgetDistributionDefinitionOutputReference
	DistributionDefinitionInput() *DashboardWidgetGroupDefinitionWidgetDistributionDefinition
	EventStreamDefinition() DashboardWidgetGroupDefinitionWidgetEventStreamDefinitionOutputReference
	EventStreamDefinitionInput() *DashboardWidgetGroupDefinitionWidgetEventStreamDefinition
	EventTimelineDefinition() DashboardWidgetGroupDefinitionWidgetEventTimelineDefinitionOutputReference
	EventTimelineDefinitionInput() *DashboardWidgetGroupDefinitionWidgetEventTimelineDefinition
	// Experimental.
	Fqn() *string
	FreeTextDefinition() DashboardWidgetGroupDefinitionWidgetFreeTextDefinitionOutputReference
	FreeTextDefinitionInput() *DashboardWidgetGroupDefinitionWidgetFreeTextDefinition
	GeomapDefinition() DashboardWidgetGroupDefinitionWidgetGeomapDefinitionOutputReference
	GeomapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetGeomapDefinition
	HeatmapDefinition() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference
	HeatmapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinition
	HostmapDefinition() DashboardWidgetGroupDefinitionWidgetHostmapDefinitionOutputReference
	HostmapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetHostmapDefinition
	Id() *float64
	IframeDefinition() DashboardWidgetGroupDefinitionWidgetIframeDefinitionOutputReference
	IframeDefinitionInput() *DashboardWidgetGroupDefinitionWidgetIframeDefinition
	ImageDefinition() DashboardWidgetGroupDefinitionWidgetImageDefinitionOutputReference
	ImageDefinitionInput() *DashboardWidgetGroupDefinitionWidgetImageDefinition
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ListStreamDefinition() DashboardWidgetGroupDefinitionWidgetListStreamDefinitionOutputReference
	ListStreamDefinitionInput() *DashboardWidgetGroupDefinitionWidgetListStreamDefinition
	LogStreamDefinition() DashboardWidgetGroupDefinitionWidgetLogStreamDefinitionOutputReference
	LogStreamDefinitionInput() *DashboardWidgetGroupDefinitionWidgetLogStreamDefinition
	ManageStatusDefinition() DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference
	ManageStatusDefinitionInput() *DashboardWidgetGroupDefinitionWidgetManageStatusDefinition
	NoteDefinition() DashboardWidgetGroupDefinitionWidgetNoteDefinitionOutputReference
	NoteDefinitionInput() *DashboardWidgetGroupDefinitionWidgetNoteDefinition
	PowerpackDefinition() DashboardWidgetGroupDefinitionWidgetPowerpackDefinitionOutputReference
	PowerpackDefinitionInput() *DashboardWidgetGroupDefinitionWidgetPowerpackDefinition
	QueryTableDefinition() DashboardWidgetGroupDefinitionWidgetQueryTableDefinitionOutputReference
	QueryTableDefinitionInput() *DashboardWidgetGroupDefinitionWidgetQueryTableDefinition
	QueryValueDefinition() DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionOutputReference
	QueryValueDefinitionInput() *DashboardWidgetGroupDefinitionWidgetQueryValueDefinition
	RunWorkflowDefinition() DashboardWidgetGroupDefinitionWidgetRunWorkflowDefinitionOutputReference
	RunWorkflowDefinitionInput() *DashboardWidgetGroupDefinitionWidgetRunWorkflowDefinition
	ScatterplotDefinition() DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionOutputReference
	ScatterplotDefinitionInput() *DashboardWidgetGroupDefinitionWidgetScatterplotDefinition
	ServiceLevelObjectiveDefinition() DashboardWidgetGroupDefinitionWidgetServiceLevelObjectiveDefinitionOutputReference
	ServiceLevelObjectiveDefinitionInput() *DashboardWidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition
	ServicemapDefinition() DashboardWidgetGroupDefinitionWidgetServicemapDefinitionOutputReference
	ServicemapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetServicemapDefinition
	SloListDefinition() DashboardWidgetGroupDefinitionWidgetSloListDefinitionOutputReference
	SloListDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSloListDefinition
	SplitGraphDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionOutputReference
	SplitGraphDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinition
	SunburstDefinition() DashboardWidgetGroupDefinitionWidgetSunburstDefinitionOutputReference
	SunburstDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSunburstDefinition
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeseriesDefinition() DashboardWidgetGroupDefinitionWidgetTimeseriesDefinitionOutputReference
	TimeseriesDefinitionInput() *DashboardWidgetGroupDefinitionWidgetTimeseriesDefinition
	ToplistDefinition() DashboardWidgetGroupDefinitionWidgetToplistDefinitionOutputReference
	ToplistDefinitionInput() *DashboardWidgetGroupDefinitionWidgetToplistDefinition
	TopologyMapDefinition() DashboardWidgetGroupDefinitionWidgetTopologyMapDefinitionOutputReference
	TopologyMapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetTopologyMapDefinition
	TraceServiceDefinition() DashboardWidgetGroupDefinitionWidgetTraceServiceDefinitionOutputReference
	TraceServiceDefinitionInput() *DashboardWidgetGroupDefinitionWidgetTraceServiceDefinition
	TreemapDefinition() DashboardWidgetGroupDefinitionWidgetTreemapDefinitionOutputReference
	TreemapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetTreemapDefinition
	WidgetLayout() DashboardWidgetGroupDefinitionWidgetWidgetLayoutOutputReference
	WidgetLayoutInput() *DashboardWidgetGroupDefinitionWidgetWidgetLayout
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
	PutAlertGraphDefinition(value *DashboardWidgetGroupDefinitionWidgetAlertGraphDefinition)
	PutAlertValueDefinition(value *DashboardWidgetGroupDefinitionWidgetAlertValueDefinition)
	PutChangeDefinition(value *DashboardWidgetGroupDefinitionWidgetChangeDefinition)
	PutCheckStatusDefinition(value *DashboardWidgetGroupDefinitionWidgetCheckStatusDefinition)
	PutDistributionDefinition(value *DashboardWidgetGroupDefinitionWidgetDistributionDefinition)
	PutEventStreamDefinition(value *DashboardWidgetGroupDefinitionWidgetEventStreamDefinition)
	PutEventTimelineDefinition(value *DashboardWidgetGroupDefinitionWidgetEventTimelineDefinition)
	PutFreeTextDefinition(value *DashboardWidgetGroupDefinitionWidgetFreeTextDefinition)
	PutGeomapDefinition(value *DashboardWidgetGroupDefinitionWidgetGeomapDefinition)
	PutHeatmapDefinition(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinition)
	PutHostmapDefinition(value *DashboardWidgetGroupDefinitionWidgetHostmapDefinition)
	PutIframeDefinition(value *DashboardWidgetGroupDefinitionWidgetIframeDefinition)
	PutImageDefinition(value *DashboardWidgetGroupDefinitionWidgetImageDefinition)
	PutListStreamDefinition(value *DashboardWidgetGroupDefinitionWidgetListStreamDefinition)
	PutLogStreamDefinition(value *DashboardWidgetGroupDefinitionWidgetLogStreamDefinition)
	PutManageStatusDefinition(value *DashboardWidgetGroupDefinitionWidgetManageStatusDefinition)
	PutNoteDefinition(value *DashboardWidgetGroupDefinitionWidgetNoteDefinition)
	PutPowerpackDefinition(value *DashboardWidgetGroupDefinitionWidgetPowerpackDefinition)
	PutQueryTableDefinition(value *DashboardWidgetGroupDefinitionWidgetQueryTableDefinition)
	PutQueryValueDefinition(value *DashboardWidgetGroupDefinitionWidgetQueryValueDefinition)
	PutRunWorkflowDefinition(value *DashboardWidgetGroupDefinitionWidgetRunWorkflowDefinition)
	PutScatterplotDefinition(value *DashboardWidgetGroupDefinitionWidgetScatterplotDefinition)
	PutServiceLevelObjectiveDefinition(value *DashboardWidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition)
	PutServicemapDefinition(value *DashboardWidgetGroupDefinitionWidgetServicemapDefinition)
	PutSloListDefinition(value *DashboardWidgetGroupDefinitionWidgetSloListDefinition)
	PutSplitGraphDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinition)
	PutSunburstDefinition(value *DashboardWidgetGroupDefinitionWidgetSunburstDefinition)
	PutTimeseriesDefinition(value *DashboardWidgetGroupDefinitionWidgetTimeseriesDefinition)
	PutToplistDefinition(value *DashboardWidgetGroupDefinitionWidgetToplistDefinition)
	PutTopologyMapDefinition(value *DashboardWidgetGroupDefinitionWidgetTopologyMapDefinition)
	PutTraceServiceDefinition(value *DashboardWidgetGroupDefinitionWidgetTraceServiceDefinition)
	PutTreemapDefinition(value *DashboardWidgetGroupDefinitionWidgetTreemapDefinition)
	PutWidgetLayout(value *DashboardWidgetGroupDefinitionWidgetWidgetLayout)
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
	ResetPowerpackDefinition()
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

// The jsii proxy struct for DashboardWidgetGroupDefinitionWidgetOutputReference
type jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) AlertGraphDefinition() DashboardWidgetGroupDefinitionWidgetAlertGraphDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetAlertGraphDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertGraphDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) AlertGraphDefinitionInput() *DashboardWidgetGroupDefinitionWidgetAlertGraphDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetAlertGraphDefinition
	_jsii_.Get(
		j,
		"alertGraphDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) AlertValueDefinition() DashboardWidgetGroupDefinitionWidgetAlertValueDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetAlertValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"alertValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) AlertValueDefinitionInput() *DashboardWidgetGroupDefinitionWidgetAlertValueDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetAlertValueDefinition
	_jsii_.Get(
		j,
		"alertValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ChangeDefinition() DashboardWidgetGroupDefinitionWidgetChangeDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetChangeDefinitionOutputReference
	_jsii_.Get(
		j,
		"changeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ChangeDefinitionInput() *DashboardWidgetGroupDefinitionWidgetChangeDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetChangeDefinition
	_jsii_.Get(
		j,
		"changeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) CheckStatusDefinition() DashboardWidgetGroupDefinitionWidgetCheckStatusDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetCheckStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"checkStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) CheckStatusDefinitionInput() *DashboardWidgetGroupDefinitionWidgetCheckStatusDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetCheckStatusDefinition
	_jsii_.Get(
		j,
		"checkStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) DistributionDefinition() DashboardWidgetGroupDefinitionWidgetDistributionDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetDistributionDefinitionOutputReference
	_jsii_.Get(
		j,
		"distributionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) DistributionDefinitionInput() *DashboardWidgetGroupDefinitionWidgetDistributionDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetDistributionDefinition
	_jsii_.Get(
		j,
		"distributionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) EventStreamDefinition() DashboardWidgetGroupDefinitionWidgetEventStreamDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetEventStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) EventStreamDefinitionInput() *DashboardWidgetGroupDefinitionWidgetEventStreamDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetEventStreamDefinition
	_jsii_.Get(
		j,
		"eventStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) EventTimelineDefinition() DashboardWidgetGroupDefinitionWidgetEventTimelineDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetEventTimelineDefinitionOutputReference
	_jsii_.Get(
		j,
		"eventTimelineDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) EventTimelineDefinitionInput() *DashboardWidgetGroupDefinitionWidgetEventTimelineDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetEventTimelineDefinition
	_jsii_.Get(
		j,
		"eventTimelineDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) FreeTextDefinition() DashboardWidgetGroupDefinitionWidgetFreeTextDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetFreeTextDefinitionOutputReference
	_jsii_.Get(
		j,
		"freeTextDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) FreeTextDefinitionInput() *DashboardWidgetGroupDefinitionWidgetFreeTextDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetFreeTextDefinition
	_jsii_.Get(
		j,
		"freeTextDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) GeomapDefinition() DashboardWidgetGroupDefinitionWidgetGeomapDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetGeomapDefinitionOutputReference
	_jsii_.Get(
		j,
		"geomapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) GeomapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetGeomapDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetGeomapDefinition
	_jsii_.Get(
		j,
		"geomapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) HeatmapDefinition() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"heatmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) HeatmapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetHeatmapDefinition
	_jsii_.Get(
		j,
		"heatmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) HostmapDefinition() DashboardWidgetGroupDefinitionWidgetHostmapDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetHostmapDefinitionOutputReference
	_jsii_.Get(
		j,
		"hostmapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) HostmapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetHostmapDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetHostmapDefinition
	_jsii_.Get(
		j,
		"hostmapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) IframeDefinition() DashboardWidgetGroupDefinitionWidgetIframeDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetIframeDefinitionOutputReference
	_jsii_.Get(
		j,
		"iframeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) IframeDefinitionInput() *DashboardWidgetGroupDefinitionWidgetIframeDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetIframeDefinition
	_jsii_.Get(
		j,
		"iframeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ImageDefinition() DashboardWidgetGroupDefinitionWidgetImageDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetImageDefinitionOutputReference
	_jsii_.Get(
		j,
		"imageDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ImageDefinitionInput() *DashboardWidgetGroupDefinitionWidgetImageDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetImageDefinition
	_jsii_.Get(
		j,
		"imageDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ListStreamDefinition() DashboardWidgetGroupDefinitionWidgetListStreamDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetListStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"listStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ListStreamDefinitionInput() *DashboardWidgetGroupDefinitionWidgetListStreamDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetListStreamDefinition
	_jsii_.Get(
		j,
		"listStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) LogStreamDefinition() DashboardWidgetGroupDefinitionWidgetLogStreamDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetLogStreamDefinitionOutputReference
	_jsii_.Get(
		j,
		"logStreamDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) LogStreamDefinitionInput() *DashboardWidgetGroupDefinitionWidgetLogStreamDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetLogStreamDefinition
	_jsii_.Get(
		j,
		"logStreamDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ManageStatusDefinition() DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetManageStatusDefinitionOutputReference
	_jsii_.Get(
		j,
		"manageStatusDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ManageStatusDefinitionInput() *DashboardWidgetGroupDefinitionWidgetManageStatusDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetManageStatusDefinition
	_jsii_.Get(
		j,
		"manageStatusDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) NoteDefinition() DashboardWidgetGroupDefinitionWidgetNoteDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetNoteDefinitionOutputReference
	_jsii_.Get(
		j,
		"noteDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) NoteDefinitionInput() *DashboardWidgetGroupDefinitionWidgetNoteDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetNoteDefinition
	_jsii_.Get(
		j,
		"noteDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PowerpackDefinition() DashboardWidgetGroupDefinitionWidgetPowerpackDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetPowerpackDefinitionOutputReference
	_jsii_.Get(
		j,
		"powerpackDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PowerpackDefinitionInput() *DashboardWidgetGroupDefinitionWidgetPowerpackDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetPowerpackDefinition
	_jsii_.Get(
		j,
		"powerpackDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) QueryTableDefinition() DashboardWidgetGroupDefinitionWidgetQueryTableDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetQueryTableDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) QueryTableDefinitionInput() *DashboardWidgetGroupDefinitionWidgetQueryTableDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetQueryTableDefinition
	_jsii_.Get(
		j,
		"queryTableDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) QueryValueDefinition() DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) QueryValueDefinitionInput() *DashboardWidgetGroupDefinitionWidgetQueryValueDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetQueryValueDefinition
	_jsii_.Get(
		j,
		"queryValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) RunWorkflowDefinition() DashboardWidgetGroupDefinitionWidgetRunWorkflowDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetRunWorkflowDefinitionOutputReference
	_jsii_.Get(
		j,
		"runWorkflowDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) RunWorkflowDefinitionInput() *DashboardWidgetGroupDefinitionWidgetRunWorkflowDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetRunWorkflowDefinition
	_jsii_.Get(
		j,
		"runWorkflowDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ScatterplotDefinition() DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionOutputReference
	_jsii_.Get(
		j,
		"scatterplotDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ScatterplotDefinitionInput() *DashboardWidgetGroupDefinitionWidgetScatterplotDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetScatterplotDefinition
	_jsii_.Get(
		j,
		"scatterplotDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ServiceLevelObjectiveDefinition() DashboardWidgetGroupDefinitionWidgetServiceLevelObjectiveDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetServiceLevelObjectiveDefinitionOutputReference
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ServiceLevelObjectiveDefinitionInput() *DashboardWidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition
	_jsii_.Get(
		j,
		"serviceLevelObjectiveDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ServicemapDefinition() DashboardWidgetGroupDefinitionWidgetServicemapDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetServicemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"servicemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ServicemapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetServicemapDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetServicemapDefinition
	_jsii_.Get(
		j,
		"servicemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) SloListDefinition() DashboardWidgetGroupDefinitionWidgetSloListDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSloListDefinitionOutputReference
	_jsii_.Get(
		j,
		"sloListDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) SloListDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSloListDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetSloListDefinition
	_jsii_.Get(
		j,
		"sloListDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) SplitGraphDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionOutputReference
	_jsii_.Get(
		j,
		"splitGraphDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) SplitGraphDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinition
	_jsii_.Get(
		j,
		"splitGraphDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) SunburstDefinition() DashboardWidgetGroupDefinitionWidgetSunburstDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSunburstDefinitionOutputReference
	_jsii_.Get(
		j,
		"sunburstDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) SunburstDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSunburstDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetSunburstDefinition
	_jsii_.Get(
		j,
		"sunburstDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) TimeseriesDefinition() DashboardWidgetGroupDefinitionWidgetTimeseriesDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetTimeseriesDefinitionOutputReference
	_jsii_.Get(
		j,
		"timeseriesDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) TimeseriesDefinitionInput() *DashboardWidgetGroupDefinitionWidgetTimeseriesDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetTimeseriesDefinition
	_jsii_.Get(
		j,
		"timeseriesDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ToplistDefinition() DashboardWidgetGroupDefinitionWidgetToplistDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetToplistDefinitionOutputReference
	_jsii_.Get(
		j,
		"toplistDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ToplistDefinitionInput() *DashboardWidgetGroupDefinitionWidgetToplistDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetToplistDefinition
	_jsii_.Get(
		j,
		"toplistDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) TopologyMapDefinition() DashboardWidgetGroupDefinitionWidgetTopologyMapDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetTopologyMapDefinitionOutputReference
	_jsii_.Get(
		j,
		"topologyMapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) TopologyMapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetTopologyMapDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetTopologyMapDefinition
	_jsii_.Get(
		j,
		"topologyMapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) TraceServiceDefinition() DashboardWidgetGroupDefinitionWidgetTraceServiceDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetTraceServiceDefinitionOutputReference
	_jsii_.Get(
		j,
		"traceServiceDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) TraceServiceDefinitionInput() *DashboardWidgetGroupDefinitionWidgetTraceServiceDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetTraceServiceDefinition
	_jsii_.Get(
		j,
		"traceServiceDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) TreemapDefinition() DashboardWidgetGroupDefinitionWidgetTreemapDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetTreemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"treemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) TreemapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetTreemapDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetTreemapDefinition
	_jsii_.Get(
		j,
		"treemapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) WidgetLayout() DashboardWidgetGroupDefinitionWidgetWidgetLayoutOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetWidgetLayoutOutputReference
	_jsii_.Get(
		j,
		"widgetLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) WidgetLayoutInput() *DashboardWidgetGroupDefinitionWidgetWidgetLayout {
	var returns *DashboardWidgetGroupDefinitionWidgetWidgetLayout
	_jsii_.Get(
		j,
		"widgetLayoutInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetGroupDefinitionWidgetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardWidgetGroupDefinitionWidgetOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetGroupDefinitionWidgetOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardWidgetGroupDefinitionWidgetOutputReference_Override(d DashboardWidgetGroupDefinitionWidgetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutAlertGraphDefinition(value *DashboardWidgetGroupDefinitionWidgetAlertGraphDefinition) {
	if err := d.validatePutAlertGraphDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAlertGraphDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutAlertValueDefinition(value *DashboardWidgetGroupDefinitionWidgetAlertValueDefinition) {
	if err := d.validatePutAlertValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAlertValueDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutChangeDefinition(value *DashboardWidgetGroupDefinitionWidgetChangeDefinition) {
	if err := d.validatePutChangeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putChangeDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutCheckStatusDefinition(value *DashboardWidgetGroupDefinitionWidgetCheckStatusDefinition) {
	if err := d.validatePutCheckStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCheckStatusDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutDistributionDefinition(value *DashboardWidgetGroupDefinitionWidgetDistributionDefinition) {
	if err := d.validatePutDistributionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDistributionDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutEventStreamDefinition(value *DashboardWidgetGroupDefinitionWidgetEventStreamDefinition) {
	if err := d.validatePutEventStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventStreamDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutEventTimelineDefinition(value *DashboardWidgetGroupDefinitionWidgetEventTimelineDefinition) {
	if err := d.validatePutEventTimelineDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventTimelineDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutFreeTextDefinition(value *DashboardWidgetGroupDefinitionWidgetFreeTextDefinition) {
	if err := d.validatePutFreeTextDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFreeTextDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutGeomapDefinition(value *DashboardWidgetGroupDefinitionWidgetGeomapDefinition) {
	if err := d.validatePutGeomapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGeomapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutHeatmapDefinition(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinition) {
	if err := d.validatePutHeatmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHeatmapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutHostmapDefinition(value *DashboardWidgetGroupDefinitionWidgetHostmapDefinition) {
	if err := d.validatePutHostmapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHostmapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutIframeDefinition(value *DashboardWidgetGroupDefinitionWidgetIframeDefinition) {
	if err := d.validatePutIframeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIframeDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutImageDefinition(value *DashboardWidgetGroupDefinitionWidgetImageDefinition) {
	if err := d.validatePutImageDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putImageDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutListStreamDefinition(value *DashboardWidgetGroupDefinitionWidgetListStreamDefinition) {
	if err := d.validatePutListStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putListStreamDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutLogStreamDefinition(value *DashboardWidgetGroupDefinitionWidgetLogStreamDefinition) {
	if err := d.validatePutLogStreamDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLogStreamDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutManageStatusDefinition(value *DashboardWidgetGroupDefinitionWidgetManageStatusDefinition) {
	if err := d.validatePutManageStatusDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putManageStatusDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutNoteDefinition(value *DashboardWidgetGroupDefinitionWidgetNoteDefinition) {
	if err := d.validatePutNoteDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNoteDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutPowerpackDefinition(value *DashboardWidgetGroupDefinitionWidgetPowerpackDefinition) {
	if err := d.validatePutPowerpackDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPowerpackDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutQueryTableDefinition(value *DashboardWidgetGroupDefinitionWidgetQueryTableDefinition) {
	if err := d.validatePutQueryTableDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryTableDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutQueryValueDefinition(value *DashboardWidgetGroupDefinitionWidgetQueryValueDefinition) {
	if err := d.validatePutQueryValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryValueDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutRunWorkflowDefinition(value *DashboardWidgetGroupDefinitionWidgetRunWorkflowDefinition) {
	if err := d.validatePutRunWorkflowDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRunWorkflowDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutScatterplotDefinition(value *DashboardWidgetGroupDefinitionWidgetScatterplotDefinition) {
	if err := d.validatePutScatterplotDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putScatterplotDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutServiceLevelObjectiveDefinition(value *DashboardWidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition) {
	if err := d.validatePutServiceLevelObjectiveDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServiceLevelObjectiveDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutServicemapDefinition(value *DashboardWidgetGroupDefinitionWidgetServicemapDefinition) {
	if err := d.validatePutServicemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServicemapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutSloListDefinition(value *DashboardWidgetGroupDefinitionWidgetSloListDefinition) {
	if err := d.validatePutSloListDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSloListDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutSplitGraphDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinition) {
	if err := d.validatePutSplitGraphDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSplitGraphDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutSunburstDefinition(value *DashboardWidgetGroupDefinitionWidgetSunburstDefinition) {
	if err := d.validatePutSunburstDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSunburstDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutTimeseriesDefinition(value *DashboardWidgetGroupDefinitionWidgetTimeseriesDefinition) {
	if err := d.validatePutTimeseriesDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeseriesDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutToplistDefinition(value *DashboardWidgetGroupDefinitionWidgetToplistDefinition) {
	if err := d.validatePutToplistDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putToplistDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutTopologyMapDefinition(value *DashboardWidgetGroupDefinitionWidgetTopologyMapDefinition) {
	if err := d.validatePutTopologyMapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTopologyMapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutTraceServiceDefinition(value *DashboardWidgetGroupDefinitionWidgetTraceServiceDefinition) {
	if err := d.validatePutTraceServiceDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTraceServiceDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutTreemapDefinition(value *DashboardWidgetGroupDefinitionWidgetTreemapDefinition) {
	if err := d.validatePutTreemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTreemapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) PutWidgetLayout(value *DashboardWidgetGroupDefinitionWidgetWidgetLayout) {
	if err := d.validatePutWidgetLayoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWidgetLayout",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetAlertGraphDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetAlertGraphDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetAlertValueDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetAlertValueDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetChangeDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetChangeDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetCheckStatusDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetCheckStatusDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetDistributionDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetDistributionDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetEventStreamDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetEventStreamDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetEventTimelineDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetEventTimelineDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetFreeTextDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetFreeTextDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetGeomapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetGeomapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetHeatmapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetHeatmapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetHostmapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetHostmapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetIframeDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetIframeDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetImageDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetImageDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetListStreamDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetListStreamDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetLogStreamDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetLogStreamDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetManageStatusDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetManageStatusDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetNoteDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetNoteDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetPowerpackDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetPowerpackDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetQueryTableDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryTableDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetQueryValueDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryValueDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetRunWorkflowDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetRunWorkflowDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetScatterplotDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetScatterplotDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetServiceLevelObjectiveDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceLevelObjectiveDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetServicemapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetServicemapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetSloListDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSloListDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetSplitGraphDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSplitGraphDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetSunburstDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSunburstDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetTimeseriesDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeseriesDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetToplistDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetToplistDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetTopologyMapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTopologyMapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetTraceServiceDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTraceServiceDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetTreemapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTreemapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ResetWidgetLayout() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetLayout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

