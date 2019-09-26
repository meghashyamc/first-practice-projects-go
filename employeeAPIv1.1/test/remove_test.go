package test


import(

"testing"
"remove"
"store"
)


func TestRemoveEmployeesFromList(t *testing.T){

initializeStuff()

removeArgs := []int{1,2}

remove.RemoveEmployeesFromList(removeArgs, &testEmployees)

if testEmployees[0].There == true{

	t.Error("Expected Pappu's There field to be false got", testEmployees[0].There)
}


if testEmployees[1].There == true{

	t.Error("Expected Rinky's There field to be false got", testEmployees[1].There)
}
}




func TestRemoveEmployeesFromIDEmpMap(t *testing.T){

	initializeStuff()
	
	removeArgs := []int{1,2}
	
	remove.RemoveEmployeesFromIdEmpMap(removeArgs, &(store.IdEmpMap))
	
	if store.IdEmpMap[1].There == true{
	
		t.Error("Expected Pappu's There field to be false in IdEmpMap got",  store.IdEmpMap[1].There)
	}
	
	
	if store.IdEmpMap[2].There == true{
	
		t.Error("Expected Rinky's There field to be false in IdEmpMap got",  store.IdEmpMap[2].There)
	}
	}


	func TestRemoveEmployeesFromDeptEmpMap(t *testing.T){

		initializeStuff()
		
		removeArgs := []int{1}
		
		remove.RemoveEmployeesFromDeptEmpMap(removeArgs, &(store.DeptEmpMap))
		
		if (*(store.DeptEmpMap["Accounts"]))[0].There == true{
		
			t.Error("Expected Pappu's there field to be false, got", true)

		}
		
		if (*(store.DeptEmpMap["Management"]))[0].There == true{
		
			t.Error("Expected Pappu's there field to be false, got", true)

		}
		}


		func TestRemoveEmployeesFromLocEmpMap(t *testing.T){

			initializeStuff()
			
			removeArgs := []int{1}
			
			remove.RemoveEmployeesFromLocEmpMap(removeArgs, &(store.LocEmpMap))
			
			if (*(store.LocEmpMap[560002]))[0].There == true{
			
				t.Error("Expected Pappu's there field to be false, got", true)
	
			}
			
			if (*(store.LocEmpMap[560103]))[0].There == true{
			
				t.Error("Expected Pappu's there field to be false, got", true)
	
			}
			}

