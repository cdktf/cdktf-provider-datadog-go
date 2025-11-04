// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dashboard

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DashboardTemplateVariableList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DashboardTemplateVariableList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DashboardTemplateVariableList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DashboardTemplateVariableList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DashboardTemplateVariableList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DashboardTemplateVariableList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DashboardTemplateVariableList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDashboardTemplateVariableListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

