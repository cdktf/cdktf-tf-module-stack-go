//go:build !no_runtime_type_checking

package tfmodulestack

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (t *jsiiProxy_TFModuleStack) validateAddDependencyParameters(dependency cdktf.TerraformStack) error {
	if dependency == nil {
		return fmt.Errorf("parameter dependency is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TFModuleStack) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TFModuleStack) validateAllocateLogicalIdParameters(tfElement interface{}) error {
	if tfElement == nil {
		return fmt.Errorf("parameter tfElement is required, but nil was provided")
	}
	switch tfElement.(type) {
	case cdktf.TerraformElement:
		// ok
	case constructs.Node:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(tfElement) {
			return fmt.Errorf("parameter tfElement must be one of the allowed types: cdktf.TerraformElement, constructs.Node; received %#v (a %T)", tfElement, tfElement)
		}
	}

	return nil
}

func (t *jsiiProxy_TFModuleStack) validateDependsOnParameters(stack cdktf.TerraformStack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TFModuleStack) validateGetLogicalIdParameters(tfElement interface{}) error {
	if tfElement == nil {
		return fmt.Errorf("parameter tfElement is required, but nil was provided")
	}
	switch tfElement.(type) {
	case cdktf.TerraformElement:
		// ok
	case constructs.Node:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(tfElement) {
			return fmt.Errorf("parameter tfElement must be one of the allowed types: cdktf.TerraformElement, constructs.Node; received %#v (a %T)", tfElement, tfElement)
		}
	}

	return nil
}

func (t *jsiiProxy_TFModuleStack) validateRegisterIncomingCrossStackReferenceParameters(fromStack cdktf.TerraformStack) error {
	if fromStack == nil {
		return fmt.Errorf("parameter fromStack is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TFModuleStack) validateRegisterOutgoingCrossStackReferenceParameters(identifier *string) error {
	if identifier == nil {
		return fmt.Errorf("parameter identifier is required, but nil was provided")
	}

	return nil
}

func validateTFModuleStack_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateTFModuleStack_IsStackParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateTFModuleStack_OfParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TFModuleStack) validateSetDependenciesParameters(val *[]cdktf.TerraformStack) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TFModuleStack) validateSetSynthesizerParameters(val cdktf.IStackSynthesizer) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTFModuleStackParameters(scope constructs.Construct, id *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

