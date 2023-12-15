// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v10/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetTimeseriesDefinitionRequestOutputReference interface {
	cdktf.ComplexObject
	ApmQuery() PowerpackWidgetTimeseriesDefinitionRequestApmQueryOutputReference
	ApmQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestApmQuery
	AuditQuery() PowerpackWidgetTimeseriesDefinitionRequestAuditQueryOutputReference
	AuditQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestAuditQuery
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
	DisplayType() *string
	SetDisplayType(val *string)
	DisplayTypeInput() *string
	Formula() PowerpackWidgetTimeseriesDefinitionRequestFormulaList
	FormulaInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() PowerpackWidgetTimeseriesDefinitionRequestLogQueryOutputReference
	LogQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestLogQuery
	Metadata() PowerpackWidgetTimeseriesDefinitionRequestMetadataList
	MetadataInput() interface{}
	NetworkQuery() PowerpackWidgetTimeseriesDefinitionRequestNetworkQueryOutputReference
	NetworkQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestNetworkQuery
	OnRightYaxis() interface{}
	SetOnRightYaxis(val interface{})
	OnRightYaxisInput() interface{}
	ProcessQuery() PowerpackWidgetTimeseriesDefinitionRequestProcessQueryOutputReference
	ProcessQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	Query() PowerpackWidgetTimeseriesDefinitionRequestQueryList
	QueryInput() interface{}
	RumQuery() PowerpackWidgetTimeseriesDefinitionRequestRumQueryOutputReference
	RumQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestRumQuery
	SecurityQuery() PowerpackWidgetTimeseriesDefinitionRequestSecurityQueryOutputReference
	SecurityQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestSecurityQuery
	Style() PowerpackWidgetTimeseriesDefinitionRequestStyleOutputReference
	StyleInput() *PowerpackWidgetTimeseriesDefinitionRequestStyle
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
	PutApmQuery(value *PowerpackWidgetTimeseriesDefinitionRequestApmQuery)
	PutAuditQuery(value *PowerpackWidgetTimeseriesDefinitionRequestAuditQuery)
	PutFormula(value interface{})
	PutLogQuery(value *PowerpackWidgetTimeseriesDefinitionRequestLogQuery)
	PutMetadata(value interface{})
	PutNetworkQuery(value *PowerpackWidgetTimeseriesDefinitionRequestNetworkQuery)
	PutProcessQuery(value *PowerpackWidgetTimeseriesDefinitionRequestProcessQuery)
	PutQuery(value interface{})
	PutRumQuery(value *PowerpackWidgetTimeseriesDefinitionRequestRumQuery)
	PutSecurityQuery(value *PowerpackWidgetTimeseriesDefinitionRequestSecurityQuery)
	PutStyle(value *PowerpackWidgetTimeseriesDefinitionRequestStyle)
	ResetApmQuery()
	ResetAuditQuery()
	ResetDisplayType()
	ResetFormula()
	ResetLogQuery()
	ResetMetadata()
	ResetNetworkQuery()
	ResetOnRightYaxis()
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

