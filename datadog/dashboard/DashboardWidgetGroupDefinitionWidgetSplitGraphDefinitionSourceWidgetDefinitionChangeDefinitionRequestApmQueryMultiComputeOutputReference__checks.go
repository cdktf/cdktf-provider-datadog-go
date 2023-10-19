// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package dashboard

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateSetAggregationParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateSetFacetParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiCompute:
		val := val.(*DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiCompute)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiCompute:
		val_ := val.(DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiCompute)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiCompute; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateSetIntervalParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestApmQueryMultiComputeOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

