// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/jsii"

	"github.com/hashicorp/cdktf-provider-datadog-go/datadog/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardWidgetHeatmapDefinitionYaxisOutputReference interface {
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
	IncludeZero() interface{}
	SetIncludeZero(val interface{})
	IncludeZeroInput() interface{}
	InternalValue() *DashboardWidgetHeatmapDefinitionYaxis
	SetInternalValue(val *DashboardWidgetHeatmapDefinitionYaxis)
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	Max() *string
	SetMax(val *string)
	MaxInput() *string
	Min() *string
	SetMin(val *string)
	MinInput() *string
	Scale() *string
	SetScale(val *string)
	ScaleInput() *string
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
	ResetIncludeZero()
	ResetLabel()
	ResetMax()
	ResetMin()
	ResetScale()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetHeatmapDefinitionYaxisOutputReference
type jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) IncludeZero() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeZero",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) IncludeZeroInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeZeroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) InternalValue() *DashboardWidgetHeatmapDefinitionYaxis {
	var returns *DashboardWidgetHeatmapDefinitionYaxis
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) Max() *string {
	var returns *string
	_jsii_.Get(
		j,
		"max",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) MaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) Min() *string {
	var returns *string
	_jsii_.Get(
		j,
		"min",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) MinInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) Scale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) ScaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetHeatmapDefinitionYaxisOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DashboardWidgetHeatmapDefinitionYaxisOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.DashboardWidgetHeatmapDefinitionYaxisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetHeatmapDefinitionYaxisOutputReference_Override(d DashboardWidgetHeatmapDefinitionYaxisOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.DashboardWidgetHeatmapDefinitionYaxisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) SetIncludeZero(val interface{}) {
	_jsii_.Set(
		j,
		"includeZero",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) SetInternalValue(val *DashboardWidgetHeatmapDefinitionYaxis) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) SetLabel(val *string) {
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) SetMax(val *string) {
	_jsii_.Set(
		j,
		"max",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) SetMin(val *string) {
	_jsii_.Set(
		j,
		"min",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) SetScale(val *string) {
	_jsii_.Set(
		j,
		"scale",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) ResetIncludeZero() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeZero",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		d,
		"resetLabel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) ResetMax() {
	_jsii_.InvokeVoid(
		d,
		"resetMax",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) ResetMin() {
	_jsii_.InvokeVoid(
		d,
		"resetMin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) ResetScale() {
	_jsii_.InvokeVoid(
		d,
		"resetScale",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionYaxisOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

