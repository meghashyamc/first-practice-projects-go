<b>Practice project that helps store data in and retrieve data from IPFS.</b>

NOTE: As I look back at this (an year later), I see that the naming conventions I've used are definitely not great. 'lepyaPract' itself doesn't make much sense. 

Supports these commands:

<b>/add</b> (Enter arg as key and upload file) 

<b>/cat</b> (Enter arg as key and hash of the file as value)

<b>/list</b> (Argument optional. If argument is there, key should be arg and value should be /folder (UNIX like path already    created)

<b>/mkdir</b> (Enter arg as key and enter name of directory as value)

<b>/status</b> (Enter arg as key and enter directory or file path eg. /jlt/cat.jpeg as value)

<b>/remove</b>  (Enter arg as key and enter directory or file path eg. /jlt/cat.jpeg as value)

<b>/copy</b>  (Enter arg1 and arg2 as keys and enter source path and destination path as values. Source path can be /ipfs/hash.. or a Unix like path like /jlt/cat.jpeg)

<b>/move</b>  (Enter arg1 and arg2 as keys and enter source path and destination path as values. Source path can be /ipfs/hash.. or a Unix like path like /jlt/cat.jpeg)