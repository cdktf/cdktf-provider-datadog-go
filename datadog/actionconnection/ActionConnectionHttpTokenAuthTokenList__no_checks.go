// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package actionconnection

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ActionConnectionHttpTokenAuthTokenList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthTokenList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ActionConnectionHttpTokenAuthTokenList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthTokenList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthTokenList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthTokenList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ActionConnectionHttpTokenAuthTokenList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewActionConnectionHttpTokenAuthTokenListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

