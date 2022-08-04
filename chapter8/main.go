package main

import (
	"chapter8/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Employee struct {
	Name string
	Age  int
	Role string
}

var EmployeeSlice []Employee = []Employee{}

func generalEmployeeHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(req.URL.Path)
	splitted_path := strings.Split(req.URL.Path, "/")
	if len(splitted_path) > 2 {
		if req.Method == "GET" {
			getEmployee(w, req)
		}
	}
	if req.Method == "GET" {
		getEmployeesHandler(w, req)
	}
	if req.Method == "POST" {
		createEmployeeHandler(w, req)
	}
}
func getEmployee(w http.ResponseWriter, req *http.Request) {
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

func main() {
	http.HandleFunc("/employees", generalEmployeeHandler)
	http.HandleFunc("/employees/", generalEmployeeHandler)
	repository := database.NewRepository()
	employee := repository.Get(1)
	fmt.Println(employee)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
