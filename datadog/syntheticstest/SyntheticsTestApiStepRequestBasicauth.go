// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestApiStepRequestBasicauth struct {
	// Access key for `SIGV4` authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#access_key SyntheticsTest#access_key}
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Access token url for `oauth-client` or `oauth-rop` authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#access_token_url SyntheticsTest#access_token_url}
	AccessTokenUrl *string `field:"optional" json:"accessTokenUrl" yaml:"accessTokenUrl"`
	// Audience for `oauth-client` or `oauth-rop` authentication. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#audience SyntheticsTest#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Client ID for `oauth-client` or `oauth-rop` authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#client_id SyntheticsTest#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Client secret for `oauth-client` or `oauth-rop` authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#client_secret SyntheticsTest#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Domain for `ntlm` authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#domain SyntheticsTest#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Password for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#password SyntheticsTest#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Region for `SIGV4` authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#region SyntheticsTest#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Resource for `oauth-client` or `oauth-rop` authentication. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#resource SyntheticsTest#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
	// Scope for `oauth-client` or `oauth-rop` authentication. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#scope SyntheticsTest#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Secret key for `SIGV4` authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#secret_key SyntheticsTest#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// Service name for `SIGV4` authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#service_name SyntheticsTest#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Session token for `SIGV4` authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#session_token SyntheticsTest#session_token}
	SessionToken *string `field:"optional" json:"sessionToken" yaml:"sessionToken"`
	// Token API Authentication for `oauth-client` or `oauth-rop` authentication. Valid values are `header`, `body`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#token_api_authentication SyntheticsTest#token_api_authentication}
	TokenApiAuthentication *string `field:"optional" json:"tokenApiAuthentication" yaml:"tokenApiAuthentication"`
	// Type of basic authentication to use when performing the test. Defaults to `"web"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Username for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#username SyntheticsTest#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Workstation for `ntlm` authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/synthetics_test#workstation SyntheticsTest#workstation}
	Workstation *string `field:"optional" json:"workstation" yaml:"workstation"`
}

