package tfmodulestack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-tf-module-stack-go/tfmodulestack/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-tf-module-stack-go/tfmodulestack/v5/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TFModuleOutput interface {
	cdktf.TerraformOutput
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	DependsOn() *[]cdktf.ITerraformDependable
	// Experimental.
	SetDependsOn(val *[]cdktf.ITerraformDependable)
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Precondition() *cdktf.Precondition
	// Experimental.
	SetPrecondition(val *cdktf.Precondition)
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Sensitive() *bool
	// Experimental.
	SetSensitive(val *bool)
	// Experimental.
	StaticId() *bool
	// Experimental.
	SetStaticId(val *bool)
	// Experimental.
	Value() interface{}
	// Experimental.
	SetValue(val interface{})
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TFModuleOutput
type jsiiProxy_TFModuleOutput struct {
	internal.Type__cdktfTerraformOutput
}

func (j *jsiiProxy_TFModuleOutput) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleOutput) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleOutput) DependsOn() *[]cdktf.ITerraformDependable {
	var returns *[]cdktf.ITerraformDependable
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleOutput) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleOutput) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleOutput) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleOutput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleOutput) Precondition() *cdktf.Precondition {
	var returns *cdktf.Precondition
	_jsii_.Get(
		j,
		"precondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleOutput) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleOutput) Sensitive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"sensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleOutput) StaticId() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"staticId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleOutput) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewTFModuleOutput(scope constructs.Construct, name *string, config *cdktf.TerraformOutputConfig) TFModuleOutput {
	_init_.Initialize()

	if err := validateNewTFModuleOutputParameters(scope, name, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TFModuleOutput{}

	_jsii_.Create(
		"@cdktf/tf-module-stack.TFModuleOutput",
		[]interface{}{scope, name, config},
		&j,
	)

	return &j
}

func NewTFModuleOutput_Override(t TFModuleOutput, scope constructs.Construct, name *string, config *cdktf.TerraformOutputConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/tf-module-stack.TFModuleOutput",
		[]interface{}{scope, name, config},
		t,
	)
}

func (j *jsiiProxy_TFModuleOutput)SetDependsOn(val *[]cdktf.ITerraformDependable) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TFModuleOutput)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TFModuleOutput)SetPrecondition(val *cdktf.Precondition) {
	if err := j.validateSetPreconditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"precondition",
		val,
	)
}

func (j *jsiiProxy_TFModuleOutput)SetSensitive(val *bool) {
	_jsii_.Set(
		j,
		"sensitive",
		val,
	)
}

func (j *jsiiProxy_TFModuleOutput)SetStaticId(val *bool) {
	if err := j.validateSetStaticIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticId",
		val,
	)
}

func (j *jsiiProxy_TFModuleOutput)SetValue(val interface{}) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func TFModuleOutput_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTFModuleOutput_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/tf-module-stack.TFModuleOutput",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TFModuleOutput_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTFModuleOutput_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/tf-module-stack.TFModuleOutput",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TFModuleOutput_IsTerraformOutput(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTFModuleOutput_IsTerraformOutputParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/tf-module-stack.TFModuleOutput",
		"isTerraformOutput",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleOutput) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TFModuleOutput) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TFModuleOutput) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TFModuleOutput) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleOutput) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleOutput) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleOutput) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleOutput) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleOutput) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

