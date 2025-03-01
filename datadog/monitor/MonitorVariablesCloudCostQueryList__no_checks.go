// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package monitor

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitorVariablesCloudCostQueryList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MonitorVariablesCloudCostQueryList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MonitorVariablesCloudCostQueryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MonitorVariablesCloudCostQueryList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MonitorVariablesCloudCostQueryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitorVariablesCloudCostQueryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MonitorVariablesCloudCostQueryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMonitorVariablesCloudCostQueryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

