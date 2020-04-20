package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func getGlobalData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	global := getAllData().Global
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(global)
}

func getAllCountryData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	countries := getCountryData()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

func getByCountry(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	country := getSpecificCountryData(vars["countryCode"])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(country)
}

func getByHighest(w http.ResponseWriter, r *http.Request) {

	country := getHighestCountry()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(country)
}

func getAverageOnXDay(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var response DayEstimation

	// err := r.ParseMultipartForm(4096)
	// if err != nil {
	// 	panic(err)
	// }

	vars := mux.Vars(r)

	// countryCode := r.FormValue("country_code")
	// chanche, _ := strconv.ParseFloat(r.FormValue("chance"), 64)
	// xDay, _ := strconv.ParseFloat(r.FormValue("day_after_today"), 64)
	// averageMeetOtherPerson, _ := strconv.ParseInt(r.FormValue("average_meet_number"), 10, 64)

	countryCode := vars["countryCode"]
	chanche, _ := strconv.ParseFloat(vars["chance"], 64)
	chanche = chanche / 100
	xDay, _ := strconv.ParseInt(vars["day"], 10, 64)
	xDay32 := int32(xDay)
	averageMeetOtherPerson, _ := strconv.ParseInt(vars["person"], 10, 64)

	estimationOnXDay, increase, dayData := getEstimationByCountry(countryCode, chanche, averageMeetOtherPerson, xDay32)

	country := getSpecificCountryData(countryCode)

	response.Country = country
	response.EstimationOnXDay = estimationOnXDay
	response.EstimationIncreasedInfectedPerson = increase
	response.TodayDate = time.Now()
	response.XDay = xDay32
	response.DayData = dayData
	response.XDayAfterTodayDate = time.Now().AddDate(0, 0, int(xDay))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
