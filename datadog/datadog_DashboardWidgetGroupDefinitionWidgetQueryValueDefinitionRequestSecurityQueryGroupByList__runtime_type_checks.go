//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestSecurityQueryGroupByList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestSecurityQueryGroupByList) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestSecurityQueryGroupByList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestSecurityQueryGroupBy:
		val := val.(*[]*DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestSecurityQueryGroupBy)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestSecurityQueryGroupBy:
		val_ := val.([]*DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestSecurityQueryGroupBy)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestSecurityQueryGroupBy; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestSecurityQueryGroupByList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestSecurityQueryGroupByList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestSecurityQueryGroupByList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestSecurityQueryGroupByListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}
