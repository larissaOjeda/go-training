# Chapter 8 - It’s your turn now

## Initial

Let's create a go project named `employee` and initialize `Go` modules.

## Working with the `sql` package

The purpose of this exercise is to get familiar with the `sql` package of the Go std library.
For this we need to spin up a MySQL instance and create a database with the following table:

```sql
CREATE TABLE `employee` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `full_name` varchar(500) NOT NULL,
  `position` int(11) NOT NULL,
  `salary` decimal(13,4) NOT NULL,
  `joined` datetime NOT NULL,
  `on_probation` bit(1) NOT NULL,
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT=''
```

```bash
docker run --name go-training -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7
```

Let's assume that we have an Employee that needs to be persisted in a RDBMS (MySQL, PostgreSQL, etc.)

```go
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
    Joined      time.Time
    OnProbation bool
    CreatedAt   time.Time
}
```

and the repository interface we would like to implement is:

```go
type Repository interface {
    Employees(ctx context.Context, pos Position) ([]Employee, error)
    Employee(ctx context.Context, id int) (*Employee, error)
    Save(ctx context.Context, e *Employee) error
}
```

What is required?

- Create the repository implementation using a [MySQL driver](https://github.com/go-sql-driver/mysql)
- Create integration test to test:
  - Inserting an Employee
  - Updating an Employee
  - Getting an employee by ID
  - Getting employees by position

Hint: In order to parse times from DB you need to add `?parseTime=true` to the DSN string.

## Working with the `http` and `encoding/json` package

Let's build up on the previous exercise and create a HTTP server that supports the above actions using JSON.
In order to learn how the http package works we should only use std http package instead of Patron, which makes a lot of things easier.

What is required?

- Setup up a http server
- HTTP POST in order to create an employee
- HTTP PUT in order to update an employee
- HTTP GET in order to get an employee by id
- HTTP GET in order to get employees based on their position
- Create tests with a stubbed repository

Hint: Watch out for HTTP content negotiation and response content type.

## Discuss about the project

- The need for a framework like Patron given the boilerplate code that needs to be written
- Do we need an ORM?
- What package layout might be a good option when we add other stuff to the repo like:
  - monitoring artifacts
  - vendor
  - documents
  - deployment
  - scripts
  - etc.

[-> Next&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;: **Resources**](../resources/README.md)  
[<- Previous&nbsp;: **Chapter 7**](../chapter7/README.md)
