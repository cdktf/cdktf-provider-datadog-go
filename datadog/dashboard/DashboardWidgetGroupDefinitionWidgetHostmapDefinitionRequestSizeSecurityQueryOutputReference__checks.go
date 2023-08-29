// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package dashboard

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validatePutComputeQueryParameters(value *DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryComputeQuery) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validatePutGroupByParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryGroupBy:
		value := value.(*[]*DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryGroupBy)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryGroupBy:
		value_ := value.([]*DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryGroupBy)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryGroupBy; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validatePutMultiComputeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryMultiCompute:
		value := value.(*[]*DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryMultiCompute)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryMultiCompute:
		value_ := value.([]*DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryMultiCompute)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryMultiCompute; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateSetIndexParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateSetInternalValueParameters(val *DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQuery) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateSetSearchQueryParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

