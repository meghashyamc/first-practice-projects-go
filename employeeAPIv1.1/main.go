package main

 import (
"os"
"fmt"
"strconv"
"store"
"remove"
"fileProcess"
"delete"
"list"
"show"
)

// indicates maximum arguments used
const ARGSNUM = 3

func main(){

	var employees []store.Employee

	file := "employees.json"
	//parses json file and builds employee objects in an array slice
	fileProcess.ParseFileAndStoreListOfEmployees(file, &employees)

	//stores employees in maps with ids, depts and locs as keys
	store.StoreEmployeesByIdDeptAndLoc(&employees)
	

for{
	//stores input arguments
	args := make([]string, ARGSNUM )

fmt.Scanf("%s", &args[0])

//exit the program
if args[0] == "exit"{
		os.Exit(0)

}

fmt.Scanf("%s", &args[1])
	
// list all shows names of all employees
// list department <NAME> shows names of all employees in one department
 if args[0] == "list"{

	if args[1] == "all"  {list.ListAllEmployees(&employees)
		
		} else if args[1] == "department"{
			
			fmt.Scanf("%s", &args[2])
			list.ListEmployeesByDept(args[2], &(store.DeptEmpMap))
			
		
		
//list location <PIN> shows employees at a given PIN code		
		} else if args[1] == "location"{

			
			fmt.Scanf("%s", &args[2])

				i,err := strconv.Atoi(args[2])

			if err != nil {argumentsError()
	}else {

		fmt.Println("Do you know the door number, street name or locality name? (y/n)")

		var yesNo string
		fmt.Scanf("%s", &yesNo)

		if (yesNo == "y"){

			fmt.Println("Enter doorno <num> or street <street name> or locality <locality name>")

			locArgs := make([]string,2)
			fmt.Scanf("%s", &locArgs[0])
			fmt.Scanf("%s", &locArgs[1])

			employeesAtLoc, ok := list.ListEmployeesByLoc(i, &(store.LocEmpMap))
			

			if(ok){

				if locArgs[0]=="doorno"{

					j,err := strconv.Atoi(locArgs[1])

					if err != nil {argumentsError()
						}
			

					list.ListEmployeesByDoorNoAtLoc(j, employeesAtLoc)

			} else if locArgs[0]=="street"{

					//locArgs[1] is name of street
				list.ListEmployeesByStreetAtLoc(locArgs[1], employeesAtLoc)

		} else if locArgs[0]=="locality"{

					//locArgs[1] is name of locality
			list.ListEmployeesByLocalityAtLoc(locArgs[1], employeesAtLoc)

	}

			}
			
		} else{
			
			list.ListEmployeesByLocPrint(i, &(store.LocEmpMap))

		}
			
	}
			 
			
		
		
		}
// show <ID> prints the name of the employee with an id <ID>
} else if args[0] == "show"{
	
	i,err := strconv.Atoi(args[1])

	if err != nil {argumentsError()
	}else {
		show.ShowEmployeeByID(i, &(store.IdEmpMap))
	 }
	 //delete all deletes all employee records
	 // delete <ID> deletes employee with a particular ID
 } else if args[0] == "delete"{

	if (args[1] == "all"){

		delete.DeleteFullEmployeeList(&employees)
		delete.DeleteFullIDempMap(&(store.IdEmpMap))
		delete.DeleteFullDeptEmpMap(&(store.DeptEmpMap))
		delete.DeleteFullLocEmpMap(&(store.LocEmpMap))
		continue

	} else{i,err := strconv.Atoi(args[1])

 	if err != nil {argumentsError()
 	}else {
		empl := (store.IdEmpMap)[i]
		 delete.DeleteEmployeeByIDFromList(i, &employees)
		   delete.DeleteEmployeeByIDFromDeptMap(i, empl.GetDept(), &(store.DeptEmpMap))
		   delete.DeleteEmployeeByIDFromLocMap(i, empl.GetPins(), &(store.LocEmpMap))
		  delete.DeleteEmployeeByIDFromIdMap(i, &(store.IdEmpMap))
		
	 }
	 
	}
 
 // remove <NUM> allows removal of NUM employees
 // then, entering list of ids leads to There key of those employees being set to false
 // removed employees don't show up in list
 } else if args[0] == "remove"{

	i,err := strconv.Atoi(args[1])
	if err != nil {argumentsError()
	} else{
	fmt.Println("Enter",i,"employee IDs to remove.")

	removeArgs := make([]int, i)

	for j:= 0; j < i; j++{

		fmt.Scanf("%d", &removeArgs[j])
	}

	remove.RemoveEmployeesFromList(removeArgs, &employees)
	remove.RemoveEmployeesFromIdEmpMap(removeArgs, &(store.IdEmpMap)) 
	remove.RemoveEmployeesFromDeptEmpMap(removeArgs, &(store.DeptEmpMap))
	remove.RemoveEmployeesFromLocEmpMap(removeArgs, &(store.LocEmpMap))
	continue
	}
 } else{

	argumentsError()	
}
}
}


func argumentsError(){
	fmt.Println("Error in arguments.")
}





