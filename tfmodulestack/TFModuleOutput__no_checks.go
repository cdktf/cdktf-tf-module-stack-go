//go:build no_runtime_type_checking

package tfmodulestack

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TFModuleOutput) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TFModuleOutput) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateTFModuleOutput_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTFModuleOutput_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateTFModuleOutput_IsTerraformOutputParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_TFModuleOutput) validateSetPreconditionParameters(val *cdktf.Precondition) error {
	return nil
}

func (j *jsiiProxy_TFModuleOutput) validateSetStaticIdParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_TFModuleOutput) validateSetValueParameters(val interface{}) error {
	return nil
}

func validateNewTFModuleOutputParameters(scope constructs.Construct, name *string, config *cdktf.TerraformOutputConfig) error {
	return nil
}

