//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionRequestLogQueryGroupByList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DashboardWidgetHeatmapDefinitionRequestLogQueryGroupByList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionRequestLogQueryGroupByList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionRequestLogQueryGroupByList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionRequestLogQueryGroupByList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetHeatmapDefinitionRequestLogQueryGroupByList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDashboardWidgetHeatmapDefinitionRequestLogQueryGroupByListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
