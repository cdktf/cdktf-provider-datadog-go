//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

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

func (j *jsiiProxy_DatadogProvider) validateSetHttpClientRetryEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatadogProvider) validateSetValidateParameters(val interface{}) error {
	return nil
}

func validateNewDatadogProviderParameters(scope constructs.Construct, id *string, config *DatadogProviderConfig) error {
	return nil
}
