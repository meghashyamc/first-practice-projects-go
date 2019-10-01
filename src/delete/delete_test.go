package delete

import (
	"store"
	"testing"
)

var testEmployees []store.Employee
var testEmployee1 store.Employee
var testEmployee2 store.Employee
var testEmployee3 store.Employee

func TestDeleteFullEmployeeList(t *testing.T) {

	initializeStuff()
	if len(testEmployees) == 0 {

		t.Error("Problem in initializing delete tests.")
	}
	err := DeleteFullEmployeeList(&testEmployees)

	if err != nil {
		t.Error("Expected 3 employees to be deleted but got error:", err)

	}
	if len(testEmployees) != 0 {

		t.Error("Expected 3 employees to be deleted but found", len(testEmployees), "left")
	}

}

func BenchmarkDeleteFullEmployeeList(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	DeleteFullEmployeeList(&testEmployees)

}

func TestDeleteFullIDempMap(t *testing.T) {

	initializeStuff()
	if len(testEmployees) == 0 {

		t.Error("Problem in initializing delete tests.")
	}

	err := DeleteFullIDempMap(&(store.IdEmpMap))
	t.Error("IdEmpMap: Expected 3 employees to be deleted but got error:", err)

	if len(store.IdEmpMap) != 0 {

		t.Error("Expected all employees to be deleted but found", len(store.IdEmpMap), "left")
	}

}

func BenchmarkDeleteFullIDempMap(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	DeleteFullIDempMap(&(store.IdEmpMap))
}

func TestDeleteFullDeptEmpMap(t *testing.T) {

	initializeStuff()
	if len(testEmployees) == 0 {

		t.Error("Problem in initializing delete tests.")
	}

	err := DeleteFullDeptEmpMap(&(store.DeptEmpMap))
	t.Error("DeptEmpMap: Expected 3 employees to be deleted but got error:", err)

	if len(store.DeptEmpMap) != 0 {

		t.Error("Expected all employees to be deleted but found", len(store.DeptEmpMap), "left")
	}

}

func BenchmarkDeleteFullDeptempMap(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	DeleteFullDeptEmpMap(&(store.DeptEmpMap))

}

func TestDeleteFullLocEmpMap(t *testing.T) {

	initializeStuff()
	if len(testEmployees) == 0 {

		t.Error("Problem in initializing delete tests.")
	}

	err := DeleteFullLocEmpMap(&(store.LocEmpMap))
	t.Error("LocEmpMap: Expected 3 employees to be deleted but got error:", err)

	if len(store.LocEmpMap) != 0 {

		t.Error("Expected all employees to be deleted but found", len(store.LocEmpMap), "left")
	}

}

func BenchmarkDeleteFullLocempMap(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	DeleteFullLocEmpMap(&(store.LocEmpMap))

}

func TestDeleteEmployeeByIDFromList(t *testing.T) {

	initializeStuff()

	if len(testEmployees) != 3 {
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}

	if (testEmployees[0]).Name != "Pappu" {
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}

	err := DeleteEmployeeByIDFromList(1, &testEmployees)
	t.Error("Expected employee with ID 1 to be deleted, got error:", err)

	if len(testEmployees) != 2 {
		t.Error("Expected size of testEmployees to be 2 after 1 deletion got", len(testEmployees))
	}

	if (testEmployees[0]).Name == "Pappu" {
		t.Error("Expected Pappu to be deleted got", (testEmployees[0]).Name)
	}

}

func BenchmarkEmployeeByIDFromList(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	for i := 0; i < 3; i++ {
		DeleteEmployeeByIDFromList(i+1, &testEmployees)
	}
}

func TestDeleteEmployeeByIDFromDeptMap(t *testing.T) {

	initializeStuff()

	if len(testEmployees) != 3 {
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}

	if (testEmployees[0]).Name != "Pappu" {
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}

	err := DeleteEmployeeByIDFromDeptMap(1, testEmployees[0].GetDept(), &(store.DeptEmpMap))
	t.Error("DeptEmpMap: Expected employee with ID 1 to be deleted, got error:", err)

	for _, empl := range *(store.DeptEmpMap["Accounts"]) {

		if empl.GetID() == 1 {

			t.Error("Expected ID 1 to be deleted from Accounts but found", empl.GetID())
		}

	}

	for _, empl := range *(store.DeptEmpMap["Management"]) {

		if empl.GetID() == 1 {

			t.Error(store.DeptEmpMap["Management"])

			t.Error("Expected ID 1 to be deleted from Management but found", empl.GetID())
		}

	}

}

