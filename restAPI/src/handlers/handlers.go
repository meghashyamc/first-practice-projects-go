package handlers

import (
	"delete"
	"fileProcess"
	"list"
	"log"
	"net/http"
	"remove"
	"show"
	"store"
	"strconv"

	"github.com/gorilla/mux"
)

func ReceiveAndProcessFile(w http.ResponseWriter, req *http.Request) {

	moreEmployees := make([]store.Employee, 0)
	ok, err := fileProcess.ParseFileAndStoreListOfEmployees(req.Body, &moreEmployees)

	if ok {
		store.StoreEmployeesByIdDeptAndLoc(&moreEmployees)

		w.Write([]byte("Successfully added these employees:"))

		for _, empl := range moreEmployees {

			w.Write([]byte("\n"))
			w.Write([]byte(empl.Name))

		}

	} else {

		w.Write([]byte(err.Error()))
		log.Println(err.Error())
	}
}

func ListSearch(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	term := params["term"]
	employees := list.ListSearchEmployees(term, &(store.Employees))
	if len(employees) == 0 {
		w.Write([]byte("No employees matched the search term "))
		w.Write([]byte(term))
		log.Println("No employees matched the search term.")
		return
	}

	w.Write([]byte("List of employees matching search term "))
	w.Write([]byte(term))
	w.Write([]byte(":\n\n"))
	writeEmployees(&w, &employees)

}

func ListAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employees := list.ListAllEmployees(&(store.Employees))

	if len(employees) == 0 {

		w.Write([]byte("No employees exist as of now."))
		log.Println("No employees exist as of now.")
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
		log.Println(err.Error())
		return
	}

	if len(employees) == 0 {

		w.Write([]byte("There are no employees in that department."))
		log.Println("There are no employees in that department.")
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

	pin, err := strconv.Atoi(pinString)

	if err != nil {
		w.Write([]byte("The pin code must be a number"))

		log.Println("The pin code must be a number")
		return
	}
	employees := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))

	if len(employees) == 0 {

		w.Write([]byte("There are no employees at that PIN code."))
		log.Println("There are no employees at that PIN code.")
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

	pin, err := strconv.Atoi(pinString)

	if err != nil {
		w.Write([]byte("The pin code must be a number."))
		log.Println("The pin code must be a number.")
		return
	}
	dnString := params["dn"]
	dn, err := strconv.Atoi(dnString)

	if err != nil {
		w.Write([]byte("The door number must be a number."))
		log.Println("The door number must be a number.")
		return
	}

	employeesAtLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))
	employeesAtDoorNo := list.ListEmployeesByDoorNoAtLoc(dn, &employeesAtLoc)

	if len(employeesAtDoorNo) == 0 {

		w.Write([]byte("There are no employees at that door number."))
		log.Println("There are no employees at that door number.")
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
		log.Println("The pin code must be a number")
		return
	}

	employeesAtLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))
	employeesAtStreet := list.ListEmployeesByStreetAtLoc(params["street"], &employeesAtLoc)
	if len(employeesAtStreet) == 0 {

		w.Write([]byte("There are no employees living in that street."))
		log.Println("There are no employees living in that street.")
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
		log.Println("The pin code must be a number")
		return

	}

	employeesAtLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))
	employeesAtLocality := list.ListEmployeesByLocalityAtLoc(params["locality"], &employeesAtLoc)
	if len(employeesAtLocality) == 0 {

		w.Write([]byte("There are no employees living in that locality."))
		log.Println("There are no employees living in that locality.")
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
		log.Println("The id must be a number")
		return
	}

	employees := make([]store.Employee, 0)
	empl := show.ShowEmployeeByID(id, &(store.IdEmpMap))
	employees = append(employees, empl)

	if employees[0].GetID() == 0 {

		w.Write([]byte("There is no employee with that ID."))
		log.Println("There is no employee with that ID.")
		return
	}

	w.Write([]byte("The employee with the ID "))
	w.Write([]byte(stringID))
	w.Write([]byte(" is: \n\n "))

	writeEmployees(&w, &employees)
}

func DeleteAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	numOfEmp := len(store.Employees)
	if numOfEmp == 0 {

		log.Println("There are no employees to delete, so none were deleted.")
		w.Write([]byte("There are no employees to delete, so none were deleted."))
		return
	}
	delete.DeleteFullEmployeeList(&(store.Employees))
	delete.DeleteFullIDempMap(&(store.IdEmpMap))
	delete.DeleteFullDeptEmpMap(&(store.DeptEmpMap))
	delete.DeleteFullLocEmpMap(&(store.LocEmpMap))

	w.Write([]byte(strconv.Itoa(numOfEmp)))
	w.Write([]byte(" employee(s) was/were deleted successfully."))

}

func DeleteByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	stringID := params["id"]
	id, err := strconv.Atoi(stringID)

	if err != nil {
		panic("The id must be a number")
	}

	empl, ok := (store.IdEmpMap)[id]
	if !ok {

		log.Println("Could not delete: an employee with that ID does not exist.")
		w.Write([]byte("Could not delete: an employee with that ID does not exist."))
		return
	}
	delete.DeleteEmployeeByIDFromList(id, &(store.Employees))
	delete.DeleteEmployeeByIDFromDeptMap(id, empl.GetDept(), &(store.DeptEmpMap))
	delete.DeleteEmployeeByIDFromLocMap(id, empl.GetPins(), &(store.LocEmpMap))
	delete.DeleteEmployeeByIDFromIdMap(id, &(store.IdEmpMap))

	w.Write([]byte("Deleted the employee with ID "))
	w.Write([]byte(stringID))
	w.Write([]byte(" successfully."))

}

func RemoveByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	stringID := params["id"]

	id, err := strconv.Atoi(stringID)

	if err != nil {
		panic("The id must be a number")
	}

	empl, ok := (store.IdEmpMap)[id]

	if !ok {

		log.Println("Could not remove: an employee with that ID does not exist.")
		w.Write([]byte("Could not remove: an employee with that ID does not exist."))
		return
	}

	if empl.There == false {
		log.Println("Did not remove this employee as he/she was already removed.")
		w.Write([]byte("Did not remove this employee as he/she was already removed."))
		return

	}
	remove.RemoveEmployeesFromList(id, &(store.Employees))
	remove.RemoveEmployeesFromIdEmpMap(id, &(store.IdEmpMap))
	remove.RemoveEmployeesFromDeptEmpMap(id, &(store.DeptEmpMap))
	remove.RemoveEmployeesFromLocEmpMap(id, &(store.LocEmpMap))
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
