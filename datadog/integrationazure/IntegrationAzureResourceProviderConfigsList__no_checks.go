// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package integrationazure

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegrationAzureResourceProviderConfigsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationAzureResourceProviderConfigsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IntegrationAzureResourceProviderConfigsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IntegrationAzureResourceProviderConfigsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationAzureResourceProviderConfigsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IntegrationAzureResourceProviderConfigsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IntegrationAzureResourceProviderConfigsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIntegrationAzureResourceProviderConfigsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

