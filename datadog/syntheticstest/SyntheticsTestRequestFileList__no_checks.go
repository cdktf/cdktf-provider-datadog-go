// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package syntheticstest

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsTestRequestFileList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SyntheticsTestRequestFileList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsTestRequestFileList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestRequestFileList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestRequestFileList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestRequestFileList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestRequestFileList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSyntheticsTestRequestFileListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

