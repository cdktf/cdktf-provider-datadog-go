//go:build no_runtime_type_checking

package dashboard

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDashboardWidgetChangeDefinitionRequestListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

