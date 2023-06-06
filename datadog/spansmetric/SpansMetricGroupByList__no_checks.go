//go:build no_runtime_type_checking

package spansmetric

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SpansMetricGroupByList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SpansMetricGroupByList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SpansMetricGroupByList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SpansMetricGroupByList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SpansMetricGroupByList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SpansMetricGroupByList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSpansMetricGroupByListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

