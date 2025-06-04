// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference interface {
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
	ComputeQuery() DashboardWidgetQueryValueDefinitionRequestAuditQueryComputeQueryOutputReference
	ComputeQueryInput() *DashboardWidgetQueryValueDefinitionRequestAuditQueryComputeQuery
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GroupBy() DashboardWidgetQueryValueDefinitionRequestAuditQueryGroupByList
	GroupByInput() interface{}
	Index() *string
	SetIndex(val *string)
	IndexInput() *string
	InternalValue() *DashboardWidgetQueryValueDefinitionRequestAuditQuery
	SetInternalValue(val *DashboardWidgetQueryValueDefinitionRequestAuditQuery)
	MultiCompute() DashboardWidgetQueryValueDefinitionRequestAuditQueryMultiComputeList
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutComputeQuery(value *DashboardWidgetQueryValueDefinitionRequestAuditQueryComputeQuery)
	PutGroupBy(value interface{})
	PutMultiCompute(value interface{})
	ResetComputeQuery()
	ResetGroupBy()
	ResetMultiCompute()
	ResetSearchQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference
type jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) ComputeQuery() DashboardWidgetQueryValueDefinitionRequestAuditQueryComputeQueryOutputReference {
	var returns DashboardWidgetQueryValueDefinitionRequestAuditQueryComputeQueryOutputReference
	_jsii_.Get(
		j,
		"computeQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) ComputeQueryInput() *DashboardWidgetQueryValueDefinitionRequestAuditQueryComputeQuery {
	var returns *DashboardWidgetQueryValueDefinitionRequestAuditQueryComputeQuery
	_jsii_.Get(
		j,
		"computeQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) GroupBy() DashboardWidgetQueryValueDefinitionRequestAuditQueryGroupByList {
	var returns DashboardWidgetQueryValueDefinitionRequestAuditQueryGroupByList
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) GroupByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) Index() *string {
	var returns *string
	_jsii_.Get(
		j,
		"index",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) IndexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) InternalValue() *DashboardWidgetQueryValueDefinitionRequestAuditQuery {
	var returns *DashboardWidgetQueryValueDefinitionRequestAuditQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) MultiCompute() DashboardWidgetQueryValueDefinitionRequestAuditQueryMultiComputeList {
	var returns DashboardWidgetQueryValueDefinitionRequestAuditQueryMultiComputeList
	_jsii_.Get(
		j,
		"multiCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) MultiComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) SearchQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) SearchQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference_Override(d DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference)SetIndex(val *string) {
	if err := j.validateSetIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"index",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference)SetInternalValue(val *DashboardWidgetQueryValueDefinitionRequestAuditQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference)SetSearchQuery(val *string) {
	if err := j.validateSetSearchQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchQuery",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) PutComputeQuery(value *DashboardWidgetQueryValueDefinitionRequestAuditQueryComputeQuery) {
	if err := d.validatePutComputeQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putComputeQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) PutGroupBy(value interface{}) {
	if err := d.validatePutGroupByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGroupBy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) PutMultiCompute(value interface{}) {
	if err := d.validatePutMultiComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMultiCompute",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) ResetComputeQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetComputeQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) ResetMultiCompute() {
	_jsii_.InvokeVoid(
		d,
		"resetMultiCompute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) ResetSearchQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSearchQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestAuditQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

