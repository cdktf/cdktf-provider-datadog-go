// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference interface {
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
	InternalValue() *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormat
	SetInternalValue(val *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormat)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnitOutputReference
	UnitInput() *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnit
	UnitScale() PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScaleOutputReference
	UnitScaleInput() *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScale
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
	PutUnit(value *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnit)
	PutUnitScale(value *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScale)
	ResetUnitScale()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference
type jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) InternalValue() *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormat {
	var returns *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormat
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) Unit() PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnitOutputReference {
	var returns PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnitOutputReference
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) UnitInput() *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnit {
	var returns *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnit
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) UnitScale() PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScaleOutputReference {
	var returns PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScaleOutputReference
	_jsii_.Get(
		j,
		"unitScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) UnitScaleInput() *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScale {
	var returns *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScale
	_jsii_.Get(
		j,
		"unitScaleInput",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference_Override(p PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference)SetInternalValue(val *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormat) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) PutUnit(value *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnit) {
	if err := p.validatePutUnitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUnit",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) PutUnitScale(value *PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScale) {
	if err := p.validatePutUnitScaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUnitScale",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) ResetUnitScale() {
	_jsii_.InvokeVoid(
		p,
		"resetUnitScale",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

