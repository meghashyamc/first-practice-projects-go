package list

import (
	"errors"
	"store"
	"strconv"
	"strings"
)

func ListAllEmployees(employees *([]store.Employee)) []store.Employee {

	list := make([]store.Employee, 0)
	for _, empl := range *employees {

		if empl.There == true {
			list = append(list, empl)
		}

	}

	return list

}

func ListSearchEmployees(term string, employees *([]store.Employee)) []store.Employee {

	// log.Println(dept)
	list := make([]store.Employee, 0)

	isNum := false
	numTerm, err := strconv.Atoi(term)

	if err == nil {
		isNum = true
	}

	for _, empl := range *employees {

		if empl.There == true {

			if isNum {

				if empl.GetID() == numTerm {
					list = append(list, empl)
					continue
				}

				for _, addr := range empl.Addresses {

					if addr.Doorno == numTerm ||
						addr.PIN == numTerm {
						list = append(list, empl)
						break

					}
				}
			} else {

				if isSubstring(term, empl.Name) {
					list = append(list, empl)
					continue
				}

				for _, addr := range empl.Addresses {

					if isSubstring(term, addr.Street) ||
						isSubstring(term, addr.Locality) {
						list = append(list, empl)
						break

					}
				}

			}

		}

	}

	return list
}
func ListEmployeesByDept(dept string, deptEmpMap *(map[string]*([]store.Employee))) ([]store.Employee, error) {

	// log.Println(dept)
	list := make([]store.Employee, 0)
	// log.Println(*deptEmpMap)
	_, ok := (*deptEmpMap)[dept]

	if ok {

		employees := (*deptEmpMap)[dept]

		for _, empl := range *employees {

			if empl.There == true {
				list = append(list, empl)
			}

		}

		return list, nil

	} else {

		return list, errors.New("This department does not exist.")
	}
}

// func ListEmployeesByLocPrint(loc int, locEmpMap *map[int]*([]store.Employee)) []store.Employee  {

// 	list := make([]store.Employee, 0)
// 	_, ok := (*locEmpMap)[loc]

// 	if(ok){

// 		employees := (*locEmpMap)[loc]
// 	for _, empl:= range *employees{

// 		 {if empl.There == true {list = append(list, empl)}}

// 	}

// 	}

// }

func ListEmployeesByLoc(loc int, locEmpMap *map[int]*([]store.Employee)) []store.Employee {
	list := make([]store.Employee, 0)

	_, ok := (*locEmpMap)[loc]

	if ok {

		employees := (*locEmpMap)[loc]

		for _, empl := range *employees {

			if empl.There == true {
				list = append(list, empl)
			}

		}
	}

	return list

}

func ListEmployeesByDoorNoAtLoc(doorno int, employeesAtLoc *([]store.Employee)) []store.Employee {
	list := make([]store.Employee, 0)

	for _, employee := range *employeesAtLoc {

		for _, thisDoorNo := range employee.GetDoorNos() {

			if thisDoorNo == doorno {

				if employee.There == true {
					list = append(list, employee)

				}
			}
		}

	}

	return list
}

func ListEmployeesByStreetAtLoc(street string, employeesAtLoc *([]store.Employee)) []store.Employee {
	list := make([]store.Employee, 0)

	for _, employee := range *employeesAtLoc {

		for _, thisStreet := range employee.GetStreets() {

			if thisStreet == street {

				if employee.There == true {
					list = append(list, employee)

				}
			}

		}

	}
	return list

}

func ListEmployeesByLocalityAtLoc(locality string, employeesAtLoc *([]store.Employee)) []store.Employee {
	list := make([]store.Employee, 0)

	for _, employee := range *employeesAtLoc {

		for _, thisLocality := range employee.GetLocalities() {

			if thisLocality == locality {

				if employee.There == true {
					list = append(list, employee)
				}
			}
		}

	}

	return list

}

func isSubstring(s1 string, s2 string) bool {

	s2 = strings.ToLower(s2)
	s1 = strings.ToLower(s1)
	for i := 0; i < len(s2)-len(s1)+1; i++ {

		if s1 == s2[i:i+len(s1)] {

			return true
		}
	}

	return false

}
