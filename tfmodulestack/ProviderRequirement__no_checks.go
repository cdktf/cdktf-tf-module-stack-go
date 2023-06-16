//go:build no_runtime_type_checking

package tfmodulestack

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProviderRequirement) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_ProviderRequirement) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateProviderRequirement_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProviderRequirement_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateProviderRequirement_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewProviderRequirementParameters(scope constructs.Construct, providerName *string) error {
	return nil
}

