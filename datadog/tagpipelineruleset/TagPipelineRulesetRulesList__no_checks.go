// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package tagpipelineruleset

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TagPipelineRulesetRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TagPipelineRulesetRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TagPipelineRulesetRulesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TagPipelineRulesetRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TagPipelineRulesetRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TagPipelineRulesetRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TagPipelineRulesetRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTagPipelineRulesetRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

