//go:build no_runtime_type_checking

// A drop-in replacement for cdktf.TerraformStack that let's you define Terraform modules as construct
package tfmodulestack

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TFModuleStack) validateAddDependencyParameters(dependency cdktf.TerraformStack) error {
	return nil
}

func (t *jsiiProxy_TFModuleStack) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TFModuleStack) validateAllocateLogicalIdParameters(tfElement interface{}) error {
	return nil
}

func (t *jsiiProxy_TFModuleStack) validateDependsOnParameters(stack cdktf.TerraformStack) error {
	return nil
}

func (t *jsiiProxy_TFModuleStack) validateGetLogicalIdParameters(tfElement interface{}) error {
	return nil
}

func (t *jsiiProxy_TFModuleStack) validateRegisterIncomingCrossStackReferenceParameters(fromStack cdktf.TerraformStack) error {
	return nil
}

func (t *jsiiProxy_TFModuleStack) validateRegisterOutgoingCrossStackReferenceParameters(identifier *string) error {
	return nil
}

func validateTFModuleStack_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTFModuleStack_IsStackParameters(x interface{}) error {
	return nil
}

func validateTFModuleStack_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_TFModuleStack) validateSetDependenciesParameters(val *[]cdktf.TerraformStack) error {
	return nil
}

func (j *jsiiProxy_TFModuleStack) validateSetSynthesizerParameters(val cdktf.IStackSynthesizer) error {
	return nil
}

func validateNewTFModuleStackParameters(scope constructs.Construct, id *string) error {
	return nil
}

