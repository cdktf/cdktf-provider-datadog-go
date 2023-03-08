//go:build no_runtime_type_checking

package monitor

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitorSchedulingOptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MonitorSchedulingOptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MonitorSchedulingOptionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MonitorSchedulingOptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitorSchedulingOptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MonitorSchedulingOptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMonitorSchedulingOptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

