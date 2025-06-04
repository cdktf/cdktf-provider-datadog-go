// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetHeatmapDefinitionRequestOutputReference interface {
	cdktf.ComplexObject
	ApmQuery() PowerpackWidgetHeatmapDefinitionRequestApmQueryOutputReference
	ApmQueryInput() *PowerpackWidgetHeatmapDefinitionRequestApmQuery
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
	Formula() PowerpackWidgetHeatmapDefinitionRequestFormulaList
	FormulaInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() PowerpackWidgetHeatmapDefinitionRequestLogQueryOutputReference
	LogQueryInput() *PowerpackWidgetHeatmapDefinitionRequestLogQuery
	ProcessQuery() PowerpackWidgetHeatmapDefinitionRequestProcessQueryOutputReference
	ProcessQueryInput() *PowerpackWidgetHeatmapDefinitionRequestProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	Query() PowerpackWidgetHeatmapDefinitionRequestQueryList
	QueryInput() interface{}
	RumQuery() PowerpackWidgetHeatmapDefinitionRequestRumQueryOutputReference
	RumQueryInput() *PowerpackWidgetHeatmapDefinitionRequestRumQuery
	SecurityQuery() PowerpackWidgetHeatmapDefinitionRequestSecurityQueryOutputReference
	SecurityQueryInput() *PowerpackWidgetHeatmapDefinitionRequestSecurityQuery
	Style() PowerpackWidgetHeatmapDefinitionRequestStyleOutputReference
	StyleInput() *PowerpackWidgetHeatmapDefinitionRequestStyle
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
	PutApmQuery(value *PowerpackWidgetHeatmapDefinitionRequestApmQuery)
	PutFormula(value interface{})
	PutLogQuery(value *PowerpackWidgetHeatmapDefinitionRequestLogQuery)
	PutProcessQuery(value *PowerpackWidgetHeatmapDefinitionRequestProcessQuery)
	PutQuery(value interface{})
	PutRumQuery(value *PowerpackWidgetHeatmapDefinitionRequestRumQuery)
	PutSecurityQuery(value *PowerpackWidgetHeatmapDefinitionRequestSecurityQuery)
	PutStyle(value *PowerpackWidgetHeatmapDefinitionRequestStyle)
	ResetApmQuery()
	ResetFormula()
	ResetLogQuery()
	ResetProcessQuery()
	ResetQ()
	ResetQuery()
	ResetRumQuery()
	ResetSecurityQuery()
	ResetStyle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetHeatmapDefinitionRequestOutputReference
type jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ApmQuery() PowerpackWidgetHeatmapDefinitionRequestApmQueryOutputReference {
	var returns PowerpackWidgetHeatmapDefinitionRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ApmQueryInput() *PowerpackWidgetHeatmapDefinitionRequestApmQuery {
	var returns *PowerpackWidgetHeatmapDefinitionRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) Formula() PowerpackWidgetHeatmapDefinitionRequestFormulaList {
	var returns PowerpackWidgetHeatmapDefinitionRequestFormulaList
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) FormulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) LogQuery() PowerpackWidgetHeatmapDefinitionRequestLogQueryOutputReference {
	var returns PowerpackWidgetHeatmapDefinitionRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) LogQueryInput() *PowerpackWidgetHeatmapDefinitionRequestLogQuery {
	var returns *PowerpackWidgetHeatmapDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ProcessQuery() PowerpackWidgetHeatmapDefinitionRequestProcessQueryOutputReference {
	var returns PowerpackWidgetHeatmapDefinitionRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ProcessQueryInput() *PowerpackWidgetHeatmapDefinitionRequestProcessQuery {
	var returns *PowerpackWidgetHeatmapDefinitionRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) Query() PowerpackWidgetHeatmapDefinitionRequestQueryList {
	var returns PowerpackWidgetHeatmapDefinitionRequestQueryList
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) RumQuery() PowerpackWidgetHeatmapDefinitionRequestRumQueryOutputReference {
	var returns PowerpackWidgetHeatmapDefinitionRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) RumQueryInput() *PowerpackWidgetHeatmapDefinitionRequestRumQuery {
	var returns *PowerpackWidgetHeatmapDefinitionRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) SecurityQuery() PowerpackWidgetHeatmapDefinitionRequestSecurityQueryOutputReference {
	var returns PowerpackWidgetHeatmapDefinitionRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) SecurityQueryInput() *PowerpackWidgetHeatmapDefinitionRequestSecurityQuery {
	var returns *PowerpackWidgetHeatmapDefinitionRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) Style() PowerpackWidgetHeatmapDefinitionRequestStyleOutputReference {
	var returns PowerpackWidgetHeatmapDefinitionRequestStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) StyleInput() *PowerpackWidgetHeatmapDefinitionRequestStyle {
	var returns *PowerpackWidgetHeatmapDefinitionRequestStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetHeatmapDefinitionRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackWidgetHeatmapDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetHeatmapDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetHeatmapDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackWidgetHeatmapDefinitionRequestOutputReference_Override(p PowerpackWidgetHeatmapDefinitionRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetHeatmapDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) PutApmQuery(value *PowerpackWidgetHeatmapDefinitionRequestApmQuery) {
	if err := p.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) PutFormula(value interface{}) {
	if err := p.validatePutFormulaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFormula",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) PutLogQuery(value *PowerpackWidgetHeatmapDefinitionRequestLogQuery) {
	if err := p.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) PutProcessQuery(value *PowerpackWidgetHeatmapDefinitionRequestProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) PutQuery(value interface{}) {
	if err := p.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) PutRumQuery(value *PowerpackWidgetHeatmapDefinitionRequestRumQuery) {
	if err := p.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) PutSecurityQuery(value *PowerpackWidgetHeatmapDefinitionRequestSecurityQuery) {
	if err := p.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) PutStyle(value *PowerpackWidgetHeatmapDefinitionRequestStyle) {
	if err := p.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putStyle",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ResetFormula() {
	_jsii_.InvokeVoid(
		p,
		"resetFormula",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		p,
		"resetQ",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		p,
		"resetStyle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

