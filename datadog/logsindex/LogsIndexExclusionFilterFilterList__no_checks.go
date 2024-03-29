// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package logsindex

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogsIndexExclusionFilterFilterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LogsIndexExclusionFilterFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogsIndexExclusionFilterFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogsIndexExclusionFilterFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogsIndexExclusionFilterFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogsIndexExclusionFilterFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogsIndexExclusionFilterFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogsIndexExclusionFilterFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

