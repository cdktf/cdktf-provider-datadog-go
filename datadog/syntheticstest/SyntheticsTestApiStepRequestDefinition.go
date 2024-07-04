// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestApiStepRequestDefinition struct {
	// Allows loading insecure content for an HTTP request in an API test or in a multistep API test step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#allow_insecure SyntheticsTest#allow_insecure}
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// The request body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#body SyntheticsTest#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Type of the request body. Valid values are `text/plain`, `application/json`, `text/xml`, `text/html`, `application/x-www-form-urlencoded`, `graphql`, `application/octet-stream`, `multipart/form-data`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#body_type SyntheticsTest#body_type}
	BodyType *string `field:"optional" json:"bodyType" yaml:"bodyType"`
	// The type of gRPC call to perform. Valid values are `healthcheck`, `unary`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#call_type SyntheticsTest#call_type}
	CallType *string `field:"optional" json:"callType" yaml:"callType"`
	// By default, the client certificate is applied on the domain of the starting URL for browser tests.
	//
	// If you want your client certificate to be applied on other domains instead, add them in `certificate_domains`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#certificate_domains SyntheticsTest#certificate_domains}
	CertificateDomains *[]*string `field:"optional" json:"certificateDomains" yaml:"certificateDomains"`
	// DNS server to use for DNS tests (`subtype = "dns"`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#dns_server SyntheticsTest#dns_server}
	DnsServer *string `field:"optional" json:"dnsServer" yaml:"dnsServer"`
	// DNS server port to use for DNS tests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#dns_server_port SyntheticsTest#dns_server_port}
	DnsServerPort *float64 `field:"optional" json:"dnsServerPort" yaml:"dnsServerPort"`
	// Determines whether or not the API HTTP test should follow redirects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#follow_redirects SyntheticsTest#follow_redirects}
	FollowRedirects interface{} `field:"optional" json:"followRedirects" yaml:"followRedirects"`
	// Host name to perform the test with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#host SyntheticsTest#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// HTTP version to use for an HTTP request in an API test or step.
	//
	// Valid values are `http1`, `http2`, `any`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#http_version SyntheticsTest#http_version}
	HttpVersion *string `field:"optional" json:"httpVersion" yaml:"httpVersion"`
	// For UDP and websocket tests, message to send with the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#message SyntheticsTest#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Either the HTTP method/verb to use or a gRPC method available on the service set in the `service` field.
	//
	// Required if `subtype` is `HTTP` or if `subtype` is `grpc` and `callType` is `unary`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#method SyntheticsTest#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Determines whether or not to save the response body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#no_saving_response_body SyntheticsTest#no_saving_response_body}
	NoSavingResponseBody interface{} `field:"optional" json:"noSavingResponseBody" yaml:"noSavingResponseBody"`
	// Number of pings to use per test for ICMP tests (`subtype = "icmp"`) between 0 and 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#number_of_packets SyntheticsTest#number_of_packets}
	NumberOfPackets *float64 `field:"optional" json:"numberOfPackets" yaml:"numberOfPackets"`
	// Persist cookies across redirects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#persist_cookies SyntheticsTest#persist_cookies}
	PersistCookies interface{} `field:"optional" json:"persistCookies" yaml:"persistCookies"`
	// The content of a proto file as a string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#plain_proto_file SyntheticsTest#plain_proto_file}
	PlainProtoFile *string `field:"optional" json:"plainProtoFile" yaml:"plainProtoFile"`
	// Port to use when performing the test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#port SyntheticsTest#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A protobuf JSON descriptor. **Deprecated.** Use `plain_proto_file` instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#proto_json_descriptor SyntheticsTest#proto_json_descriptor}
	ProtoJsonDescriptor *string `field:"optional" json:"protoJsonDescriptor" yaml:"protoJsonDescriptor"`
	// For SSL tests, it specifies on which server you want to initiate the TLS handshake, allowing the server to present one of multiple possible certificates on the same IP address and TCP port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#servername SyntheticsTest#servername}
	Servername *string `field:"optional" json:"servername" yaml:"servername"`
	// The gRPC service on which you want to perform the gRPC call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#service SyntheticsTest#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// This will turn on a traceroute probe to discover all gateways along the path to the host destination.
	//
	// For ICMP tests (`subtype = "icmp"`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#should_track_hops SyntheticsTest#should_track_hops}
	ShouldTrackHops interface{} `field:"optional" json:"shouldTrackHops" yaml:"shouldTrackHops"`
	// Timeout in seconds for the test. Defaults to `60`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#timeout SyntheticsTest#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The URL to send the request to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#url SyntheticsTest#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

