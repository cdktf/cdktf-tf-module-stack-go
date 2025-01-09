package tfmodulestack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-tf-module-stack-go/tfmodulestack/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-tf-module-stack-go/tfmodulestack/v6/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TFModuleVariable interface {
	cdktf.TerraformVariable
	// Experimental.
	BooleanValue() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Default() interface{}
	// Experimental.
	Description() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	ListValue() *[]*string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Nullable() *bool
	// Experimental.
	NumberValue() *float64
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Sensitive() *bool
	// Experimental.
	StringValue() *string
	// Experimental.
	Type() *string
	// Experimental.
	Validation() *[]*cdktf.TerraformVariableValidationConfig
	// Experimental.
	Value() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddValidation(validation *cdktf.TerraformVariableValidationConfig)
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
	//
	// Returns: a string token referencing the value of this variable.
	// Experimental.
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TFModuleVariable
type jsiiProxy_TFModuleVariable struct {
	internal.Type__cdktfTerraformVariable
}

func (j *jsiiProxy_TFModuleVariable) BooleanValue() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"booleanValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) Default() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) ListValue() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) Nullable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"nullable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) NumberValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) Sensitive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"sensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) Validation() *[]*cdktf.TerraformVariableValidationConfig {
	var returns *[]*cdktf.TerraformVariableValidationConfig
	_jsii_.Get(
		j,
		"validation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleVariable) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewTFModuleVariable(scope constructs.Construct, name *string, config *cdktf.TerraformVariableConfig) TFModuleVariable {
	_init_.Initialize()

	if err := validateNewTFModuleVariableParameters(scope, name, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TFModuleVariable{}

	_jsii_.Create(
		"@cdktf/tf-module-stack.TFModuleVariable",
		[]interface{}{scope, name, config},
		&j,
	)

	return &j
}

func NewTFModuleVariable_Override(t TFModuleVariable, scope constructs.Construct, name *string, config *cdktf.TerraformVariableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/tf-module-stack.TFModuleVariable",
		[]interface{}{scope, name, config},
		t,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func TFModuleVariable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTFModuleVariable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/tf-module-stack.TFModuleVariable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TFModuleVariable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTFModuleVariable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/tf-module-stack.TFModuleVariable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleVariable) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TFModuleVariable) AddValidation(validation *cdktf.TerraformVariableValidationConfig) {
	if err := t.validateAddValidationParameters(validation); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addValidation",
		[]interface{}{validation},
	)
}

func (t *jsiiProxy_TFModuleVariable) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TFModuleVariable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TFModuleVariable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleVariable) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleVariable) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleVariable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleVariable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleVariable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

