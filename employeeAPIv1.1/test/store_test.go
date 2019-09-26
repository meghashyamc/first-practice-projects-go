package test

import (
"testing"
"store"

)


func TestStoreEmployeesByIdDeptAndLoc(t *testing.T){
	initializeStuff()

store.StoreEmployeesByIdDeptAndLoc(&testEmployees)

if len(store.IdEmpMap) != 3{

t.Error("Expected IdEmpMap to have size 3 got", len(store.IdEmpMap))}

if len(store.DeptEmpMap) != 4{

	t.Error("Expected IdEmpMap to have size 4 got", len(store.DeptEmpMap))}


if len(store.LocEmpMap) != 4{

	t.Error("Expected LocEmpMap to have size 4 got", len(store.LocEmpMap))}
	
if (store.IdEmpMap[2]).Name != testEmployee2.Name{

		t.Error("Expected employee with ID2 to be Rinky got", testEmployee2.Name)}

		e := (store.IdEmpMap[3])
		if (e.GetLocalities())[0] != "Hosur"{

			t.Error("Expected location of employee with ID3 as Hosur got", e.GetLocalities()[0])}

if len(*(store.DeptEmpMap["Accounts"])) != 2{

	t.Error(store.DeptEmpMap["Accounts"])

				t.Error("Expected number of employees in Accounts as 2 got", len(*(store.DeptEmpMap["Accounts"])))}
		
if ((*(store.DeptEmpMap["IT"]))[0]).Name != "Rinky"{

					t.Error("Expected first IT employee to be Rinky got", ((*(store.DeptEmpMap["IT"]))[0]).Name)}
			
if len(*(store.LocEmpMap[560203])) != 1{

				t.Error("Expected number of people living at 560203 to be 1 got", len(*(store.LocEmpMap[560203])))}
		
if ((*(store.LocEmpMap[560003]))[0]).Name != "Rinky"{

					t.Error("Expected first employee living at 56003 to be Rinky got", ((*(store.LocEmpMap[560003]))[0]).Name)}
			
}

