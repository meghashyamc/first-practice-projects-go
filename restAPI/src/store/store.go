package store

import (
	"sync"
	"time"
)

var MaxID int
var IDset map[int]bool
var Employees []Employee
var IdEmpMap map[int]Employee
var DeptEmpMap map[string]*([]Employee)
var LocEmpMap map[int]*([]Employee)
var wg sync.WaitGroup

type Address struct {
	Doorno   int
	Street   string
	Locality string
	PIN      int
}

//helper methods for Address
func (a *Address) getPin() int {

	return a.PIN
}

func (a *Address) GetDoorNo() int {

	return a.Doorno
}

func (a *Address) GetStreet() string {

	return a.Street
}

func (a *Address) GetLocality() string {

	return a.Locality
}

type Employee struct {
	EmpID      int
	Name       string
	Department []string
	Addresses  []Address
	There      bool
	Timestamp  time.Time
}

//helper methods for Employee

func (e *Employee) GetID() int {

	return e.EmpID
}

func (e *Employee) GetPins() []int {

	var pins []int

	for _, addr := range e.Addresses {

		pins = append(pins, addr.getPin())

	}

	return pins
}

func (e *Employee) GetDoorNos() []int {

	var doornos []int

	for _, addr := range e.Addresses {

		doornos = append(doornos, addr.GetDoorNo())

	}

	return doornos
}

func (e *Employee) GetStreets() []string {

	var streets []string

	for _, addr := range e.Addresses {

		streets = append(streets, addr.GetStreet())

	}

	return streets
}

func (e *Employee) GetLocalities() []string {

	var localities []string

	for _, addr := range e.Addresses {

		localities = append(localities, addr.GetLocality())

	}

	return localities
}

func (e *Employee) GetDept() []string {

	return e.Department
}

func InitializeEmployeesAndMaps() {

	Employees = make([]Employee, 0)

	IdEmpMap = make(map[int]Employee)
	DeptEmpMap = make(map[string]*([]Employee))
	LocEmpMap = make(map[int]*([]Employee))
	IDset = make(map[int]bool)
	MaxID = 0

}
func StoreEmployeesByIdDeptAndLoc(moreEmployees *([]Employee)) {

	wg.Add(4)
	go storeIDs(moreEmployees)
	//stores ids as keys and employees as values

	go storeEmployeesByIDMap(moreEmployees)

	//stores department names (strings) as keys and reference to array of employees as values
	go storeEmployeesByDeptMap(moreEmployees)

	//stores PIN codes (int) as keys and reference to array of employees as values

	go storeEmployeesByLocMap(moreEmployees)

	wg.Wait()
}

func storeIDs(moreEmployees *([]Employee)) {

	defer wg.Done()
	for _, empl := range *moreEmployees {

		IDset[empl.GetID()] = true

		if empl.GetID() > MaxID {
			MaxID = empl.GetID()
		}

	}

}

func storeEmployeesByDeptMap(moreEmployees *([]Employee)) {
	defer wg.Done()
	for _, empl := range *moreEmployees {

		deptList := empl.GetDept()

		for _, dept := range deptList {

			_, ok := DeptEmpMap[dept]

			var emplswithThisDept *([]Employee)

			if ok {

				emplswithThisDept = DeptEmpMap[dept]

			} else {

				emptySlice := make([]Employee, 0)
				emplswithThisDept = &(emptySlice)
			}

			*emplswithThisDept = append(*emplswithThisDept, empl)
			DeptEmpMap[dept] = emplswithThisDept

		}
	}

}

func storeEmployeesByIDMap(moreEmployees *([]Employee)) {
	defer wg.Done()
	for _, empl := range *moreEmployees {

		IdEmpMap[empl.GetID()] = empl

	}

}

func storeEmployeesByLocMap(moreEmployees *([]Employee)) {
	defer wg.Done()
	for _, empl := range *moreEmployees {

		pins := empl.GetPins()

		for _, pin := range pins {

			_, ok := LocEmpMap[pin]

			var emplswithThisPin *([]Employee)

			if ok {

				emplswithThisPin = LocEmpMap[pin]

			} else {

				emptySlice := (make([]Employee, 0))

				emplswithThisPin = &emptySlice
			}

			*emplswithThisPin = append(*emplswithThisPin, empl)
			LocEmpMap[pin] = emplswithThisPin

		}

	}

}
