// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package actionconnection

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ActionConnectionHttpTokenAuthUrlParameterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthUrlParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthUrlParameterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthUrlParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthUrlParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthUrlParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthUrlParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewActionConnectionHttpTokenAuthUrlParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

