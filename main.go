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

func sendPostRequest(endpoint string, postData string) string {
	url := base_url + endpoint
	method := "POST"
	payload := strings.NewReader(postData)
	client := &http.Client{}
	var req *http.Request
	var err error
	req, err = http.NewRequest(method, url, payload)
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
func sendGetRequest(endpoint string) string {
	url := base_url + endpoint
	method := "GET"

	client := &http.Client{}
	var req *http.Request
	var err error
	req, err = http.NewRequest(method, url, nil)
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
	return sendPostRequest("/api/flights/shop/", dataJsonStr)
}
// Update the code of api.
func GetAirattributesByFlightNumber(dataJsonStr string) string {
	return sendPostRequest("/search/attributes/flight/", dataJsonStr)
}

// Update the code of api.
func GetAirattributesByRoute(dataJsonStr string) string {
	return sendPostRequest("/search/attributes/route/", dataJsonStr)
}

func GetSchedulesByFlightNumber(dataJsonStr string) string {
	return sendPostRequest("/api/schedule-flight", dataJsonStr)
}

func GetSchedulesByRoute(dataJsonStr string) string {
	return sendPostRequest("/api/schedule/", dataJsonStr)
}

func GetSeatMaps(dataJsonStr string) string {
	return sendPostRequest("/api/seat-maps", dataJsonStr)
}

func GetAircrafts() string {
	return sendGetRequest("/api/aircraft/")
}

func GetAircraft(iata_code string) string {
	return sendGetRequest("/api/aircraft/" + iata_code + "/")
}

func GetAirlines() string {
	return sendGetRequest("/api/airlines/")
}

func GetAirline(iata_code string) string {
	return sendGetRequest("/api/airlines/" + iata_code + "/")
}

func GetAirports() string {
	return sendGetRequest("/api/airports/")
}

func GetAirport(iata_code string) string {
	return sendGetRequest("/api/airports/" + iata_code + "/")
}

func GetAirportByCity(iata_code string) string {
	return sendGetRequest("/api/cities/" + iata_code + "/airports/")
}

func GetCities() string {
	return sendGetRequest("/api/cities")
}

func GetCity(iata_code string) string {
	return sendGetRequest("/api/cities/" + iata_code + "/")
}

func GetCabinClassMapping(carrier string, cabin_class string) string {
	if carrier != "" && cabin_class != "" {
		return sendGetRequest("/api/cabin-booking?carrier=" + carrier + "&cabin_class" + cabin_class)
	} else {
		return sendGetRequest("/api/cabin-booking/")
	}
}

func GetSeatTypes() string {
	return sendGetRequest("/api/seats/")
}

func GetSeatLayouts() string {
	return sendGetRequest("/api/layouts/")
}

func GetFoods() string {
	return sendGetRequest("/api/foods/")
}

func GetBeverages() string {
	return sendGetRequest("/api/beverages/")
}

func GetEntertainments() string {
	return sendGetRequest("/api/entertainments/")
}

func GetWifis() string {
	return sendGetRequest("/api/wifis/")
}

func GetPowers() string {
	return sendGetRequest("/api/powers")
}

func main() {
	fmt.Println("!oG ,olleH")
	fmt.Println("Hello World", "Hello Go")
}
