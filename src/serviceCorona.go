package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

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

func getHighestCountry() Country {
	countries := getCountryData()
	var temp Country
	temp.TotalConfirmed = 0
	for i := range countries {
		if temp.TotalConfirmed < countries[i].TotalConfirmed {
			temp = countries[i]
		}
	}

	return temp
}

func getEstimationByCountry(countryCode string, chance float64, averageMeetOtherPerson int64, dayAfterToday float64) (int64, int64) {

	country := getSpecificCountryData(countryCode)

	infectedTotal := country.TotalConfirmed

	chanceMultiplePersonMeet := float64(averageMeetOtherPerson) * chance
	estimationMultiplicationIncrease := math.Pow((1 + chanceMultiplePersonMeet), dayAfterToday)
	estimationOnXDay := estimationMultiplicationIncrease * float64(infectedTotal)

	estimationincreased := int64(estimationOnXDay) - infectedTotal

	log.Print(chanceMultiplePersonMeet)
	log.Print(estimationincreased)

	return int64(estimationOnXDay), estimationincreased
}
