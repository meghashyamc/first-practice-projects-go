package remove

import (
	"errors"
	"store"
	"sync"
)

var (
	mu  sync.Mutex
	err error
)
var wg sync.WaitGroup
var removeErrorRemoved error = errors.New("Could not remove: This employee has already been removed.")
var removeErrorNotFound error = errors.New("Could not remove: An employee with this id does not exist.")

func RemoveEmployeesByIDEverywhere(id int) error {

	_, ok := (store.IdEmpMap)[id]

	if !ok {

		return removeErrorRemoved
	}

	err = nil

	wg.Add(4)

	go RemoveEmployeesFromList(id, &(store.Employees))

	go RemoveEmployeesFromIdEmpMap(id, &(store.IdEmpMap))

	go RemoveEmployeesFromDeptEmpMap(id, &(store.DeptEmpMap))

	go RemoveEmployeesFromLocEmpMap(id, &(store.LocEmpMap))

	wg.Wait()
	return err

}
func RemoveEmployeesFromList(id int, employees *([]store.Employee)) {

	defer wg.Done()
	mu.Lock()
	if err != nil {
		mu.Unlock()
		return
	}
	mu.Unlock()
	for i := 0; i < len(*employees); i++ {

		if ((*employees)[i]).GetID() == id {

			if ((*employees)[i]).There == false {
				mu.Lock()

				err = removeErrorRemoved
				mu.Unlock()

			} else {
				((*employees)[i]).There = false
			}
		}

	}

}

func RemoveEmployeesFromIdEmpMap(id int, idEmpMap *map[int](store.Employee)) {

	defer wg.Done()

	mu.Lock()
	if err != nil {
		mu.Unlock()
		return
	}
	mu.Unlock()
	for mapId := range *idEmpMap {

		if mapId == id {

			newEmpl := (*idEmpMap)[mapId]

			if newEmpl.There == false {
				mu.Lock()

				err = removeErrorRemoved
				mu.Unlock()

			}
			newEmpl.There = false
			(*idEmpMap)[mapId] = newEmpl
		}
	}

}

func RemoveEmployeesFromDeptEmpMap(id int, deptEmpMap *map[string]*([](store.Employee))) {
	defer wg.Done()
	mu.Lock()
	if err != nil {
		mu.Unlock()
		return
	}
	mu.Unlock()
	for dept := range *deptEmpMap {

		for i := 0; i < len(*((*deptEmpMap)[dept])); i++ {

			if (*((*deptEmpMap)[dept]))[i].GetID() == id {

				if (*((*deptEmpMap)[dept]))[i].There == false {
					mu.Lock()

					err = removeErrorRemoved
					mu.Unlock()

				}
				(*((*deptEmpMap)[dept]))[i].There = false
			}

		}
	}

}

func RemoveEmployeesFromLocEmpMap(id int, locEmpMap *map[int]*([]store.Employee)) {
	defer wg.Done()
	mu.Lock()
	if err != nil {
		mu.Unlock()
		return
	}
	mu.Unlock()
	for pin := range *locEmpMap {

		for i := 0; i < len(*((*locEmpMap)[pin])); i++ {

			if (*((*locEmpMap)[pin]))[i].GetID() == id {

				if (*((*locEmpMap)[pin]))[i].There == false {
					mu.Lock()

					err = removeErrorRemoved
					mu.Unlock()

				}
				(*((*locEmpMap)[pin]))[i].There = false
			}

		}
	}

}
