package test


import(

"testing"
"remove"
"store"
)


func TestRemoveEmployeesFromList(t *testing.T){

initializeStuff()



remove.RemoveEmployeesFromList(1, &testEmployees)

if testEmployees[0].There == true{

	t.Error("Expected Pappu's There field to be false got", testEmployees[0].There)
}

}


func TestRemoveEmployeesFromIDEmpMap(t *testing.T){

	initializeStuff()
	
	
	
	remove.RemoveEmployeesFromIdEmpMap(2, &(store.IdEmpMap))
	
		
	if store.IdEmpMap[2].There == true{
	
		t.Error("Expected Rinky's There field to be false in IdEmpMap got",  store.IdEmpMap[2].There)
	}
	}


	func TestRemoveEmployeesFromDeptEmpMap(t *testing.T){

		initializeStuff()
		

		
		remove.RemoveEmployeesFromDeptEmpMap(1, &(store.DeptEmpMap))
		
		if (*(store.DeptEmpMap["Accounts"]))[0].There == true{
		
			t.Error("Expected Pappu's there field to be false, got", true)

		}
		
		if (*(store.DeptEmpMap["Management"]))[0].There == true{
		
			t.Error("Expected Pappu's there field to be false, got", true)

		}
		}


		func TestRemoveEmployeesFromLocEmpMap(t *testing.T){

			initializeStuff()
			
		
			
			remove.RemoveEmployeesFromLocEmpMap(1, &(store.LocEmpMap))
			
			if (*(store.LocEmpMap[560002]))[0].There == true{
			
				t.Error("Expected Pappu's there field to be false, got", true)
	
			}
			
			if (*(store.LocEmpMap[560103]))[0].There == true{
			
				t.Error("Expected Pappu's there field to be false, got", true)
	
			}
			}

