// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package costbudget

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CostBudgetEntriesTagFiltersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CostBudgetEntriesTagFiltersList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CostBudgetEntriesTagFiltersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CostBudgetEntriesTagFiltersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CostBudgetEntriesTagFiltersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CostBudgetEntriesTagFiltersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CostBudgetEntriesTagFiltersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCostBudgetEntriesTagFiltersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

