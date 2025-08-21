// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Dataset) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateMoveToIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (d *jsiiProxy_Dataset) validatePutProductFiltersParameters(value interface{}) error {
	return nil
}

func validateDataset_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDataset_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDataset_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDataset_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Dataset) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Dataset) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Dataset) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Dataset) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Dataset) validateSetPrincipalsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Dataset) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewDatasetParameters(scope constructs.Construct, id *string, config *DatasetConfig) error {
	return nil
}

