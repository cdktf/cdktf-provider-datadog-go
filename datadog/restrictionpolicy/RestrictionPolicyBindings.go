package restrictionpolicy


type RestrictionPolicyBindings struct {
	// An array of principals.
	//
	// A principal is a subject or group of subjects. Each principal is formatted as `type:id`. Supported types: `role` and `org`. The org ID can be obtained through the api/v2/users API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/restriction_policy#principals RestrictionPolicy#principals}
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
	// The role/level of access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/restriction_policy#relation RestrictionPolicy#relation}
	Relation *string `field:"optional" json:"relation" yaml:"relation"`
}

