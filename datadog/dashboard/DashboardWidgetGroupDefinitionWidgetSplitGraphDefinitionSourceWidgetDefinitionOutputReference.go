// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference interface {
	cdktf.ComplexObject
	ChangeDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionOutputReference
	ChangeDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinition
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
	// Experimental.
	Fqn() *string
	GeomapDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionOutputReference
	GeomapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinition
	InternalValue() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinition
	SetInternalValue(val *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinition)
	QueryTableDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryTableDefinitionOutputReference
	QueryTableDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryTableDefinition
	QueryValueDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionOutputReference
	QueryValueDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinition
	ScatterplotDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionOutputReference
	ScatterplotDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinition
	SunburstDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionOutputReference
	SunburstDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinition
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeseriesDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference
	TimeseriesDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition
	ToplistDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionOutputReference
	ToplistDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinition
	TreemapDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionOutputReference
	TreemapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinition
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
	PutChangeDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinition)
	PutGeomapDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinition)
	PutQueryTableDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryTableDefinition)
	PutQueryValueDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinition)
	PutScatterplotDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinition)
	PutSunburstDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinition)
	PutTimeseriesDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition)
	PutToplistDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinition)
	PutTreemapDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinition)
	ResetChangeDefinition()
	ResetGeomapDefinition()
	ResetQueryTableDefinition()
	ResetQueryValueDefinition()
	ResetScatterplotDefinition()
	ResetSunburstDefinition()
	ResetTimeseriesDefinition()
	ResetToplistDefinition()
	ResetTreemapDefinition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference
type jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ChangeDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionOutputReference
	_jsii_.Get(
		j,
		"changeDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ChangeDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinition
	_jsii_.Get(
		j,
		"changeDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) GeomapDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionOutputReference
	_jsii_.Get(
		j,
		"geomapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) GeomapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinition
	_jsii_.Get(
		j,
		"geomapDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) InternalValue() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) QueryTableDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryTableDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryTableDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) QueryTableDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryTableDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryTableDefinition
	_jsii_.Get(
		j,
		"queryTableDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) QueryValueDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionOutputReference
	_jsii_.Get(
		j,
		"queryValueDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) QueryValueDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinition
	_jsii_.Get(
		j,
		"queryValueDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ScatterplotDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionOutputReference
	_jsii_.Get(
		j,
		"scatterplotDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ScatterplotDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinition
	_jsii_.Get(
		j,
		"scatterplotDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) SunburstDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionOutputReference
	_jsii_.Get(
		j,
		"sunburstDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) SunburstDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinition
	_jsii_.Get(
		j,
		"sunburstDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) TimeseriesDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionOutputReference
	_jsii_.Get(
		j,
		"timeseriesDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) TimeseriesDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition
	_jsii_.Get(
		j,
		"timeseriesDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ToplistDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionOutputReference
	_jsii_.Get(
		j,
		"toplistDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ToplistDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinition
	_jsii_.Get(
		j,
		"toplistDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) TreemapDefinition() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionOutputReference
	_jsii_.Get(
		j,
		"treemapDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) TreemapDefinitionInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinition {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinition
	_jsii_.Get(
		j,
		"treemapDefinitionInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference_Override(d DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference)SetInternalValue(val *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) PutChangeDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinition) {
	if err := d.validatePutChangeDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putChangeDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) PutGeomapDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinition) {
	if err := d.validatePutGeomapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGeomapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) PutQueryTableDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryTableDefinition) {
	if err := d.validatePutQueryTableDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryTableDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) PutQueryValueDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinition) {
	if err := d.validatePutQueryValueDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryValueDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) PutScatterplotDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinition) {
	if err := d.validatePutScatterplotDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putScatterplotDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) PutSunburstDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinition) {
	if err := d.validatePutSunburstDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSunburstDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) PutTimeseriesDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition) {
	if err := d.validatePutTimeseriesDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeseriesDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) PutToplistDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinition) {
	if err := d.validatePutToplistDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putToplistDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) PutTreemapDefinition(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinition) {
	if err := d.validatePutTreemapDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTreemapDefinition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ResetChangeDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetChangeDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ResetGeomapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetGeomapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ResetQueryTableDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryTableDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ResetQueryValueDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryValueDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ResetScatterplotDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetScatterplotDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ResetSunburstDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetSunburstDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ResetTimeseriesDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeseriesDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ResetToplistDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetToplistDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ResetTreemapDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetTreemapDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

