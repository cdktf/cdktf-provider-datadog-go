//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DashboardWidgetQueryTableDefinitionRequestFormulaList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DashboardWidgetQueryTableDefinitionRequestFormulaList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetQueryTableDefinitionRequestFormulaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetQueryTableDefinitionRequestFormulaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetQueryTableDefinitionRequestFormulaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetQueryTableDefinitionRequestFormulaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDashboardWidgetQueryTableDefinitionRequestFormulaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
