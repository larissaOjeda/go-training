package database

import (
	"chapter8/domain"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MySqlRepository struct {
	db *sql.DB
}

func NewRepository() domain.EmployeeRepository {
	db, err := sql.Open("mysql", "go-training-user:123456@tcp(localhost:3306)/go-training")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &MySqlRepository{db: db}
}

func (repository *MySqlRepository) Get(ID int) domain.Employee {
	statement, err := repository.db.Prepare("Select * from employees where id = ?")
	if err != nil {
		log.Fatal(err.Error())
	}
	var employee domain.Employee
	err = statement.QueryRow(ID).Scan(&employee.ID, &employee.FullName, &employee.Position, &employee.Salary, &employee.Joined, &employee.OnProbation, &employee.CreatedAt)
	if err != nil {
		log.Fatal(err.Error())
	}ss
	return employee

}
