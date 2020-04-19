package main

import "time"

type AllDataCorona struct {
	Global    Global
	Countries []Country
}

type Global struct {
	NewConfirmed   int64
	TotalConfirmed int64
	NewDeaths      int64
	TotalDeaths    int64
	NewRecovered   int64
	TotalRecovered int64
}

type Country struct {
	Country        string
	CountryCode    string
	Slug           string
	NewConfirmed   int64
	TotalConfirmed int64
	NewDeaths      int64
	TotalDeaths    int64
	NewRecovered   int64
	TotalRecovered int64
	Date           time.Time
}

type DayEstimation struct {
	Country                           Country
	EstimationOnXDay                  int64
	EstimationIncreasedInfectedPerson int64
	XDay                              float64
	TodayDate                         time.Time
	XDayAfterTodayDate                time.Time
}
