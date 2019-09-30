package fileProcess

import (
	"store"
	"testing"
)

var testEmployees []store.Employee
var testEmployee1 store.Employee
var testEmployee2 store.Employee
var testEmployee3 store.Employee

func TestGenerateIDsIfNotThere(t *testing.T) {

	initializeStuff()

	if testEmployees[2].GetID() != 0 {
		t.Error("Problem initializing generate ID test.")
	}

	generateIDsIfNotThere(&testEmployees)

	if testEmployees[2].GetID() == 0 {
		t.Error("Expected a non-zero ID for an employee with no initial ID, got", 0)
	}

}

func BenchmarkGenerateIDsIfNotThere(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		generateIDsIfNotThere(&testEmployees)

	}

}

func TestAreCompulsoryFieldsThere(t *testing.T) {

	initializeStuff()

	fieldsThere := areCompulsoryFieldsThere(&testEmployees)

	if fieldsThere == true {
		t.Error("Expected compulsory fields to not be there, got", true)
	}

}

func BenchmarkAreCompulsoryFieldsThere(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		areCompulsoryFieldsThere(&testEmployees)

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

		Name:      "Tinkadi",
		Addresses: [](store.Address){address4},
		There:     true,
	}

	testEmployees = []store.Employee{
		testEmployee1, testEmployee2, testEmployee3}

}
