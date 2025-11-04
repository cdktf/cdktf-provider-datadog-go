// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptions
	SetInternalValue(val *PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptions)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrendType() *string
	SetTrendType(val *string)
	TrendTypeInput() *string
	YScale() *string
	SetYScale(val *string)
	YScaleInput() *string
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
	ResetTrendType()
	ResetYScale()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference
type jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) InternalValue() *PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptions {
	var returns *PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) TrendType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trendType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) TrendTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trendTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) YScale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"yScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) YScaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"yScaleInput",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference_Override(p PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference)SetInternalValue(val *PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference)SetTrendType(val *string) {
	if err := j.validateSetTrendTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trendType",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference)SetYScale(val *string) {
	if err := j.validateSetYScaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yScale",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) ResetTrendType() {
	_jsii_.InvokeVoid(
		p,
		"resetTrendType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) ResetYScale() {
	_jsii_.InvokeVoid(
		p,
		"resetYScale",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

