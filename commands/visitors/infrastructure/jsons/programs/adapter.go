package programs

import (
	"encoding/json"

	"steve.care/network/commands/visitors/domain/programs"
)

type adapter struct {
	builder             programs.Builder
	instructionsBuilder programs.InstructionsBuilder
	instructionBuilder  programs.InstructionBuilder
	assignmentBuilder   programs.AssignmentBuilder
	assignableBuilder   programs.AssignableBuilder
	credentialsBuilder  programs.CredentialsBuilder
}

func createAdapter(
	builder programs.Builder,
	instructionsBuilder programs.InstructionsBuilder,
	instructionBuilder programs.InstructionBuilder,
	assignmentBuilder programs.AssignmentBuilder,
	assignableBuilder programs.AssignableBuilder,
	credentialsBuilder programs.CredentialsBuilder,
) programs.Adapter {
	out := adapter{
		builder:             builder,
		instructionsBuilder: instructionsBuilder,
		instructionBuilder:  instructionBuilder,
		assignmentBuilder:   assignmentBuilder,
		assignableBuilder:   assignableBuilder,
		credentialsBuilder:  credentialsBuilder,
	}

	return &out
}

// ToBytes converts a Program instance to json bytes
func (app *adapter) ToBytes(ins programs.Program) ([]byte, error) {
	stList := []Instruction{}
	list := ins.Instructions().List()
	for _, oneInstruction := range list {
		stInstruction := app.instructionToStruct(oneInstruction)
		stList = append(stList, stInstruction)
	}

	stProgram := Program{
		Instructions: stList,
	}

	return json.Marshal(stProgram)
}

func (app *adapter) instructionToStruct(ins programs.Instruction) Instruction {
	assignment := ins.Assignment()
	stAssignment := app.assignmentToStruct(assignment)
	return Instruction{
		Assignment: &stAssignment,
	}
}

func (app *adapter) assignmentToStruct(ins programs.Assignment) Assignment {
	assignable := ins.Assignable()
	stAssignable := app.assignableToStruct(assignable)
	return Assignment{
		Name:       ins.Name(),
		Assignable: stAssignable,
	}
}

func (app *adapter) assignableToStruct(ins programs.Assignable) Assignable {
	if ins.IsListNames() {
		return Assignable{
			IsListNames: true,
		}
	}

	if ins.IsCreate() {
		create := ins.Create()
		stCredentials := app.credentialsToStruct(create)
		return Assignable{
			Create: &stCredentials,
		}
	}

	authorize := ins.Authorize()
	stAuthorize := app.credentialsToStruct(authorize)
	return Assignable{
		Authorize: &stAuthorize,
	}
}

func (app *adapter) credentialsToStruct(ins programs.Credentials) Credentials {
	return Credentials{
		Username: ins.Username(),
		Password: ins.Password(),
	}
}

// ToInstance converts json bytes to a program instance
func (app *adapter) ToInstance(bytes []byte) (programs.Program, error) {
	ins := new(Program)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.structsToProgram(*ins)
}

func (app *adapter) structsToProgram(ins Program) (programs.Program, error) {
	instructions, err := app.structsToInstruction(ins.Instructions)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithInstructions(instructions).
		Now()
}

func (app *adapter) structsToInstruction(stList []Instruction) (programs.Instructions, error) {
	list := []programs.Instruction{}
	for _, oneStIns := range stList {
		ins, err := app.structToInstruction(oneStIns)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.instructionsBuilder.Create().
		WithList(list).
		Now()
}

func (app *adapter) structToInstruction(ins Instruction) (programs.Instruction, error) {
	builder := app.instructionBuilder.Create()
	if ins.Assignment != nil {
		assignment, err := app.structToAssignment(*ins.Assignment)
		if err != nil {
			return nil, err
		}

		builder.WithAssignment(assignment)
	}

	return builder.Now()
}

func (app *adapter) structToAssignment(ins Assignment) (programs.Assignment, error) {
	assignable, err := app.structToAssignable(ins.Assignable)
	if err != nil {
		return nil, err
	}

	return app.assignmentBuilder.Create().
		WithName(ins.Name).
		WithAssignable(assignable).
		Now()
}

func (app *adapter) structToAssignable(ins Assignable) (programs.Assignable, error) {
	builder := app.assignableBuilder.Create()
	if ins.IsListNames {
		builder.IsListNames()
	}

	if ins.Create != nil {
		credentials, err := app.structToCredentials(*ins.Create)
		if err != nil {
			return nil, err
		}

		builder.WithCreate(credentials)
	}

	if ins.Authorize != nil {
		credentials, err := app.structToCredentials(*ins.Authorize)
		if err != nil {
			return nil, err
		}

		builder.WithAuthorize(credentials)
	}

	return builder.Now()
}

func (app *adapter) structToCredentials(ins Credentials) (programs.Credentials, error) {
	return app.credentialsBuilder.Create().
		WithUsername(ins.Username).
		WithPassword(ins.Password).
		Now()
}
