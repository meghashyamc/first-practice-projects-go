package main

 import (
 "encoding/json"
"os"
"fmt"
 "io/ioutil"
"strconv"
)

// indicates maximum arguments used
const ARGSNUM = 3


type Address struct{
	DoorNo int
	Street string 
	Locality string 
	PIN int 
	
	}

	//helper method for Addres
	func (a *Address) getPin() int{

		return a.PIN
	  }


	type Employee struct {
		EmpID int 
		Name string
		Department []string
		Addresses []Address
		There bool
	  }

//helper methods for Employee
	  func (e *Employee) printName(){

		fmt.Println(e.Name)
	  }

	  func (e *Employee) getID() int{

		return e.EmpID
	  }

	  func (e *Employee) getPins() []int{

		var pins []int
		
		for _, addr := range e.Addresses{

				pins = append(pins, addr.getPin())

		}

		return pins
	  }

	  func (e *Employee) getDept() []string{

		return e.Department
	  }



func main(){


	var employees []Employee
	// returns a byte slice
	file, err := ioutil.ReadFile("employees.json") 

   if err != nil {fmt.Println("Error parsing json file.")}

   //stores corresponding fields in employees slice
   json.Unmarshal([]byte(file), &employees)

	//stores ids as keys and employees as values
   idEmpMap := storeEmployeesByIDMap(&employees)

	//stores department names (strings) as keys and reference to array of employees as values
   deptEmpMap := storeEmployeesByDeptMap(&employees)

	//stores PIN codes (int) as keys and reference to array of employees as values

   locEmpMap := storeEmployeesByLocMap(&employees)


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

	if args[1] == "all"  {listAllEmployees(&employees)
		
		} else if args[1] == "department"{
			
			fmt.Scanf("%s", &args[2])
			listEmployeesByDept(args[2], &deptEmpMap)
			
		
		
//list location <PIN> shows employees at a given PIN code		
		} else if args[1] == "location"{
			
			fmt.Scanf("%s", &args[2])
			i,err := strconv.Atoi(args[2])

			if err != nil {argumentsError()
	}else {
			
			listEmployeesByLoc(i, &locEmpMap)


	}
		
		
		}
// show <ID> prints the name of the employee with 
} else if args[0] == "show"{
	
	i,err := strconv.Atoi(args[1])

	if err != nil {argumentsError()
	}else {
		showEmployeeByID(i, &idEmpMap)
	 }
	 //delete all deletes all employee records
	 // delete <ID> deletes employee with a particular ID
 } else if args[0] == "delete"{

	if (args[1] == "all"){

		deleteFullEmployeeList(&employees)
		deleteFullIDempMap(&idEmpMap)
		deleteFullDeptEmpMap(&deptEmpMap)
		deleteFullLocEmpMap(&locEmpMap)
		continue

	} else{i,err := strconv.Atoi(args[1])

 	if err != nil {argumentsError()
 	}else {
		empl := idEmpMap[i]
		 deleteEmployeeByIDFromList(i, &employees)
		
		   deleteEmployeeByIDFromDeptMap(i, empl.getDept(), &deptEmpMap)
		   deleteEmployeeByIDFromLocMap(i, empl.getPins(), &locEmpMap)
		  deleteEmployeeByIDFromIdMap(i, &idEmpMap)
		
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

	removeEmployeesFromList(removeArgs, &employees)
	removeEmployeesFromIdEmpMap(removeArgs, &idEmpMap) 
	removeEmployeesFromDeptEmpMap(removeArgs, &deptEmpMap)
	removeEmployeesFromLocEmpMap(removeArgs, &locEmpMap)
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

func removeEmployeesFromList(ids []int, employees *([]Employee)){


	for _, id  := range ids{

		for i:=0; i < len(*employees); i++{

				if ((*employees)[i]).getID() == id{

					((*employees)[i]).There = false
				}
		}

	}
}


func removeEmployeesFromIdEmpMap(ids []int, idEmpMap *map[int]Employee){


	for _, id  := range ids{

		for mapId:= range *idEmpMap{

				if mapId == id{

					newEmpl := (*idEmpMap)[mapId]

					newEmpl.There = false
					(*idEmpMap)[mapId] = newEmpl
				}
		}

	}
}

func removeEmployeesFromDeptEmpMap(ids []int, deptEmpMap *map[string]*([]Employee)){
	

	for _, id  := range ids{

		for dept:= range *deptEmpMap{

				for i:= 0; i < len(*((*deptEmpMap)[dept])); i++{

					if (*((*deptEmpMap)[dept]))[i].getID() == id{
						
						(*((*deptEmpMap)[dept]))[i].There = false
					}

				}
		}

	}
}


func removeEmployeesFromLocEmpMap(ids []int, locEmpMap *map[int]*([]Employee)){
	

	for _, id  := range ids{

		for pin:= range *locEmpMap{

				for i:= 0; i < len(*((*locEmpMap)[pin])); i++{

					if (*((*locEmpMap)[pin]))[i].getID() == id{
						
						(*((*locEmpMap)[pin]))[i].There = false
					}

				}
		}

	}
}


func deleteFullEmployeeList(employees *[]Employee){

*employees = (*employees)[:0]

}

func deleteFullIDempMap(idEmpMap *map[int]Employee){

	for key := range *idEmpMap {
		delete(*idEmpMap, key)
	}
	
	}

	func deleteFullDeptEmpMap(deptEmpMap *map[string]*([]Employee)){

		for key := range *deptEmpMap {
			delete(*deptEmpMap, key)
		}
		
		}

		func deleteFullLocEmpMap(locEmpMap *map[int]*([]Employee)){

			for key := range *locEmpMap {
				delete(*locEmpMap, key)
			}
			
			}


func deleteEmployeeByIDFromList(id int, employees *([]Employee)){

	for i:= 0; i < len(*employees); i++{

	if id == (*employees)[i].getID(){

		(*employees)[i] = (*employees)[len(*employees)-1]
		*employees = (*employees)[:len(*employees)-1]

		}

}

}

func deleteEmployeeByIDFromIdMap(id int, idEmpMap *map[int]Employee){


	 delete (*idEmpMap, id)
}

 func deleteEmployeeByIDFromDeptMap(id int, depts []string, deptIDMap *map[string]*([]Employee)){

 for _, dept:= range depts{

	
 	 employees  := (*deptIDMap)[dept]

	
 for i:= 0; i < len(*employees); i++{

 if id == (*employees)[i].getID(){

 	(*employees)[i] = (*employees)[len(*employees)-1]
 	*employees = (*employees)[:len(*employees)-1]
 	return

 }
 } 
 }

 }



 func deleteEmployeeByIDFromLocMap(id int, pins []int, locIDMap *map[int](*[]Employee)){

 	for _, pin := range pins{
	
 		employees := (*locIDMap)[pin]
 	for i:= 0; i < len(*employees); i++{
	
 	if id == (*employees)[i].getID(){
	
 		(*employees)[i] = (*employees)[len(*employees)-1]
		 *employees = (*employees)[:len(*employees)-1]
		
 		return
	
 	}
 	} 
	 }
	 
	
 	}


func storeEmployeesByDeptMap(employees *([]Employee)) map[string]*([]Employee){

	deptEmpMap := make(map[string]*([]Employee))

	for _, empl:= range *employees{

		deptList  := empl.getDept()


		for _, dept := range deptList{

			_, ok := deptEmpMap[dept]

			var emplswithThisDept *([]Employee)

			if(ok){

			emplswithThisDept = deptEmpMap[dept]

			} else{

				
				emptySlice := make([]Employee, 1)
				emplswithThisDept = &(emptySlice)
			}

			*emplswithThisDept = append(*emplswithThisDept, empl)
			deptEmpMap[dept] = emplswithThisDept

		}
	}

		return deptEmpMap
}

func storeEmployeesByIDMap(employees *([]Employee)) map[int]Employee{

	idEmpMap := make(map[int]Employee)

	for _, empl:= range *employees{

	
		idEmpMap[empl.getID()] = empl
		

	}


		return idEmpMap
}

func storeEmployeesByLocMap(employees *([]Employee)) map[int]*([]Employee){

	locEmpMap := make(map[int]*[](Employee))

	for _, empl:= range *employees{

		pins := empl.getPins()

		for _, pin := range pins{

			_, ok := locEmpMap[pin]

			var emplswithThisPin *([]Employee)

			if(ok){

			emplswithThisPin = locEmpMap[pin]

			} else{

				emptySlice := (make([]Employee, 1))

				emplswithThisPin = &emptySlice
			}

			*emplswithThisPin = append(*emplswithThisPin, empl)
			locEmpMap[pin] = emplswithThisPin


		}

	}

		return locEmpMap
}


func showEmployeeByID(id int, idEmpMap *map[int](Employee)){

		empl := (*idEmpMap)[id]
		if empl.There == true {empl.printName()}
}

func listAllEmployees(employees *([]Employee)) {

for _, empl:= range *employees{

if empl.There == true {empl.printName()}


}

}

func listEmployeesByDept(dept string, deptEmpMap *(map[string]*([]Employee)))  {


	_, ok := (*deptEmpMap)[dept]

	if(ok){

		employees := (*deptEmpMap)[dept]

	for _, empl:= range *employees{
	
	if empl.There == true {empl.printName()}
	
	
	} 
	
	}else{

		fmt.Println("No such department exists.")

	}


}


func listEmployeesByLoc(loc int, locEmpMap *map[int]*([]Employee))  {


	_, ok := (*locEmpMap)[loc]

	if(ok){

		employees := (*locEmpMap)[loc]
	for _, empl:= range *employees{
	
		 {if empl.There == true {empl.printName()}}
	
	
	} 
	
	}else{

		fmt.Println("No employee at that location.")

	}


}

