//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataDatadogMonitorMonitorThresholdsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataDatadogMonitorMonitorThresholdsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataDatadogMonitorMonitorThresholdsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataDatadogMonitorMonitorThresholdsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataDatadogMonitorMonitorThresholdsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataDatadogMonitorMonitorThresholdsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
