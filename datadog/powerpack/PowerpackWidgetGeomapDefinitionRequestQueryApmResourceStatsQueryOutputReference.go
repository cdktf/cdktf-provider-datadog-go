// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference interface {
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
	Env() *string
	SetEnv(val *string)
	EnvInput() *string
	// Experimental.
	Fqn() *string
	GroupBy() *[]*string
	SetGroupBy(val *[]*string)
	GroupByInput() *[]*string
	InternalValue() *PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQuery
	SetInternalValue(val *PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQuery)
	Name() *string
	SetName(val *string)
	NameInput() *string
	OperationName() *string
	SetOperationName(val *string)
	OperationNameInput() *string
	PrimaryTagName() *string
	SetPrimaryTagName(val *string)
	PrimaryTagNameInput() *string
	PrimaryTagValue() *string
	SetPrimaryTagValue(val *string)
	PrimaryTagValueInput() *string
	ResourceName() *string
	SetResourceName(val *string)
	ResourceNameInput() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	Stat() *string
	SetStat(val *string)
	StatInput() *string
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
	ResetCrossOrgUuids()
	ResetGroupBy()
	ResetOperationName()
	ResetPrimaryTagName()
	ResetPrimaryTagValue()
	ResetResourceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference
type jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) CrossOrgUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) CrossOrgUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) Env() *string {
	var returns *string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) EnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) GroupBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) GroupByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) InternalValue() *PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQuery {
	var returns *PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) OperationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) PrimaryTagName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) PrimaryTagNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) PrimaryTagValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) PrimaryTagValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) Stat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) StatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference_Override(p PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetCrossOrgUuids(val *[]*string) {
	if err := j.validateSetCrossOrgUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossOrgUuids",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetEnv(val *string) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetGroupBy(val *[]*string) {
	if err := j.validateSetGroupByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupBy",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetInternalValue(val *PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetOperationName(val *string) {
	if err := j.validateSetOperationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetPrimaryTagName(val *string) {
	if err := j.validateSetPrimaryTagNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryTagName",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetPrimaryTagValue(val *string) {
	if err := j.validateSetPrimaryTagValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryTagValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetResourceName(val *string) {
	if err := j.validateSetResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceName",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetStat(val *string) {
	if err := j.validateSetStatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stat",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetCrossOrgUuids() {
	_jsii_.InvokeVoid(
		p,
		"resetCrossOrgUuids",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetOperationName() {
	_jsii_.InvokeVoid(
		p,
		"resetOperationName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetPrimaryTagName() {
	_jsii_.InvokeVoid(
		p,
		"resetPrimaryTagName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetPrimaryTagValue() {
	_jsii_.InvokeVoid(
		p,
		"resetPrimaryTagValue",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetResourceName() {
	_jsii_.InvokeVoid(
		p,
		"resetResourceName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetGeomapDefinitionRequestQueryApmResourceStatsQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

