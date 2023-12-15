// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package powerpack

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PowerpackWidgetList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PowerpackWidgetList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPowerpackWidgetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

