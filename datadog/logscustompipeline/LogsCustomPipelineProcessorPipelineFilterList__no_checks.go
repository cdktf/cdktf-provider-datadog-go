// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package logscustompipeline

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogsCustomPipelineProcessorPipelineFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogsCustomPipelineProcessorPipelineFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogsCustomPipelineProcessorPipelineFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

