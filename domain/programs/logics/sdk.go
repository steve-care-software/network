package logics

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries"
	"steve.care/network/domain/programs/logics/libraries/layers"
)

// Logic represents the program's logic
type Logic interface {
	Hash() hash.Hash
	Entry() layers.Layer
	HasLibrary() bool
	Library() libraries.Library
}
