package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
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

func getEstimationByCountry(countryCode string, chance float64, averageMeetOtherPerson int64, dayAfterToday int32) (int64, int64, []TotalPerDay) {

	country := getSpecificCountryData(countryCode)

	infectedTotal := country.TotalConfirmed
	var dayData []TotalPerDay

	for i := 0; i <= int(dayAfterToday); i++ {
		chanceMultiplePersonMeet := float64(averageMeetOtherPerson) * chance
		estimationMultiplicationIncrease := math.Pow((1 + chanceMultiplePersonMeet), float64(i))
		estimationOnXDay := estimationMultiplicationIncrease * float64(infectedTotal)

		date := time.Now().AddDate(0, 0, int(i))

		var totalDay TotalPerDay
		totalDay.Date = date
		totalDay.DateStr = strconv.Itoa(date.Day()) + " - " + strconv.Itoa(int(date.Month()))
		totalDay.Total = int64(estimationOnXDay)
		dayData = append(dayData, totalDay)
	}

	// chanceMultiplePersonMeet := float64(averageMeetOtherPerson) * chance
	// estimationMultiplicationIncrease := math.Pow((1 + chanceMultiplePersonMeet), float64(dayAfterToday))
	// estimationOnXDay := estimationMultiplicationIncrease * float64(infectedTotal)

	estimationOnLastDay := dayData[len(dayData)-1].Total
	estimationincreased := estimationOnLastDay - infectedTotal

	// log.Print(chanceMultiplePersonMeet)
	log.Print(estimationincreased)

	return estimationOnLastDay, estimationincreased, dayData
}
