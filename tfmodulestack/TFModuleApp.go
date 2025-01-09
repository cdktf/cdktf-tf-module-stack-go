package tfmodulestack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-tf-module-stack-go/tfmodulestack/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-tf-module-stack-go/tfmodulestack/v6/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TFModuleApp interface {
	cdktf.App
	// Experimental.
	HclOutput() *bool
	// Experimental.
	Manifest() cdktf.Manifest
	// The tree node.
	Node() constructs.Node
	// The output directory into which resources will be synthesized.
	// Experimental.
	Outdir() *string
	// Whether to skip backend validation during synthesis of the app.
	// Experimental.
	SkipBackendValidation() *bool
	// Whether to skip all validations during synthesis of the app.
	// Experimental.
	SkipValidation() *bool
	// The stack which will be synthesized.
	//
	// If not set, all stacks will be synthesized.
	// Experimental.
	TargetStackId() *string
	// Creates a reference from one stack to another, invoked on prepareStack since it creates extra resources.
	// Experimental.
	CrossStackReference(fromStack cdktf.TerraformStack, toStack cdktf.TerraformStack, identifier *string) *string
	// Synthesizes all resources to the output directory.
	// Experimental.
	Synth()
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TFModuleApp
type jsiiProxy_TFModuleApp struct {
	internal.Type__cdktfApp
}

func (j *jsiiProxy_TFModuleApp) HclOutput() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hclOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleApp) Manifest() cdktf.Manifest {
	var returns cdktf.Manifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleApp) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleApp) SkipBackendValidation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipBackendValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleApp) SkipValidation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleApp) TargetStackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetStackId",
		&returns,
	)
	return returns
}


func NewTFModuleApp(options *cdktf.AppConfig) TFModuleApp {
	_init_.Initialize()

	if err := validateNewTFModuleAppParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_TFModuleApp{}

	_jsii_.Create(
		"@cdktf/tf-module-stack.TFModuleApp",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewTFModuleApp_Override(t TFModuleApp, options *cdktf.AppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/tf-module-stack.TFModuleApp",
		[]interface{}{options},
		t,
	)
}

// Experimental.
func TFModuleApp_IsApp(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTFModuleApp_IsAppParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/tf-module-stack.TFModuleApp",
		"isApp",
		[]interface{}{x},
		&returns,
	)

	return returns
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
func TFModuleApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTFModuleApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/tf-module-stack.TFModuleApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TFModuleApp_Of(construct constructs.IConstruct) cdktf.App {
	_init_.Initialize()

	if err := validateTFModuleApp_OfParameters(construct); err != nil {
		panic(err)
	}
	var returns cdktf.App

	_jsii_.StaticInvoke(
		"@cdktf/tf-module-stack.TFModuleApp",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleApp) CrossStackReference(fromStack cdktf.TerraformStack, toStack cdktf.TerraformStack, identifier *string) *string {
	if err := t.validateCrossStackReferenceParameters(fromStack, toStack, identifier); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"crossStackReference",
		[]interface{}{fromStack, toStack, identifier},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleApp) Synth() {
	_jsii_.InvokeVoid(
		t,
		"synth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TFModuleApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

