//go:build !no_runtime_type_checking

// A drop-in replacement for cdktf.TerraformStack that let's you define Terraform modules as construct
package tfmodulestack

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (t *jsiiProxy_TFModuleApp) validateCrossStackReferenceParameters(fromStack cdktf.TerraformStack, toStack cdktf.TerraformStack, identifier *string) error {
	if fromStack == nil {
		return fmt.Errorf("parameter fromStack is required, but nil was provided")
	}

	if toStack == nil {
		return fmt.Errorf("parameter toStack is required, but nil was provided")
	}

	if identifier == nil {
		return fmt.Errorf("parameter identifier is required, but nil was provided")
	}

	return nil
}

func validateTFModuleApp_IsAppParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateTFModuleApp_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateTFModuleApp_OfParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewTFModuleAppParameters(options *cdktf.AppOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

