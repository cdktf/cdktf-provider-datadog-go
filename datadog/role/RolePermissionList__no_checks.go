// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package role

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RolePermissionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RolePermissionList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RolePermissionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RolePermissionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RolePermissionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RolePermissionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RolePermissionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRolePermissionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

