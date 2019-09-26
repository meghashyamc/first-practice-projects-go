package show

import "store"

func ShowEmployeeByID(id int, idEmpMap *map[int](store.Employee)){

	empl := (*idEmpMap)[id]
	if empl.There == true {empl.PrintName()}
}