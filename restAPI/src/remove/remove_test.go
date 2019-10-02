package remove

import (
	"store"
	"testing"
)

var testEmployees []store.Employee
var testEmployee1 store.Employee
var testEmployee2 store.Employee
var testEmployee3 store.Employee

func TestRemoveEmployeesByIDEverywhere(t *testing.T) {
	initializeStuff()

	err := RemoveEmployeesByIDEverywhere(2)

	if err != nil {

		t.Error("Expected Rinky's There field to be false got error:", err)

	}

	if store.IdEmpMap[2].There == true {

		t.Error("Expected Rinky's There field to be false in IdEmpMap got", store.IdEmpMap[2].There)
	}

	if (*(store.DeptEmpMap["IT"]))[0].There == true {

		t.Error("Expected Rinky's there field to be false, got", true)

	}

	if (*(store.LocEmpMap[560103]))[1].There == true {

		t.Error("Expected Rinky's there field to be false, got", true)

	}

}

func BenchmarkRemoveEmployeesByIDEverywhere(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	for i := 0; i < 3; i++ {

		RemoveEmployeesByIDEverywhere(i + 1)

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
