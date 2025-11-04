// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference interface {
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
	ComputeQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryComputeQueryOutputReference
	ComputeQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryComputeQuery
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GroupBy() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryGroupByList
	GroupByInput() interface{}
	Index() *string
	SetIndex(val *string)
	IndexInput() *string
	InternalValue() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQuery
	SetInternalValue(val *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQuery)
	MultiCompute() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryMultiComputeList
	MultiComputeInput() interface{}
	SearchQuery() *string
	SetSearchQuery(val *string)
	SearchQueryInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutComputeQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryComputeQuery)
	PutGroupBy(value interface{})
	PutMultiCompute(value interface{})
	ResetComputeQuery()
	ResetGroupBy()
	ResetMultiCompute()
	ResetSearchQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference
type jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) ComputeQuery() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryComputeQueryOutputReference {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryComputeQueryOutputReference
	_jsii_.Get(
		j,
		"computeQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) ComputeQueryInput() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryComputeQuery {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryComputeQuery
	_jsii_.Get(
		j,
		"computeQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) GroupBy() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryGroupByList {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryGroupByList
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) GroupByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) IndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) InternalValue() *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQuery {
	var returns *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) MultiCompute() DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryMultiComputeList {
	var returns DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryMultiComputeList
	_jsii_.Get(
		j,
		"multiCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) MultiComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) SearchQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) SearchQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference_Override(d DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference)SetIndex(val *string) {
	if err := j.validateSetIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference)SetInternalValue(val *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference)SetSearchQuery(val *string) {
	if err := j.validateSetSearchQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchQuery",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) PutComputeQuery(value *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryComputeQuery) {
	if err := d.validatePutComputeQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putComputeQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) PutGroupBy(value interface{}) {
	if err := d.validatePutGroupByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGroupBy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) PutMultiCompute(value interface{}) {
	if err := d.validatePutMultiComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMultiCompute",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) ResetComputeQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetComputeQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) ResetMultiCompute() {
	_jsii_.InvokeVoid(
		d,
		"resetMultiCompute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) ResetSearchQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSearchQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionRequestLogQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

