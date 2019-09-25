# goLangPractice
Practice project in Golang to read and process employee data from JSON file

* Run the file (go run main.go).
* Employee data is read from the json file.

* The following commands are supported.


<b>LIST</b>
* list all : lists names of all employees in the json file
* list department <i>NAME</i>: lists employees working in that department (eg. Accounts, IT, Management, Admin)
 * list location <i>PIN</i>: lists employees living at a given PIN code (eg. 560042, 560086, 560003)
  
  <b>SHOW</b>
  
  * show <i>ID</i> shows employee name with a given ID.
  
  <b>DELETE</b>
  * delete all: deletes all employee data
  * delete <i>ID</i>: deletes data of a particular employee
  
  <b>REMOVE</b>
  * remove <i>NUM</i>: asks you to enter NUM ids to remove (eg. remove 5 means you now have to enter 5 ids to remove)
  * id1 id2 id3 etc. sets There key of the employees with those ids to false. They don't show up in lists then.
  
  <b>EXIT</b>
  * exit exits the program
  
  
