//go:build no_runtime_type_checking

package syntheticstest

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsTestBrowserVariableList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsTestBrowserVariableList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestBrowserVariableList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestBrowserVariableList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestBrowserVariableList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestBrowserVariableList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSyntheticsTestBrowserVariableListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

