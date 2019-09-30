
<b>This is an updated version of the rest API to upload/view/delete employee details.</b>

Here are the supported operations:

<b>/add </b>: adds employee list in JSON array format // NEW: when another list is added, it gets appended to the previously added list

<b>/list </b>: Lists all employees    // NEW: this replaces list/all, also the list is now in a more user friendly format.

<b>/list/{term}</b>: Lists all employees whose fields include the given search term (case insensitive) // NEW

<b>/list/dept/{name}</b>: Lists all employees with given department name

<b>/list/loc/{pin}</b>: Lists employees living at that pin code

<b>/list/loc/{pin}/dn/{doorno}</b>: Lists employees living at that door no.

<b>/list/loc/{pin}/street/{streetname}</b>: Lists employees living at that street

<b>/list/loc/{pin}/locality/{locality}</b>: Lists employees living at that locality

<b>/show/{id}</b>: shows employee with that id

<b>/delete/all</b>: deletes all employees

<b>/delete/{id}</b>: deletes employee with given id

<b>/remove/{id}</b>: removes employee with given id

// NEW: Only employee details with name, department and address are accepted.

// NEW: If the ID is not given, it is generated automatically
