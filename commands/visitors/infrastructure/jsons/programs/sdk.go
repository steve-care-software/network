package programs

// Program represents a program
type Program struct {
	Instructions []Instruction `json:"instructions"`
}

// Instruction represents an instruction
type Instruction struct {
	Assignment *Assignment `json:"assignment"`
}

// Assignment represents an assignment
type Assignment struct {
	Name       string     `json:"name"`
	Assignable Assignable `json:"assignable"`
}

// Assignable represents an assignable
type Assignable struct {
	IsListNames bool         `json:"is_list_names"`
	Authorize   *Credentials `json:"authorize"`
	Create      *Credentials `json:"create"`
}

// Credentials represents credentials
type Credentials struct {
	Username string `json:"username"`
	Password []byte `json:"password"`
}
