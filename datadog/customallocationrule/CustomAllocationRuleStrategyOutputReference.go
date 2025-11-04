// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customallocationrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/customallocationrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomAllocationRuleStrategyOutputReference interface {
	cdktf.ComplexObject
	AllocatedBy() CustomAllocationRuleStrategyAllocatedByList
	AllocatedByFilters() CustomAllocationRuleStrategyAllocatedByFiltersList
	AllocatedByFiltersInput() interface{}
	AllocatedByInput() interface{}
	AllocatedByTagKeys() *[]*string
	SetAllocatedByTagKeys(val *[]*string)
	AllocatedByTagKeysInput() *[]*string
	BasedOnCosts() CustomAllocationRuleStrategyBasedOnCostsList
	BasedOnCostsInput() interface{}
	BasedOnTimeseries() CustomAllocationRuleStrategyBasedOnTimeseriesOutputReference
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
	EvaluateGroupedByFilters() CustomAllocationRuleStrategyEvaluateGroupedByFiltersList
	EvaluateGroupedByFiltersInput() interface{}
	EvaluateGroupedByTagKeys() *[]*string
	SetEvaluateGroupedByTagKeys(val *[]*string)
	EvaluateGroupedByTagKeysInput() *[]*string
	// Experimental.
	Fqn() *string
	Granularity() *string
	SetGranularity(val *string)
	GranularityInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Method() *string
	SetMethod(val *string)
	MethodInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutAllocatedBy(value interface{})
	PutAllocatedByFilters(value interface{})
	PutBasedOnCosts(value interface{})
	PutBasedOnTimeseries(value *CustomAllocationRuleStrategyBasedOnTimeseries)
	PutEvaluateGroupedByFilters(value interface{})
	ResetAllocatedBy()
	ResetAllocatedByFilters()
	ResetAllocatedByTagKeys()
	ResetBasedOnCosts()
	ResetEvaluateGroupedByFilters()
	ResetEvaluateGroupedByTagKeys()
	ResetGranularity()
	ResetMethod()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomAllocationRuleStrategyOutputReference
type jsiiProxy_CustomAllocationRuleStrategyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) AllocatedBy() CustomAllocationRuleStrategyAllocatedByList {
	var returns CustomAllocationRuleStrategyAllocatedByList
	_jsii_.Get(
		j,
		"allocatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) AllocatedByFilters() CustomAllocationRuleStrategyAllocatedByFiltersList {
	var returns CustomAllocationRuleStrategyAllocatedByFiltersList
	_jsii_.Get(
		j,
		"allocatedByFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) AllocatedByFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allocatedByFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) AllocatedByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allocatedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) AllocatedByTagKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allocatedByTagKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) AllocatedByTagKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allocatedByTagKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) BasedOnCosts() CustomAllocationRuleStrategyBasedOnCostsList {
	var returns CustomAllocationRuleStrategyBasedOnCostsList
	_jsii_.Get(
		j,
		"basedOnCosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) BasedOnCostsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basedOnCostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) BasedOnTimeseries() CustomAllocationRuleStrategyBasedOnTimeseriesOutputReference {
	var returns CustomAllocationRuleStrategyBasedOnTimeseriesOutputReference
	_jsii_.Get(
		j,
		"basedOnTimeseries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) BasedOnTimeseriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basedOnTimeseriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) EvaluateGroupedByFilters() CustomAllocationRuleStrategyEvaluateGroupedByFiltersList {
	var returns CustomAllocationRuleStrategyEvaluateGroupedByFiltersList
	_jsii_.Get(
		j,
		"evaluateGroupedByFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) EvaluateGroupedByFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluateGroupedByFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) EvaluateGroupedByTagKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"evaluateGroupedByTagKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) EvaluateGroupedByTagKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"evaluateGroupedByTagKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) Granularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) GranularityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granularityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) MethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCustomAllocationRuleStrategyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CustomAllocationRuleStrategyOutputReference {
	_init_.Initialize()

	if err := validateNewCustomAllocationRuleStrategyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomAllocationRuleStrategyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.customAllocationRule.CustomAllocationRuleStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomAllocationRuleStrategyOutputReference_Override(c CustomAllocationRuleStrategyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.customAllocationRule.CustomAllocationRuleStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference)SetAllocatedByTagKeys(val *[]*string) {
	if err := j.validateSetAllocatedByTagKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedByTagKeys",
		val,
	)
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference)SetEvaluateGroupedByTagKeys(val *[]*string) {
	if err := j.validateSetEvaluateGroupedByTagKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluateGroupedByTagKeys",
		val,
	)
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference)SetGranularity(val *string) {
	if err := j.validateSetGranularityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"granularity",
		val,
	)
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomAllocationRuleStrategyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) PutAllocatedBy(value interface{}) {
	if err := c.validatePutAllocatedByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAllocatedBy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) PutAllocatedByFilters(value interface{}) {
	if err := c.validatePutAllocatedByFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAllocatedByFilters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) PutBasedOnCosts(value interface{}) {
	if err := c.validatePutBasedOnCostsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBasedOnCosts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) PutBasedOnTimeseries(value *CustomAllocationRuleStrategyBasedOnTimeseries) {
	if err := c.validatePutBasedOnTimeseriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBasedOnTimeseries",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) PutEvaluateGroupedByFilters(value interface{}) {
	if err := c.validatePutEvaluateGroupedByFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEvaluateGroupedByFilters",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) ResetAllocatedBy() {
	_jsii_.InvokeVoid(
		c,
		"resetAllocatedBy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) ResetAllocatedByFilters() {
	_jsii_.InvokeVoid(
		c,
		"resetAllocatedByFilters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) ResetAllocatedByTagKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetAllocatedByTagKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) ResetBasedOnCosts() {
	_jsii_.InvokeVoid(
		c,
		"resetBasedOnCosts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) ResetEvaluateGroupedByFilters() {
	_jsii_.InvokeVoid(
		c,
		"resetEvaluateGroupedByFilters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) ResetEvaluateGroupedByTagKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetEvaluateGroupedByTagKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) ResetGranularity() {
	_jsii_.InvokeVoid(
		c,
		"resetGranularity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetMethod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomAllocationRuleStrategyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

