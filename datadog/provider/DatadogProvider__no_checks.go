// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatadogProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DatadogProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateDatadogProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatadogProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDatadogProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewDatadogProviderParameters(scope constructs.Construct, id *string, config *DatadogProviderConfig) error {
	return nil
}

