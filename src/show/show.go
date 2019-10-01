package show

import (
	"errors"
	"store"
)

func ShowEmployeeByID(id int, idEmpMap *map[int](store.Employee)) (store.Employee, error) {
	returnEmpl := new(store.Employee)
	empl := (*idEmpMap)[id]
	if empl.There == true {
		returnEmpl = &empl
	}

	if (*returnEmpl).GetID() == 0 {

		return (*returnEmpl), errors.New("There is no employee with that ID.")
	} else {
		return *returnEmpl, nil
	}
}
