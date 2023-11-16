package applications

import (
	"time"

	"steve.care/network/domain/programs"
	"steve.care/network/libraries/blockchains"
	"steve.care/network/libraries/blockchains/blocks"
	"steve.care/network/libraries/blockchains/blocks/queues"
	"steve.care/network/libraries/blockchains/roots"
	"steve.care/network/libraries/blockchains/roots/resolutions"
	"steve.care/network/libraries/commands"
	"steve.care/network/libraries/stacks"
)

type application struct {
	frameFactory           stacks.FrameFactory
	frameBuilder           stacks.FrameBuilder
	frameAssignableBuilder stacks.AssignableBuilder
	blockchainBuilder      blockchains.Builder
	blockchainRepository   blockchains.Repository
	blockchainService      blockchains.Service
	blockBuilder           blocks.Builder
	rootBuilder            roots.Builder
	resolutionBuilder      resolutions.Builder
	queueBuilder           queues.Builder
	queueRepository        queues.Repository
	queueService           queues.Service
	commandsBuilder        commands.Builder
}

func createApplication(
	frameFactory stacks.FrameFactory,
	frameBuilder stacks.FrameBuilder,
	blockchainBuilder blockchains.Builder,
	blockchainRepository blockchains.Repository,
	blockchainService blockchains.Service,
	blockBuilder blocks.Builder,
	rootBuilder roots.Builder,
	resolutionBuilder resolutions.Builder,
	queueBuilder queues.Builder,
	queueRepository queues.Repository,
	queueService queues.Service,
	commandsBuilder commands.Builder,
) Application {
	out := application{
		frameFactory:         frameFactory,
		frameBuilder:         frameBuilder,
		blockchainBuilder:    blockchainBuilder,
		blockchainRepository: blockchainRepository,
		blockchainService:    blockchainService,
		blockBuilder:         blockBuilder,
		rootBuilder:          rootBuilder,
		resolutionBuilder:    resolutionBuilder,
		queueBuilder:         queueBuilder,
		queueRepository:      queueRepository,
		queueService:         queueService,
		commandsBuilder:      commandsBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(programm programs.Program, stack stacks.Stack) (stacks.Stack, error) {
	innerFrame, err := app.createInnerFrame(programm, stack)
	if err != nil {
		return nil, err
	}

	instructions := programm.Instructions()
	err = app.executeInstructions(instructions, innerFrame)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (app *application) executeInstructions(instructions programs.Instructions, frame stacks.Frame) error {
	list := instructions.List()
	for _, oneInstruction := range list {
		err := app.executeInstruction(oneInstruction, frame)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *application) executeInstruction(instruction programs.Instruction, frame stacks.Frame) error {
	if instruction.IsAssignment() {

	}

	if instruction.IsDelete() {

	}

	if instruction.IsClear() {

	}

	if instruction.IsBack() {
		variable := instruction.Back()
		return app.executeBack(variable, frame)
	}

	if instruction.IsCommit() {
		commit := instruction.Commit()
		return app.executeCommit(commit, frame)
	}

	if instruction.IsRollback() {

	}

	init := instruction.Init()
	return app.executeInit(init)
}

func (app *application) executeAssignment(assignment programs.Assignment, frame stacks.Frame) (stacks.Assignable, error) {
	assignable := assignment.Assignable()

	if assignable.IsBegin() {

	}

	if assignable.IsExists() {

	}

	if assignable.IsQueue() {
		variable := assignable.Queue()
		return app.fetchQueue(variable, frame)
	}

	if assignable.IsTransact() {

	}

	return nil, nil
}

func (app *application) fetchQueue(variable string, frame stacks.Frame) (stacks.Assignable, error) {
	pContext, err := frame.FetchUint(variable)
	if err != nil {
		return nil, err
	}

	queue, err := app.queueRepository.Retrieve(*pContext)
	if err != nil {
		return nil, err
	}

	return app.frameAssignableBuilder.Create().
		WithQueue(queue).
		Now()
}

func (app *application) executeCommit(commit programs.Commit, frame stacks.Frame) error {
	contextVariable := commit.Context()
	pContext, err := frame.FetchUint(contextVariable)
	if err != nil {
		return err
	}

	message := commit.Message()

	queue, err := app.queueRepository.Retrieve(*pContext)
	if err != nil {
		return err
	}

	path := queue.Path()
	blockchain, err := app.blockchainRepository.Retrieve(path)
	if err != nil {
		return err
	}

	createdOn := time.Now().UTC()
	commands := queue.Commands()
	rootHash := blockchain.Root().Hash()
	blockBuilder := app.blockBuilder.Create().
		WithMessage(message).
		WithCommands(commands).
		CreatedOn(createdOn).
		WithParent(rootHash)

	if blockchain.HasHead() {
		headHash := blockchain.Head().Hash()
		blockBuilder.WithParent(headHash)
	}

	block, err := blockBuilder.Now()
	if err != nil {
		return err
	}

	return app.blockchainService.Chain(
		blockchain,
		block,
		func() error {
			return app.queueService.Clear(*pContext)
		},
	)
}

func (app *application) executeBack(variable string, frame stacks.Frame) error {
	pContext, err := frame.FetchUint(variable)
	if err != nil {
		return err
	}

	queue, err := app.queueRepository.Retrieve(*pContext)
	if err != nil {
		return err
	}

	commandsList := queue.Commands().List()
	if len(commandsList) <= 1 {
		return app.queueService.Clear(*pContext)
	}

	updatedCommands, err := app.commandsBuilder.Create().
		WithList(commandsList[:len(commandsList)-1]).
		Now()

	if err != nil {
		return err
	}

	path := queue.Path()
	updatedQueue, err := app.queueBuilder.Create().
		WithPath(path).Create().
		WithCommands(updatedCommands).
		Now()

	if err != nil {
		return err
	}

	return app.queueService.Replace(
		*pContext,
		updatedQueue,
	)
}

func (app *application) executeInit(init programs.Init) error {
	rootIns := init.Root()
	affiliate := rootIns.Affiliate()
	fees := rootIns.Fees()
	resolution, err := app.resolutionBuilder.Create().
		WithAffiliate(affiliate).
		WithFees(fees).
		Now()

	if err != nil {
		return err
	}

	createdOn := time.Now().UTC()
	root, err := app.rootBuilder.Create().WithResolution(resolution).CreatedOn(createdOn).Now()
	if err != nil {
		return err
	}

	blockchain, err := app.blockchainBuilder.Create().
		WithRoot(root).
		Now()

	if err != nil {
		return err
	}

	path := init.Path()
	return app.blockchainService.Insert(
		path,
		blockchain,
	)
}

func (app *application) createInnerFrame(programm programs.Program, stack stacks.Stack) (stacks.Frame, error) {
	if !programm.HasParameters() {
		return app.frameFactory.Create(), nil
	}

	last := stack.Last()
	list := []stacks.Assignable{}
	parameters := programm.Parameters()
	for _, oneParameter := range parameters {
		assignable, err := last.Fetch(oneParameter)
		if err != nil {
			// error, param is not in frame
		}

		list = append(list, assignable)
	}

	return app.frameBuilder.Create().
		WihtList(list).
		Now()
}
