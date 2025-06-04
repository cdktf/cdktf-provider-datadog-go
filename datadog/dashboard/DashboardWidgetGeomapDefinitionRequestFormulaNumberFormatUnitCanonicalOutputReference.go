// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference interface {
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
	InternalValue() *DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonical
	SetInternalValue(val *DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonical)
	PerUnitName() *string
	SetPerUnitName(val *string)
	PerUnitNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnitName() *string
	SetUnitName(val *string)
	UnitNameInput() *string
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
	ResetPerUnitName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference
type jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) InternalValue() *DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonical {
	var returns *DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonical
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) PerUnitName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perUnitName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) PerUnitNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perUnitNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) UnitName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) UnitNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitNameInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference_Override(d DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetInternalValue(val *DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonical) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetPerUnitName(val *string) {
	if err := j.validateSetPerUnitNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perUnitName",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetUnitName(val *string) {
	if err := j.validateSetUnitNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitName",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) ResetPerUnitName() {
	_jsii_.InvokeVoid(
		d,
		"resetPerUnitName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

