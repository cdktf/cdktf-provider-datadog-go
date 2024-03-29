// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package logsindex

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogsIndexExclusionFilterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LogsIndexExclusionFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogsIndexExclusionFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogsIndexExclusionFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogsIndexExclusionFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogsIndexExclusionFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogsIndexExclusionFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogsIndexExclusionFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

