package peers

import (
	"bytes"

	applications_blocks "steve.care/network/applications/applications/programs/blocks"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/peers"
	"steve.care/network/domain/programs"
	"steve.care/network/domain/programs/blocks"
)

type application struct {
	blockApplication        applications_blocks.Application
	repository              peers.Repository
	blockApplicationBuilder BlockApplicationBuilder
}

func createApplication(
	blockApplication applications_blocks.Application,
	repository peers.Repository,
	blockApplicationBuilder BlockApplicationBuilder,
) Application {
	out := application{
		blockApplication:        blockApplication,
		repository:              repository,
		blockApplicationBuilder: blockApplicationBuilder,
	}

	return &out
}

// Execute executes the peers application on a program
func (app *application) Execute(program programs.Program) error {
	programHash := program.Hash()
	peers, err := app.repository.Retrieve(programHash)
	if err != nil {
		return err
	}

	peersList := peers.Peers()
	for _, onePeer := range peersList {
		peerBlockApp, err := app.blockApplicationBuilder.Create().WithURL(onePeer).Now()
		if err != nil {
			return err
		}

		// retrieve the peer's program's head block:
		peerHeadBlock, err := peerBlockApp.RetrieveHeadByProgram(programHash)
		if err != nil {
			return err
		}

		// retrieve our program's head block:
		ourHeadBlock, err := app.blockApplication.RetrieveHeadByProgram(programHash)
		if err != nil {
			return err
		}

		// sync the block with our peer:
		return app.syncHeadBlockWithPeer(
			programHash,
			peerHeadBlock,
			ourHeadBlock,
			peerBlockApp,
			[]blocks.Block{},
		)
	}

	return nil
}

func (app *application) syncHeadBlockWithPeer(
	program hash.Hash,
	peerBlock blocks.Block,
	ourBlock blocks.Block,
	peerBlockApp applications_blocks.Application,
	currentList []blocks.Block,
) error {

	// if the block hash is the same:
	if bytes.Equal(
		peerBlock.Hash().Bytes(),
		ourBlock.Hash().Bytes(),
	) {
		// save the list of block if not empty:
		return app.saveBlocks(currentList)
	}

	// if the peer block exists in our database - means we reached our own block spot:
	peerHeadBlockHash := peerBlock.Hash()
	exists := app.blockApplication.Exists(peerHeadBlockHash)
	if exists {
		// save the list of block if not empty:
		return app.saveBlocks(currentList)
	}

	// add the peer block to the list:
	currentList = append(currentList, peerBlock)

	// if the peer has parent:
	peerBlockContent := peerBlock.Content()
	if !peerBlockContent.HasParent() {
		// if we have no block for the program, save the list:
		_, err := app.blockApplication.RetrieveHeadByProgram(program)
		if err != nil {
			// we reached the root without reaching our program's head block so return without updating, we have a different branch:
			return nil
		}

		// our program has no block, so save the list:
		return app.saveBlocks(currentList)
	}

	peerParentBlockHash := peerBlockContent.Parent()
	peerParentBlock, err := peerBlockApp.Retrieve(peerParentBlockHash)
	if err != nil {
		return err
	}

	return app.syncHeadBlockWithPeer(
		program,
		peerParentBlock,
		ourBlock,
		peerBlockApp,
		currentList,
	)
}

func (app *application) saveBlocks(list []blocks.Block) error {
	if len(list) > 0 {
		return nil
	}

	// reverse our blocks list:
	lastIndex := len(list) - 1
	reversed := []blocks.Block{}
	for i := 0; i < len(list); i++ {
		reversed = append(reversed, list[lastIndex-i])
	}

	for _, oneBlock := range reversed {
		err := app.blockApplication.Insert(oneBlock)
		if err != nil {
			return err
		}
	}

	return nil
}
