
This is a rest API to upload/view/delete employee details.

/add : add employee list in JSON format
/list/all : Lists all employees
/list/dept/{name}: Lists all employees with given department name
/list/loc/{pin}: Lists employees living at that pin code
/list/loc/{pin}/dn/{doorno}: Lists employees living at that door no.
/list/loc/{pin}/street/{streetname}: Lists employees living at that street
/list/loc/{pin}/locality/{locality}: Lists employees living at that locality

/show/{id}: shows employee with that id

/delete/all: deletes all employees
/delete/{id}: deletes employee with given id

/remove/{id}: removes employee with given id
