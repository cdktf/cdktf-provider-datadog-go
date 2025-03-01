// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package actionconnection

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ActionConnection) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateImportFromParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateMoveToIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validatePutAwsParameters(value *ActionConnectionAws) error {
	return nil
}

func (a *jsiiProxy_ActionConnection) validatePutHttpParameters(value *ActionConnectionHttp) error {
	return nil
}

func validateActionConnection_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateActionConnection_IsConstructParameters(x interface{}) error {
	return nil
}

func validateActionConnection_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateActionConnection_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ActionConnection) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ActionConnection) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ActionConnection) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ActionConnection) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ActionConnection) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewActionConnectionParameters(scope constructs.Construct, id *string, config *ActionConnectionConfig) error {
	return nil
}

