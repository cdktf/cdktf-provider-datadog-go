//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dashboard

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DashboardWidgetTimeseriesDefinitionRequestFormulaList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DashboardWidgetTimeseriesDefinitionRequestFormulaList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetTimeseriesDefinitionRequestFormulaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetTimeseriesDefinitionRequestFormulaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetTimeseriesDefinitionRequestFormulaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetTimeseriesDefinitionRequestFormulaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDashboardWidgetTimeseriesDefinitionRequestFormulaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

