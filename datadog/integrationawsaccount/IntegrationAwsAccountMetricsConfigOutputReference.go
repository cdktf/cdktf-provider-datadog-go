// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/integrationawsaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationAwsAccountMetricsConfigOutputReference interface {
	cdktf.ComplexObject
	AutomuteEnabled() interface{}
	SetAutomuteEnabled(val interface{})
	AutomuteEnabledInput() interface{}
	CollectCloudwatchAlarms() interface{}
	SetCollectCloudwatchAlarms(val interface{})
	CollectCloudwatchAlarmsInput() interface{}
	CollectCustomMetrics() interface{}
	SetCollectCustomMetrics(val interface{})
	CollectCustomMetricsInput() interface{}
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NamespaceFilters() IntegrationAwsAccountMetricsConfigNamespaceFiltersOutputReference
	NamespaceFiltersInput() interface{}
	TagFilters() IntegrationAwsAccountMetricsConfigTagFiltersList
	TagFiltersInput() interface{}
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
	PutNamespaceFilters(value *IntegrationAwsAccountMetricsConfigNamespaceFilters)
	PutTagFilters(value interface{})
	ResetAutomuteEnabled()
	ResetCollectCloudwatchAlarms()
	ResetCollectCustomMetrics()
	ResetEnabled()
	ResetNamespaceFilters()
	ResetTagFilters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationAwsAccountMetricsConfigOutputReference
type jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) AutomuteEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automuteEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) AutomuteEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automuteEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) CollectCloudwatchAlarms() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectCloudwatchAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) CollectCloudwatchAlarmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectCloudwatchAlarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) CollectCustomMetrics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectCustomMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) CollectCustomMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectCustomMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) NamespaceFilters() IntegrationAwsAccountMetricsConfigNamespaceFiltersOutputReference {
	var returns IntegrationAwsAccountMetricsConfigNamespaceFiltersOutputReference
	_jsii_.Get(
		j,
		"namespaceFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) NamespaceFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"namespaceFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) TagFilters() IntegrationAwsAccountMetricsConfigTagFiltersList {
	var returns IntegrationAwsAccountMetricsConfigTagFiltersList
	_jsii_.Get(
		j,
		"tagFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) TagFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIntegrationAwsAccountMetricsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IntegrationAwsAccountMetricsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationAwsAccountMetricsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.integrationAwsAccount.IntegrationAwsAccountMetricsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIntegrationAwsAccountMetricsConfigOutputReference_Override(i IntegrationAwsAccountMetricsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.integrationAwsAccount.IntegrationAwsAccountMetricsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference)SetAutomuteEnabled(val interface{}) {
	if err := j.validateSetAutomuteEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automuteEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference)SetCollectCloudwatchAlarms(val interface{}) {
	if err := j.validateSetCollectCloudwatchAlarmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectCloudwatchAlarms",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference)SetCollectCustomMetrics(val interface{}) {
	if err := j.validateSetCollectCustomMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectCustomMetrics",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) PutNamespaceFilters(value *IntegrationAwsAccountMetricsConfigNamespaceFilters) {
	if err := i.validatePutNamespaceFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putNamespaceFilters",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) PutTagFilters(value interface{}) {
	if err := i.validatePutTagFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTagFilters",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) ResetAutomuteEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetAutomuteEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) ResetCollectCloudwatchAlarms() {
	_jsii_.InvokeVoid(
		i,
		"resetCollectCloudwatchAlarms",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) ResetCollectCustomMetrics() {
	_jsii_.InvokeVoid(
		i,
		"resetCollectCustomMetrics",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) ResetNamespaceFilters() {
	_jsii_.InvokeVoid(
		i,
		"resetNamespaceFilters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) ResetTagFilters() {
	_jsii_.InvokeVoid(
		i,
		"resetTagFilters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

