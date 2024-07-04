// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package syntheticstest

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsTestApiStepRequestFileList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestFileList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsTestApiStepRequestFileList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestFileList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestFileList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestFileList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestApiStepRequestFileList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSyntheticsTestApiStepRequestFileListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

