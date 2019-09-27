package handlers

import(

	"net/http"
	"fileProcess"
	"store"
	"encoding/json"
	"list"
	"github.com/gorilla/mux"
	"strconv"
	"show"
	"delete"
	"remove"

)
var employees []store.Employee

func ReceiveAndProcessFile(w http.ResponseWriter, req *http.Request){

	fileProcess.ParseFileAndStoreListOfEmployees(req.Body, &employees)
	store.StoreEmployeesByIdDeptAndLoc(&employees)
}


func ListAll(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
 	json.NewEncoder(w).Encode(list.ListAllEmployees(&employees))
	

}

func ListByDept(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	// log.Println(list.ListEmployeesByDept(params["name"], &(store.DeptEmpMap)))
	json.NewEncoder(w).Encode(list.ListEmployeesByDept(params["name"], &(store.DeptEmpMap)))

}

func ListByLoc(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	pin, err := strconv.Atoi(params["pin"])

	if (err!= nil) {panic("The pin code must be a number")}
	
	json.NewEncoder(w).Encode(list.ListEmployeesByLoc(pin, &(store.LocEmpMap)))
}

func ListByLocDoorNo(w http.ResponseWriter, req *http.Request){
	
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	pin, err := strconv.Atoi(params["pin"])

	if (err!= nil) {panic("The pin code must be a number")}
	
	dn, err := strconv.Atoi(params["dn"])

	if (err!= nil) {panic("The door number must be a number")}
	
	employeesAtLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))
	// log.Println(employeesAtLoc)
	json.NewEncoder(w).Encode(list.ListEmployeesByDoorNoAtLoc(dn, &employeesAtLoc))

}

func ListByLocStreet(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pin, err := strconv.Atoi(params["pin"])

	if (err!= nil) {panic("The pin code must be a number")}
	
	employeesAtLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))
	{json.NewEncoder(w).Encode(list.ListEmployeesByStreetAtLoc(params["street"], &employeesAtLoc))}

}

func ListByLocLocality(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	pin, err := strconv.Atoi(params["pin"])

	if (err!= nil) {panic("The pin code must be a number")}
	
	employeesAtLoc := list.ListEmployeesByLoc(pin, &(store.LocEmpMap))
	json.NewEncoder(w).Encode(list.ListEmployeesByLocalityAtLoc(params["locality"], &employeesAtLoc))

}

func ShowByID(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])

	if (err!= nil) {panic("The id must be a number")}
	
	json.NewEncoder(w).Encode(show.ShowEmployeeByID(id, &(store.IdEmpMap)))
}

func DeleteAll(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")


		delete.DeleteFullEmployeeList(&employees)
		delete.DeleteFullIDempMap(&(store.IdEmpMap))
		delete.DeleteFullDeptEmpMap(&(store.DeptEmpMap))
		delete.DeleteFullLocEmpMap(&(store.LocEmpMap))
}

func DeleteByID(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])

	if (err!= nil) {panic("The id must be a number")}
	
	empl := (store.IdEmpMap)[id]
	delete.DeleteEmployeeByIDFromList(id, &employees)
	  delete.DeleteEmployeeByIDFromDeptMap(id, empl.GetDept(), &(store.DeptEmpMap))
	  delete.DeleteEmployeeByIDFromLocMap(id, empl.GetPins(), &(store.LocEmpMap))
	 delete.DeleteEmployeeByIDFromIdMap(id, &(store.IdEmpMap))
   
}

func RemoveByID(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])

	if (err!= nil) {panic("The id must be a number")}
	
	remove.RemoveEmployeesFromList(id, &employees)
	remove.RemoveEmployeesFromIdEmpMap(id, &(store.IdEmpMap)) 
	remove.RemoveEmployeesFromDeptEmpMap(id, &(store.DeptEmpMap))
	remove.RemoveEmployeesFromLocEmpMap(id, &(store.LocEmpMap))
	
}