//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DashboardWidgetDistributionDefinitionRequestSecurityQueryGroupByList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DashboardWidgetDistributionDefinitionRequestSecurityQueryGroupByList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetDistributionDefinitionRequestSecurityQueryGroupByList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetDistributionDefinitionRequestSecurityQueryGroupByList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetDistributionDefinitionRequestSecurityQueryGroupByList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetDistributionDefinitionRequestSecurityQueryGroupByList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDashboardWidgetDistributionDefinitionRequestSecurityQueryGroupByListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
