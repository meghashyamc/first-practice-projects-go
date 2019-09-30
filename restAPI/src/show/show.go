package show

import "store"

func ShowEmployeeByID(id int, idEmpMap *map[int](store.Employee)) store.Employee {
	returnEmpl := new(store.Employee)
	empl := (*idEmpMap)[id]
	if empl.There == true {
		returnEmpl = &empl
	}

	return *returnEmpl
}
