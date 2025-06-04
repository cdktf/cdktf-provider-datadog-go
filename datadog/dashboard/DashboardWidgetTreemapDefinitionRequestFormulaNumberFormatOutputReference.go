// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference interface {
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
	InternalValue() *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormat
	SetInternalValue(val *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormat)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnitOutputReference
	UnitInput() *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnit
	UnitScale() DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScaleOutputReference
	UnitScaleInput() *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScale
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
	PutUnit(value *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnit)
	PutUnitScale(value *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScale)
	ResetUnitScale()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference
type jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) InternalValue() *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormat {
	var returns *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormat
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) Unit() DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnitOutputReference {
	var returns DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnitOutputReference
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) UnitInput() *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnit {
	var returns *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnit
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) UnitScale() DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScaleOutputReference {
	var returns DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScaleOutputReference
	_jsii_.Get(
		j,
		"unitScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) UnitScaleInput() *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScale {
	var returns *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScale
	_jsii_.Get(
		j,
		"unitScaleInput",
		&returns,
	)
	return returns
}


func NewDashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference_Override(d DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference)SetInternalValue(val *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormat) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) PutUnit(value *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnit) {
	if err := d.validatePutUnitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUnit",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) PutUnitScale(value *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnitScale) {
	if err := d.validatePutUnitScaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUnitScale",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) ResetUnitScale() {
	_jsii_.InvokeVoid(
		d,
		"resetUnitScale",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

