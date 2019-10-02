package handlers

import (
	"delete"
	"encoding/json"
	"fileProcess"
	"github.com/gorilla/mux"
	"kiplog"
	"list"
	"net/http"
	"remove"
	"show"
	"store"
	"strconv"
)

const SUCCESS = "Success"

type successResponseList struct {
	Message string
	Data    *([]store.Employee)
}

type successResponse struct {
	Message string
	Data    string
}

type failureResponse struct {
	Message string
	Data    string
}

func ReceiveAndProcessFile(w http.ResponseWriter, req *http.Request) {

	moreEmployees := make([]store.Employee, 0)
	_, err := fileProcess.ParseFileAndStoreListOfEmployees(req.Body, &moreEmployees)

	if err != nil {

		writeFailureResponse(&w, http.StatusBadRequest, err)

	} else {
		store.StoreEmployeesByIdDeptAndLoc(&moreEmployees)

		writeSuccessResponseList(&w, &moreEmployees)
	}
}

func ListSearch(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	term := params["term"]
	employees, err := list.ListSearchEmployees(term, &(store.Employees))
	if err != nil {
		writeFailureResponse(&w, http.StatusNotFound, err)

		return
	}
	writeSuccessResponseList(&w, &employees)

}

func ListAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employees, err := list.ListAllEmployees(&(store.Employees))

	if err != nil {

		writeFailureResponse(&w, http.StatusNotFound, err)

		return

	}

	writeSuccessResponseList(&w, &employees)

}

func ListByDept(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	dept := params["name"]
	employees, err := list.ListEmployeesByDept(dept, &(store.DeptEmpMap))

	if err != nil {

		writeFailureResponse(&w, http.StatusNotFound, err)

		return
	}

	writeSuccessResponseList(&w, &employees)

}

func ListByLoc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pinString := params["pin"]

	pin, err1 := strconv.Atoi(pinString)

	if err1 != nil {
		writeFailureResponseStr(&w, http.StatusBadRequest, "Pin code must be a number")
		return
	}
	employees, err := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))

	if err != nil {

		writeFailureResponse(&w, http.StatusNotFound, err)

		return

	}

	writeSuccessResponseList(&w, &employees)

}

func ListByLocDoorNo(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pinString := params["pin"]

	pin, err1 := strconv.Atoi(pinString)

	if err1 != nil {
		writeFailureResponseStr(&w, http.StatusBadRequest, "Pin code must be a number")

		return
	}
	dnString := params["dn"]
	dn, err2 := strconv.Atoi(dnString)

	if err2 != nil {
		writeFailureResponseStr(&w, http.StatusBadRequest, "Door number must be a number")
		return
	}

	employeesAtLoc, errLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))
	if errLoc != nil {

		writeFailureResponse(&w, http.StatusNotFound, errLoc)

		return

	}

	employeesAtDoorNo, errDoor := list.ListEmployeesByDoorNoAtLoc(dn, &employeesAtLoc)

	if errDoor != nil {

		writeFailureResponse(&w, http.StatusNotFound, errDoor)

		return

	}

	writeSuccessResponseList(&w, &employeesAtDoorNo)

}

func ListByLocStreet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pinString := params["pin"]

	pin, err := strconv.Atoi(pinString)

	if err != nil {
		writeFailureResponseStr(&w, http.StatusBadRequest, "Pin code must be a number")
		return
	}

	employeesAtLoc, errLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))

	if errLoc != nil {

		writeFailureResponse(&w, http.StatusNotFound, errLoc)

		return

	}

	employeesAtStreet, errStreet := list.ListEmployeesByStreetAtLoc(params["street"], &employeesAtLoc)

	if errStreet != nil {

		writeFailureResponse(&w, http.StatusNotFound, errStreet)

		return

	}
	writeSuccessResponseList(&w, &employeesAtStreet)

}

func ListByLocLocality(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pinString := params["pin"]

	pin, err := strconv.Atoi(pinString)

	if err != nil {
		writeFailureResponseStr(&w, http.StatusBadRequest, "Pin code must be a number")
		return

	}

	employeesAtLoc, errLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))

	if errLoc != nil {

		writeFailureResponse(&w, http.StatusNotFound, errLoc)

		return

	}
	employeesAtLocality, errLocality := list.ListEmployeesByLocalityAtLoc(params["locality"], &employeesAtLoc)
	if errLocality != nil {

		writeFailureResponse(&w, http.StatusNotFound, errLocality)
		return

	}
	writeSuccessResponseList(&w, &employeesAtLocality)

}

func ShowByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	stringID := params["id"]

	id, err := strconv.Atoi(stringID)

	if err != nil {
		writeFailureResponseStr(&w, http.StatusBadRequest, "The ID must be a number")
		return
	}

	employees := make([]store.Employee, 0)
	empl, errShow := show.ShowEmployeeByID(id, &(store.IdEmpMap))

	if errShow != nil {

		writeFailureResponse(&w, http.StatusNotFound, errShow)
		return

	}

	employees = append(employees, empl)
	writeSuccessResponseList(&w, &employees)

}

func DeleteAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	numOfEmp := len(store.Employees)

	err := delete.DeleteFromEverywhere()

	if err != nil {
		writeFailureResponse(&w, http.StatusNotFound, err)

		return

	}

	writeSuccessResponse(&w, strconv.Itoa(numOfEmp)+" employee(s) was/were deleted successfully.")

}

func DeleteByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	stringID := params["id"]
	id, err1 := strconv.Atoi(stringID)

	if err1 != nil {
		writeFailureResponseStr(&w, http.StatusBadRequest, "The ID must be a number")
	}
	err := delete.DeleteByIDFromEverywhere(id)

	if err != nil {

		writeFailureResponse(&w, http.StatusNotFound, err)
		return
	}

	writeSuccessResponse(&w, "Deleted the employee with ID "+stringID+" successfully.")

}

func RemoveByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	stringID := params["id"]

	id, err1 := strconv.Atoi(stringID)

	if err1 != nil {
		writeFailureResponseStr(&w, http.StatusBadRequest, "The ID must be a number")
	}

	err := remove.RemoveEmployeesByIDEverywhere(id)

	if err != nil {

		writeFailureResponse(&w, http.StatusNotFound, err)
		return
	}

	writeSuccessResponse(&w, "Removed the employee with ID "+stringID+" successfully.")

}

func writeFailureResponse(w *(http.ResponseWriter), code int, err error) {
	r := failureResponse{(http.StatusText(code)), err.Error()}
	jsonObj, _ := json.MarshalIndent(r, "", "    ")
	(*w).WriteHeader(code)
	(*w).Write(jsonObj)
	kiplog.HTTPLog(err.Error())
}

func writeFailureResponseStr(w *(http.ResponseWriter), code int, errString string) {
	r := failureResponse{(http.StatusText(code)), errString}
	jsonObj, _ := json.MarshalIndent(r, "", "    ")
	(*w).WriteHeader(code)
	(*w).Write(jsonObj)
	kiplog.HTTPLog(errString)
}

func writeSuccessResponseList(w *(http.ResponseWriter), moreEmployees *([]store.Employee)) {

	r := successResponseList{SUCCESS, moreEmployees}

	jsonObj, _ := json.MarshalIndent(r, "", "    ")

	(*w).Write(jsonObj)
}

func writeSuccessResponse(w *(http.ResponseWriter), data string) {

	r := successResponse{SUCCESS, data}

	jsonObj, _ := json.MarshalIndent(r, "", "    ")

	(*w).Write(jsonObj)
}
