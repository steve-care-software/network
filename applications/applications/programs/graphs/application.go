package graphs

import (
	"errors"
	"fmt"
	"strings"

	"steve.care/network/applications/applications/programs/blocks"
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs"
)

type application struct {
	blockApplication blocks.Application
	repository       programs.Repository
	service          programs.Service
	builder          programs.Builder
	metadataBuilder  programs.MetaDataBuilder
}

func createApplication(
	blockApplication blocks.Application,
	repository programs.Repository,
	service programs.Service,
	builder programs.Builder,
	metadataBuilder programs.MetaDataBuilder,
) Application {
	out := application{
		blockApplication: blockApplication,
		repository:       repository,
		service:          service,
		builder:          builder,
		metadataBuilder:  metadataBuilder,
	}

	return &out
}

// List returns the list of all program hashes at their latest height
func (app *application) List(pIsActive *bool) ([]hash.Hash, error) {
	return app.repository.List(pIsActive)
}

// Children returns the children's program name by its path
func (app *application) Children(path []string, pIsActive *bool) ([]string, error) {
	return app.repository.Children(path, pIsActive)
}

// Retrieve returns the program by hash
func (app *application) Retrieve(hash hash.Hash) (programs.Program, error) {
	return app.repository.Retrieve(hash)
}

// Revision returns the program based on its path and height
func (app *application) Revision(path []string, height uint) (programs.Program, error) {
	return app.repository.Revision(path, height)
}

// Current returns the program based on its path at the latest height
func (app *application) Current(path []string) (programs.Program, error) {
	return app.repository.Current(path)
}

// Insert inserts a new program
func (app *application) Insert(path []string, description string) error {
	builder := app.builder.Create().WithDescription(description)
	if len(path) > 0 {
		lastIndex := len(path) - 1
		retProgram, err := app.Current(path[0:lastIndex])
		if err != nil {
			return err
		}

		metaData, err := app.metadataBuilder.Create().
			WithName(path[lastIndex]).
			WithParent(retProgram).
			Now()

		if err != nil {
			return err
		}

		builder.WithMetaData(metaData)
	}

	newProgram, err := builder.Now()
	if err != nil {
		return err
	}

	return app.service.Insert(newProgram)
}

// Rewind rewinds the program
func (app *application) Rewind(path []string) error {
	retProgram, err := app.Current(path)
	if err != nil {
		return err
	}

	if !retProgram.HasHead() {
		str := fmt.Sprintf("the program (path: %s) cannot rewind because it does NOT contain an head", strings.Join(path, "/"))
		return errors.New(str)
	}

	headContent := retProgram.Head().Content()
	if !headContent.HasParent() {
		str := fmt.Sprintf("the program (path: %s) cannot rewind because it's head block does NOT contain a parent block", strings.Join(path, "/"))
		return errors.New(str)
	}

	parentHash := headContent.Parent()
	newHeadBlock, err := app.blockApplication.Retrieve(parentHash)
	if err != nil {
		return err
	}

	description := retProgram.Description()
	builder := app.builder.Create().WithDescription(description).WithHead(newHeadBlock)
	if retProgram.HasLogic() {
		logic := retProgram.Logic()
		builder.WithLogic(logic)
	}

	if retProgram.HasMetaData() {
		metaData := retProgram.MetaData()
		builder.WithMetaData(metaData)
	}

	ins, err := builder.Now()
	if err != nil {
		return err
	}

	return app.service.Insert(ins)
}

// Delete deletes a program
func (app *application) Delete(hash hash.Hash) error {
	retProgram, err := app.Retrieve(hash)
	if err != nil {
		return err
	}

	return app.service.Delete(retProgram)
}
