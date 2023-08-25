// A drop-in replacement for cdktf.TerraformStack that lets you define Terraform modules as constructs
package tfmodulestack

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/tf-module-stack.ProviderRequirement",
		reflect.TypeOf((*ProviderRequirement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ProviderRequirement{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/tf-module-stack.TFModuleApp",
		reflect.TypeOf((*TFModuleApp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "crossStackReference", GoMethod: "CrossStackReference"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "skipBackendValidation", GoGetter: "SkipBackendValidation"},
			_jsii_.MemberProperty{JsiiProperty: "skipValidation", GoGetter: "SkipValidation"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberProperty{JsiiProperty: "targetStackId", GoGetter: "TargetStackId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TFModuleApp{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfApp)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/tf-module-stack.TFModuleOutput",
		reflect.TypeOf((*TFModuleOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "precondition", GoGetter: "Precondition"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "sensitive", GoGetter: "Sensitive"},
			_jsii_.MemberProperty{JsiiProperty: "staticId", GoGetter: "StaticId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_TFModuleOutput{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformOutput)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/tf-module-stack.TFModuleStack",
		reflect.TypeOf((*TFModuleStack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "allocateLogicalId", GoMethod: "AllocateLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "allProviders", GoMethod: "AllProviders"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberMethod{JsiiMethod: "dependsOn", GoMethod: "DependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "ensureBackendExists", GoMethod: "EnsureBackendExists"},
			_jsii_.MemberMethod{JsiiMethod: "getLogicalId", GoMethod: "GetLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "prepareStack", GoMethod: "PrepareStack"},
			_jsii_.MemberMethod{JsiiMethod: "registerIncomingCrossStackReference", GoMethod: "RegisterIncomingCrossStackReference"},
			_jsii_.MemberMethod{JsiiMethod: "registerOutgoingCrossStackReference", GoMethod: "RegisterOutgoingCrossStackReference"},
			_jsii_.MemberMethod{JsiiMethod: "runAllValidations", GoMethod: "RunAllValidations"},
			_jsii_.MemberProperty{JsiiProperty: "synthesizer", GoGetter: "Synthesizer"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_TFModuleStack{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformStack)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/tf-module-stack.TFModuleVariable",
		reflect.TypeOf((*TFModuleVariable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addValidation", GoMethod: "AddValidation"},
			_jsii_.MemberProperty{JsiiProperty: "booleanValue", GoGetter: "BooleanValue"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "default", GoGetter: "Default"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "listValue", GoGetter: "ListValue"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nullable", GoGetter: "Nullable"},
			_jsii_.MemberProperty{JsiiProperty: "numberValue", GoGetter: "NumberValue"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "sensitive", GoGetter: "Sensitive"},
			_jsii_.MemberProperty{JsiiProperty: "stringValue", GoGetter: "StringValue"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "validation", GoGetter: "Validation"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_TFModuleVariable{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformVariable)
			return &j
		},
	)
}
