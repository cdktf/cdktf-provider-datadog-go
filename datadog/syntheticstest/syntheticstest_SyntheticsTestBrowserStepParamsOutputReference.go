package syntheticstest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v5/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v5/syntheticstest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsTestBrowserStepParamsOutputReference interface {
	cdktf.ComplexObject
	Attribute() *string
	SetAttribute(val *string)
	AttributeInput() *string
	Check() *string
	SetCheck(val *string)
	CheckInput() *string
	ClickType() *string
	SetClickType(val *string)
	ClickTypeInput() *string
	Code() *string
	SetCode(val *string)
	CodeInput() *string
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
	Delay() *float64
	SetDelay(val *float64)
	DelayInput() *float64
	Element() *string
	SetElement(val *string)
	ElementInput() *string
	ElementUserLocator() SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference
	ElementUserLocatorInput() *SyntheticsTestBrowserStepParamsElementUserLocator
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	File() *string
	SetFile(val *string)
	FileInput() *string
	Files() *string
	SetFiles(val *string)
	FilesInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SyntheticsTestBrowserStepParams
	SetInternalValue(val *SyntheticsTestBrowserStepParams)
	Modifiers() *[]*string
	SetModifiers(val *[]*string)
	ModifiersInput() *[]*string
	PlayingTabId() *string
	SetPlayingTabId(val *string)
	PlayingTabIdInput() *string
	Request() *string
	SetRequest(val *string)
	RequestInput() *string
	SubtestPublicId() *string
	SetSubtestPublicId(val *string)
	SubtestPublicIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *string
	SetValue(val *string)
	ValueInput() *string
	Variable() SyntheticsTestBrowserStepParamsVariableOutputReference
	VariableInput() *SyntheticsTestBrowserStepParamsVariable
	WithClick() interface{}
	SetWithClick(val interface{})
	WithClickInput() interface{}
	X() *float64
	SetX(val *float64)
	XInput() *float64
	Y() *float64
	SetY(val *float64)
	YInput() *float64
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
	PutElementUserLocator(value *SyntheticsTestBrowserStepParamsElementUserLocator)
	PutVariable(value *SyntheticsTestBrowserStepParamsVariable)
	ResetAttribute()
	ResetCheck()
	ResetClickType()
	ResetCode()
	ResetDelay()
	ResetElement()
	ResetElementUserLocator()
	ResetEmail()
	ResetFile()
	ResetFiles()
	ResetModifiers()
	ResetPlayingTabId()
	ResetRequest()
	ResetSubtestPublicId()
	ResetValue()
	ResetVariable()
	ResetWithClick()
	ResetX()
	ResetY()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SyntheticsTestBrowserStepParamsOutputReference
type jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Attribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) AttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Check() *string {
	var returns *string
	_jsii_.Get(
		j,
		"check",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) CheckInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ClickType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clickType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ClickTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clickTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Code() *string {
	var returns *string
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) CodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Delay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) DelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Element() *string {
	var returns *string
	_jsii_.Get(
		j,
		"element",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ElementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ElementUserLocator() SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference {
	var returns SyntheticsTestBrowserStepParamsElementUserLocatorOutputReference
	_jsii_.Get(
		j,
		"elementUserLocator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ElementUserLocatorInput() *SyntheticsTestBrowserStepParamsElementUserLocator {
	var returns *SyntheticsTestBrowserStepParamsElementUserLocator
	_jsii_.Get(
		j,
		"elementUserLocatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) File() *string {
	var returns *string
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) FileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Files() *string {
	var returns *string
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) FilesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) InternalValue() *SyntheticsTestBrowserStepParams {
	var returns *SyntheticsTestBrowserStepParams
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Modifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"modifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ModifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"modifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) PlayingTabId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playingTabId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) PlayingTabIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"playingTabIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Request() *string {
	var returns *string
	_jsii_.Get(
		j,
		"request",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) RequestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) SubtestPublicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtestPublicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) SubtestPublicIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtestPublicIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Variable() SyntheticsTestBrowserStepParamsVariableOutputReference {
	var returns SyntheticsTestBrowserStepParamsVariableOutputReference
	_jsii_.Get(
		j,
		"variable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) VariableInput() *SyntheticsTestBrowserStepParamsVariable {
	var returns *SyntheticsTestBrowserStepParamsVariable
	_jsii_.Get(
		j,
		"variableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) WithClick() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withClick",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) WithClickInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withClickInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) XInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"xInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) YInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yInput",
		&returns,
	)
	return returns
}


func NewSyntheticsTestBrowserStepParamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SyntheticsTestBrowserStepParamsOutputReference {
	_init_.Initialize()

	if err := validateNewSyntheticsTestBrowserStepParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestBrowserStepParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSyntheticsTestBrowserStepParamsOutputReference_Override(s SyntheticsTestBrowserStepParamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.syntheticsTest.SyntheticsTestBrowserStepParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetAttribute(val *string) {
	if err := j.validateSetAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetCheck(val *string) {
	if err := j.validateSetCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"check",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetClickType(val *string) {
	if err := j.validateSetClickTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clickType",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetCode(val *string) {
	if err := j.validateSetCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetDelay(val *float64) {
	if err := j.validateSetDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delay",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetElement(val *string) {
	if err := j.validateSetElementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"element",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetFile(val *string) {
	if err := j.validateSetFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"file",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetFiles(val *string) {
	if err := j.validateSetFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"files",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetInternalValue(val *SyntheticsTestBrowserStepParams) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetModifiers(val *[]*string) {
	if err := j.validateSetModifiersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modifiers",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetPlayingTabId(val *string) {
	if err := j.validateSetPlayingTabIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"playingTabId",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetRequest(val *string) {
	if err := j.validateSetRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"request",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetSubtestPublicId(val *string) {
	if err := j.validateSetSubtestPublicIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subtestPublicId",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetWithClick(val interface{}) {
	if err := j.validateSetWithClickParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withClick",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetX(val *float64) {
	if err := j.validateSetXParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference)SetY(val *float64) {
	if err := j.validateSetYParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) PutElementUserLocator(value *SyntheticsTestBrowserStepParamsElementUserLocator) {
	if err := s.validatePutElementUserLocatorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putElementUserLocator",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) PutVariable(value *SyntheticsTestBrowserStepParamsVariable) {
	if err := s.validatePutVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVariable",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetAttribute() {
	_jsii_.InvokeVoid(
		s,
		"resetAttribute",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetCheck() {
	_jsii_.InvokeVoid(
		s,
		"resetCheck",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetClickType() {
	_jsii_.InvokeVoid(
		s,
		"resetClickType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetCode() {
	_jsii_.InvokeVoid(
		s,
		"resetCode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetDelay() {
	_jsii_.InvokeVoid(
		s,
		"resetDelay",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetElement() {
	_jsii_.InvokeVoid(
		s,
		"resetElement",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetElementUserLocator() {
	_jsii_.InvokeVoid(
		s,
		"resetElementUserLocator",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		s,
		"resetEmail",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		s,
		"resetFile",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetFiles() {
	_jsii_.InvokeVoid(
		s,
		"resetFiles",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetModifiers() {
	_jsii_.InvokeVoid(
		s,
		"resetModifiers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetPlayingTabId() {
	_jsii_.InvokeVoid(
		s,
		"resetPlayingTabId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetRequest() {
	_jsii_.InvokeVoid(
		s,
		"resetRequest",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetSubtestPublicId() {
	_jsii_.InvokeVoid(
		s,
		"resetSubtestPublicId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		s,
		"resetValue",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetVariable() {
	_jsii_.InvokeVoid(
		s,
		"resetVariable",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetWithClick() {
	_jsii_.InvokeVoid(
		s,
		"resetWithClick",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetX() {
	_jsii_.InvokeVoid(
		s,
		"resetX",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ResetY() {
	_jsii_.InvokeVoid(
		s,
		"resetY",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsTestBrowserStepParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

