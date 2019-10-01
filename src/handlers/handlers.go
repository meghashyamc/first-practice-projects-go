package handlers

import (
	"delete"
	"encoding/json"
	"fileProcess"
	"kiplog"
	"list"
	"net/http"
	"remove"
	"show"
	"store"
	"strconv"

	"github.com/gorilla/mux"
)

func ReceiveAndProcessFile(w http.ResponseWriter, req *http.Request) {

	moreEmployees := make([]store.Employee, 0)
	_, err := fileProcess.ParseFileAndStoreListOfEmployees(req.Body, &moreEmployees)

	if err != nil {
		http.Error(w, err.Error(), 400)
		kiplog.HTTPLog(err.Error())
	} else {
		store.StoreEmployeesByIdDeptAndLoc(&moreEmployees)
		w.Write([]byte("Success: The following employees were added:\n\n"))
		bytes, _ := json.MarshalIndent(&moreEmployees, "", "    ")

		w.Write(bytes)

	}
}

func ListSearch(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	term := params["term"]
	employees, err := list.ListSearchEmployees(term, &(store.Employees))
	if err != nil {
		http.Error(w, err.Error(), 400)
		kiplog.HTTPLog(err.Error())
		return
	}
	json.NewEncoder(w).Encode(&employees)

}

func ListAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employees, err := list.ListAllEmployees(&(store.Employees))

	if err != nil {

		http.Error(w, err.Error(), 400)
		kiplog.HTTPLog(err.Error())
		return

	}

	json.NewEncoder(w).Encode(&employees)

}

func ListByDept(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	dept := params["name"]
	employees, err := list.ListEmployeesByDept(dept, &(store.DeptEmpMap))

	if err != nil {

		http.Error(w, err.Error(), 400)
		kiplog.HTTPLog(err.Error())
		return
	}

	json.NewEncoder(w).Encode(&employees)

}

func ListByLoc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pinString := params["pin"]

	pin, err1 := strconv.Atoi(pinString)

	if err1 != nil {
		http.Error(w, "The pin code must be a number", 400)

		kiplog.HTTPLog("The pin code must be a number")
		return
	}
	employees, err := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))

	if err != nil {

		http.Error(w, err.Error(), 400)
		kiplog.HTTPLog(err.Error())
		return

	}

	json.NewEncoder(w).Encode(&employees)

}

func ListByLocDoorNo(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pinString := params["pin"]

	pin, err1 := strconv.Atoi(pinString)

	if err1 != nil {
		http.Error(w, "The pin code must be a number.", 400)
		kiplog.HTTPLog("The pin code must be a number.")
		return
	}
	dnString := params["dn"]
	dn, err2 := strconv.Atoi(dnString)

	if err2 != nil {
		http.Error(w, "The door number must be a number.", 400)
		kiplog.HTTPLog("The door number must be a number.")
		return
	}

	employeesAtLoc, errLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))
	if errLoc != nil {

		http.Error(w, errLoc.Error(), 400)

		kiplog.HTTPLog(errLoc.Error())
		return

	}

	employeesAtDoorNo, errDoor := list.ListEmployeesByDoorNoAtLoc(dn, &employeesAtLoc)

	if errDoor != nil {

		http.Error(w, errDoor.Error(), 400)
		kiplog.HTTPLog(errDoor.Error())

		return

	}

	json.NewEncoder(w).Encode(&employeesAtDoorNo)

}

func ListByLocStreet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pinString := params["pin"]

	pin, err := strconv.Atoi(pinString)

	if err != nil {
		http.Error(w, "The pin code must be a number", 400)
		kiplog.HTTPLog("The pin code must be a number")
		return
	}

	employeesAtLoc, errLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))

	if errLoc != nil {

		http.Error(w, errLoc.Error(), 400)
		kiplog.HTTPLog(errLoc.Error())
		return

	}

	employeesAtStreet, errStreet := list.ListEmployeesByStreetAtLoc(params["street"], &employeesAtLoc)

	if errStreet != nil {

		http.Error(w, errStreet.Error(), 400)
		kiplog.HTTPLog(errStreet.Error())
		return

	}
	json.NewEncoder(w).Encode(&employeesAtStreet)

}

func ListByLocLocality(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pinString := params["pin"]

	pin, err := strconv.Atoi(pinString)

	if err != nil {
		http.Error(w, "The pin code must be a number", 400)
		kiplog.HTTPLog("The pin code must be a number")
		return

	}

	employeesAtLoc, errLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))

	if errLoc != nil {

		http.Error(w, errLoc.Error(), 400)
		kiplog.HTTPLog(errLoc.Error())
		return

	}
	employeesAtLocality, errLocality := list.ListEmployeesByLocalityAtLoc(params["locality"], &employeesAtLoc)
	if errLocality != nil {

		http.Error(w, errLocality.Error(), 400)
		kiplog.HTTPLog(errLocality.Error())
		return

	}
	json.NewEncoder(w).Encode(&employeesAtLocality)

}

func ShowByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	stringID := params["id"]

	id, err := strconv.Atoi(stringID)

	if err != nil {
		http.Error(w, "The id must be a number", 400)
		kiplog.HTTPLog("The id must be a number")
		return
	}

	employees := make([]store.Employee, 0)
	empl, errShow := show.ShowEmployeeByID(id, &(store.IdEmpMap))

	if errShow != nil {

		http.Error(w, errShow.Error(), 400)
		kiplog.HTTPLog(errShow.Error())
		return

	}

	employees = append(employees, empl)
	json.NewEncoder(w).Encode(&employees)

}

func DeleteAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	numOfEmp := len(store.Employees)

	err := delete.DeleteFromEverywhere()

	if err != nil {
		http.Error(w, err.Error(), 400)
		kiplog.HTTPLog(err.Error())
		return

	}

	w.Write([]byte(strconv.Itoa(numOfEmp)))
	w.Write([]byte(" employee(s) was/were deleted successfully."))

}

func DeleteByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	stringID := params["id"]
	id, err1 := strconv.Atoi(stringID)

	if err1 != nil {
		http.Error(w, "The id must be a number", 400)
		kiplog.HTTPLog("The id must be a number")
	}
	err := delete.DeleteByIDFromEverywhere(id)

	if err != nil {

		kiplog.HTTPLog(err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("Deleted the employee with ID "))
	w.Write([]byte(stringID))
	w.Write([]byte(" successfully."))

}

func RemoveByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	stringID := params["id"]

	id, err1 := strconv.Atoi(stringID)

	if err1 != nil {
		http.Error(w, "The id must be a number", 400)
		kiplog.HTTPLog("The id must be a number")
	}

	err := remove.RemoveEmployeesByIDEverywhere(id)

	if err != nil {

		kiplog.HTTPLog(err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("Removed the employee with ID "))
	w.Write([]byte(stringID))
	w.Write([]byte(" successfully."))
}
