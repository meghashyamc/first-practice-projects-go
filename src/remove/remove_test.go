package remove

import (
	"store"
	"testing"
)

var testEmployees []store.Employee
var testEmployee1 store.Employee
var testEmployee2 store.Employee
var testEmployee3 store.Employee

func TestRemoveEmployeesFromList(t *testing.T) {

	initializeStuff()

	err := RemoveEmployeesFromList(1, &testEmployees)

	if err != nil {
		t.Error("Expected Pappu's There field to be false got error:", err)

	}
	if testEmployees[0].There == true {

		t.Error("Expected Pappu's There field to be false got", testEmployees[0].There)
	}

}

func BenchmarkRemoveEmployeesFromList(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	for i := 0; i < 3; i++ {

		RemoveEmployeesFromList(i+1, &testEmployees)

	}

}

func TestRemoveEmployeesFromIDEmpMap(t *testing.T) {

	initializeStuff()

	err := RemoveEmployeesFromIdEmpMap(2, &(store.IdEmpMap))

	if err != nil {
		t.Error("Expected Rinky's There field to be false got error:", err)

	}
	if store.IdEmpMap[2].There == true {

		t.Error("Expected Rinky's There field to be false in IdEmpMap got", store.IdEmpMap[2].There)
	}
}

func BenchmarkRemoveEmployeesFromIDEmpMap(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	for i := 0; i < 3; i++ {

		RemoveEmployeesFromIdEmpMap(i+1, &(store.IdEmpMap))

	}

}

func TestRemoveEmployeesFromDeptEmpMap(t *testing.T) {

	initializeStuff()

	err := RemoveEmployeesFromDeptEmpMap(1, &(store.DeptEmpMap))

	if err != nil {
		t.Error("Expected Pappu's there field to be false got error:", err)

	}
	if (*(store.DeptEmpMap["Accounts"]))[0].There == true {

		t.Error("Expected Pappu's there field to be false, got", true)

	}

	if (*(store.DeptEmpMap["Management"]))[0].There == true {

		t.Error("Expected Pappu's there field to be false, got", true)

	}
}

func BenchmarkRemoveEmployeesFromDeptEmpMap(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	for i := 0; i < 3; i++ {

		RemoveEmployeesFromDeptEmpMap(i+1, &(store.DeptEmpMap))

	}

}

func TestRemoveEmployeesFromLocEmpMap(t *testing.T) {

	initializeStuff()

	err := RemoveEmployeesFromLocEmpMap(1, &(store.LocEmpMap))

	if err != nil {
		t.Error("Expected Pappu's there field to be false got error:", err)

	}
	if (*(store.LocEmpMap[560002]))[0].There == true {

		t.Error("Expected Pappu's there field to be false, got", true)

	}

	if (*(store.LocEmpMap[560103]))[0].There == true {

		t.Error("Expected Pappu's there field to be false, got", true)

	}
}

func BenchmarkRemoveEmployeesFromLocEmpMap(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	for i := 0; i < 3; i++ {

		RemoveEmployeesFromLocEmpMap(i+1, &(store.LocEmpMap))

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
