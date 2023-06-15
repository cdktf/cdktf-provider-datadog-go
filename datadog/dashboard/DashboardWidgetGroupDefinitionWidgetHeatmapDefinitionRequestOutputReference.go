package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v7/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v7/dashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference interface {
	cdktf.ComplexObject
	ApmQuery() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryOutputReference
	ApmQueryInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQuery
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQueryOutputReference
	LogQueryInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQuery
	ProcessQuery() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQueryOutputReference
	ProcessQueryInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	RumQuery() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQueryOutputReference
	RumQueryInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQuery
	SecurityQuery() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQueryOutputReference
	SecurityQueryInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQuery
	Style() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyleOutputReference
	StyleInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyle
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
	PutApmQuery(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQuery)
	PutLogQuery(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQuery)
	PutProcessQuery(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQuery)
	PutRumQuery(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQuery)
	PutSecurityQuery(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQuery)
	PutStyle(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyle)
	ResetApmQuery()
	ResetLogQuery()
	ResetProcessQuery()
	ResetQ()
	ResetRumQuery()
	ResetSecurityQuery()
	ResetStyle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference
type jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ApmQuery() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ApmQueryInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) LogQuery() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) LogQueryInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ProcessQuery() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ProcessQueryInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) RumQuery() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) RumQueryInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) SecurityQuery() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQueryOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) SecurityQueryInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQuery {
	var returns *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) Style() DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyleOutputReference {
	var returns DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) StyleInput() *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyle {
	var returns *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference_Override(d DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.dashboard.DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutApmQuery(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQuery) {
	if err := d.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutLogQuery(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQuery) {
	if err := d.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutProcessQuery(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutRumQuery(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQuery) {
	if err := d.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutSecurityQuery(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQuery) {
	if err := d.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutStyle(value *DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyle) {
	if err := d.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStyle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		d,
		"resetQ",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		d,
		"resetStyle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

