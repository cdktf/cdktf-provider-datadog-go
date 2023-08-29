// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package servicelevelobjective

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceLevelObjectiveThresholdsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceLevelObjectiveThresholdsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceLevelObjectiveThresholdsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceLevelObjectiveThresholdsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceLevelObjectiveThresholdsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceLevelObjectiveThresholdsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceLevelObjectiveThresholdsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

