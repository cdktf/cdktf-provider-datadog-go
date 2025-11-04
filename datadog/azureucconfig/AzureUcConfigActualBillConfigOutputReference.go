// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package azureucconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/azureucconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AzureUcConfigActualBillConfigOutputReference interface {
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
	ExportName() *string
	SetExportName(val *string)
	ExportNameInput() *string
	ExportPath() *string
	SetExportPath(val *string)
	ExportPathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StorageAccount() *string
	SetStorageAccount(val *string)
	StorageAccountInput() *string
	StorageContainer() *string
	SetStorageContainer(val *string)
	StorageContainerInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AzureUcConfigActualBillConfigOutputReference
type jsiiProxy_AzureUcConfigActualBillConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) ExportName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) ExportNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) ExportPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) ExportPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) StorageAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) StorageAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) StorageContainer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) StorageContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAzureUcConfigActualBillConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AzureUcConfigActualBillConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAzureUcConfigActualBillConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AzureUcConfigActualBillConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.azureUcConfig.AzureUcConfigActualBillConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAzureUcConfigActualBillConfigOutputReference_Override(a AzureUcConfigActualBillConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.azureUcConfig.AzureUcConfigActualBillConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference)SetExportName(val *string) {
	if err := j.validateSetExportNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportName",
		val,
	)
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference)SetExportPath(val *string) {
	if err := j.validateSetExportPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportPath",
		val,
	)
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference)SetStorageAccount(val *string) {
	if err := j.validateSetStorageAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccount",
		val,
	)
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference)SetStorageContainer(val *string) {
	if err := j.validateSetStorageContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageContainer",
		val,
	)
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AzureUcConfigActualBillConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AzureUcConfigActualBillConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

