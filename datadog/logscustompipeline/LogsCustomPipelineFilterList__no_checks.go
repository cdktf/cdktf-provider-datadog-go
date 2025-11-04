// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package logscustompipeline

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogsCustomPipelineFilterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LogsCustomPipelineFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogsCustomPipelineFilterList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogsCustomPipelineFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogsCustomPipelineFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogsCustomPipelineFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogsCustomPipelineFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogsCustomPipelineFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

