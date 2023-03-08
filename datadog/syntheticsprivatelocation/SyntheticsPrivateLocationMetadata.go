package syntheticsprivatelocation


type SyntheticsPrivateLocationMetadata struct {
	// A list of role identifiers pulled from the Roles API to restrict read and write access.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/synthetics_private_location#restricted_roles SyntheticsPrivateLocation#restricted_roles}
	RestrictedRoles *[]*string `field:"optional" json:"restrictedRoles" yaml:"restrictedRoles"`
}

