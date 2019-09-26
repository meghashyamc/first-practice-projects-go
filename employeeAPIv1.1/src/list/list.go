package list

import (

"store"
"fmt"	
)



func ListAllEmployees(employees *([]store.Employee)) {

	for _, empl:= range *employees{
	
	if empl.There == true {empl.PrintName()}
	
	
	}
	
	}
	
	func ListEmployeesByDept(dept string, deptEmpMap *(map[string]*([]store.Employee)))  {
	
	
		_, ok := (*deptEmpMap)[dept]
	
		if(ok){
	
			employees := (*deptEmpMap)[dept]
	
		for _, empl:= range *employees{
		
		if empl.There == true {empl.PrintName()}
		
		
		} 
		
		}else{
	
			fmt.Println("No such department exists.")
	
		}
	
	
	}
	
	
	func ListEmployeesByLocPrint(loc int, locEmpMap *map[int]*([]store.Employee))  {
	
	
		_, ok := (*locEmpMap)[loc]
	
		if(ok){
	
			employees := (*locEmpMap)[loc]
		for _, empl:= range *employees{
		
			 {if empl.There == true {empl.PrintName()}}
		
		
		} 
		
		}else{
	
			fmt.Println("No employee at that location.")
	
		}
	
	
	}
	
	func ListEmployeesByLoc(loc int, locEmpMap *map[int]*([]store.Employee)) (*([]store.Employee), bool)  {
	
	
		_, ok := (*locEmpMap)[loc]
	
		if(ok){
	
			return (*locEmpMap)[loc], true
		
		
		}else{
	
			fmt.Println("No employee at that location.")
			return nil, false
	
		}
	
	
	}
	
	func ListEmployeesByDoorNoAtLoc(doorno int, employeesAtLoc *([]store.Employee)){
	
		var count int
	
		for _, employee := range *employeesAtLoc{
	
	
			for _, thisDoorNo := range employee.GetDoorNos(){
	
	
				if thisDoorNo==doorno{
	
					if employee.There == true {employee.PrintName()
					count++
					}
				}
			}
	
	}
	
	if count == 0 {fmt.Println("No employee with this door number")}
	}
	
	func ListEmployeesByStreetAtLoc(street string, employeesAtLoc *([]store.Employee)){
		var count int
	
		for _, employee := range *employeesAtLoc{
		
		
				for _, thisStreet := range employee.GetStreets(){
		
		
					if thisStreet==street{
		
						if employee.There == true {employee.PrintName()
						count++}
					}
				}
		
		}
	
		if count == 0 {fmt.Println("No employee has this street listed in his address")}
	
		
		}
	
	
		func ListEmployeesByLocalityAtLoc(locality string, employeesAtLoc *([]store.Employee)){
	
			var count int
			for _, employee := range *employeesAtLoc{
			
			
					for _, thisLocality := range employee.GetLocalities(){
			
			
						if thisLocality==locality{
			
							if employee.There == true {employee.PrintName()
							count++}
						}
					}
			
			}
	
			if count == 0 {fmt.Println("No employee has this locality listed in his address")}
	
			
			}
	
	
	
	
	