package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getGlobalData(w http.ResponseWriter, r *http.Request) {
	global := getAllData().Global
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(global)
}

func getAllCountryData(w http.ResponseWriter, r *http.Request) {
	countries := getCountryData()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

func getByCountry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	country := getSpecificCountryData(vars["countryCode"])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(country)
}

func getAllData() AllDataCorona {
	var allData AllDataCorona

	resp, err := http.Get("https://api.covid19api.com/summary")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &allData)
	return allData
}

func getCountryData() []Country {
	countries := getAllData().Countries
	return countries
}

func getSpecificCountryData(countryCode string) Country {
	var country Country

	countries := getCountryData()
	for i := range countries {
		if countries[i].CountryCode == countryCode {
			country = countries[i]
		}
	}

	return country
}
