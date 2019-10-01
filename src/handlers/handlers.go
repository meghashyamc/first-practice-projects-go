package handlers

import (
	"delete"
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

		w.Write([]byte(err.Error()))
		kiplog.HTTPLog(err.Error())
	} else {
		store.StoreEmployeesByIdDeptAndLoc(&moreEmployees)

		w.Write([]byte("Successfully added these employees:"))

		for _, empl := range moreEmployees {

			w.Write([]byte("\n"))
			w.Write([]byte(empl.Name))

		}

	}
}

func ListSearch(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	term := params["term"]
	employees, err := list.ListSearchEmployees(term, &(store.Employees))
	if err != nil {
		w.Write([]byte(err.Error()))
		kiplog.HTTPLog(err.Error())
		return
	}
	w.Write([]byte("List of employees matching search term "))
	w.Write([]byte(term))
	w.Write([]byte(":\n\n"))
	writeEmployees(&w, &employees)

}

func ListAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employees, err := list.ListAllEmployees(&(store.Employees))

	if err != nil {

		w.Write([]byte(err.Error()))
		kiplog.HTTPLog(err.Error())
		return

	}
	w.Write([]byte("List of all employees:\n\n"))

	writeEmployees(&w, &employees)

}

func ListByDept(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	dept := params["name"]
	employees, err := list.ListEmployeesByDept(dept, &(store.DeptEmpMap))

	if err != nil {

		w.Write([]byte(err.Error()))
		kiplog.HTTPLog(err.Error())
		return
	}

	w.Write([]byte("List of all employees in the department "))
	w.Write([]byte(dept))
	w.Write([]byte(":\n\n"))

	writeEmployees(&w, &employees)

}

func ListByLoc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pinString := params["pin"]

	pin, err1 := strconv.Atoi(pinString)

	if err1 != nil {
		w.Write([]byte("The pin code must be a number"))

		kiplog.HTTPLog("The pin code must be a number")
		return
	}
	employees, err := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))

	if err != nil {

		w.Write([]byte(err.Error()))
		kiplog.HTTPLog(err.Error())
		return

	}
	w.Write([]byte("List of all employees at the PIN code "))
	w.Write([]byte(pinString))
	w.Write([]byte(":\n\n"))
	writeEmployees(&w, &employees)

}

func ListByLocDoorNo(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pinString := params["pin"]

	pin, err1 := strconv.Atoi(pinString)

	if err1 != nil {
		w.Write([]byte("The pin code must be a number."))
		kiplog.HTTPLog("The pin code must be a number.")
		return
	}
	dnString := params["dn"]
	dn, err2 := strconv.Atoi(dnString)

	if err2 != nil {
		w.Write([]byte("The door number must be a number."))
		kiplog.HTTPLog("The door number must be a number.")
		return
	}

	employeesAtLoc, errLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))
	if errLoc != nil {

		w.Write([]byte(errLoc.Error()))
		kiplog.HTTPLog(errLoc.Error())
		return

	}

	employeesAtDoorNo, errDoor := list.ListEmployeesByDoorNoAtLoc(dn, &employeesAtLoc)

	if errDoor != nil {

		w.Write([]byte(errDoor.Error()))
		kiplog.HTTPLog(errDoor.Error())

		return

	}

	w.Write([]byte("List of all employees at the PIN code "))
	w.Write([]byte(pinString))
	w.Write([]byte(", and door number "))
	w.Write([]byte(dnString))
	w.Write([]byte(":\n\n"))
	writeEmployees(&w, &employeesAtDoorNo)
}

func ListByLocStreet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pinString := params["pin"]

	pin, err := strconv.Atoi(pinString)

	if err != nil {
		w.Write([]byte("The pin code must be a number"))
		kiplog.HTTPLog("The pin code must be a number")
		return
	}

	employeesAtLoc, errLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))

	if errLoc != nil {

		w.Write([]byte(errLoc.Error()))
		kiplog.HTTPLog(errLoc.Error())
		return

	}

	employeesAtStreet, errStreet := list.ListEmployeesByStreetAtLoc(params["street"], &employeesAtLoc)

	if errStreet != nil {

		w.Write([]byte(errStreet.Error()))
		kiplog.HTTPLog(errStreet.Error())
		return

	}
	w.Write([]byte("List of all employees at the PIN code "))
	w.Write([]byte(pinString))
	w.Write([]byte(", and door number "))
	w.Write([]byte(params["street"]))
	w.Write([]byte(":\n\n"))
	writeEmployees(&w, &employeesAtStreet)

}

