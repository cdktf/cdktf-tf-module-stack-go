package tfmodulestack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-tf-module-stack-go/tfmodulestack/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-tf-module-stack-go/tfmodulestack/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TFModuleStack interface {
	cdktf.TerraformStack
	// Experimental.
	Dependencies() *[]cdktf.TerraformStack
	// Experimental.
	SetDependencies(val *[]cdktf.TerraformStack)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Synthesizer() cdktf.IStackSynthesizer
	// Experimental.
	SetSynthesizer(val cdktf.IStackSynthesizer)
	// Experimental.
	AddDependency(dependency cdktf.TerraformStack)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Returns the naming scheme used to allocate logical IDs.
	//
	// By default, uses
	// the `HashedAddressingScheme` but this method can be overridden to customize
	// this behavior.
	// Experimental.
	AllocateLogicalId(tfElement interface{}) *string
	// Experimental.
	AllProviders() *[]cdktf.TerraformProvider
	// Experimental.
	DependsOn(stack cdktf.TerraformStack) *bool
	// Experimental.
	EnsureBackendExists() cdktf.TerraformBackend
	// Experimental.
	GetLogicalId(tfElement interface{}) *string
	// Experimental.
	PrepareStack()
	// Experimental.
	RegisterIncomingCrossStackReference(fromStack cdktf.TerraformStack) cdktf.TerraformRemoteState
	// Experimental.
	RegisterOutgoingCrossStackReference(identifier *string) cdktf.TerraformOutput
	// Run all validations on the stack.
	// Experimental.
	RunAllValidations()
	// Returns a string representation of this construct.
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TFModuleStack
type jsiiProxy_TFModuleStack struct {
	internal.Type__cdktfTerraformStack
}

func (j *jsiiProxy_TFModuleStack) Dependencies() *[]cdktf.TerraformStack {
	var returns *[]cdktf.TerraformStack
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TFModuleStack) Synthesizer() cdktf.IStackSynthesizer {
	var returns cdktf.IStackSynthesizer
	_jsii_.Get(
		j,
		"synthesizer",
		&returns,
	)
	return returns
}


// Experimental.
func NewTFModuleStack(scope constructs.Construct, id *string) TFModuleStack {
	_init_.Initialize()

	if err := validateNewTFModuleStackParameters(scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_TFModuleStack{}

	_jsii_.Create(
		"@cdktf/tf-module-stack.TFModuleStack",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewTFModuleStack_Override(t TFModuleStack, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/tf-module-stack.TFModuleStack",
		[]interface{}{scope, id},
		t,
	)
}

func (j *jsiiProxy_TFModuleStack)SetDependencies(val *[]cdktf.TerraformStack) {
	if err := j.validateSetDependenciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependencies",
		val,
	)
}

func (j *jsiiProxy_TFModuleStack)SetSynthesizer(val cdktf.IStackSynthesizer) {
	if err := j.validateSetSynthesizerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"synthesizer",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func TFModuleStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTFModuleStack_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/tf-module-stack.TFModuleStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TFModuleStack_IsStack(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTFModuleStack_IsStackParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/tf-module-stack.TFModuleStack",
		"isStack",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TFModuleStack_Of(construct constructs.IConstruct) cdktf.TerraformStack {
	_init_.Initialize()

	if err := validateTFModuleStack_OfParameters(construct); err != nil {
		panic(err)
	}
	var returns cdktf.TerraformStack

	_jsii_.StaticInvoke(
		"@cdktf/tf-module-stack.TFModuleStack",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleStack) AddDependency(dependency cdktf.TerraformStack) {
	if err := t.validateAddDependencyParameters(dependency); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addDependency",
		[]interface{}{dependency},
	)
}

func (t *jsiiProxy_TFModuleStack) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TFModuleStack) AllocateLogicalId(tfElement interface{}) *string {
	if err := t.validateAllocateLogicalIdParameters(tfElement); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"allocateLogicalId",
		[]interface{}{tfElement},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleStack) AllProviders() *[]cdktf.TerraformProvider {
	var returns *[]cdktf.TerraformProvider

	_jsii_.Invoke(
		t,
		"allProviders",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleStack) DependsOn(stack cdktf.TerraformStack) *bool {
	if err := t.validateDependsOnParameters(stack); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		t,
		"dependsOn",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleStack) EnsureBackendExists() cdktf.TerraformBackend {
	var returns cdktf.TerraformBackend

	_jsii_.Invoke(
		t,
		"ensureBackendExists",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleStack) GetLogicalId(tfElement interface{}) *string {
	if err := t.validateGetLogicalIdParameters(tfElement); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getLogicalId",
		[]interface{}{tfElement},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleStack) PrepareStack() {
	_jsii_.InvokeVoid(
		t,
		"prepareStack",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TFModuleStack) RegisterIncomingCrossStackReference(fromStack cdktf.TerraformStack) cdktf.TerraformRemoteState {
	if err := t.validateRegisterIncomingCrossStackReferenceParameters(fromStack); err != nil {
		panic(err)
	}
	var returns cdktf.TerraformRemoteState

	_jsii_.Invoke(
		t,
		"registerIncomingCrossStackReference",
		[]interface{}{fromStack},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleStack) RegisterOutgoingCrossStackReference(identifier *string) cdktf.TerraformOutput {
	if err := t.validateRegisterOutgoingCrossStackReferenceParameters(identifier); err != nil {
		panic(err)
	}
	var returns cdktf.TerraformOutput

	_jsii_.Invoke(
		t,
		"registerOutgoingCrossStackReference",
		[]interface{}{identifier},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleStack) RunAllValidations() {
	_jsii_.InvokeVoid(
		t,
		"runAllValidations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TFModuleStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TFModuleStack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

