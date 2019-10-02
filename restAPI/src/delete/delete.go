package delete

import (
	"errors"
	"store"
	"sync"
)

var delError error = errors.New("Could not delete: there were no employees to delete.")
var delErrorID error = errors.New("Could not delete: there was no employee with that ID.")

var (
	mu  sync.Mutex
	err error
)

var wg sync.WaitGroup

func DeleteFromEverywhere() error {
	err = nil
	wg.Add(4)
	go DeleteFullEmployeeList(&(store.Employees))

	go DeleteFullIDempMap(&(store.IdEmpMap))

	go DeleteFullDeptEmpMap(&(store.DeptEmpMap))

	go DeleteFullLocEmpMap(&(store.LocEmpMap))

	wg.Wait()
	return err

}

func DeleteByIDFromEverywhere(id int) error {

	err = nil

	empl, ok := (store.IdEmpMap)[id]

	if !ok {

		return delErrorID
	}
	wg.Add(4)
	go DeleteEmployeeByIDFromList(id, &(store.Employees))

	go DeleteEmployeeByIDFromDeptMap(id, empl.GetDept(), &(store.DeptEmpMap))

	go DeleteEmployeeByIDFromLocMap(id, empl.GetPins(), &(store.LocEmpMap))

	go DeleteEmployeeByIDFromIdMap(id, &(store.IdEmpMap))

	wg.Wait()
	return err
}

func DeleteFullEmployeeList(employees *[]store.Employee) {
	defer wg.Done()

	mu.Lock()
	if err != nil {
		mu.Unlock()
		return
	}
	mu.Unlock()

	if len(*employees) == 0 {
		mu.Lock()
		err = delError
		mu.Unlock()
	}
	*employees = (*employees)[:0]

}

func DeleteFullIDempMap(idEmpMap *map[int]store.Employee) {
	defer wg.Done()
	mu.Lock()
	if err != nil {
		mu.Unlock()
		return
	}
	mu.Unlock()
	if len(*idEmpMap) == 0 {
		mu.Lock()
		err = delError

		mu.Unlock()
	}
	for key := range *idEmpMap {
		delete(*idEmpMap, key)
	}

}

func DeleteFullDeptEmpMap(deptEmpMap *map[string]*([]store.Employee)) {
	defer wg.Done()
	mu.Lock()
	if err != nil {
		mu.Unlock()
		return
	}
	mu.Unlock()

	if len(*deptEmpMap) == 0 {
		mu.Lock()
		err = delError

		mu.Unlock()
	}
	for key := range *deptEmpMap {
		delete(*deptEmpMap, key)
	}

}

func DeleteFullLocEmpMap(locEmpMap *map[int]*([]store.Employee)) {
	defer wg.Done()
	mu.Lock()
	if err != nil {
		mu.Unlock()
		return
	}
	mu.Unlock()

	if len(*locEmpMap) == 0 {
		mu.Lock()
		err = delError
		mu.Unlock()
	}
	for key := range *locEmpMap {
		delete(*locEmpMap, key)
	}

}

func DeleteEmployeeByIDFromList(id int, employees *([]store.Employee)) {
	defer wg.Done()
	mu.Lock()
	if err != nil {
		mu.Unlock()
		return
	}
	mu.Unlock()
	count := 0
	for i := 0; i < len(*employees); i++ {

		if id == (*employees)[i].GetID() {

			(*employees)[i] = (*employees)[len(*employees)-1]
			*employees = (*employees)[:len(*employees)-1]
			count++

		}

	}

	if count == 0 {
		mu.Lock()
		err = delErrorID

		mu.Unlock()

	}
}
func DeleteEmployeeByIDFromIdMap(id int, idEmpMap *map[int]store.Employee) {
	defer wg.Done()
	mu.Lock()
	if err != nil {
		mu.Unlock()
		return
	}
	mu.Unlock()
	_, ok := (*idEmpMap)[id]

	if !ok {
		mu.Lock()
		err = delErrorID

		mu.Unlock()
	}
	delete(*idEmpMap, id)

}

func DeleteEmployeeByIDFromDeptMap(id int, depts []string, deptIDMap *map[string]*([]store.Employee)) {
	defer wg.Done()
	mu.Lock()
	if err != nil {
		mu.Unlock()
		return
	}
	mu.Unlock()
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
		mu.Lock()
		err = delErrorID

		mu.Unlock()
	}

}

func DeleteEmployeeByIDFromLocMap(id int, pins []int, locIDMap *map[int](*[]store.Employee)) {
	defer wg.Done()
	mu.Lock()
	if err != nil {
		mu.Unlock()
		return
	}
	mu.Unlock()
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
		mu.Lock()
		err = delErrorID

		mu.Unlock()
	}

}
