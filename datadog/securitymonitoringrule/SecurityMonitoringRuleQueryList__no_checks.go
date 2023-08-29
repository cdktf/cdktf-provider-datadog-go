// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package securitymonitoringrule

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityMonitoringRuleQueryList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecurityMonitoringRuleQueryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecurityMonitoringRuleQueryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecurityMonitoringRuleQueryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecurityMonitoringRuleQueryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecurityMonitoringRuleQueryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecurityMonitoringRuleQueryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

