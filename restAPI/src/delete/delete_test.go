package delete

import (
	"store"
	"testing"
)

var testEmployees []store.Employee
var testEmployee1 store.Employee
var testEmployee2 store.Employee
var testEmployee3 store.Employee

func TestDeleteFromEverywhere(t *testing.T) {

	initializeStuff()

	err := DeleteFromEverywhere()

	if err != nil {

		t.Error("Expected all employees to be deleted, got", err)

	}

	if len(store.IdEmpMap) != 0 {

		t.Error("Expected length of IdEmpMap to be 0 after deletion, got", len(store.IdEmpMap))
	}

	if len(store.DeptEmpMap) != 0 {

		t.Error("Expected length of DeptEmpMap to be 0 after deletion, got", len(store.DeptEmpMap))
	}

	if len(store.LocEmpMap) != 0 {

		t.Error("Expected length of LocEmpMap to be 0 after deletion, got", len(store.LocEmpMap))
	}

}

func BenchmarkDeleteFromEverywhere(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	DeleteFromEverywhere()

}

func TestDeleteByIDFromEverywhere(t *testing.T) {

	initializeStuff()

	err := DeleteByIDFromEverywhere(1)

	if err != nil {

		t.Error("Expected employee with ID 1 to be deleted, got", err)

	}

	if len(store.Employees) != 2 {
		t.Error("Expected size of Employees list to be 2 after 1 deletion got", len(store.Employees))
	}

	if (store.Employees[0]).Name == "Pappu" {
		t.Error("Expected Pappu to be deleted got", (store.Employees[0]).Name)
	}

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

	if (store.IdEmpMap[1]).Name == "Pappu" {

		t.Error("Expected Pappu to be deleted from ID Emp map but found", (store.IdEmpMap[1]).Name)
	}

}

func BenchmarkDeleteByIDFromEverywhere(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	for i := 0; i < 3; i++ {
		DeleteByIDFromEverywhere(i + 1)
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
	store.Employees = append(store.Employees, testEmployees...)
	store.StoreEmployeesByIdDeptAndLoc(&testEmployees)

}
