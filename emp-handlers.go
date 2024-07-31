package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

//  GET METHOD   [GetEmployees] --------------->
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var  emp Employee 
	json.NewDecoder(r.Body).Decode(&emp)
	// DATABASE SAVE 
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}

//  POST METHOD   [CreateEmployee] ------------------->
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var employees [] Employee
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)
	
}
//  GET METHOD [GetEmployeeByID ] ------------------------->
func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var employee  Employee
	Database.First(&employee, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(employee)
	
}
//  PUT METHOD [UpdateEmployee ] ----------------------->
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var employee  Employee
	Database.First(&employee, mux.Vars(r)["id"])
	json.NewDecoder(r.Body).Decode(&employee)
	Database.Save(&employee)
	json.NewEncoder(w).Encode(employee)

}
//  Delete METHOD [ DeleteEmployee ] ------------------>
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var  emp Employee 
	Database.Delete(&emp, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode("Employee is deleted !!")
}