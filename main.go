package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var token string
var base_url string

func init() {
	base_url = "https://api.flyline.io"
}

func SetToken(f_token string) {
	token = f_token
}

func sendRequest(endpoint string, postData string) string {
	url := base_url + endpoint
	method := "POST"
	payload := strings.NewReader(postData)
	client := &http.Client{}
	var req *http.Request
	var err error
	if postData == "" {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, payload)
	}
	if err != nil {
		fmt.Println(err)
		return ("Http Request Failed")
	}

	req.Header.Add("Authorization", "FToken "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ("Request Client Failed")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ("Faild to Get Response")
	}
	return string(body)

}

func GetAirfares(dataJsonStr string) string {
	return sendRequest("/api/flights/shop/", dataJsonStr)
}

func GetAirattributesByFlightNumber(dataJsonStr string) string {
	return sendRequest("/api/search/amenities/", dataJsonStr)
}

func GetAirattributesByRoute(dataJsonStr string) string {
	return sendRequest("/api/amenities/search/route", dataJsonStr)
}

func GetSchedulesByFlightNumber(dataJsonStr string) string {
	return sendRequest("/api/schedule-flight", dataJsonStr)
}

func GetSchedulesByRoute(dataJsonStr string) string {
	return sendRequest("/api/schedule/", dataJsonStr)
}

func GetSeatMaps(dataJsonStr string) string {
	return sendRequest("/api/seat-maps", dataJsonStr)
}

func GetAircrafts() string {
	return sendRequest("/api/aircraft/", "")
}

func GetAircraft(iata_code string) string {
	return sendRequest("/api/aircraft/"+iata_code+"/", "")
}

func GetAirlines() string {
	return sendRequest("/api/airlines/", "")
}

func GetAirline(iata_code string) string {
	return sendRequest("/api/airlines/"+iata_code+"/", "")
}

func GetAirports() string {
	return sendRequest("/api/airports/", "")
}

func GetAirport(iata_code string) string {
	return sendRequest("/api/airports/"+iata_code+"/", "")
}

func GetAirportByCity(iata_code string) string {
	return sendRequest("/api/cities/"+iata_code+"/airports/", "")
}

func GetCities() string {
	return sendRequest("/api/cities", "")
}

func GetCity(iata_code string) string {
	return sendRequest("/api/cities/"+iata_code+"/", "")
}

func GetCabinClassMapping(carrier string, cabin_class string) string {
	if carrier != "" && cabin_class != "" {
		return sendRequest("/api/cabin-booking?carrier="+carrier+"&cabin_class"+cabin_class, "")
	} else {
		return sendRequest("/api/cabin-booking/", "")
	}
}

func GetSeatTypes() string {
	return sendRequest("/api/seats/", "")
}

func GetSeatLayouts() string {
	return sendRequest("/api/layouts/", "")
}

func GetFoods() string {
	return sendRequest("/api/foods/", "")
}

func GetBeverages() string {
	return sendRequest("/api/beverages/", "")
}

func GetEntertainments() string {
	return sendRequest("/api/entertainments/", "")
}

func GetWifis() string {
	return sendRequest("/api/wifis/", "")
}

func GetPowers() string {
	return sendRequest("/api/powers", "")
}
