// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/powerpack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference interface {
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
	InternalValue() *PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace
	SetInternalValue(val *PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace)
	Substring() *string
	SetSubstring(val *string)
	SubstringInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	With() *string
	SetWith(val *string)
	WithInput() *string
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
	ResetSubstring()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference
type jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) InternalValue() *PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace {
	var returns *PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) Substring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"substring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) SubstringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"substringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) With() *string {
	var returns *string
	_jsii_.Get(
		j,
		"with",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) WithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"withInput",
		&returns,
	)
	return returns
}


func NewPowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference_Override(p PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.powerpack.PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference)SetInternalValue(val *PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference)SetSubstring(val *string) {
	if err := j.validateSetSubstringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"substring",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference)SetWith(val *string) {
	if err := j.validateSetWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"with",
		val,
	)
}

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) ResetSubstring() {
	_jsii_.InvokeVoid(
		p,
		"resetSubstring",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

