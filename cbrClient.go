package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AllData struct {
	Valute struct{
		USD Values `json:"USD"`
		EUR Values `json:"EUR"`
		CNY Values `json:"CNY"`
	}
}

type Values struct {
	Name string `json:"Name"`
	Value float64 `json:"Value"`
}

var answer AllData

const cbrURL = "https://www.cbr-xml-daily.ru/daily_json.js"

func courseRequest() string  {
	client := http.Client{}

	resp, err := client.Get(cbrURL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(buf, &answer); err != nil {
		log.Fatal(err)
	}

	formattedResponse := fmt.Sprintf("%s: %.2f\n%s: %.2f\n%s: %.2f",
		answer.Valute.EUR.Name, answer.Valute.EUR.Value,
		answer.Valute.USD.Name, answer.Valute.USD.Value,
		answer.Valute.CNY.Name, answer.Valute.CNY.Value)

	return formattedResponse
}
