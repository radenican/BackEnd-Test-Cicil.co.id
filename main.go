package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux
var db *sql.DB

const (
	dbName = "cicil_phonebook"
	dbPass = "mysql"
	dbHost = "localhost"
	dbPort = "3306"
)

// Phonebook type details
type Phonebook struct {
	ID           int    `json:"id,omitempty"`
	FullName     string `json:"fullname,omitempty"`
	MobileNumber string `json:"mobilenumber,omitempty"`
	HomeNumber   string `json:"homenumber,omitempty"`
}

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	var err error
	//sql connection
	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName)
	db, err = sql.Open("mysql", dbSource)
	catch(err)
}

//routing
func routers() *chi.Mux {
	router.Get("/", ping)
	router.Get("/phonebook", AllData)
	router.Get("/phonebook/{id}", SelectedData)
	router.Post("/phonebook/create", Create)
	router.Put("/phonebook/update/{id}", Update)
	router.Delete("/phonebook/delete/{id}", Delete)

	return router
}

// server starting point
func ping(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, map[string]string{"message": "Pong"})
}

//-------------- API ENDPOINT ------------------//
// Alldata
func AllData(w http.ResponseWriter, r *http.Request) {
	errors := []error{}
	payload := []Phonebook{}

	rows, err := db.Query("Select id, fullname , mobilenumber , homenumber From phonebook")
	catch(err)

	defer rows.Close()

	for rows.Next() {
		data := Phonebook{}

		er := rows.Scan(&data.ID, &data.FullName, &data.MobileNumber, &data.HomeNumber)

		if er != nil {
			errors = append(errors, er)
		}
		payload = append(payload, data)
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Create
func Create(w http.ResponseWriter, r *http.Request) {
	var phonebook Phonebook
	json.NewDecoder(r.Body).Decode(&phonebook)
	if _, err := strconv.Atoi(string(phonebook.HomeNumber)); err != nil {
		respondwithJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Home Number is not valid"})
		return
	} //check home number is not Numbers return not valid

	if _, err := strconv.Atoi(string(phonebook.MobileNumber)); err != nil {
		respondwithJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Mobile Number is not valid"})
		return
	} //check mobile number is not Numbers return not valid

	query, err := db.Prepare("Insert phonebook SET fullname=? , mobilenumber=? , homenumber=? ")
	catch(err)

	_, er := query.Exec(phonebook.FullName, phonebook.MobileNumber, phonebook.HomeNumber)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

// Selected
func SelectedData(w http.ResponseWriter, r *http.Request) {
	selected := Phonebook{}
	id := chi.URLParam(r, "id") //get id
	row := db.QueryRow("Select id, fullname, mobilenumber , homenumber  From phonebook where id=?", id)
	err := row.Scan(
		&selected.ID,
		&selected.FullName,
		&selected.MobileNumber,
		&selected.HomeNumber,
	)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "no data with id =  "+id)
		return
	}

	respondwithJSON(w, http.StatusOK, selected)
}

// Update
func Update(w http.ResponseWriter, r *http.Request) {
	var updated Phonebook
	id := chi.URLParam(r, "id") //get id
	json.NewDecoder(r.Body).Decode(&updated)

	if _, err := strconv.Atoi(string(updated.HomeNumber)); err != nil {
		respondwithJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Home Number is not valid"})
		return
	} //check home number is not Numbers return not valid
	if _, err := strconv.Atoi(string(updated.MobileNumber)); err != nil {
		respondwithJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Mobile Number is not valid"})
		return
	} //check mobile number is not Numbers return not valid

	query, err := db.Prepare("Update phonebook set fullname=?, mobilenumber=? , homenumber=? where id=?")
	catch(err)
	_, er := query.Exec(updated.FullName, updated.MobileNumber, updated.HomeNumber, id)
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})

}

// Delete
func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") //getid
	query, err := db.Prepare("delete from phonebook where id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "delete successfully"})
}

func main() {
	routers() //get routing
	http.ListenAndServe(":8089", Logger()) //http run in port 8089
}
