package types

import "steve.care/network/domain/schemas/roots/groups/resources/fields/types/dependencies"

type typ struct {
	pKind      *uint8
	dependency dependencies.Dependency
}

func createTypeWithKind(
	pKind *uint8,
) Type {
	return createTypeInternally(pKind, nil)
}

func createTypeWithDependency(
	dependency dependencies.Dependency,
) Type {
	return createTypeInternally(nil, dependency)
}

func createTypeInternally(
	pKind *uint8,
	dependency dependencies.Dependency,
) Type {
	out := typ{
		pKind:      pKind,
		dependency: dependency,
	}

	return &out
}

// IsKind returns true if there is a kind, false otherwise
func (obj *typ) IsKind() bool {
	return obj.pKind != nil
}

// Kind returns the kind, if any
func (obj *typ) Kind() *uint8 {
	return obj.pKind
}

// IsDependency returns true if there is a dependency, false otherwise
func (obj *typ) IsDependency() bool {
	return obj.dependency != nil
}

// Dependency returns the dependency, if any
func (obj *typ) Dependency() dependencies.Dependency {
	return obj.dependency
}
