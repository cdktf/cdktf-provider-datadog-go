// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package monitor

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMonitorSchedulingOptionsEvaluationWindowListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

