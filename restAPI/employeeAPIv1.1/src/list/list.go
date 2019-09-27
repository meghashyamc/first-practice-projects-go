package list

import ("store"
)





func ListAllEmployees(employees *([]store.Employee)) []store.Employee {

	list := make([]store.Employee, 0)
	for _, empl:= range *employees{
	
	if empl.There == true {list = append(list, empl)}
	
	}

	return list
	
	}
	
	func ListEmployeesByDept(dept string, deptEmpMap *(map[string]*([]store.Employee))) []store.Employee {
	
		// log.Println(dept)
		list := make([]store.Employee, 0)
		// log.Println(*deptEmpMap)
		_, ok := (*deptEmpMap)[dept]
	
		if(ok){
			

			employees := (*deptEmpMap)[dept]
	
		for _, empl:= range *employees{
		
		if empl.There == true {list = append(list, empl)
		}
	
		} 
		
		}

		return list
	}
	
	
	// func ListEmployeesByLocPrint(loc int, locEmpMap *map[int]*([]store.Employee)) []store.Employee  {
	
	// 	list := make([]store.Employee, 0)
	// 	_, ok := (*locEmpMap)[loc]
	
	// 	if(ok){
	
	// 		employees := (*locEmpMap)[loc]
	// 	for _, empl:= range *employees{
		
	// 		 {if empl.There == true {list = append(list, empl)}}
		
		
	// 	} 
		
	// 	}

	
	
	// }
	
	func ListEmployeesByLoc(loc int, locEmpMap *map[int]*([]store.Employee)) []store.Employee  {
		list := make([]store.Employee, 0)

		_, ok := (*locEmpMap)[loc]
	
			if(ok){
	
				employees := (*locEmpMap)[loc]
		
			for _, empl:= range *employees{
			
			if empl.There == true {list = append(list, empl)}
			
				 
		}
	}
	
		return list

	}
	
	func ListEmployeesByDoorNoAtLoc(doorno int, employeesAtLoc *([]store.Employee)) []store.Employee{
		list := make([]store.Employee, 0)

		for _, employee := range *employeesAtLoc{
	
	
			for _, thisDoorNo := range employee.GetDoorNos(){
	
	
				if thisDoorNo==doorno{
	
					if employee.There == true {list = append(list, employee)
			
					}
				}
			}
	
	}
	
	return list
	}
	
	func ListEmployeesByStreetAtLoc(street string, employeesAtLoc *([]store.Employee)) []store.Employee{
		list := make([]store.Employee, 0)

	
		for _, employee := range *employeesAtLoc{
		
		
				for _, thisStreet := range employee.GetStreets(){
		
		
					if thisStreet==street{
		
						if employee.There == true {list = append(list, employee)
						
					}
				}
		
		}
	
	}
		return list
		
		}
	
	
		func ListEmployeesByLocalityAtLoc(locality string, employeesAtLoc *([]store.Employee)) []store.Employee{
			list := make([]store.Employee, 0)

			
			for _, employee := range *employeesAtLoc{
			
			
					for _, thisLocality := range employee.GetLocalities(){
			
			
						if thisLocality==locality{
			
							if employee.There == true {list = append(list, employee)
							}
						}
					}
			
			}
	
			return list
			
			}
	
	
	
	
	