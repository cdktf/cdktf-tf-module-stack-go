//go:build no_runtime_type_checking

// A drop-in replacement for cdktf.TerraformStack that let's you define Terraform modules as construct
package tfmodulestack

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TFModuleApp) validateCrossStackReferenceParameters(fromStack cdktf.TerraformStack, toStack cdktf.TerraformStack, identifier *string) error {
	return nil
}

func validateTFModuleApp_IsAppParameters(x interface{}) error {
	return nil
}

func validateTFModuleApp_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTFModuleApp_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTFModuleAppParameters(options *cdktf.AppOptions) error {
	return nil
}

