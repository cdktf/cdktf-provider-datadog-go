package organizationsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-datadog-go/datadog/v3/jsii"

	"github.com/cdktf/cdktf-provider-datadog-go/datadog/v3/organizationsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationSettingsSettingsOutputReference interface {
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
	InternalValue() *OrganizationSettingsSettings
	SetInternalValue(val *OrganizationSettingsSettings)
	PrivateWidgetShare() interface{}
	SetPrivateWidgetShare(val interface{})
	PrivateWidgetShareInput() interface{}
	Saml() OrganizationSettingsSettingsSamlOutputReference
	SamlAutocreateAccessRole() *string
	SetSamlAutocreateAccessRole(val *string)
	SamlAutocreateAccessRoleInput() *string
	SamlAutocreateUsersDomains() OrganizationSettingsSettingsSamlAutocreateUsersDomainsOutputReference
	SamlAutocreateUsersDomainsInput() *OrganizationSettingsSettingsSamlAutocreateUsersDomains
	SamlCanBeEnabled() cdktf.IResolvable
	SamlIdpEndpoint() *string
	SamlIdpInitiatedLogin() OrganizationSettingsSettingsSamlIdpInitiatedLoginOutputReference
	SamlIdpInitiatedLoginInput() *OrganizationSettingsSettingsSamlIdpInitiatedLogin
	SamlIdpMetadataUploaded() cdktf.IResolvable
	SamlInput() *OrganizationSettingsSettingsSaml
	SamlLoginUrl() *string
	SamlStrictMode() OrganizationSettingsSettingsSamlStrictModeOutputReference
	SamlStrictModeInput() *OrganizationSettingsSettingsSamlStrictMode
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
	PutSaml(value *OrganizationSettingsSettingsSaml)
	PutSamlAutocreateUsersDomains(value *OrganizationSettingsSettingsSamlAutocreateUsersDomains)
	PutSamlIdpInitiatedLogin(value *OrganizationSettingsSettingsSamlIdpInitiatedLogin)
	PutSamlStrictMode(value *OrganizationSettingsSettingsSamlStrictMode)
	ResetPrivateWidgetShare()
	ResetSamlAutocreateAccessRole()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrganizationSettingsSettingsOutputReference
type jsiiProxy_OrganizationSettingsSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) InternalValue() *OrganizationSettingsSettings {
	var returns *OrganizationSettingsSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) PrivateWidgetShare() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateWidgetShare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) PrivateWidgetShareInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateWidgetShareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) Saml() OrganizationSettingsSettingsSamlOutputReference {
	var returns OrganizationSettingsSettingsSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) SamlAutocreateAccessRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlAutocreateAccessRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) SamlAutocreateAccessRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlAutocreateAccessRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) SamlAutocreateUsersDomains() OrganizationSettingsSettingsSamlAutocreateUsersDomainsOutputReference {
	var returns OrganizationSettingsSettingsSamlAutocreateUsersDomainsOutputReference
	_jsii_.Get(
		j,
		"samlAutocreateUsersDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) SamlAutocreateUsersDomainsInput() *OrganizationSettingsSettingsSamlAutocreateUsersDomains {
	var returns *OrganizationSettingsSettingsSamlAutocreateUsersDomains
	_jsii_.Get(
		j,
		"samlAutocreateUsersDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) SamlCanBeEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"samlCanBeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) SamlIdpEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlIdpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) SamlIdpInitiatedLogin() OrganizationSettingsSettingsSamlIdpInitiatedLoginOutputReference {
	var returns OrganizationSettingsSettingsSamlIdpInitiatedLoginOutputReference
	_jsii_.Get(
		j,
		"samlIdpInitiatedLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) SamlIdpInitiatedLoginInput() *OrganizationSettingsSettingsSamlIdpInitiatedLogin {
	var returns *OrganizationSettingsSettingsSamlIdpInitiatedLogin
	_jsii_.Get(
		j,
		"samlIdpInitiatedLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) SamlIdpMetadataUploaded() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"samlIdpMetadataUploaded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) SamlInput() *OrganizationSettingsSettingsSaml {
	var returns *OrganizationSettingsSettingsSaml
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) SamlLoginUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlLoginUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) SamlStrictMode() OrganizationSettingsSettingsSamlStrictModeOutputReference {
	var returns OrganizationSettingsSettingsSamlStrictModeOutputReference
	_jsii_.Get(
		j,
		"samlStrictMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) SamlStrictModeInput() *OrganizationSettingsSettingsSamlStrictMode {
	var returns *OrganizationSettingsSettingsSamlStrictMode
	_jsii_.Get(
		j,
		"samlStrictModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOrganizationSettingsSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrganizationSettingsSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewOrganizationSettingsSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrganizationSettingsSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-datadog.organizationSettings.OrganizationSettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrganizationSettingsSettingsOutputReference_Override(o OrganizationSettingsSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-datadog.organizationSettings.OrganizationSettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference)SetInternalValue(val *OrganizationSettingsSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference)SetPrivateWidgetShare(val interface{}) {
	if err := j.validateSetPrivateWidgetShareParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateWidgetShare",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference)SetSamlAutocreateAccessRole(val *string) {
	if err := j.validateSetSamlAutocreateAccessRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlAutocreateAccessRole",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettingsSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) PutSaml(value *OrganizationSettingsSettingsSaml) {
	if err := o.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSaml",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) PutSamlAutocreateUsersDomains(value *OrganizationSettingsSettingsSamlAutocreateUsersDomains) {
	if err := o.validatePutSamlAutocreateUsersDomainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSamlAutocreateUsersDomains",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) PutSamlIdpInitiatedLogin(value *OrganizationSettingsSettingsSamlIdpInitiatedLogin) {
	if err := o.validatePutSamlIdpInitiatedLoginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSamlIdpInitiatedLogin",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) PutSamlStrictMode(value *OrganizationSettingsSettingsSamlStrictMode) {
	if err := o.validatePutSamlStrictModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSamlStrictMode",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) ResetPrivateWidgetShare() {
	_jsii_.InvokeVoid(
		o,
		"resetPrivateWidgetShare",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) ResetSamlAutocreateAccessRole() {
	_jsii_.InvokeVoid(
		o,
		"resetSamlAutocreateAccessRole",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettingsSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

