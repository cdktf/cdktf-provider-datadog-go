//go:build no_runtime_type_checking

package childorganization

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChildOrganizationSettingsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ChildOrganizationSettingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ChildOrganizationSettingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ChildOrganizationSettingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ChildOrganizationSettingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewChildOrganizationSettingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
