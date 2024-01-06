package logics

import (
	"steve.care/network/domain/hash"
	"steve.care/network/domain/programs/logics/libraries"
	"steve.care/network/domain/programs/logics/libraries/layers"
	"steve.care/network/domain/programs/logics/suites"
	"steve.care/network/domain/programs/logics/threads"
)

type logic struct {
	hash    hash.Hash
	entry   layers.Layer
	library libraries.Library
	suites  suites.Suites
	threads threads.Threads
}

func createLogic(
	hash hash.Hash,
	entry layers.Layer,
	library libraries.Library,
) Logic {
	return createLogicInternally(hash, entry, library, nil, nil)
}

func createLogicWithSuites(
	hash hash.Hash,
	entry layers.Layer,
	library libraries.Library,
	suites suites.Suites,
) Logic {
	return createLogicInternally(hash, entry, library, suites, nil)
}

func createLogicWithThreads(
	hash hash.Hash,
	entry layers.Layer,
	library libraries.Library,
	threads threads.Threads,
) Logic {
	return createLogicInternally(hash, entry, library, nil, threads)
}

func createLogicWithSuitesAndThreads(
	hash hash.Hash,
	entry layers.Layer,
	library libraries.Library,
	suites suites.Suites,
	threads threads.Threads,
) Logic {
	return createLogicInternally(hash, entry, library, suites, threads)
}

func createLogicInternally(
	hash hash.Hash,
	entry layers.Layer,
	library libraries.Library,
	suites suites.Suites,
	threads threads.Threads,
) Logic {
	out := logic{
		hash:    hash,
		entry:   entry,
		library: library,
		suites:  suites,
		threads: threads,
	}

	return &out
}

// Hash returns the hash
func (obj *logic) Hash() hash.Hash {
	return obj.hash
}

// Entry returns the entry
func (obj *logic) Entry() layers.Layer {
	return obj.entry
}

// Library returns the library
func (obj *logic) Library() libraries.Library {
	return obj.library
}

// HasSuites returns true if there is suites, false otherwise
func (obj *logic) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *logic) Suites() suites.Suites {
	return obj.suites
}

// HasThreads returns true if there is threads, false otherwise
func (obj *logic) HasThreads() bool {
	return obj.threads != nil
}

// Threads returns the threads, if any
func (obj *logic) Threads() threads.Threads {
	return obj.threads
}
