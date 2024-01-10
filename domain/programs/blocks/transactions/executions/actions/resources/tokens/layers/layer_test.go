package layers

import (
	"reflect"
	"testing"

	"steve.care/network/domain/programs/logics/libraries/layers"
)

func TestLayer_withLayer_Success(t *testing.T) {
	layer := layers.NewLayerForTests(
		layers.NewInstructionsForTests([]layers.Instruction{
			layers.NewInstructionWithStopForTests(),
		}),
		layers.NewOutputForTests(
			"myVariable",
			layers.NewKindWithPromptForTests(),
		),
		"myInput",
	)

	ins := NewLayerWithLayerForTests(layer)

	if !ins.IsLayer() {
		t.Errorf("the layer was expected to contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retLayer := ins.Layer()
	if !reflect.DeepEqual(layer, retLayer) {
		t.Errorf("the returned layer is invalid")
		return
	}
}

func TestLayer_withOutput_Success(t *testing.T) {
	output := layers.NewOutputForTests(
		"myVariable",
		layers.NewKindWithPromptForTests(),
	)

	ins := NewLayerWithOutputForTests(output)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if !ins.IsOutput() {
		t.Errorf("the layer was expected to contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retOutput := ins.Output()
	if !reflect.DeepEqual(output, retOutput) {
		t.Errorf("the returned output is invalid")
		return
	}
}

func TestLayer_withKind_Success(t *testing.T) {
	kind := layers.NewKindWithPromptForTests()

	ins := NewLayerWithKindForTests(kind)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if !ins.IsKind() {
		t.Errorf("the layer was expected to contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retKind := ins.Kind()
	if !reflect.DeepEqual(kind, retKind) {
		t.Errorf("the returned kind is invalid")
		return
	}
}

func TestLayer_withInstruction_Success(t *testing.T) {
	instruction := layers.NewInstructionWithStopForTests()

	ins := NewLayerWithInstructionForTests(instruction)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if !ins.IsInstruction() {
		t.Errorf("the layer was expected to contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retInstruction := ins.Instruction()
	if !reflect.DeepEqual(instruction, retInstruction) {
		t.Errorf("the returned instruction is invalid")
		return
	}
}

func TestLayer_withCondition_Success(t *testing.T) {
	condition := layers.NewConditionForTest(
		"myName",
		layers.NewInstructionsForTests([]layers.Instruction{
			layers.NewInstructionWithStopForTests(),
		}),
	)

	ins := NewLayerWithConditionForTests(condition)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if !ins.IsCondition() {
		t.Errorf("the layer was expected to contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retCondition := ins.Condition()
	if !reflect.DeepEqual(condition, retCondition) {
		t.Errorf("the returned condition is invalid")
		return
	}
}

func TestLayer_withAssignment_Success(t *testing.T) {
	assignment := layers.NewAssignmentForTests(
		"myName",
		layers.NewAssignableWithBytesForTests(
			layers.NewBytesWithJoinForTests([]string{
				"first",
				"second",
			}),
		),
	)

	ins := NewLayerWithAssignmentForTests(assignment)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if !ins.IsAssignment() {
		t.Errorf("the layer was expected to contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retAssignment := ins.Assignment()
	if !reflect.DeepEqual(assignment, retAssignment) {
		t.Errorf("the returned assignment is invalid")
		return
	}
}

func TestLayer_withAssignable_Success(t *testing.T) {
	assignable := layers.NewAssignableWithBytesForTests(
		layers.NewBytesWithJoinForTests([]string{
			"first",
			"second",
		}),
	)

	ins := NewLayerWithAssignableForTests(assignable)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if !ins.IsAssignable() {
		t.Errorf("the layer was expected to contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retAssignable := ins.Assignable()
	if !reflect.DeepEqual(assignable, retAssignable) {
		t.Errorf("the returned assignable is invalid")
		return
	}
}

func TestLayer_withEngine_Success(t *testing.T) {
	engine := layers.NewEngineWithExecutionForTests(
		layers.NewExecutionForTests("myInput"),
	)

	ins := NewLayerWithEngineForTests(engine)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if !ins.IsEngine() {
		t.Errorf("the layer was expected to contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retEngine := ins.Engine()
	if !reflect.DeepEqual(engine, retEngine) {
		t.Errorf("the returned engine is invalid")
		return
	}
}

func TestLayer_withExecution_Success(t *testing.T) {
	execution := layers.NewExecutionForTests("myInput")

	ins := NewLayerWithExecutionForTests(execution)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if !ins.IsExecution() {
		t.Errorf("the layer was expected to contain an execution")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retExecution := ins.Execution()
	if !reflect.DeepEqual(execution, retExecution) {
		t.Errorf("the returned execution is invalid")
		return
	}
}

func TestLayer_withAssignableResource_Success(t *testing.T) {
	assignableResource := layers.NewAssignableResourceWithCompileForTests(
		"myVariable",
	)

	ins := NewLayerWithAssignableResourceForTests(assignableResource)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if !ins.IsAssignableResource() {
		t.Errorf("the layer was expected to contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retAssignableResource := ins.AssignableResource()
	if !reflect.DeepEqual(assignableResource, retAssignableResource) {
		t.Errorf("the returned assignableResource is invalid")
		return
	}
}

func TestLayer_withBytes_Success(t *testing.T) {
	bytesIns := layers.NewBytesWithJoinForTests([]string{
		"first",
		"second",
	})

	ins := NewLayerWithBytesForTests(bytesIns)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if !ins.IsBytes() {
		t.Errorf("the layer was expected to contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retBytes := ins.Bytes()
	if !reflect.DeepEqual(bytesIns, retBytes) {
		t.Errorf("the returned bytes is invalid")
		return
	}
}

func TestLayer_withIdentity_Success(t *testing.T) {
	identity := layers.NewIdentityWithSignerForTests(
		layers.NewSignerWithSignForTests("mySign"),
	)

	ins := NewLayerWithIdentityForTests(identity)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if !ins.IsIdentity() {
		t.Errorf("the layer was expected to contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retIdentity := ins.Identity()
	if !reflect.DeepEqual(identity, retIdentity) {
		t.Errorf("the returned identity is invalid")
		return
	}
}

func TestLayer_withEncryptor_Success(t *testing.T) {
	encryptor := layers.NewEncryptorWithDecryptForTests("myVariable")
	ins := NewLayerWithEncryptorForTests(encryptor)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if !ins.IsEncryptor() {
		t.Errorf("the layer was expected to contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retEncryptor := ins.Encryptor()
	if !reflect.DeepEqual(encryptor, retEncryptor) {
		t.Errorf("the returned encryptor is invalid")
		return
	}
}

func TestLayer_withSigner_Success(t *testing.T) {
	signer := layers.NewSignerWithSignForTests("mySign")
	ins := NewLayerWithSignerForTests(signer)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if !ins.IsSigner() {
		t.Errorf("the layer was expected to contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retSigner := ins.Signer()
	if !reflect.DeepEqual(signer, retSigner) {
		t.Errorf("the returned signer is invalid")
		return
	}
}

func TestLayer_withSignatureVerify_Success(t *testing.T) {
	signatureVerify := layers.NewSignatureVerifyForTests(
		"mySignature",
		"myMessage",
	)

	ins := NewLayerWithSignatureVerifyForTests(signatureVerify)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if !ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retSignatureVerify := ins.SignatureVerify()
	if !reflect.DeepEqual(signatureVerify, retSignatureVerify) {
		t.Errorf("the returned signatureVerify is invalid")
		return
	}
}

func TestLayer_withVoteVerify_Success(t *testing.T) {
	voteVerify := layers.NewVoteVerifyForTests(
		"myVote",
		"myMessage",
		"myHashedRingVariable",
	)

	ins := NewLayerWithVoteVerifyForTests(voteVerify)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if !ins.IsVoteVerify() {
		t.Errorf("the layer was expected to contain a voteVerify")
		return
	}

	if ins.IsVote() {
		t.Errorf("the layer was expected to NOT contain a vote")
		return
	}

	retVoteVerify := ins.VoteVerify()
	if !reflect.DeepEqual(voteVerify, retVoteVerify) {
		t.Errorf("the returned voteVerify is invalid")
		return
	}
}

func TestLayer_withVote_Success(t *testing.T) {
	vote := layers.NewVoteForTests(
		"myRingVariable",
		"myMessage",
	)

	ins := NewLayerWithVoteForTests(vote)

	if ins.IsLayer() {
		t.Errorf("the layer was expected to NOT contain a layer")
		return
	}

	if ins.IsOutput() {
		t.Errorf("the layer was expected to NOT contain an output")
		return
	}

	if ins.IsKind() {
		t.Errorf("the layer was expected to NOT contain a kind")
		return
	}

	if ins.IsInstruction() {
		t.Errorf("the layer was expected to NOT contain an instruction")
		return
	}

	if ins.IsCondition() {
		t.Errorf("the layer was expected to NOT contain a condition")
		return
	}

	if ins.IsAssignment() {
		t.Errorf("the layer was expected to NOT contain an assignment")
		return
	}

	if ins.IsAssignable() {
		t.Errorf("the layer was expected to NOT contain an assignable")
		return
	}

	if ins.IsEngine() {
		t.Errorf("the layer was expected to NOT contain an engine")
		return
	}

	if ins.IsAssignableResource() {
		t.Errorf("the layer was expected to NOT contain an assignableResource")
		return
	}

	if ins.IsBytes() {
		t.Errorf("the layer was expected to NOT contain a bytes")
		return
	}

	if ins.IsIdentity() {
		t.Errorf("the layer was expected to NOT contain an identity")
		return
	}

	if ins.IsEncryptor() {
		t.Errorf("the layer was expected to NOT contain an encryptor")
		return
	}

	if ins.IsSigner() {
		t.Errorf("the layer was expected to NOT contain a signer")
		return
	}

	if ins.IsSignatureVerify() {
		t.Errorf("the layer was expected to NOT contain a signatureVerify")
		return
	}

	if ins.IsVoteVerify() {
		t.Errorf("the layer was expected to NOT contain a voteVerify")
		return
	}

	if !ins.IsVote() {
		t.Errorf("the layer was expected to contain a vote")
		return
	}

	retVote := ins.Vote()
	if !reflect.DeepEqual(vote, retVote) {
		t.Errorf("the returned vote is invalid")
		return
	}
}

func TestLayer_withoutParam_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
