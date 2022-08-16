// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/jsii"

	"github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetChangeDefinitionRequestOutputReference interface {
	cdktf.ComplexObject
	ApmQuery() DashboardWidgetChangeDefinitionRequestApmQueryOutputReference
	ApmQueryInput() *DashboardWidgetChangeDefinitionRequestApmQuery
	ChangeType() *string
	SetChangeType(val *string)
	ChangeTypeInput() *string
	CompareTo() *string
	SetCompareTo(val *string)
	CompareToInput() *string
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
	Formula() DashboardWidgetChangeDefinitionRequestFormulaList
	FormulaInput() interface{}
	// Experimental.
	Fqn() *string
	IncreaseGood() interface{}
	SetIncreaseGood(val interface{})
	IncreaseGoodInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() DashboardWidgetChangeDefinitionRequestLogQueryOutputReference
	LogQueryInput() *DashboardWidgetChangeDefinitionRequestLogQuery
	OrderBy() *string
	SetOrderBy(val *string)
	OrderByInput() *string
	OrderDir() *string
	SetOrderDir(val *string)
	OrderDirInput() *string
	ProcessQuery() DashboardWidgetChangeDefinitionRequestProcessQueryOutputReference
	ProcessQueryInput() *DashboardWidgetChangeDefinitionRequestProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	Query() DashboardWidgetChangeDefinitionRequestQueryList
	QueryInput() interface{}
	RumQuery() DashboardWidgetChangeDefinitionRequestRumQueryOutputReference
	RumQueryInput() *DashboardWidgetChangeDefinitionRequestRumQuery
	SecurityQuery() DashboardWidgetChangeDefinitionRequestSecurityQueryOutputReference
	SecurityQueryInput() *DashboardWidgetChangeDefinitionRequestSecurityQuery
	ShowPresent() interface{}
	SetShowPresent(val interface{})
	ShowPresentInput() interface{}
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
	PutApmQuery(value *DashboardWidgetChangeDefinitionRequestApmQuery)
	PutFormula(value interface{})
	PutLogQuery(value *DashboardWidgetChangeDefinitionRequestLogQuery)
	PutProcessQuery(value *DashboardWidgetChangeDefinitionRequestProcessQuery)
	PutQuery(value interface{})
	PutRumQuery(value *DashboardWidgetChangeDefinitionRequestRumQuery)
	PutSecurityQuery(value *DashboardWidgetChangeDefinitionRequestSecurityQuery)
	ResetApmQuery()
	ResetChangeType()
	ResetCompareTo()
	ResetFormula()
	ResetIncreaseGood()
	ResetLogQuery()
	ResetOrderBy()
	ResetOrderDir()
	ResetProcessQuery()
	ResetQ()
	ResetQuery()
	ResetRumQuery()
	ResetSecurityQuery()
	ResetShowPresent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetChangeDefinitionRequestOutputReference
type jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ApmQuery() DashboardWidgetChangeDefinitionRequestApmQueryOutputReference {
	var returns DashboardWidgetChangeDefinitionRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ApmQueryInput() *DashboardWidgetChangeDefinitionRequestApmQuery {
	var returns *DashboardWidgetChangeDefinitionRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ChangeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ChangeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) CompareTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) CompareToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) Formula() DashboardWidgetChangeDefinitionRequestFormulaList {
	var returns DashboardWidgetChangeDefinitionRequestFormulaList
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) FormulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) IncreaseGood() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"increaseGood",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) IncreaseGoodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"increaseGoodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) LogQuery() DashboardWidgetChangeDefinitionRequestLogQueryOutputReference {
	var returns DashboardWidgetChangeDefinitionRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) LogQueryInput() *DashboardWidgetChangeDefinitionRequestLogQuery {
	var returns *DashboardWidgetChangeDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) OrderBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) OrderByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) OrderDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) OrderDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ProcessQuery() DashboardWidgetChangeDefinitionRequestProcessQueryOutputReference {
	var returns DashboardWidgetChangeDefinitionRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ProcessQueryInput() *DashboardWidgetChangeDefinitionRequestProcessQuery {
	var returns *DashboardWidgetChangeDefinitionRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) Query() DashboardWidgetChangeDefinitionRequestQueryList {
	var returns DashboardWidgetChangeDefinitionRequestQueryList
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) RumQuery() DashboardWidgetChangeDefinitionRequestRumQueryOutputReference {
	var returns DashboardWidgetChangeDefinitionRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) RumQueryInput() *DashboardWidgetChangeDefinitionRequestRumQuery {
	var returns *DashboardWidgetChangeDefinitionRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SecurityQuery() DashboardWidgetChangeDefinitionRequestSecurityQueryOutputReference {
	var returns DashboardWidgetChangeDefinitionRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SecurityQueryInput() *DashboardWidgetChangeDefinitionRequestSecurityQuery {
	var returns *DashboardWidgetChangeDefinitionRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ShowPresent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPresent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ShowPresentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPresentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetChangeDefinitionRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardWidgetChangeDefinitionRequestOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.DashboardWidgetChangeDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardWidgetChangeDefinitionRequestOutputReference_Override(d DashboardWidgetChangeDefinitionRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.DashboardWidgetChangeDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SetChangeType(val *string) {
	_jsii_.Set(
		j,
		"changeType",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SetCompareTo(val *string) {
	_jsii_.Set(
		j,
		"compareTo",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SetIncreaseGood(val interface{}) {
	_jsii_.Set(
		j,
		"increaseGood",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SetOrderBy(val *string) {
	_jsii_.Set(
		j,
		"orderBy",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SetOrderDir(val *string) {
	_jsii_.Set(
		j,
		"orderDir",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SetQ(val *string) {
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SetShowPresent(val interface{}) {
	_jsii_.Set(
		j,
		"showPresent",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) PutApmQuery(value *DashboardWidgetChangeDefinitionRequestApmQuery) {
	_jsii_.InvokeVoid(
		d,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) PutFormula(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putFormula",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) PutLogQuery(value *DashboardWidgetChangeDefinitionRequestLogQuery) {
	_jsii_.InvokeVoid(
		d,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) PutProcessQuery(value *DashboardWidgetChangeDefinitionRequestProcessQuery) {
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) PutQuery(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) PutRumQuery(value *DashboardWidgetChangeDefinitionRequestRumQuery) {
	_jsii_.InvokeVoid(
		d,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) PutSecurityQuery(value *DashboardWidgetChangeDefinitionRequestSecurityQuery) {
	_jsii_.InvokeVoid(
		d,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetChangeType() {
	_jsii_.InvokeVoid(
		d,
		"resetChangeType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetCompareTo() {
	_jsii_.InvokeVoid(
		d,
		"resetCompareTo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetFormula() {
	_jsii_.InvokeVoid(
		d,
		"resetFormula",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetIncreaseGood() {
	_jsii_.InvokeVoid(
		d,
		"resetIncreaseGood",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetOrderBy() {
	_jsii_.InvokeVoid(
		d,
		"resetOrderBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetOrderDir() {
	_jsii_.InvokeVoid(
		d,
		"resetOrderDir",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		d,
		"resetQ",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ResetShowPresent() {
	_jsii_.InvokeVoid(
		d,
		"resetShowPresent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

