

package delete


import ("store"

)

func DeleteFullEmployeeList(employees *[]store.Employee){

	*employees = (*employees)[:0]
	
	}
	
	func DeleteFullIDempMap(idEmpMap *map[int]store.Employee){
	
		for key := range *idEmpMap {
			delete(*idEmpMap, key)
		}
		
		}
	
		func DeleteFullDeptEmpMap(deptEmpMap *map[string]*([]store.Employee)){
	
			for key := range *deptEmpMap {
				delete(*deptEmpMap, key)
			}
			
			}
	
			func DeleteFullLocEmpMap(locEmpMap *map[int]*([]store.Employee)){
	
				for key := range *locEmpMap {
					delete(*locEmpMap, key)
				}
				
				}
	
	
	func DeleteEmployeeByIDFromList(id int, employees *([]store.Employee)){
	
		for i:= 0; i < len(*employees); i++{
	
		if id == (*employees)[i].GetID(){
	
			(*employees)[i] = (*employees)[len(*employees)-1]
			*employees = (*employees)[:len(*employees)-1]
	
			}
	
	}
	
	}
	
	func DeleteEmployeeByIDFromIdMap(id int, idEmpMap *map[int]store.Employee){
	
	
		 delete (*idEmpMap, id)
	}
	
	 func DeleteEmployeeByIDFromDeptMap(id int, depts []string, deptIDMap *map[string]*([]store.Employee)){
	
	 for _, dept:= range depts{
	
		
		  employees  := (*deptIDMap)[dept]
	
		
	 for i:= 0; i < len(*employees); i++{
	
	 if id == (*employees)[i].GetID(){
	
		 (*employees)[i] = (*employees)[len(*employees)-1]
		 *employees = (*employees)[:len(*employees)-1]
		 
	
	 }
	 } 
	 }
	
	 }
	
	
	
	 func DeleteEmployeeByIDFromLocMap(id int, pins []int, locIDMap *map[int](*[]store.Employee)){
	
		 for _, pin := range pins{
		
			 employees := (*locIDMap)[pin]
		 for i:= 0; i < len(*employees); i++{
		
		 if id == (*employees)[i].GetID(){
		
			 (*employees)[i] = (*employees)[len(*employees)-1]
			 *employees = (*employees)[:len(*employees)-1]
			
			 
		
		 }
		 } 
		 }
		 
		
		 }
	
	
	