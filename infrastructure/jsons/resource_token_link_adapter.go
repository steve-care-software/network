package jsons

import (
	"steve.care/network/domain/receipts/commands/links"
	structs_links "steve.care/network/infrastructure/jsons/resources/tokens/links"
)

type resourceTokenLinkAdapter struct {
}

// LinkToStruct converts a link to struct
func (app *resourceTokenLinkAdapter) LinkToStruct(
	ins links.Link,
) structs_links.Link {
	return structs_links.Link{}
}

// StructToLink converts a struct to link
func (app *resourceTokenLinkAdapter) StructToLink(
	ins structs_links.Link,
) (links.Link, error) {
	return nil, nil
}
