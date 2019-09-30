package store

import (
	"testing"
)

var testEmployees []Employee
var testEmployee1 Employee
var testEmployee2 Employee
var testEmployee3 Employee

func TestStoreEmployeesByIdDeptAndLoc(t *testing.T) {
	initializeStuff()

	StoreEmployeesByIdDeptAndLoc(&testEmployees)

	if len(IdEmpMap) != 3 {

		t.Error("Expected IdEmpMap to have size 3 got", len(IdEmpMap))
	}

	if len(DeptEmpMap) != 4 {

		t.Error("Expected IdEmpMap to have size 4 got", len(DeptEmpMap))
	}

	if len(LocEmpMap) != 4 {

		t.Error("Expected LocEmpMap to have size 4 got", len(LocEmpMap))
	}

	if (IdEmpMap[2]).Name != testEmployee2.Name {

		t.Error("Expected employee with ID2 to be Rinky got", testEmployee2.Name)
	}

	e := (IdEmpMap[3])
	if (e.GetLocalities())[0] != "Hosur" {

		t.Error("Expected location of employee with ID3 as Hosur got", e.GetLocalities()[0])
	}

	if len(*(DeptEmpMap["Accounts"])) != 2 {

		t.Error(DeptEmpMap["Accounts"])

		t.Error("Expected number of employees in Accounts as 2 got", len(*(DeptEmpMap["Accounts"])))
	}

	if ((*(DeptEmpMap["IT"]))[0]).Name != "Rinky" {

		t.Error("Expected first IT employee to be Rinky got", ((*(DeptEmpMap["IT"]))[0]).Name)
	}

	if len(*(LocEmpMap[560203])) != 1 {

		t.Error("Expected number of people living at 560203 to be 1 got", len(*(LocEmpMap[560203])))
	}

	if ((*(LocEmpMap[560003]))[0]).Name != "Rinky" {

		t.Error("Expected first employee living at 56003 to be Rinky got", ((*(LocEmpMap[560003]))[0]).Name)
	}

}

func initializeStuff() {

	address1 := Address{Doorno: 25,
		Street:   "Adepalli Street",
		Locality: "Basvangudi",
		PIN:      560002}
	address2 := Address{Doorno: 32,
		Street:   "Pilli Street",
		Locality: "Balagere",
		PIN:      560103}

	address3 := Address{Doorno: 564,
		Street:   "White House Street",
		Locality: "Malleshwaram",
		PIN:      560003}

	address4 := Address{Doorno: 43,
		Street:   "Vishveshwarya Street",
		Locality: "Hosur",
		PIN:      560203}
	testEmployee1 = Employee{
		EmpID:      1,
		Name:       "Pappu",
		Department: []string{"Accounts", "Management"},
		Addresses:  [](Address){address1, address2},
		There:      true,
	}

	testEmployee2 = Employee{
		EmpID:      2,
		Name:       "Rinky",
		Department: []string{"IT", "Admin"},
		Addresses:  [](Address){address2, address3},
		There:      true,
	}

	testEmployee3 = Employee{
		EmpID:      3,
		Name:       "Tinkadi",
		Department: []string{"Accounts", "Admin"},
		Addresses:  [](Address){address4},
		There:      true,
	}

	testEmployees = []Employee{
		testEmployee1, testEmployee2, testEmployee3}
	InitializeEmployeesAndMaps()

}
