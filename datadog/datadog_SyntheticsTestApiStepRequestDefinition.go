// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type SyntheticsTestApiStepRequestDefinition struct {
	// Allows loading insecure content for an HTTP test.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#allow_insecure SyntheticsTest#allow_insecure}
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// The request body.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#body SyntheticsTest#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// DNS server to use for DNS tests (`subtype = "dns"`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#dns_server SyntheticsTest#dns_server}
	DnsServer *string `field:"optional" json:"dnsServer" yaml:"dnsServer"`
	// DNS server port to use for DNS tests.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#dns_server_port SyntheticsTest#dns_server_port}
	DnsServerPort *float64 `field:"optional" json:"dnsServerPort" yaml:"dnsServerPort"`
	// Determines whether or not the API HTTP test should follow redirects.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#follow_redirects SyntheticsTest#follow_redirects}
	FollowRedirects interface{} `field:"optional" json:"followRedirects" yaml:"followRedirects"`
	// Host name to perform the test with.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#host SyntheticsTest#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// For UDP and websocket tests, message to send with the request.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#message SyntheticsTest#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// The HTTP method. Valid values are `GET`, `POST`, `PATCH`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#method SyntheticsTest#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Determines whether or not to save the response body.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#no_saving_response_body SyntheticsTest#no_saving_response_body}
	NoSavingResponseBody interface{} `field:"optional" json:"noSavingResponseBody" yaml:"noSavingResponseBody"`
	// Number of pings to use per test for ICMP tests (`subtype = "icmp"`) between 0 and 10.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#number_of_packets SyntheticsTest#number_of_packets}
	NumberOfPackets *float64 `field:"optional" json:"numberOfPackets" yaml:"numberOfPackets"`
	// Port to use when performing the test.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#port SyntheticsTest#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// For SSL tests, it specifies on which server you want to initiate the TLS handshake, allowing the server to present one of multiple possible certificates on the same IP address and TCP port number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#servername SyntheticsTest#servername}
	Servername *string `field:"optional" json:"servername" yaml:"servername"`
	// For gRPC tests, service to target for healthcheck.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#service SyntheticsTest#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// This will turn on a traceroute probe to discover all gateways along the path to the host destination.
	//
	// For ICMP tests (`subtype = "icmp"`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#should_track_hops SyntheticsTest#should_track_hops}
	ShouldTrackHops interface{} `field:"optional" json:"shouldTrackHops" yaml:"shouldTrackHops"`
	// Timeout in seconds for the test. Defaults to `60`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#timeout SyntheticsTest#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The URL to send the request to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_test#url SyntheticsTest#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}
