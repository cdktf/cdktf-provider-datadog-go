//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChildOrganizationApiKeyList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ChildOrganizationApiKeyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ChildOrganizationApiKeyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ChildOrganizationApiKeyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ChildOrganizationApiKeyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewChildOrganizationApiKeyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

