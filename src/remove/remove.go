package remove

import (
	"errors"
	"store"
)

func RemoveEmployeesByIDEverywhere(id int) error {

	_, ok := (store.IdEmpMap)[id]

	if !ok {

		return errors.New(doesNotExist())
	}

	var err error
	err = RemoveEmployeesFromList(id, &(store.Employees))
	if err != nil {

		return err
	}
	err = RemoveEmployeesFromIdEmpMap(id, &(store.IdEmpMap))
	if err != nil {

		return err
	}
	err = RemoveEmployeesFromDeptEmpMap(id, &(store.DeptEmpMap))
	if err != nil {

		return err
	}
	err = RemoveEmployeesFromLocEmpMap(id, &(store.LocEmpMap))
	if err != nil {

		return err
	}

	return nil

}
func RemoveEmployeesFromList(id int, employees *([]store.Employee)) error {

	for i := 0; i < len(*employees); i++ {

		if ((*employees)[i]).GetID() == id {

			if ((*employees)[i]).There == false {

				return errors.New(alreadyRemoved())
			} else {
				((*employees)[i]).There = false
			}
		}

	}

	return nil
}

func RemoveEmployeesFromIdEmpMap(id int, idEmpMap *map[int](store.Employee)) error {

	for mapId := range *idEmpMap {

		if mapId == id {

			newEmpl := (*idEmpMap)[mapId]

			if newEmpl.There == false {
				return errors.New(alreadyRemoved())

			}
			newEmpl.There = false
			(*idEmpMap)[mapId] = newEmpl
		}
	}

	return nil

}

func RemoveEmployeesFromDeptEmpMap(id int, deptEmpMap *map[string]*([](store.Employee))) error {

	for dept := range *deptEmpMap {

		for i := 0; i < len(*((*deptEmpMap)[dept])); i++ {

			if (*((*deptEmpMap)[dept]))[i].GetID() == id {

				if (*((*deptEmpMap)[dept]))[i].There == false {

					return errors.New(alreadyRemoved())
				}
				(*((*deptEmpMap)[dept]))[i].There = false
			}

		}
	}

	return nil
}

func RemoveEmployeesFromLocEmpMap(id int, locEmpMap *map[int]*([]store.Employee)) error {

	for pin := range *locEmpMap {

		for i := 0; i < len(*((*locEmpMap)[pin])); i++ {

			if (*((*locEmpMap)[pin]))[i].GetID() == id {

				if (*((*locEmpMap)[pin]))[i].There == false {

					return errors.New(alreadyRemoved())
				}
				(*((*locEmpMap)[pin]))[i].There = false
			}

		}
	}

	return nil

}

func doesNotExist() string {

	return "Could not remove: An employee with this id does not exist."

}

func alreadyRemoved() string {

	return "Could not remove: This employee has already been removed."

}
