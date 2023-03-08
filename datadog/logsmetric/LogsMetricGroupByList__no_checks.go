//go:build no_runtime_type_checking

package logsmetric

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogsMetricGroupByList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogsMetricGroupByList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogsMetricGroupByList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogsMetricGroupByList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogsMetricGroupByList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogsMetricGroupByList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogsMetricGroupByListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

