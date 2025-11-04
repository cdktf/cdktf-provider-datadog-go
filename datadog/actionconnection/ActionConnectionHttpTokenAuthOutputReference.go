// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/actionconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionConnectionHttpTokenAuthOutputReference interface {
	cdktf.ComplexObject
	Body() ActionConnectionHttpTokenAuthBodyOutputReference
	BodyInput() interface{}
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
	Header() ActionConnectionHttpTokenAuthHeaderList
	HeaderInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Token() ActionConnectionHttpTokenAuthTokenList
	TokenInput() interface{}
	UrlParameter() ActionConnectionHttpTokenAuthUrlParameterList
	UrlParameterInput() interface{}
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
	PutBody(value *ActionConnectionHttpTokenAuthBody)
	PutHeader(value interface{})
	PutToken(value interface{})
	PutUrlParameter(value interface{})
	ResetBody()
	ResetHeader()
	ResetToken()
	ResetUrlParameter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ActionConnectionHttpTokenAuthOutputReference
type jsiiProxy_ActionConnectionHttpTokenAuthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) Body() ActionConnectionHttpTokenAuthBodyOutputReference {
	var returns ActionConnectionHttpTokenAuthBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) BodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) Header() ActionConnectionHttpTokenAuthHeaderList {
	var returns ActionConnectionHttpTokenAuthHeaderList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) Token() ActionConnectionHttpTokenAuthTokenList {
	var returns ActionConnectionHttpTokenAuthTokenList
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) TokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) UrlParameter() ActionConnectionHttpTokenAuthUrlParameterList {
	var returns ActionConnectionHttpTokenAuthUrlParameterList
	_jsii_.Get(
		j,
		"urlParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) UrlParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlParameterInput",
		&returns,
	)
	return returns
}


func NewActionConnectionHttpTokenAuthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ActionConnectionHttpTokenAuthOutputReference {
	_init_.Initialize()

	if err := validateNewActionConnectionHttpTokenAuthOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActionConnectionHttpTokenAuthOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.actionConnection.ActionConnectionHttpTokenAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewActionConnectionHttpTokenAuthOutputReference_Override(a ActionConnectionHttpTokenAuthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.actionConnection.ActionConnectionHttpTokenAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) PutBody(value *ActionConnectionHttpTokenAuthBody) {
	if err := a.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBody",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) PutHeader(value interface{}) {
	if err := a.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) PutToken(value interface{}) {
	if err := a.validatePutTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putToken",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) PutUrlParameter(value interface{}) {
	if err := a.validatePutUrlParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUrlParameter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		a,
		"resetBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) ResetToken() {
	_jsii_.InvokeVoid(
		a,
		"resetToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) ResetUrlParameter() {
	_jsii_.InvokeVoid(
		a,
		"resetUrlParameter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

