//go:build !no_runtime_type_checking

// A drop-in replacement for cdktf.TerraformStack that let's you define Terraform modules as construct
package tfmodulestack

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (p *jsiiProxy_ProviderRequirement) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_ProviderRequirement) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func validateProviderRequirement_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewProviderRequirementParameters(scope constructs.Construct, providerName *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if providerName == nil {
		return fmt.Errorf("parameter providerName is required, but nil was provided")
	}

	return nil
}

