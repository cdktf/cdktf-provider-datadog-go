// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package powerpack

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PowerpackWidgetTopologyMapDefinitionRequestList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PowerpackWidgetTopologyMapDefinitionRequestList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PowerpackWidgetTopologyMapDefinitionRequestList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetTopologyMapDefinitionRequestList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetTopologyMapDefinitionRequestList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetTopologyMapDefinitionRequestList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PowerpackWidgetTopologyMapDefinitionRequestList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPowerpackWidgetTopologyMapDefinitionRequestListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

