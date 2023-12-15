// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package powerpack

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PowerpackWidgetToplistDefinitionRequestFormulaList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PowerpackWidgetToplistDefinitionRequestFormulaList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetToplistDefinitionRequestFormulaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetToplistDefinitionRequestFormulaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetToplistDefinitionRequestFormulaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetToplistDefinitionRequestFormulaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPowerpackWidgetToplistDefinitionRequestFormulaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

