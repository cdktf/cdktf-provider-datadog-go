// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package childorganization

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChildOrganizationSettingsSamlList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ChildOrganizationSettingsSamlList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ChildOrganizationSettingsSamlList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ChildOrganizationSettingsSamlList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ChildOrganizationSettingsSamlList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ChildOrganizationSettingsSamlList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewChildOrganizationSettingsSamlListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

