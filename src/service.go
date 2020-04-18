package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arr_user []Users
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select id, first_name, last_name from person")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_user = append(arr_user, users)
		}
	}

	response.Status = 500
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func insertUsersMultipart(w http.ResponseWriter, r *http.Request) {
	// var users Users
	// var arr_user []Users
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")

	_, err = db.Exec("INSERT INTO person (first_name, last_name) values (?,?)",
		first_name,
		last_name,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 500
	response.Message = "Added Successfully"
	log.Print("Insert data to database")

	w.Header().Set("Contetnt-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updateUserMultipart(w http.ResponseWriter, r *http.Request) {
	var response ReponseSingle
	var user Users

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	user.Id = r.FormValue("user_id")
	user.FirstName = r.FormValue("first_name")
	user.LastName = r.FormValue("last_name")

	_, err = db.Exec("UPDATE person set first_name = ?, last_name = ? where id = ?",
		user.FirstName,
		user.LastName,
		user.Id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 500
	response.Message = "Success Update Data"
	response.Data = user
	log.Print("Update data to Database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func deleteUserData(w http.ResponseWriter, r *http.Request) {
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("user_id")

	_, err = db.Exec("DELETE from person where id = ?",
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 500
	response.Message = "Success Delete Data"
	log.Print("Delete data from database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
