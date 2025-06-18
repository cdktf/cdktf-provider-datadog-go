// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package costbudget

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CostBudgetEntriesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CostBudgetEntriesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CostBudgetEntriesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CostBudgetEntriesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CostBudgetEntriesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CostBudgetEntriesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CostBudgetEntriesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCostBudgetEntriesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

