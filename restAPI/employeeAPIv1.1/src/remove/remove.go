package remove

import "store"

func RemoveEmployeesFromList(id int, employees *([]store.Employee)){


		for i:=0; i < len(*employees); i++{

				if ((*employees)[i]).GetID() == id{

					((*employees)[i]).There = false
				}
		}

}

func RemoveEmployeesFromIdEmpMap(id int, idEmpMap *map[int](store.Employee)){


		for mapId:= range *idEmpMap{

				if mapId == id{

					newEmpl := (*idEmpMap)[mapId]

					newEmpl.There = false
					(*idEmpMap)[mapId] = newEmpl
				}
		}

}

func RemoveEmployeesFromDeptEmpMap(id int, deptEmpMap *map[string]*([](store.Employee))){
	



		for dept:= range *deptEmpMap{

				for i:= 0; i < len(*((*deptEmpMap)[dept])); i++{

					if (*((*deptEmpMap)[dept]))[i].GetID() == id{
						
						(*((*deptEmpMap)[dept]))[i].There = false
					}

				}
		}

	
}


func RemoveEmployeesFromLocEmpMap(id int, locEmpMap *map[int]*([]store.Employee)){
	


		for pin:= range *locEmpMap{

				for i:= 0; i < len(*((*locEmpMap)[pin])); i++{

					if (*((*locEmpMap)[pin]))[i].GetID() == id{
						
						(*((*locEmpMap)[pin]))[i].There = false
					}

				}
		}

}
