package programs

// Program represents a program
type Program interface {
}

// Assignable represents the assignable
type Assignable interface {
	IsExists() bool
	Exists() []string
	IsList() bool
	List() []string
	IsDir() bool
	Dir() []string
	IsExecute() bool
	Execute() []string
}
