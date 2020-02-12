# phone book cicil test backend 

# External package
go-chi <br/>
go-sql-driver

**run external package on default src folder ex:**  ``c:/go/src/`` **  <br/>
$ go get -u github.com/go-chi/chi <br/>
$ go get -u github.com/go-sql-driver/mysql



# Create DB
**cicil_phonebook** 

``CREATE TABLE `phonebook` (
  `ID` int(11) NOT NULL,
  `fullname` varchar(100) NOT NULL,
  `mobilenumber` varchar(20) NOT NULL,
  `homenumber` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
``


# edit config db
go to main.go line 18 <br/>
``const (
	dbName = "cicil_phonebook"
	dbPass = "mysql"
	dbHost = "localhost"
	dbPort = "3306"
)``

# Identity your dependency
Open Terminal go to your Workspace Location ex: ``c:/go/src/cicil`` <br/>
**dep init** <br/>
Enter


# Running
Open Terminal go to your Workspace Location ex: ``c:/go/src/cicil`` <br/>
**go build** <br/>
Enter <br/>
Wait until finish <br/>
Open ``cicil.exe``

# Testing Open Rest application
ex: postman / insomnia

# Create
/post http://localhost:8089/phonebook/create
<br/>
**body** 
`{
"fullname" : "Ikhsan3" ,
"mobilenumber" : "+6282297383474" , 
"homenumber" : "+6222942131231"
}`

# GET ALL DATA
/get http://localhost:8089/phonebook

<br/>
**response**
`[
     {
        "id": 14,
        "fullname": "Ikhsan3",
        "mobilenumber": "082297383474",
        "homenumber": "022942131231"
    },
    {
        "id": 15,
        "fullname": "Ikhsan3",
        "mobilenumber": "+6282297383474",
        "homenumber": "+6222942131231"
    }
]`

# GET SELECTED DATA
/get http://localhost:8089/phonebook/1
<br/>
**response**
`{
    "id": 15,
    "fullname": "Ikhsan3",
    "mobilenumber": "+6282297383474",
    "homenumber": "+6222942131231"
}`


# UPDATE
/put http://localhost:8089/phonebook/update/15
 <br/> **body** <br/>
`{
    "fullname": "Ikhsan",
    "mobilenumber": "+6282297383474",
    "homenumber": "+219123109"
}`
 <br/> **response** <br/>

`{
    "message": "update successfully"
}`

# DELETE
 /delete http://localhost:8089/phonebook/delete/1
 <br/>**response** <br/>

`{
    "message": "delete successfully"
}`
