package test

import ("testing"
"store"
"fileProcess"

)

func TestParseFileAndStoreListOfEmployees(t *testing.T){

	var testEmployees []store.Employee

	fileProcess.ParseFileAndStoreListOfEmployees("employees.json", &testEmployees)

	if len(testEmployees) != 3{
		t.Error("Expected size of testEmployees: 3 got", len(testEmployees))}

if testEmployees[0].Name != "Pappu"{
	t.Error("Expected name of first employee: Pappu got", testEmployees[0].Name)}
	
if (testEmployees[1].GetDoorNos())[0] != 34{
	t.Error("Expected first doorno of second employee: 34 got", (testEmployees[1].GetDoorNos())[0])}

if testEmployees[2].GetID() != 3{
	t.Error("Expected id of third employee: 3 got", testEmployees[2].GetID())}

if 	testEmployees[1].GetDept()[1] != "Admin"{
	t.Error("Expected second department of second employee: Admin got", testEmployees[1].GetDept())}
if (testEmployees[2].GetLocalities())[1] != "Yelahanka"{
	t.Error("Expected locality of third employee's second address: Yelahanka got", (testEmployees[2].GetLocalities())[1])}

if (testEmployees[0].GetPins())[1] != 560042{
	t.Error("Expected pin of first employee's second address: 560042 got", (testEmployees[0].GetPins())[1])}

if  testEmployees[0].There == false{
	t.Error("Expected There value of first employee: true got", testEmployees[0].There)}



	
}