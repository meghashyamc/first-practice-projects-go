package fileProcess

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"kiplog"
	"math/rand"
	"store"
	"time"
)

func ParseFileAndStoreListOfEmployees(reqBody io.Reader, moreEmployees *([]store.Employee)) (bool, error) {

	// returns a byte slice
	body, err := ioutil.ReadAll(reqBody)

	if err != nil {
		kiplog.HTTPLog(err.Error())
		return false, err

	}

	//stores corresponding fields in moreEmployees slice
	err1 := json.Unmarshal([]byte(body), moreEmployees)

	if err1 != nil {

		kiplog.HTTPLog("Could not read JSON file.")
		return false, errors.New("Could not read JSON file.")
	}

	err2 := areCompulsoryFieldsThere(moreEmployees)

	if err2 != nil {
		return false, err2
	}

	generateIDsIfNotThere(moreEmployees)
	addTimeStamps(moreEmployees)

	store.Employees = append(store.Employees, *moreEmployees...)
	return true, nil
}

func addTimeStamps(moreEmployees *([]store.Employee)) {

	for i := 0; i < len(*moreEmployees); i++ {

		if (*moreEmployees)[i].Timestamp == 0 {

			(*moreEmployees)[i].Timestamp = time.Now().Unix()
		}

	}
}

func generateIDsIfNotThere(moreEmployees *([]store.Employee)) {

	for i := 0; i < len(*moreEmployees); i++ {

		if (*moreEmployees)[i].GetID() == 0 {

			(*moreEmployees)[i].EmpID = generateID(store.MaxID, &store.IDset)
		}

	}

}

func areCompulsoryFieldsThere(moreEmployees *([]store.Employee)) error {

	for _, empl := range *moreEmployees {
		if empl.Name == "" {

			return errors.New("Name of employee not specified. Please add it.")
		}

		if len(empl.Addresses) == 0 {

			return errors.New("Zero addresses were specified for " + empl.Name + ". Atleast one address must be specified.")
		}

		if len(empl.GetDept()) == 0 {

			return errors.New("Zero departments were specified for " + empl.Name + ". Atleast one department must be specified.")
		}
	}

	return nil
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
