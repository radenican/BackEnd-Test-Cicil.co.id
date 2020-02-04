Create Table <br/>
**cicil_phonebook** 


``CREATE TABLE `phonebook` (
  `ID` int(11) NOT NULL,
  `fullname` varchar(100) NOT NULL,
  `mobilenumber` int(11) NOT NULL,
  `homenumber` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
``

edit config db in main.go

open terminal use
<br/>dep init

and build go with command
<br/> go build

//create / Post
http://localhost:8089/phonebook/create
**body**
`{
"fullname" : "Ikhsan" ,
"mobilenumber" : "082297383474" , 
"homenumber" : "022942131231"
}`

//getdata /get
**response**
http://localhost:8089/phonebook
`[
    {
        "id": 1,
        "fullname": "Ikhsan",
        "mobilenumber": 0,
        "homenumber": 0
    },
    {
        "id": 2,
        "fullname": "Ikhsan2",
        "mobilenumber": 0,
        "homenumber": 0
    }
]`

//get data with selected id /get http://localhost:8089/phonebook/1
response
`{
    "id": 1,
    "fullname": "Ikhsan",
    "mobilenumber": 0,
    "homenumber": 0
}`

//update data /put http://localhost:8089/phonebook/update/1

body 
`{
"fullname" : "raden m Ikhsan " 
}`
response
`{
    "message": "update successfully"
}`

//delete data  /delete http://localhost:8089/phonebook/delete/1
response
`{
    "message": "delete successfully"
}`
