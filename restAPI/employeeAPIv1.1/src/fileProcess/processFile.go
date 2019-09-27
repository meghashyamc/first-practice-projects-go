package fileProcess

import ("log"
"store"
"io/ioutil"
"os"
"encoding/json"
"io"

)

func ParseFileAndStoreListOfEmployees(reqBody io.Reader, employees *([]store.Employee)){

// returns a byte slice
body, err := ioutil.ReadAll(reqBody) 

if err != nil {log.Println("Error reading request body.")

os.Exit(0)}

//stores corresponding fields in employees slice
json.Unmarshal([]byte(body), employees)

}
