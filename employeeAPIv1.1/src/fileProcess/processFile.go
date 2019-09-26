package fileProcess

import ("fmt"
"store"
"io/ioutil"
"os"
"encoding/json"

)

func ParseFileAndStoreListOfEmployees(fileJson string, employees *([]store.Employee)){

// returns a byte slice
file, err := ioutil.ReadFile(fileJson) 

if err != nil {fmt.Println("Error parsing json file.")

os.Exit(0)}

//stores corresponding fields in employees slice
json.Unmarshal([]byte(file), employees)

}
