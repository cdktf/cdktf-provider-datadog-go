// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package logscustompipeline

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogsCustomPipelineProcessorList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LogsCustomPipelineProcessorList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LogsCustomPipelineProcessorList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LogsCustomPipelineProcessorList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LogsCustomPipelineProcessorList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LogsCustomPipelineProcessorList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LogsCustomPipelineProcessorList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLogsCustomPipelineProcessorListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

