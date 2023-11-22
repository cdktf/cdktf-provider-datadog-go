// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package monitor

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitorSchedulingOptionsCustomScheduleList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MonitorSchedulingOptionsCustomScheduleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MonitorSchedulingOptionsCustomScheduleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MonitorSchedulingOptionsCustomScheduleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitorSchedulingOptionsCustomScheduleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MonitorSchedulingOptionsCustomScheduleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMonitorSchedulingOptionsCustomScheduleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

