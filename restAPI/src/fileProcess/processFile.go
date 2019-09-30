package fileProcess

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"store"
)

func ParseFileAndStoreListOfEmployees(reqBody io.Reader, moreEmployees *([]store.Employee)) (bool, error) {

	// returns a byte slice
	body, err := ioutil.ReadAll(reqBody)

	if err != nil {
		log.Println(err)
		return false, err

	}

	//stores corresponding fields in moreEmployees slice
	err1 := json.Unmarshal([]byte(body), moreEmployees)

	if err1 != nil {

		log.Println("Could not read JSON file.")
		return false, errors.New("Could not read JSON file.")
	}

	if !areCompulsoryFieldsThere(moreEmployees) {

		log.Println("Compulsory fields Name and Address were empty for some employees. Did not add any employee. Please complete the Name and Address fields and resubmit.")
		return false, errors.New("Compulsory fields Name and Address were empty for some employees. Did not add any employee. Please complete the Name and Address fields and resubmit.")
	}

	generateIDsIfNotThere(moreEmployees)

	store.Employees = append(store.Employees, *moreEmployees...)
	return true, nil
}

func generateIDsIfNotThere(moreEmployees *([]store.Employee)) {

	for i := 0; i < len(*moreEmployees); i++ {

		if (*moreEmployees)[i].GetID() == 0 {

			(*moreEmployees)[i].EmpID = generateID(store.MaxID, &store.IDset)
		}

	}

}

func areCompulsoryFieldsThere(moreEmployees *([]store.Employee)) bool {

	for _, empl := range *moreEmployees {
		if empl.Name == "" {

			return false
		}

		if len(empl.Addresses) == 0 {

			return false
		}

		if len(empl.GetDept()) == 0 {

			return false
		}
	}

	return true
}

func generateID(max int, set *(map[int]bool)) int {

	var randID int
	for {

		randID = rand.Intn(max + 100)
		if randID == 0 {
			continue
		}

		if (*set)[randID] == false {
			break
		}
	}

	return randID

}
