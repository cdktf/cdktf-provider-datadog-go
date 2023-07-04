//go:build no_runtime_type_checking

package restrictionpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RestrictionPolicyBindingsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RestrictionPolicyBindingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RestrictionPolicyBindingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RestrictionPolicyBindingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RestrictionPolicyBindingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RestrictionPolicyBindingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRestrictionPolicyBindingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

