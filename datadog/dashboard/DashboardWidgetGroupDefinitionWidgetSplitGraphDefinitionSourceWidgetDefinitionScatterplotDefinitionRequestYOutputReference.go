// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference interface {
	cdktf.ComplexObject
	Aggregator() *string
	SetAggregator(val *string)
	AggregatorInput() *string
	ApmQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYApmQueryOutputReference
	ApmQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYApmQuery
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYLogQueryOutputReference
	LogQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYLogQuery
	ProcessQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYProcessQueryOutputReference
	ProcessQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	RumQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYRumQueryOutputReference
	RumQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYRumQuery
	SecurityQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYSecurityQueryOutputReference
	SecurityQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYSecurityQuery
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
	PutApmQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYApmQuery)
	PutLogQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYLogQuery)
	PutProcessQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYProcessQuery)
	PutRumQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYRumQuery)
	PutSecurityQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYSecurityQuery)
	ResetAggregator()
	ResetApmQuery()
	ResetLogQuery()
	ResetProcessQuery()
	ResetQ()
	ResetRumQuery()
	ResetSecurityQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference
type jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) Aggregator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) AggregatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ApmQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYApmQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ApmQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYApmQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) LogQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYLogQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) LogQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYLogQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ProcessQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYProcessQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ProcessQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYProcessQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) RumQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYRumQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) RumQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYRumQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) SecurityQuery() DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYSecurityQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) SecurityQueryInput() *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYSecurityQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference_Override(d DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference)SetAggregator(val *string) {
	if err := j.validateSetAggregatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregator",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) PutApmQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYApmQuery) {
	if err := d.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) PutLogQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYLogQuery) {
	if err := d.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) PutProcessQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) PutRumQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYRumQuery) {
	if err := d.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) PutSecurityQuery(value *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYSecurityQuery) {
	if err := d.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ResetAggregator() {
	_jsii_.InvokeVoid(
		d,
		"resetAggregator",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		d,
		"resetQ",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

