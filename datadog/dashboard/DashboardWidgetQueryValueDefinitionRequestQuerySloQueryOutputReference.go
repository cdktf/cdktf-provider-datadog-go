// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference interface {
	cdktf.ComplexObject
	AdditionalQueryFilters() *string
	SetAdditionalQueryFilters(val *string)
	AdditionalQueryFiltersInput() *string
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
	CrossOrgUuids() *[]*string
	SetCrossOrgUuids(val *[]*string)
	CrossOrgUuidsInput() *[]*string
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	// Experimental.
	Fqn() *string
	GroupMode() *string
	SetGroupMode(val *string)
	GroupModeInput() *string
	InternalValue() *DashboardWidgetQueryValueDefinitionRequestQuerySloQuery
	SetInternalValue(val *DashboardWidgetQueryValueDefinitionRequestQuerySloQuery)
	Measure() *string
	SetMeasure(val *string)
	MeasureInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	SloId() *string
	SetSloId(val *string)
	SloIdInput() *string
	SloQueryType() *string
	SetSloQueryType(val *string)
	SloQueryTypeInput() *string
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
	ResetAdditionalQueryFilters()
	ResetCrossOrgUuids()
	ResetGroupMode()
	ResetName()
	ResetSloQueryType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference
type jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) AdditionalQueryFilters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalQueryFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) AdditionalQueryFiltersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalQueryFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) CrossOrgUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) CrossOrgUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) GroupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) GroupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) InternalValue() *DashboardWidgetQueryValueDefinitionRequestQuerySloQuery {
	var returns *DashboardWidgetQueryValueDefinitionRequestQuerySloQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) Measure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) MeasureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) SloId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) SloIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) SloQueryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloQueryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) SloQueryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloQueryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference_Override(d DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference)SetAdditionalQueryFilters(val *string) {
	if err := j.validateSetAdditionalQueryFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalQueryFilters",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference)SetCrossOrgUuids(val *[]*string) {
	if err := j.validateSetCrossOrgUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossOrgUuids",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference)SetGroupMode(val *string) {
	if err := j.validateSetGroupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupMode",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference)SetInternalValue(val *DashboardWidgetQueryValueDefinitionRequestQuerySloQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference)SetMeasure(val *string) {
	if err := j.validateSetMeasureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measure",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference)SetSloId(val *string) {
	if err := j.validateSetSloIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sloId",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference)SetSloQueryType(val *string) {
	if err := j.validateSetSloQueryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sloQueryType",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) ResetAdditionalQueryFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetAdditionalQueryFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) ResetCrossOrgUuids() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossOrgUuids",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) ResetGroupMode() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) ResetSloQueryType() {
	_jsii_.InvokeVoid(
		d,
		"resetSloQueryType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetQueryValueDefinitionRequestQuerySloQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