// The jsii proxy struct for PowerpackWidgetTimeseriesDefinitionRequestOutputReference
type jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ApmQuery() PowerpackWidgetTimeseriesDefinitionRequestApmQueryOutputReference {
	var returns PowerpackWidgetTimeseriesDefinitionRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ApmQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestApmQuery {
	var returns *PowerpackWidgetTimeseriesDefinitionRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) AuditQuery() PowerpackWidgetTimeseriesDefinitionRequestAuditQueryOutputReference {
	var returns PowerpackWidgetTimeseriesDefinitionRequestAuditQueryOutputReference
	_jsii_.Get(
		j,
		"auditQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) AuditQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestAuditQuery {
	var returns *PowerpackWidgetTimeseriesDefinitionRequestAuditQuery
	_jsii_.Get(
		j,
		"auditQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) DisplayType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) DisplayTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) Formula() PowerpackWidgetTimeseriesDefinitionRequestFormulaList {
	var returns PowerpackWidgetTimeseriesDefinitionRequestFormulaList
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) FormulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) LogQuery() PowerpackWidgetTimeseriesDefinitionRequestLogQueryOutputReference {
	var returns PowerpackWidgetTimeseriesDefinitionRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) LogQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestLogQuery {
	var returns *PowerpackWidgetTimeseriesDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) Metadata() PowerpackWidgetTimeseriesDefinitionRequestMetadataList {
	var returns PowerpackWidgetTimeseriesDefinitionRequestMetadataList
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) NetworkQuery() PowerpackWidgetTimeseriesDefinitionRequestNetworkQueryOutputReference {
	var returns PowerpackWidgetTimeseriesDefinitionRequestNetworkQueryOutputReference
	_jsii_.Get(
		j,
		"networkQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) NetworkQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestNetworkQuery {
	var returns *PowerpackWidgetTimeseriesDefinitionRequestNetworkQuery
	_jsii_.Get(
		j,
		"networkQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) OnRightYaxis() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onRightYaxis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) OnRightYaxisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onRightYaxisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ProcessQuery() PowerpackWidgetTimeseriesDefinitionRequestProcessQueryOutputReference {
	var returns PowerpackWidgetTimeseriesDefinitionRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ProcessQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestProcessQuery {
	var returns *PowerpackWidgetTimeseriesDefinitionRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) Query() PowerpackWidgetTimeseriesDefinitionRequestQueryList {
	var returns PowerpackWidgetTimeseriesDefinitionRequestQueryList
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) RumQuery() PowerpackWidgetTimeseriesDefinitionRequestRumQueryOutputReference {
	var returns PowerpackWidgetTimeseriesDefinitionRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) RumQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestRumQuery {
	var returns *PowerpackWidgetTimeseriesDefinitionRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) SecurityQuery() PowerpackWidgetTimeseriesDefinitionRequestSecurityQueryOutputReference {
	var returns PowerpackWidgetTimeseriesDefinitionRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) SecurityQueryInput() *PowerpackWidgetTimeseriesDefinitionRequestSecurityQuery {
	var returns *PowerpackWidgetTimeseriesDefinitionRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) Style() PowerpackWidgetTimeseriesDefinitionRequestStyleOutputReference {
	var returns PowerpackWidgetTimeseriesDefinitionRequestStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) StyleInput() *PowerpackWidgetTimeseriesDefinitionRequestStyle {
	var returns *PowerpackWidgetTimeseriesDefinitionRequestStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetTimeseriesDefinitionRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackWidgetTimeseriesDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetTimeseriesDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetTimeseriesDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackWidgetTimeseriesDefinitionRequestOutputReference_Override(p PowerpackWidgetTimeseriesDefinitionRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetTimeseriesDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference)SetDisplayType(val *string) {
	if err := j.validateSetDisplayTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayType",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference)SetOnRightYaxis(val interface{}) {
	if err := j.validateSetOnRightYaxisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onRightYaxis",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) PutApmQuery(value *PowerpackWidgetTimeseriesDefinitionRequestApmQuery) {
	if err := p.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) PutAuditQuery(value *PowerpackWidgetTimeseriesDefinitionRequestAuditQuery) {
	if err := p.validatePutAuditQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAuditQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) PutFormula(value interface{}) {
	if err := p.validatePutFormulaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFormula",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) PutLogQuery(value *PowerpackWidgetTimeseriesDefinitionRequestLogQuery) {
	if err := p.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) PutMetadata(value interface{}) {
	if err := p.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetadata",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) PutNetworkQuery(value *PowerpackWidgetTimeseriesDefinitionRequestNetworkQuery) {
	if err := p.validatePutNetworkQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNetworkQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) PutProcessQuery(value *PowerpackWidgetTimeseriesDefinitionRequestProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) PutQuery(value interface{}) {
	if err := p.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) PutRumQuery(value *PowerpackWidgetTimeseriesDefinitionRequestRumQuery) {
	if err := p.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) PutSecurityQuery(value *PowerpackWidgetTimeseriesDefinitionRequestSecurityQuery) {
	if err := p.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) PutStyle(value *PowerpackWidgetTimeseriesDefinitionRequestStyle) {
	if err := p.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putStyle",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetAuditQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetAuditQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetDisplayType() {
	_jsii_.InvokeVoid(
		p,
		"resetDisplayType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetFormula() {
	_jsii_.InvokeVoid(
		p,
		"resetFormula",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		p,
		"resetMetadata",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetNetworkQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetNetworkQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetOnRightYaxis() {
	_jsii_.InvokeVoid(
		p,
		"resetOnRightYaxis",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		p,
		"resetQ",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		p,
		"resetStyle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetTimeseriesDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

