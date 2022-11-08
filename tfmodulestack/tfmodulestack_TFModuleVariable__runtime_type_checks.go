//go:build !no_runtime_type_checking

// A drop-in replacement for cdktf.TerraformStack that let's you define Terraform modules as construct
package tfmodulestack

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (t *jsiiProxy_TFModuleVariable) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TFModuleVariable) validateAddValidationParameters(validation *cdktf.TerraformVariableValidationConfig) error {
	if validation == nil {
		return fmt.Errorf("parameter validation is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(validation, func() string { return "parameter validation" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TFModuleVariable) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func validateTFModuleVariable_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewTFModuleVariableParameters(scope constructs.Construct, name *string, config *cdktf.TerraformVariableConfig) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