func BenchmarkDeleteEmployeeIDFromDeptMap(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	for i := 0; i < 3; i++ {
		DeleteEmployeeByIDFromDeptMap(i+1, testEmployees[i].GetDept(), &(store.DeptEmpMap))
	}
}

func TestDeleteEmployeeByIDFromLocMap(t *testing.T) {

	initializeStuff()

	if len(testEmployees) != 3 {
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}

	if (testEmployees[0]).Name != "Pappu" {
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}

	err := DeleteEmployeeByIDFromLocMap(1, testEmployees[0].GetPins(), &(store.LocEmpMap))
	t.Error("LocEmpMap: Expected employee with ID 1 to be deleted, got error:", err)

	for _, empl := range *(store.LocEmpMap[560002]) {

		if empl.GetID() == 1 {

			t.Error("Expected ID 1 to be deleted from 560002 but found", empl.GetID())
		}

	}

	for _, empl := range *(store.LocEmpMap[560103]) {

		if empl.GetID() == 1 {

			t.Error(store.LocEmpMap[560103])

			t.Error("Expected ID 1 to be deleted from 560103 but found", empl.GetID())
		}

	}

}

func BenchmarkDeleteEmployeeIDFromLocMap(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	for i := 0; i < 3; i++ {
		DeleteEmployeeByIDFromLocMap(i+1, testEmployees[i].GetPins(), &(store.LocEmpMap))
	}
}

func TestDeleteEmployeeByIDFromIDMap(t *testing.T) {

	initializeStuff()

	if len(testEmployees) != 3 {
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}

	if (testEmployees[0]).Name != "Pappu" {
		t.Error("Problem with initializing stuff for DeleteEmployeeByIDFromlist")
	}

	err := DeleteEmployeeByIDFromIdMap(1, &(store.IdEmpMap))
	t.Error("IdEmpMap: Expected employee with ID 1 to be deleted, got error:", err)

	if (store.IdEmpMap[1]).Name == "Pappu" {

		t.Error("Expected Pappu to be deleted from ID Emp map but found", (store.IdEmpMap[1]).Name)
	}

}

func BenchmarkDeleteEmployeeByIDFromIdMap(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	for i := 0; i < 3; i++ {
		DeleteEmployeeByIDFromIdMap(i+1, &(store.IdEmpMap))
	}
}

func initializeStuff() {

	address1 := store.Address{Doorno: 25,
		Street:   "Adepalli Street",
		Locality: "Basvangudi",
		PIN:      560002}
	address2 := store.Address{Doorno: 32,
		Street:   "Pilli Street",
		Locality: "Balagere",
		PIN:      560103}

	address3 := store.Address{Doorno: 564,
		Street:   "White House Street",
		Locality: "Malleshwaram",
		PIN:      560003}

	address4 := store.Address{Doorno: 43,
		Street:   "Vishveshwarya Street",
		Locality: "Hosur",
		PIN:      560203}
	testEmployee1 = store.Employee{
		EmpID:      1,
		Name:       "Pappu",
		Department: []string{"Accounts", "Management"},
		Addresses:  [](store.Address){address1, address2},
		There:      true,
	}

	testEmployee2 = store.Employee{
		EmpID:      2,
		Name:       "Rinky",
		Department: []string{"IT", "Admin"},
		Addresses:  [](store.Address){address2, address3},
		There:      true,
	}

	testEmployee3 = store.Employee{
		EmpID:      3,
		Name:       "Tinkadi",
		Department: []string{"Accounts", "Admin"},
		Addresses:  [](store.Address){address4},
		There:      true,
	}

	testEmployees = []store.Employee{
		testEmployee1, testEmployee2, testEmployee3}
	store.InitializeEmployeesAndMaps()
	store.StoreEmployeesByIdDeptAndLoc(&testEmployees)

}
