package provider


type DatadogProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog#alias DatadogProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// (Required unless validate is false) Datadog API key. This can also be set via the DD_API_KEY environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog#api_key DatadogProvider#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// The API URL.
	//
	// This can also be set via the DD_HOST environment variable. Note that this URL must not end with the `/api/` path. For example, `https://api.datadoghq.com/` is a correct value, while `https://api.datadoghq.com/api/` is not. And if you're working with "EU" version of Datadog, use `https://api.datadoghq.eu/`. Other Datadog region examples: `https://api.us5.datadoghq.com/`, `https://api.us3.datadoghq.com/` and `https://api.ddog-gov.com/`. See https://docs.datadoghq.com/getting_started/site/ for all available regions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog#api_url DatadogProvider#api_url}
	ApiUrl *string `field:"optional" json:"apiUrl" yaml:"apiUrl"`
	// (Required unless validate is false) Datadog APP key. This can also be set via the DD_APP_KEY environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog#app_key DatadogProvider#app_key}
	AppKey *string `field:"optional" json:"appKey" yaml:"appKey"`
	// The HTTP request retry back off base. Defaults to 2.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog#http_client_retry_backoff_base DatadogProvider#http_client_retry_backoff_base}
	HttpClientRetryBackoffBase *float64 `field:"optional" json:"httpClientRetryBackoffBase" yaml:"httpClientRetryBackoffBase"`
	// The HTTP request retry back off multiplier. Defaults to 2.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog#http_client_retry_backoff_multiplier DatadogProvider#http_client_retry_backoff_multiplier}
	HttpClientRetryBackoffMultiplier *float64 `field:"optional" json:"httpClientRetryBackoffMultiplier" yaml:"httpClientRetryBackoffMultiplier"`
	// Enables request retries on HTTP status codes 429 and 5xx. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog#http_client_retry_enabled DatadogProvider#http_client_retry_enabled}
	HttpClientRetryEnabled interface{} `field:"optional" json:"httpClientRetryEnabled" yaml:"httpClientRetryEnabled"`
	// The HTTP request maximum retry number. Defaults to 3.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog#http_client_retry_max_retries DatadogProvider#http_client_retry_max_retries}
	HttpClientRetryMaxRetries *float64 `field:"optional" json:"httpClientRetryMaxRetries" yaml:"httpClientRetryMaxRetries"`
	// The HTTP request retry timeout period. Defaults to 60 seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog#http_client_retry_timeout DatadogProvider#http_client_retry_timeout}
	HttpClientRetryTimeout *float64 `field:"optional" json:"httpClientRetryTimeout" yaml:"httpClientRetryTimeout"`
	// Enables validation of the provided API and APP keys during provider initialization.
	//
	// Default is true. When false, api_key and app_key won't be checked.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog#validate DatadogProvider#validate}
	Validate interface{} `field:"optional" json:"validate" yaml:"validate"`
}

