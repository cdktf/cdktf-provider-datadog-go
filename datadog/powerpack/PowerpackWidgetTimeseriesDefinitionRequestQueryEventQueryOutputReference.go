// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference interface {
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
	Compute() PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryComputeList
	ComputeInput() interface{}
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
	GroupBy() PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryGroupByList
	GroupByInput() interface{}
	Indexes() *[]*string
	SetIndexes(val *[]*string)
	IndexesInput() *[]*string
	InternalValue() *PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuery
	SetInternalValue(val *PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuery)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Search() PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuerySearchOutputReference
	SearchInput() *PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuerySearch
	Storage() *string
	SetStorage(val *string)
	StorageInput() *string
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
	PutCompute(value interface{})
	PutGroupBy(value interface{})
	PutSearch(value *PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuerySearch)
	ResetCrossOrgUuids()
	ResetGroupBy()
	ResetIndexes()
	ResetSearch()
	ResetStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference
type jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) Compute() PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryComputeList {
	var returns PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryComputeList
	_jsii_.Get(
		j,
		"compute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) ComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) CrossOrgUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) CrossOrgUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) GroupBy() PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryGroupByList {
	var returns PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryGroupByList
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) GroupByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) Indexes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) IndexesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) InternalValue() *PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuery {
	var returns *PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) Search() PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuerySearchOutputReference {
	var returns PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuerySearchOutputReference
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) SearchInput() *PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuerySearch {
	var returns *PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuerySearch
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) Storage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) StorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference_Override(p PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference)SetCrossOrgUuids(val *[]*string) {
	if err := j.validateSetCrossOrgUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossOrgUuids",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference)SetIndexes(val *[]*string) {
	if err := j.validateSetIndexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexes",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference)SetInternalValue(val *PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference)SetStorage(val *string) {
	if err := j.validateSetStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storage",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) PutCompute(value interface{}) {
	if err := p.validatePutComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCompute",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) PutGroupBy(value interface{}) {
	if err := p.validatePutGroupByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGroupBy",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) PutSearch(value *PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuerySearch) {
	if err := p.validatePutSearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSearch",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) ResetCrossOrgUuids() {
	_jsii_.InvokeVoid(
		p,
		"resetCrossOrgUuids",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) ResetIndexes() {
	_jsii_.InvokeVoid(
		p,
		"resetIndexes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) ResetSearch() {
	_jsii_.InvokeVoid(
		p,
		"resetSearch",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) ResetStorage() {
	_jsii_.InvokeVoid(
		p,
		"resetStorage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestQueryEventQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

