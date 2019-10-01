package delete

import (
	"errors"
	"store"
)

var delError error = errors.New("dfgfdgfd")

func DeleteFromEverywhere() error {

	var err error
	err = DeleteFullEmployeeList(&(store.Employees))
	if err != nil {
		return err
	}
	err = DeleteFullIDempMap(&(store.IdEmpMap))
	if err != nil {
		return err
	}
	err = DeleteFullDeptEmpMap(&(store.DeptEmpMap))
	if err != nil {
		return err
	}
	err = DeleteFullLocEmpMap(&(store.LocEmpMap))
	if err != nil {
		return err
	}

	return nil

}

func DeleteByIDFromEverywhere(id int) error {

	var err error
	empl, ok := (store.IdEmpMap)[id]

	if !ok {

		return errors.New(cantDeleteID())
	}

	err = DeleteEmployeeByIDFromList(id, &(store.Employees))
	if err != nil {
		return err
	}
	err = DeleteEmployeeByIDFromDeptMap(id, empl.GetDept(), &(store.DeptEmpMap))
	if err != nil {
		return err
	}

	err = DeleteEmployeeByIDFromLocMap(id, empl.GetPins(), &(store.LocEmpMap))
	if err != nil {
		return err
	}

	err = DeleteEmployeeByIDFromIdMap(id, &(store.IdEmpMap))
	if err != nil {
		return err
	}

	return nil
}

func DeleteFullEmployeeList(employees *[]store.Employee) error {

	if len(*employees) == 0 {

		return errors.New(cantDelete())
	}
	*employees = (*employees)[:0]
	return nil

}

func DeleteFullIDempMap(idEmpMap *map[int]store.Employee) error {

	if len(*idEmpMap) == 0 {

		return errors.New(cantDelete())
	}
	for key := range *idEmpMap {
		delete(*idEmpMap, key)
	}

	return nil

}

func DeleteFullDeptEmpMap(deptEmpMap *map[string]*([]store.Employee)) error {
	if len(*deptEmpMap) == 0 {

		return errors.New(cantDelete())
	}
	for key := range *deptEmpMap {
		delete(*deptEmpMap, key)
	}

	return nil

}

func DeleteFullLocEmpMap(locEmpMap *map[int]*([]store.Employee)) error {
	if len(*locEmpMap) == 0 {

		return errors.New(cantDelete())
	}
	for key := range *locEmpMap {
		delete(*locEmpMap, key)
	}

	return nil

}

func DeleteEmployeeByIDFromList(id int, employees *([]store.Employee)) error {

	count := 0
	for i := 0; i < len(*employees); i++ {

		if id == (*employees)[i].GetID() {

			(*employees)[i] = (*employees)[len(*employees)-1]
			*employees = (*employees)[:len(*employees)-1]
			count++
		}

	}

	if count == 0 {
		return errors.New(cantDeleteID())

	} else {
		return nil

	}
}
func DeleteEmployeeByIDFromIdMap(id int, idEmpMap *map[int]store.Employee) error {

	_, ok := (*idEmpMap)[id]

	if !ok {
		return errors.New(cantDeleteID())
	}
	delete(*idEmpMap, id)
	return nil
}

func DeleteEmployeeByIDFromDeptMap(id int, depts []string, deptIDMap *map[string]*([]store.Employee)) error {

	count := 0
	for _, dept := range depts {

		employees := (*deptIDMap)[dept]

		for i := 0; i < len(*employees); i++ {

			if id == (*employees)[i].GetID() {

				(*employees)[i] = (*employees)[len(*employees)-1]
				*employees = (*employees)[:len(*employees)-1]
				count++
			}
		}
	}
	if count == 0 {
		return errors.New(cantDeleteID())

	} else {
		return nil

	}

}

func DeleteEmployeeByIDFromLocMap(id int, pins []int, locIDMap *map[int](*[]store.Employee)) error {

	count := 0
	for _, pin := range pins {

		employees := (*locIDMap)[pin]
		for i := 0; i < len(*employees); i++ {

			if id == (*employees)[i].GetID() {

				(*employees)[i] = (*employees)[len(*employees)-1]
				*employees = (*employees)[:len(*employees)-1]
				count++
			}
		}
	}

	if count == 0 {
		return errors.New(cantDeleteID())

	} else {
		return nil

	}

}

func cantDelete() string {

	return ("Could not delete: there were no employees to delete.")
}

func cantDeleteID() string {

	return ("Could not delete: there were no employee with that ID.")
}
