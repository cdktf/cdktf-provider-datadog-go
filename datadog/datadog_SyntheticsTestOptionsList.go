// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type SyntheticsTestOptionsList struct {
	// How often the test should run (in seconds).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#tick_every SyntheticsTest#tick_every}
	TickEvery *float64 `field:"required" json:"tickEvery" yaml:"tickEvery"`
	// For SSL test, whether or not the test should allow self signed certificates.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#accept_self_signed SyntheticsTest#accept_self_signed}
	AcceptSelfSigned interface{} `field:"optional" json:"acceptSelfSigned" yaml:"acceptSelfSigned"`
	// Allows loading insecure content for an HTTP test.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#allow_insecure SyntheticsTest#allow_insecure}
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// For SSL test, whether or not the test should fail on revoked certificate in stapled OCSP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#check_certificate_revocation SyntheticsTest#check_certificate_revocation}
	CheckCertificateRevocation interface{} `field:"optional" json:"checkCertificateRevocation" yaml:"checkCertificateRevocation"`
	// ci block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#ci SyntheticsTest#ci}
	Ci *SyntheticsTestOptionsListCi `field:"optional" json:"ci" yaml:"ci"`
	// Determines whether or not the API HTTP test should follow redirects.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#follow_redirects SyntheticsTest#follow_redirects}
	FollowRedirects interface{} `field:"optional" json:"followRedirects" yaml:"followRedirects"`
	// Minimum amount of time in failure required to trigger an alert. Default is `0`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#min_failure_duration SyntheticsTest#min_failure_duration}
	MinFailureDuration *float64 `field:"optional" json:"minFailureDuration" yaml:"minFailureDuration"`
	// Minimum number of locations in failure required to trigger an alert. Default is `1`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#min_location_failed SyntheticsTest#min_location_failed}
	MinLocationFailed *float64 `field:"optional" json:"minLocationFailed" yaml:"minLocationFailed"`
	// The monitor name is used for the alert title as well as for all monitor dashboard widgets and SLOs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#monitor_name SyntheticsTest#monitor_name}
	MonitorName *string `field:"optional" json:"monitorName" yaml:"monitorName"`
	// monitor_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#monitor_options SyntheticsTest#monitor_options}
	MonitorOptions *SyntheticsTestOptionsListMonitorOptions `field:"optional" json:"monitorOptions" yaml:"monitorOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#monitor_priority SyntheticsTest#monitor_priority}.
	MonitorPriority *float64 `field:"optional" json:"monitorPriority" yaml:"monitorPriority"`
	// Prevents saving screenshots of the steps.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#no_screenshot SyntheticsTest#no_screenshot}
	NoScreenshot interface{} `field:"optional" json:"noScreenshot" yaml:"noScreenshot"`
	// A list of role identifiers pulled from the Roles API to restrict read and write access.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#restricted_roles SyntheticsTest#restricted_roles}
	RestrictedRoles *[]*string `field:"optional" json:"restrictedRoles" yaml:"restrictedRoles"`
	// retry block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#retry SyntheticsTest#retry}
	Retry *SyntheticsTestOptionsListRetry `field:"optional" json:"retry" yaml:"retry"`
	// rum_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#rum_settings SyntheticsTest#rum_settings}
	RumSettings *SyntheticsTestOptionsListRumSettings `field:"optional" json:"rumSettings" yaml:"rumSettings"`
}