func ListByLocLocality(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pinString := params["pin"]

	pin, err := strconv.Atoi(pinString)

	if err != nil {
		w.Write([]byte("The pin code must be a number"))
		kiplog.HTTPLog("The pin code must be a number")
		return

	}

	employeesAtLoc, errLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))

	if errLoc != nil {

		w.Write([]byte(errLoc.Error()))
		kiplog.HTTPLog(errLoc.Error())
		return

	}
	employeesAtLocality, errLocality := list.ListEmployeesByLocalityAtLoc(params["locality"], &employeesAtLoc)
	if errLocality != nil {

		w.Write([]byte(errLocality.Error()))
		kiplog.HTTPLog(errLocality.Error())
		return

	}
	w.Write([]byte("List of all employees at the PIN code "))
	w.Write([]byte(pinString))
	w.Write([]byte(", and locality "))
	w.Write([]byte(params["locality"]))
	w.Write([]byte(":\n\n"))
	writeEmployees(&w, &employeesAtLocality)

}

func ShowByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	stringID := params["id"]

	id, err := strconv.Atoi(stringID)

	if err != nil {
		w.Write([]byte("The id must be a number"))
		kiplog.HTTPLog("The id must be a number")
		return
	}

	employees := make([]store.Employee, 0)
	empl, errShow := show.ShowEmployeeByID(id, &(store.IdEmpMap))

	if errShow != nil {

		w.Write([]byte(errShow.Error()))
		kiplog.HTTPLog(errShow.Error())
		return

	}

	employees = append(employees, empl)

	w.Write([]byte("The employee with the ID "))
	w.Write([]byte(stringID))
	w.Write([]byte(" is: \n\n "))

	writeEmployees(&w, &employees)
}

func DeleteAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	numOfEmp := len(store.Employees)

	err := delete.DeleteFromEverywhere()

	if err != nil {
		w.Write([]byte(err.Error()))
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
		w.Write([]byte("The id must be a number"))
		kiplog.HTTPLog("The id must be a number")
	}
	err := delete.DeleteByIDFromEverywhere(id)

	if err != nil {

		kiplog.HTTPLog(err.Error())
		w.Write([]byte(err.Error()))
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
		w.Write([]byte("The id must be a number"))
		kiplog.HTTPLog("The id must be a number")
	}

	err := remove.RemoveEmployeesByIDEverywhere(id)

	if err != nil {

		kiplog.HTTPLog(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Removed the employee with ID "))
	w.Write([]byte(stringID))
	w.Write([]byte(" successfully."))
}

func writeEmployees(w *http.ResponseWriter, employees *([]store.Employee)) {

	for _, empl := range *employees {
		(*w).Write([]byte("ID: "))
		(*w).Write([]byte(strconv.Itoa(empl.GetID())))
		(*w).Write([]byte("\n"))
		(*w).Write([]byte("Name: "))
		(*w).Write([]byte(empl.Name))
		(*w).Write([]byte("\n"))
		(*w).Write([]byte("Departments: "))

		for _, dept := range empl.GetDept() {
			(*w).Write([]byte(dept))
			(*w).Write([]byte(" "))
		}

		(*w).Write([]byte("\n"))
		(*w).Write([]byte("Addresses: "))
		(*w).Write([]byte("\n"))

		for _, addr := range empl.Addresses {

			(*w).Write([]byte(strconv.Itoa(addr.GetDoorNo())))
			(*w).Write([]byte(", "))
			(*w).Write([]byte(addr.GetStreet()))
			(*w).Write([]byte(", "))
			(*w).Write([]byte(addr.GetLocality()))
			(*w).Write([]byte(", "))
			(*w).Write([]byte(strconv.Itoa(addr.PIN)))
			(*w).Write([]byte("\n"))

		}

		(*w).Write([]byte("\n\n"))
	}
}
