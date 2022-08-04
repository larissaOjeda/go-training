package domain

type EmployeeRepository interface {
	Get(ID int) Employee
}
