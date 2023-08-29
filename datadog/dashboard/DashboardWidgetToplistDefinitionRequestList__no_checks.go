// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dashboard

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DashboardWidgetToplistDefinitionRequestList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DashboardWidgetToplistDefinitionRequestList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetToplistDefinitionRequestList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetToplistDefinitionRequestList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetToplistDefinitionRequestList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetToplistDefinitionRequestList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDashboardWidgetToplistDefinitionRequestListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

