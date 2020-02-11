# phone book cicil test backend 

# External package
go-chi <br/>
go-sql-driver

**run** 
$ go get -u github.com/go-chi/chi <br/>
$ go get -u github.com/go-sql-driver/mysql



# Create DB
**cicil_phonebook** 

``CREATE TABLE `phonebook` (
  `ID` int(11) NOT NULL,
  `fullname` varchar(100) NOT NULL,
  `mobilenumber` int(11) NOT NULL,
  `homenumber` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
``


# edit config db
go to main.go line 18
const (
	dbName = "cicil_phonebook"
	dbPass = ""
	dbHost = "localhost"
	dbPort = "3306"
)

# Identity your dependency
Open Terminal go to your Workspace Location ex: c:/go/src/cicil
**dep init**
Enter


# Running
Open Terminal go to your Workspace Location ex: c:/go/src/cicil
**go build**
Enter
Wait until finish 
Open cicil.exe

# Testing Open Rest application
ex: postman / insomnia

# Create
/post http://localhost:8089/phonebook/create
**body**
`{
"fullname" : "Ikhsan" ,
"mobilenumber" : "082297383474" , 
"homenumber" : "022942131231"
}`

# GET ALL DATA
**response**
/get http://localhost:8089/phonebook
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

# GET SELECTED DATA
/get http://localhost:8089/phonebook/1
**response**
`{
    "id": 1,
    "fullname": "Ikhsan",
    "mobilenumber": 0,
    "homenumber": 0
}`


# UPDATE
/put http://localhost:8089/phonebook/update/1
**body**
`{
    "fullname": "Ikhsan",
    "mobilenumber": "222",
    "homenumber": "111"
}`
**response**

`{
    "message": "update successfully"
}`

# DELETE
 /delete http://localhost:8089/phonebook/delete/1
**response**

`{
    "message": "delete successfully"
}`
