// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatasetProductFiltersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatasetProductFiltersList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatasetProductFiltersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatasetProductFiltersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatasetProductFiltersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatasetProductFiltersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatasetProductFiltersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatasetProductFiltersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

