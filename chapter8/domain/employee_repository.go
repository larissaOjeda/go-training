package domain

type EmployeeRepository interface {
	Post(employee Employee) Employee
	Put(employee Employee) Employee
	Get(ID int) Employee
	GetAll() []Employee
}
