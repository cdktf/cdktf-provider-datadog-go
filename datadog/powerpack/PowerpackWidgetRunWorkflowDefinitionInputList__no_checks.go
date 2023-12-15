// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package powerpack

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PowerpackWidgetRunWorkflowDefinitionInputList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PowerpackWidgetRunWorkflowDefinitionInputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetRunWorkflowDefinitionInputList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetRunWorkflowDefinitionInputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetRunWorkflowDefinitionInputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetRunWorkflowDefinitionInputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPowerpackWidgetRunWorkflowDefinitionInputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

