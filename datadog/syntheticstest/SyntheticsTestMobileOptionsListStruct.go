// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestMobileOptionsListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#device_ids SyntheticsTest#device_ids}.
	DeviceIds *[]*string `field:"required" json:"deviceIds" yaml:"deviceIds"`
	// mobile_application block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#mobile_application SyntheticsTest#mobile_application}
	MobileApplication *SyntheticsTestMobileOptionsListMobileApplication `field:"required" json:"mobileApplication" yaml:"mobileApplication"`
	// How often the test should run (in seconds). Valid range is `300-604800` for mobile tests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#tick_every SyntheticsTest#tick_every}
	TickEvery *float64 `field:"required" json:"tickEvery" yaml:"tickEvery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#allow_application_crash SyntheticsTest#allow_application_crash}.
	AllowApplicationCrash interface{} `field:"optional" json:"allowApplicationCrash" yaml:"allowApplicationCrash"`
	// bindings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#bindings SyntheticsTest#bindings}
	Bindings interface{} `field:"optional" json:"bindings" yaml:"bindings"`
	// ci block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#ci SyntheticsTest#ci}
	Ci *SyntheticsTestMobileOptionsListCi `field:"optional" json:"ci" yaml:"ci"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#default_step_timeout SyntheticsTest#default_step_timeout}.
	DefaultStepTimeout *float64 `field:"optional" json:"defaultStepTimeout" yaml:"defaultStepTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#disable_auto_accept_alert SyntheticsTest#disable_auto_accept_alert}.
	DisableAutoAcceptAlert interface{} `field:"optional" json:"disableAutoAcceptAlert" yaml:"disableAutoAcceptAlert"`
	// Minimum amount of time in failure required to trigger an alert (in seconds). Default is `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#min_failure_duration SyntheticsTest#min_failure_duration}
	MinFailureDuration *float64 `field:"optional" json:"minFailureDuration" yaml:"minFailureDuration"`
	// The monitor name is used for the alert title as well as for all monitor dashboard widgets and SLOs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#monitor_name SyntheticsTest#monitor_name}
	MonitorName *string `field:"optional" json:"monitorName" yaml:"monitorName"`
	// monitor_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#monitor_options SyntheticsTest#monitor_options}
	MonitorOptions *SyntheticsTestMobileOptionsListMonitorOptions `field:"optional" json:"monitorOptions" yaml:"monitorOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#monitor_priority SyntheticsTest#monitor_priority}.
	MonitorPriority *float64 `field:"optional" json:"monitorPriority" yaml:"monitorPriority"`
	// Prevents saving screenshots of the steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#no_screenshot SyntheticsTest#no_screenshot}
	NoScreenshot interface{} `field:"optional" json:"noScreenshot" yaml:"noScreenshot"`
	// A list of role identifiers pulled from the Roles API to restrict read and write access.
	//
	// **Deprecated.** This field is no longer supported by the Datadog API. Please use `datadog_restriction_policy` instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#restricted_roles SyntheticsTest#restricted_roles}
	RestrictedRoles *[]*string `field:"optional" json:"restrictedRoles" yaml:"restrictedRoles"`
	// retry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#retry SyntheticsTest#retry}
	Retry *SyntheticsTestMobileOptionsListRetry `field:"optional" json:"retry" yaml:"retry"`
	// scheduling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#scheduling SyntheticsTest#scheduling}
	Scheduling *SyntheticsTestMobileOptionsListScheduling `field:"optional" json:"scheduling" yaml:"scheduling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#verbosity SyntheticsTest#verbosity}.
	Verbosity *float64 `field:"optional" json:"verbosity" yaml:"verbosity"`
}

