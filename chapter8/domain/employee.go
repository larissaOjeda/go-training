package domain

type Position int

const (
	Undetermined Position = iota
	Junior
	Senior
	Manager
	CEO
)

type Employee struct {
	ID          int
	FullName    string
	Position    Position
	Salary      float64
	Joined      string
	OnProbation bool
	CreatedAt   string
}
