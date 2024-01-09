package layers

import (
	"errors"

	"steve.care/network/domain/programs/logics/libraries/layers"
)

type builder struct {
	layer              layers.Layer
	output             layers.Output
	kind               layers.Kind
	instruction        layers.Instruction
	condition          layers.Condition
	assignment         layers.Assignment
	assignable         layers.Assignable
	engine             layers.Engine
	execution          layers.Execution
	assignableResource layers.AssignableResource
	bytes              layers.Bytes
	identity           layers.Identity
	encryptor          layers.Encryptor
	signer             layers.Signer
	signatureVerify    layers.SignatureVerify
	voteVerify         layers.VoteVerify
	vote               layers.Vote
}

func createBuilder() Builder {
	out := builder{
		layer:              nil,
		output:             nil,
		kind:               nil,
		instruction:        nil,
		condition:          nil,
		assignment:         nil,
		assignable:         nil,
		engine:             nil,
		execution:          nil,
		assignableResource: nil,
		bytes:              nil,
		identity:           nil,
		encryptor:          nil,
		signer:             nil,
		signatureVerify:    nil,
		voteVerify:         nil,
		vote:               nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithLayer adds a layer to the builder
func (app *builder) WithLayer(layer layers.Layer) Builder {
	app.layer = layer
	return app
}

// WithOutput adds an output to the builder
func (app *builder) WithOutput(output layers.Output) Builder {
	app.output = output
	return app
}

// WithKind adds a kind to the builder
func (app *builder) WithKind(kind layers.Kind) Builder {
	app.kind = kind
	return app
}

// WithInstruction adds an instruction to the builder
func (app *builder) WithInstruction(ins layers.Instruction) Builder {
	app.instruction = ins
	return app
}

// WithCondition adds a condition to the builder
func (app *builder) WithCondition(condition layers.Condition) Builder {
	app.condition = condition
	return app
}

// WithAssignment adds an assignment to the builder
func (app *builder) WithAssignment(assignment layers.Assignment) Builder {
	app.assignment = assignment
	return app
}

// WithAssignable adds an assignable to the builder
func (app *builder) WithAssignable(assignable layers.Assignable) Builder {
	app.assignable = assignable
	return app
}

// WithEngine adds an engine to the builder
func (app *builder) WithEngine(engine layers.Engine) Builder {
	app.engine = engine
	return app
}

// WithExecution adds an execution to the builder
func (app *builder) WithExecution(execution layers.Execution) Builder {
	app.execution = execution
	return app
}

// WithAssignableResource adds an assignableResource to the builder
func (app *builder) WithAssignableResource(assignableResource layers.AssignableResource) Builder {
	app.assignableResource = assignableResource
	return app
}

// WithBytes add bytes to the builder
func (app *builder) WithBytes(bytes layers.Bytes) Builder {
	app.bytes = bytes
	return app
}

// WithIdentity adds identity to the builder
func (app *builder) WithIdentity(identity layers.Identity) Builder {
	app.identity = identity
	return app
}

// WithEncryptor adds encryptor to the builder
func (app *builder) WithEncryptor(encryptor layers.Encryptor) Builder {
	app.encryptor = encryptor
	return app
}

// WithSigner adds signer to the builder
func (app *builder) WithSigner(signer layers.Signer) Builder {
	app.signer = signer
	return app
}

// WithSignatureVerify adds signature verify to the builder
func (app *builder) WithSignatureVerify(sigVerify layers.SignatureVerify) Builder {
	app.signatureVerify = sigVerify
	return app
}

// WithVoteVerify adds voteVerify to the builder
func (app *builder) WithVoteVerify(voteVerify layers.VoteVerify) Builder {
	app.voteVerify = voteVerify
	return app
}

// WithVote adds vote to the builder
func (app *builder) WithVote(vote layers.Vote) Builder {
	app.vote = vote
	return app
}

// Now builds a new Layer instance
func (app *builder) Now() (Layer, error) {
	if app.layer != nil {
		return createLayerWithLayer(app.layer), nil
	}

	if app.output != nil {
		return createLayerWithOutput(app.output), nil
	}

	if app.kind != nil {
		return createLayerWithKind(app.kind), nil
	}

	if app.instruction != nil {
		return createLayerWithInstruction(app.instruction), nil
	}

	if app.condition != nil {
		return createLayerWithCondition(app.condition), nil
	}

	if app.assignment != nil {
		return createLayerWithAssignment(app.assignment), nil
	}

	if app.assignable != nil {
		return createLayerWithAssignable(app.assignable), nil
	}

	if app.engine != nil {
		return createLayerWithEngine(app.engine), nil
	}

	if app.execution != nil {
		return createLayerWithExecution(app.execution), nil
	}

	if app.assignableResource != nil {
		return createLayerWithAssignableResource(app.assignableResource), nil
	}

	if app.bytes != nil {
		return createLayerWithBytes(app.bytes), nil
	}

	if app.identity != nil {
		return createLayerWithIdentity(app.identity), nil
	}

	if app.encryptor != nil {
		return createLayerWithEncryptor(app.encryptor), nil
	}

	if app.signer != nil {
		return createLayerWithSigner(app.signer), nil
	}

	if app.signatureVerify != nil {
		return createLayerWithSignatureVerify(app.signatureVerify), nil
	}

	if app.voteVerify != nil {
		return createLayerWithVoteVerify(app.voteVerify), nil
	}

	if app.vote != nil {
		return createLayerWithVote(app.vote), nil
	}

	return nil, errors.New("the Layer resource is invalid")
}
