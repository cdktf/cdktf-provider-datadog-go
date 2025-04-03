// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appbuilderapp

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppBuilderApp) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateImportFromParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateMoveToIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_AppBuilderApp) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateAppBuilderApp_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateAppBuilderApp_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAppBuilderApp_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateAppBuilderApp_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AppBuilderApp) validateSetActionQueryNamesToConnectionIdsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_AppBuilderApp) validateSetAppJsonParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppBuilderApp) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppBuilderApp) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppBuilderApp) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppBuilderApp) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_AppBuilderApp) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppBuilderApp) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_AppBuilderApp) validateSetPublishedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppBuilderApp) validateSetRootInstanceNameParameters(val *string) error {
	return nil
}

func validateNewAppBuilderAppParameters(scope constructs.Construct, id *string, config *AppBuilderAppConfig) error {
	return nil
}

