package list

import (
	"testing"

	"store"
)

var employees []store.Employee

func TestListAllEmployees(t *testing.T) {

	initializeStuff()
	list, err := ListAllEmployees(&employees)

	if err != nil {
		t.Error("Expected 2 employees to get added, got error:", err)

	}
	if len(list) != 2 {

		t.Error("Expected 2 employees to get added, got", len(list))

	}

	if list[0].Name != "Pappu" {
		t.Error("Expected Pappu to be the first added employee, got", list[0].Name)

	}

}

func BenchmarkListAllEmployees(b *testing.B) {

	initializeStuff()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ListAllEmployees(&employees)
	}

}

func TestListSearchEmployees(t *testing.T) {

	term := "pap"
	list, err := ListSearchEmployees(term, &employees)

	if err != nil {
		t.Error("Expected 1 employee to match term 'pap', got error:", err)

	}
	if len(list) != 1 {

		t.Error("Expected 1 employee to match term 'pap' got", len(list))

	}

	if list[0].Name != "Pappu" {
		t.Error("Expected Pappu to show up on the found list, got", list[0].Name)

	}

}

func BenchmarkListSearchEmployees(b *testing.B) {

	initializeStuff()
	b.ResetTimer()
	term := "pap"
	for i := 0; i < b.N; i++ {

		ListSearchEmployees(term, &employees)
	}

}
func TestListEmployeesByLoc(t *testing.T) {

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

	testEmployee1 := store.Employee{
		EmpID:      1,
		Name:       "Pappu",
		Department: []string{"Accounts", "Management"},
		Addresses:  [](store.Address){address1, address2},
		There:      true,
	}

	testEmployee2 := store.Employee{
		EmpID:      2,
		Name:       "Rinky",
		Department: []string{"IT", "Admin"},
		Addresses:  [](store.Address){address2, address3},
		There:      true,
	}

	locEmpMap := make(map[int]*([]store.Employee))

	list1 := []store.Employee{testEmployee1}
	list2 := []store.Employee{testEmployee2}
	list3 := []store.Employee{testEmployee1, testEmployee2}
	locEmpMap[560002] = &list1
	locEmpMap[560003] = &list2
	locEmpMap[560103] = &list3

	listAddress, err := ListEmployeesByLoc(560002, &locEmpMap)

	if err != nil {

		t.Error("Expected 560002 to have employee Pappu, got error:", err)
	}
	if listAddress[0].Name != list1[0].Name {

		t.Error("Expected 560002 to have employee Pappu got the address(es) of", listAddress)

	}

	listAddress, err = ListEmployeesByLoc(560103, &locEmpMap)

	if err != nil {

		t.Error("Expected 560103 to have  employees Pappu and Rinky, got error:", err)

	}

	if listAddress[0].Name != list3[0].Name &&
		listAddress[1].Name != list3[1].Name {

		t.Error("Expected 560103 to have employees Pappu and Rinky got the address(es) of", listAddress)

	}

}

func BenchmarkListEmployeesByLoc(b *testing.B) {

	initializeStuff()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		ListEmployeesByLoc(560002, &(store.LocEmpMap))
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

	testEmployee1 := store.Employee{
		EmpID:      1,
		Name:       "Pappu",
		Department: []string{"Accounts", "Management"},
		Addresses:  [](store.Address){address1, address2},
		There:      true,
	}

	testEmployee2 := store.Employee{
		EmpID:      2,
		Name:       "Rinky",
		Department: []string{"IT", "Admin"},
		Addresses:  [](store.Address){address2, address3},
		There:      true,
	}

	employees = []store.Employee{testEmployee1, testEmployee2}
	store.InitializeEmployeesAndMaps()
	store.StoreEmployeesByIdDeptAndLoc(&employees)
}
