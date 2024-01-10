// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference interface {
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
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	// Experimental.
	Fqn() *string
	GroupMode() *string
	SetGroupMode(val *string)
	GroupModeInput() *string
	InternalValue() *PowerpackWidgetGeomapDefinitionRequestQuerySloQuery
	SetInternalValue(val *PowerpackWidgetGeomapDefinitionRequestQuerySloQuery)
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

// The jsii proxy struct for PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference
type jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) AdditionalQueryFilters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalQueryFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) AdditionalQueryFiltersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalQueryFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GroupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GroupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) InternalValue() *PowerpackWidgetGeomapDefinitionRequestQuerySloQuery {
	var returns *PowerpackWidgetGeomapDefinitionRequestQuerySloQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) Measure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) MeasureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) SloId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) SloIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) SloQueryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloQueryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) SloQueryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloQueryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference_Override(p PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetAdditionalQueryFilters(val *string) {
	if err := j.validateSetAdditionalQueryFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalQueryFilters",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetGroupMode(val *string) {
	if err := j.validateSetGroupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupMode",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetInternalValue(val *PowerpackWidgetGeomapDefinitionRequestQuerySloQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetMeasure(val *string) {
	if err := j.validateSetMeasureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measure",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetSloId(val *string) {
	if err := j.validateSetSloIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sloId",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetSloQueryType(val *string) {
	if err := j.validateSetSloQueryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sloQueryType",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ResetAdditionalQueryFilters() {
	_jsii_.InvokeVoid(
		p,
		"resetAdditionalQueryFilters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ResetGroupMode() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ResetSloQueryType() {
	_jsii_.InvokeVoid(
		p,
		"resetSloQueryType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

