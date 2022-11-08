//go:build no_runtime_type_checking

// A drop-in replacement for cdktf.TerraformStack that let's you define Terraform modules as construct
package tfmodulestack

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TFModuleVariable) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TFModuleVariable) validateAddValidationParameters(validation *cdktf.TerraformVariableValidationConfig) error {
	return nil
}

func (t *jsiiProxy_TFModuleVariable) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTFModuleVariable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTFModuleVariableParameters(scope constructs.Construct, name *string, config *cdktf.TerraformVariableConfig) error {
	return nil
}

