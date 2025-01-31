// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package monitor

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Monitor) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateImportFromParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateMoveToIdParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (m *jsiiProxy_Monitor) validatePutMonitorThresholdsParameters(value *MonitorMonitorThresholds) error {
	return nil
}

func (m *jsiiProxy_Monitor) validatePutMonitorThresholdWindowsParameters(value *MonitorMonitorThresholdWindows) error {
	return nil
}

func (m *jsiiProxy_Monitor) validatePutSchedulingOptionsParameters(value *MonitorSchedulingOptions) error {
	return nil
}

func (m *jsiiProxy_Monitor) validatePutVariablesParameters(value *MonitorVariables) error {
	return nil
}

func validateMonitor_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateMonitor_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMonitor_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateMonitor_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetEnableLogsSampleParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetEnableSamplesParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetEscalationMessageParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetEvaluationDelayParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetForceDeleteParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetGroupbySimpleMonitorParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetGroupRetentionDurationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetIncludeTagsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetLockedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetMessageParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetNewGroupDelayParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetNewHostDelayParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetNoDataTimeframeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetNotificationPresetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetNotifyAuditParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetNotifyByParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetNotifyNoDataParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetOnMissingDataParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetPriorityParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetQueryParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetRenotifyIntervalParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetRenotifyOccurrencesParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetRenotifyStatusesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetRequireFullWindowParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetRestrictedRolesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetTagsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetTimeoutHParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Monitor) validateSetValidateParameters(val interface{}) error {
	return nil
}

func validateNewMonitorParameters(scope constructs.Construct, id *string, config *MonitorConfig) error {
	return nil
}

