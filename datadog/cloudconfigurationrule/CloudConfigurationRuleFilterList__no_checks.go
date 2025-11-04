// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudconfigurationrule

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudConfigurationRuleFilterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudConfigurationRuleFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudConfigurationRuleFilterList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudConfigurationRuleFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudConfigurationRuleFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudConfigurationRuleFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudConfigurationRuleFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudConfigurationRuleFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

