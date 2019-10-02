package list

import (
	"errors"
	"sort"
	"store"
	"strconv"
	"strings"
)

func ListAllEmployees(employees *([]store.Employee)) ([]store.Employee, error) {

	list := make([]store.Employee, 0)
	for _, empl := range *employees {

		if empl.There == true {
			list = append(list, empl)
		}

	}

	if len(list) == 0 {
		return list, errors.New("No employees exist as of now.")
	} else {

		sortByTimestamp(&list)
		return list, nil
	}

}

func ListSearchEmployees(term string, employees *([]store.Employee)) ([]store.Employee, error) {

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

				for _, dept := range empl.GetDept() {

					if isSubstring(term, dept) {
						list = append(list, empl)
						break

					}
				}

			}

		}

	}

	if len(list) == 0 {

		return list, errors.New("No employees matched the search term " + term)
	} else {

		if len(list) > 1 {

			sortByTimestamp(&list)

		}
		return list, nil
	}
}
func ListEmployeesByDept(dept string, deptEmpMap *(map[string]*([]store.Employee))) ([]store.Employee, error) {

	list := make([]store.Employee, 0)

	_, ok := (*deptEmpMap)[dept]

	if ok {

		employees := (*deptEmpMap)[dept]

		for _, empl := range *employees {

			if empl.There == true {
				list = append(list, empl)
			}

		}

		if len(list) == 0 {

			return list, errors.New("This department exists but it has no employees.")

		} else {

			if len(list) > 1 {

				sortByTimestamp(&list)

			}
			return list, nil

		}
	} else {

		return list, errors.New("This department does not exist.")
	}
}

func ListEmployeesByLoc(loc int, locEmpMap *map[int]*([]store.Employee)) ([]store.Employee, error) {
	list := make([]store.Employee, 0)

	_, ok := (*locEmpMap)[loc]

	if ok {

		employees := (*locEmpMap)[loc]

		for _, empl := range *employees {

			if empl.There == true {
				list = append(list, empl)
			}

		}
	} else {
		return list, errors.New("No employee lives at this pincode")

	}

	if len(list) == 0 {
		return list, errors.New("No employee lives at this pincode")
	} else {

		if len(list) > 1 {

			sortByTimestamp(&list)

		}
		return list, nil

	}
}
func ListEmployeesByDoorNoAtLoc(doorno int, employeesAtLoc *([]store.Employee)) ([]store.Employee, error) {
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

	if len(list) == 0 {
		return list, errors.New("No employee has this door number.")
	} else {
		if len(list) > 1 {

			sortByTimestamp(&list)

		}
		return list, nil

	}
}

func ListEmployeesByStreetAtLoc(street string, employeesAtLoc *([]store.Employee)) ([]store.Employee, error) {
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
	if len(list) == 0 {
		return list, errors.New("No employee lives in this street.")

	} else {

		if len(list) > 1 {

			sortByTimestamp(&list)

		}
		return list, nil

	}
}

func ListEmployeesByLocalityAtLoc(locality string, employeesAtLoc *([]store.Employee)) ([]store.Employee, error) {
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

	if len(list) == 0 {
		return list, errors.New("No employee lives in this locality.")

	} else {

		if len(list) > 1 {

			sortByTimestamp(&list)

		}
		return list, nil

	}

}

func isSubstring(s1 string, s2 string) bool {

	s2 = strings.ToLower(s2)
	s1 = strings.ToLower(s1)

	return strings.Contains(s2, s1)

}

func sortByTimestamp(list *([]store.Employee)) {

	sort.Slice(*list, func(i, j int) bool {
		return ((*list)[i]).Timestamp.Unix() < ((*list)[j]).Timestamp.Unix()
	})

}
