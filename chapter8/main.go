package main

import (
	"chapter8/database"
	"chapter8/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Employee struct {
	Name string
	Age  int
	Role string
}

var EmployeeSlice []Employee = []Employee{}

func generalEmployeeHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	splitted_path := strings.Split(req.URL.Path, "/")
	if len(splitted_path) > 2 {
		if req.Method == "GET" {
			getEmployeeHandler(w, req)
		}
	}
	if req.Method == "GET" {
		getEmployeesHandler(w, req)
	}
	if req.Method == "POST" {
		createEmployeeHandler(w, req)
	}
}
func getEmployeeHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(EmployeeSlice)
	//get the id again
	//repository := database.NewRepository()
	//employee := repository.Get(id)
}

func getEmployeesHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(EmployeeSlice)
}
func createEmployeeHandler(w http.ResponseWriter, req *http.Request) {
	request, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var employee Employee
	err = json.Unmarshal(request, &employee)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	EmployeeSlice = append(EmployeeSlice, employee)
	fmt.Println(EmployeeSlice)
	w.WriteHeader(http.StatusCreated) //change the http status
}

func updateEmployeeHandler(w http.ResponseWriter, req *http.Request) {
	request, err := ioutil.ReadAll(req.Body)

}

func main() {
	http.HandleFunc("/employees", generalEmployeeHandler)
	http.HandleFunc("/employees/", generalEmployeeHandler)
	repository := database.NewRepository()
	employee := repository.Get(1)
	fmt.Println(employee)
	employees := repository.GetAll()
	fmt.Println(employees)
	newEmployee := domain.Employee{ID: 1234, FullName: "Juan Arroyo", Position: 1, Salary: 4343.43, Joined: time.Now(), OnProbation: 0, CreatedAt: time.Now()}
	newEmployee := repository.Post(newEmployee)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
