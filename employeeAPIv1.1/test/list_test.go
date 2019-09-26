package test

import ("testing"

"store"
"list"



)

func TestListEmployeesByLoc(t *testing.T){


	address1 := store.Address{Doorno: 25,
		Street: "Adepalli Street",
		Locality: "Basvangudi",
		PIN: 560002}
		address2 := store.Address{Doorno: 32,
			Street: "Pilli Street",
			Locality: "Balagere",
			PIN: 560103}
		
		address3 := store.Address{Doorno: 564,
			Street: "White House Street",
			Locality: "Malleshwaram",
			PIN: 560003}
		
		
			testEmployee1 := store.Employee{
		EmpID: 1,
		Name: "Pappu",
		Department: []string{"Accounts", "Management"},
		Addresses: [](store.Address){address1, address2},
		There: true,
		
			}
		
			testEmployee2 := store.Employee{
				EmpID: 2,
				Name: "Rinky",
				Department: []string{"IT", "Admin"},
				Addresses: [](store.Address){address2, address3},
				There: true,
				
					}

locEmpMap := make(map[int]*([]store.Employee))

list1 := []store.Employee{testEmployee1}
list2 := []store.Employee{testEmployee2}
list3 := []store.Employee{testEmployee1, testEmployee2}
locEmpMap[560002] = &list1
locEmpMap[560003] = &list2
locEmpMap[560103] = &list3

listAddress, _ := list.ListEmployeesByLoc(560002, &locEmpMap)
if listAddress != &list1{

t.Error("Expected 560002 to have employee Pappu got the address(es) of", *listAddress)

}


listAddress, _ = list.ListEmployeesByLoc(560103, &locEmpMap)
if listAddress != &list3{

t.Error("Expected 560002 to have employees Pappu and Rinku got the address(es) of", *listAddress)

}

}
