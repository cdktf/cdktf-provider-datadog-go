// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogcustomallocationrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/datadatadogcustomallocationrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogCustomAllocationRuleStrategyOutputReference interface {
	cdktf.ComplexObject
	AllocatedBy() DataDatadogCustomAllocationRuleStrategyAllocatedByList
	AllocatedByFilters() DataDatadogCustomAllocationRuleStrategyAllocatedByFiltersList
	AllocatedByFiltersInput() interface{}
	AllocatedByInput() interface{}
	AllocatedByTagKeys() *[]*string
	BasedOnCosts() DataDatadogCustomAllocationRuleStrategyBasedOnCostsList
	BasedOnCostsInput() interface{}
	BasedOnTimeseries() DataDatadogCustomAllocationRuleStrategyBasedOnTimeseriesOutputReference
	BasedOnTimeseriesInput() interface{}
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
	EvaluateGroupedByFilters() DataDatadogCustomAllocationRuleStrategyEvaluateGroupedByFiltersList
	EvaluateGroupedByFiltersInput() interface{}
	EvaluateGroupedByTagKeys() *[]*string
	// Experimental.
	Fqn() *string
	Granularity() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Method() *string
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
	PutAllocatedBy(value interface{})
	PutAllocatedByFilters(value interface{})
	PutBasedOnCosts(value interface{})
	PutBasedOnTimeseries(value *DataDatadogCustomAllocationRuleStrategyBasedOnTimeseries)
	PutEvaluateGroupedByFilters(value interface{})
	ResetAllocatedBy()
	ResetAllocatedByFilters()
	ResetBasedOnCosts()
	ResetEvaluateGroupedByFilters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatadogCustomAllocationRuleStrategyOutputReference
type jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) AllocatedBy() DataDatadogCustomAllocationRuleStrategyAllocatedByList {
	var returns DataDatadogCustomAllocationRuleStrategyAllocatedByList
	_jsii_.Get(
		j,
		"allocatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) AllocatedByFilters() DataDatadogCustomAllocationRuleStrategyAllocatedByFiltersList {
	var returns DataDatadogCustomAllocationRuleStrategyAllocatedByFiltersList
	_jsii_.Get(
		j,
		"allocatedByFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) AllocatedByFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allocatedByFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) AllocatedByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allocatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) AllocatedByTagKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allocatedByTagKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) BasedOnCosts() DataDatadogCustomAllocationRuleStrategyBasedOnCostsList {
	var returns DataDatadogCustomAllocationRuleStrategyBasedOnCostsList
	_jsii_.Get(
		j,
		"basedOnCosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) BasedOnCostsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basedOnCostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) BasedOnTimeseries() DataDatadogCustomAllocationRuleStrategyBasedOnTimeseriesOutputReference {
	var returns DataDatadogCustomAllocationRuleStrategyBasedOnTimeseriesOutputReference
	_jsii_.Get(
		j,
		"basedOnTimeseries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) BasedOnTimeseriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basedOnTimeseriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) EvaluateGroupedByFilters() DataDatadogCustomAllocationRuleStrategyEvaluateGroupedByFiltersList {
	var returns DataDatadogCustomAllocationRuleStrategyEvaluateGroupedByFiltersList
	_jsii_.Get(
		j,
		"evaluateGroupedByFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) EvaluateGroupedByFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluateGroupedByFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) EvaluateGroupedByTagKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"evaluateGroupedByTagKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) Granularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataDatadogCustomAllocationRuleStrategyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataDatadogCustomAllocationRuleStrategyOutputReference {
	_init_.Initialize()

	if err := validateNewDataDatadogCustomAllocationRuleStrategyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dataDatadogCustomAllocationRule.DataDatadogCustomAllocationRuleStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataDatadogCustomAllocationRuleStrategyOutputReference_Override(d DataDatadogCustomAllocationRuleStrategyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dataDatadogCustomAllocationRule.DataDatadogCustomAllocationRuleStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) PutAllocatedBy(value interface{}) {
	if err := d.validatePutAllocatedByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAllocatedBy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) PutAllocatedByFilters(value interface{}) {
	if err := d.validatePutAllocatedByFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAllocatedByFilters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) PutBasedOnCosts(value interface{}) {
	if err := d.validatePutBasedOnCostsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBasedOnCosts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) PutBasedOnTimeseries(value *DataDatadogCustomAllocationRuleStrategyBasedOnTimeseries) {
	if err := d.validatePutBasedOnTimeseriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBasedOnTimeseries",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) PutEvaluateGroupedByFilters(value interface{}) {
	if err := d.validatePutEvaluateGroupedByFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEvaluateGroupedByFilters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) ResetAllocatedBy() {
	_jsii_.InvokeVoid(
		d,
		"resetAllocatedBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) ResetAllocatedByFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetAllocatedByFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) ResetBasedOnCosts() {
	_jsii_.InvokeVoid(
		d,
		"resetBasedOnCosts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) ResetEvaluateGroupedByFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetEvaluateGroupedByFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatadogCustomAllocationRuleStrategyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

