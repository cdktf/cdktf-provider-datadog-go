// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package powerpack

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetHeatmapDefinitionRequestList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPowerpackWidgetHeatmapDefinitionRequestListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

