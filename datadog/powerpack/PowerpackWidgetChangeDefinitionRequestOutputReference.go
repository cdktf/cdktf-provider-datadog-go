// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v11/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetChangeDefinitionRequestOutputReference interface {
	cdktf.ComplexObject
	ApmQuery() PowerpackWidgetChangeDefinitionRequestApmQueryOutputReference
	ApmQueryInput() *PowerpackWidgetChangeDefinitionRequestApmQuery
	ChangeType() *string
	SetChangeType(val *string)
	ChangeTypeInput() *string
	CompareTo() *string
	SetCompareTo(val *string)
	CompareToInput() *string
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
	Formula() PowerpackWidgetChangeDefinitionRequestFormulaList
	FormulaInput() interface{}
	// Experimental.
	Fqn() *string
	IncreaseGood() interface{}
	SetIncreaseGood(val interface{})
	IncreaseGoodInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() PowerpackWidgetChangeDefinitionRequestLogQueryOutputReference
	LogQueryInput() *PowerpackWidgetChangeDefinitionRequestLogQuery
	OrderBy() *string
	SetOrderBy(val *string)
	OrderByInput() *string
	OrderDir() *string
	SetOrderDir(val *string)
	OrderDirInput() *string
	ProcessQuery() PowerpackWidgetChangeDefinitionRequestProcessQueryOutputReference
	ProcessQueryInput() *PowerpackWidgetChangeDefinitionRequestProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	Query() PowerpackWidgetChangeDefinitionRequestQueryList
	QueryInput() interface{}
	RumQuery() PowerpackWidgetChangeDefinitionRequestRumQueryOutputReference
	RumQueryInput() *PowerpackWidgetChangeDefinitionRequestRumQuery
	SecurityQuery() PowerpackWidgetChangeDefinitionRequestSecurityQueryOutputReference
	SecurityQueryInput() *PowerpackWidgetChangeDefinitionRequestSecurityQuery
	ShowPresent() interface{}
	SetShowPresent(val interface{})
	ShowPresentInput() interface{}
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
	PutApmQuery(value *PowerpackWidgetChangeDefinitionRequestApmQuery)
	PutFormula(value interface{})
	PutLogQuery(value *PowerpackWidgetChangeDefinitionRequestLogQuery)
	PutProcessQuery(value *PowerpackWidgetChangeDefinitionRequestProcessQuery)
	PutQuery(value interface{})
	PutRumQuery(value *PowerpackWidgetChangeDefinitionRequestRumQuery)
	PutSecurityQuery(value *PowerpackWidgetChangeDefinitionRequestSecurityQuery)
	ResetApmQuery()
	ResetChangeType()
	ResetCompareTo()
	ResetFormula()
	ResetIncreaseGood()
	ResetLogQuery()
	ResetOrderBy()
	ResetOrderDir()
	ResetProcessQuery()
	ResetQ()
	ResetQuery()
	ResetRumQuery()
	ResetSecurityQuery()
	ResetShowPresent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetChangeDefinitionRequestOutputReference
type jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ApmQuery() PowerpackWidgetChangeDefinitionRequestApmQueryOutputReference {
	var returns PowerpackWidgetChangeDefinitionRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ApmQueryInput() *PowerpackWidgetChangeDefinitionRequestApmQuery {
	var returns *PowerpackWidgetChangeDefinitionRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ChangeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ChangeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) CompareTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) CompareToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) Formula() PowerpackWidgetChangeDefinitionRequestFormulaList {
	var returns PowerpackWidgetChangeDefinitionRequestFormulaList
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) FormulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) IncreaseGood() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"increaseGood",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) IncreaseGoodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"increaseGoodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) LogQuery() PowerpackWidgetChangeDefinitionRequestLogQueryOutputReference {
	var returns PowerpackWidgetChangeDefinitionRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) LogQueryInput() *PowerpackWidgetChangeDefinitionRequestLogQuery {
	var returns *PowerpackWidgetChangeDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) OrderBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) OrderByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) OrderDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) OrderDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ProcessQuery() PowerpackWidgetChangeDefinitionRequestProcessQueryOutputReference {
	var returns PowerpackWidgetChangeDefinitionRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ProcessQueryInput() *PowerpackWidgetChangeDefinitionRequestProcessQuery {
	var returns *PowerpackWidgetChangeDefinitionRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) Query() PowerpackWidgetChangeDefinitionRequestQueryList {
	var returns PowerpackWidgetChangeDefinitionRequestQueryList
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) RumQuery() PowerpackWidgetChangeDefinitionRequestRumQueryOutputReference {
	var returns PowerpackWidgetChangeDefinitionRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) RumQueryInput() *PowerpackWidgetChangeDefinitionRequestRumQuery {
	var returns *PowerpackWidgetChangeDefinitionRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) SecurityQuery() PowerpackWidgetChangeDefinitionRequestSecurityQueryOutputReference {
	var returns PowerpackWidgetChangeDefinitionRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) SecurityQueryInput() *PowerpackWidgetChangeDefinitionRequestSecurityQuery {
	var returns *PowerpackWidgetChangeDefinitionRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ShowPresent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPresent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ShowPresentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPresentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetChangeDefinitionRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackWidgetChangeDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetChangeDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetChangeDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackWidgetChangeDefinitionRequestOutputReference_Override(p PowerpackWidgetChangeDefinitionRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetChangeDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference)SetChangeType(val *string) {
	if err := j.validateSetChangeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeType",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference)SetCompareTo(val *string) {
	if err := j.validateSetCompareToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compareTo",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference)SetIncreaseGood(val interface{}) {
	if err := j.validateSetIncreaseGoodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"increaseGood",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference)SetOrderBy(val *string) {
	if err := j.validateSetOrderByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orderBy",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference)SetOrderDir(val *string) {
	if err := j.validateSetOrderDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orderDir",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference)SetShowPresent(val interface{}) {
	if err := j.validateSetShowPresentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showPresent",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) PutApmQuery(value *PowerpackWidgetChangeDefinitionRequestApmQuery) {
	if err := p.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) PutFormula(value interface{}) {
	if err := p.validatePutFormulaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFormula",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) PutLogQuery(value *PowerpackWidgetChangeDefinitionRequestLogQuery) {
	if err := p.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) PutProcessQuery(value *PowerpackWidgetChangeDefinitionRequestProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) PutQuery(value interface{}) {
	if err := p.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) PutRumQuery(value *PowerpackWidgetChangeDefinitionRequestRumQuery) {
	if err := p.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) PutSecurityQuery(value *PowerpackWidgetChangeDefinitionRequestSecurityQuery) {
	if err := p.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetChangeType() {
	_jsii_.InvokeVoid(
		p,
		"resetChangeType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetCompareTo() {
	_jsii_.InvokeVoid(
		p,
		"resetCompareTo",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetFormula() {
	_jsii_.InvokeVoid(
		p,
		"resetFormula",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetIncreaseGood() {
	_jsii_.InvokeVoid(
		p,
		"resetIncreaseGood",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetOrderBy() {
	_jsii_.InvokeVoid(
		p,
		"resetOrderBy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetOrderDir() {
	_jsii_.InvokeVoid(
		p,
		"resetOrderDir",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		p,
		"resetQ",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ResetShowPresent() {
	_jsii_.InvokeVoid(
		p,
		"resetShowPresent",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetChangeDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

