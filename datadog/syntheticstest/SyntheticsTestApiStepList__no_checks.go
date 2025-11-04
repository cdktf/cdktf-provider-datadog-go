// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package syntheticstest

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsTestApiStepList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SyntheticsTestApiStepList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsTestApiStepList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestApiStepList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestApiStepList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestApiStepList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestApiStepList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSyntheticsTestApiStepListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

