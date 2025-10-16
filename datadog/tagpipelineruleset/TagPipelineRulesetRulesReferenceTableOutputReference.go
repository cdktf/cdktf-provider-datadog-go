// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagpipelineruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v12/tagpipelineruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TagPipelineRulesetRulesReferenceTableOutputReference interface {
	cdktf.ComplexObject
	CaseInsensitivity() interface{}
	SetCaseInsensitivity(val interface{})
	CaseInsensitivityInput() interface{}
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
	FieldPairs() TagPipelineRulesetRulesReferenceTableFieldPairsList
	FieldPairsInput() interface{}
	// Experimental.
	Fqn() *string
	IfNotExists() interface{}
	SetIfNotExists(val interface{})
	IfNotExistsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SourceKeys() *[]*string
	SetSourceKeys(val *[]*string)
	SourceKeysInput() *[]*string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
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
	PutFieldPairs(value interface{})
	ResetCaseInsensitivity()
	ResetFieldPairs()
	ResetIfNotExists()
	ResetSourceKeys()
	ResetTableName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TagPipelineRulesetRulesReferenceTableOutputReference
type jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) CaseInsensitivity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseInsensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) CaseInsensitivityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseInsensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) FieldPairs() TagPipelineRulesetRulesReferenceTableFieldPairsList {
	var returns TagPipelineRulesetRulesReferenceTableFieldPairsList
	_jsii_.Get(
		j,
		"fieldPairs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) FieldPairsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldPairsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) IfNotExists() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ifNotExists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) IfNotExistsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ifNotExistsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) SourceKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) SourceKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTagPipelineRulesetRulesReferenceTableOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TagPipelineRulesetRulesReferenceTableOutputReference {
	_init_.Initialize()

	if err := validateNewTagPipelineRulesetRulesReferenceTableOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.tagPipelineRuleset.TagPipelineRulesetRulesReferenceTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTagPipelineRulesetRulesReferenceTableOutputReference_Override(t TagPipelineRulesetRulesReferenceTableOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.tagPipelineRuleset.TagPipelineRulesetRulesReferenceTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference)SetCaseInsensitivity(val interface{}) {
	if err := j.validateSetCaseInsensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caseInsensitivity",
		val,
	)
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference)SetIfNotExists(val interface{}) {
	if err := j.validateSetIfNotExistsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ifNotExists",
		val,
	)
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference)SetSourceKeys(val *[]*string) {
	if err := j.validateSetSourceKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceKeys",
		val,
	)
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) PutFieldPairs(value interface{}) {
	if err := t.validatePutFieldPairsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFieldPairs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) ResetCaseInsensitivity() {
	_jsii_.InvokeVoid(
		t,
		"resetCaseInsensitivity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) ResetFieldPairs() {
	_jsii_.InvokeVoid(
		t,
		"resetFieldPairs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) ResetIfNotExists() {
	_jsii_.InvokeVoid(
		t,
		"resetIfNotExists",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) ResetSourceKeys() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceKeys",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) ResetTableName() {
	_jsii_.InvokeVoid(
		t,
		"resetTableName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagPipelineRulesetRulesReferenceTableOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

