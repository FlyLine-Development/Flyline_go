package client

import (
	"encoding/json"
	"errors"
)

type Param struct {
	CabinClass    string `json:"cabin_class"`
	Departure     string `json:"departure"`
	Arrival       string `json:"arrival"`
	DepartureDate string `json:"departure_date"`
	FlightNo      string `json:"flight_no"`
	Carrier       string `json:"carrier"`
}
type PermittedTime struct {
	Kind      string `json:"kind"`
	Startdate string `json:"start_date"`
	Enddate   string `json:"end_date"`
}

type PermittedDate struct {
	Kind      string `json:"kind"`
	Starttime string `json:"start_time"`
	Endtime   string `json:"end_time"`
}

type Place struct {
	Code           string          `json:"code"`
	Date           string          `json:"date"`
	PermittedTimes []PermittedTime `json:"permitted_times"`
	PermittedDates []PermittedDate `json:"permitted_dates"`
}

type RequestSlice struct {
	Departure     Place    `json:"departure"`
	Arrival       Place    `json:"arrival"`
	FlightNumbers []string `json:"flight_numbers"`
}
type Passenger struct {
	Id   string `json:"id"`
	Age  string `json:"age"`
	Type string `json:"type"`
}

type Price struct {
	Currencty string `json:"currency"`
	Price     string `json:"price"`
}

type Route struct {
	CabinClass        string         `json:"cabin_class, omitempty"`
	Slices            []RequestSlice `json:"slices, omitempty"`
	Passengers        []Passenger    `json:"passengers, omitempty"`
	PermittedCarriers []string       `json:"permitted_carriers, omitempty"`
	MaxPrice          Price          `json:"max_price"`
}

type ResponseCarrier struct {
	IataCode string `json:"iata_code"`
	Name     string `json:"name"`
}

type ResponseAircraft struct {
	IataCode string `json:"iata_code"`
	Name     string `json:"name"`
}

type GetAirAttributeResponse struct {
	CabinClass string             `json:"cabin_class"`
	Carriers   []ResponseCarrier  `json:"carriers"`
	Aircrafts  []ResponseAircraft `json:"aircraft"`
}

type GetAirAttributeOptions struct {
	Params Param
	Route  Route
}

type getAirAttributeRequestOptions struct {
	Params Param `json:"param, omitempty"`
	Route  Route `json:"route, omitempty"`
}

type getAirAttributeRequest struct {
	Options getAirAttributeRequestOptions `json:"options, omitempty"`
}

func (c *Client) GetAirAttributeWithParams(f_token string, Params Param) (resp GetAirAttributeResponse, err error) {
	options := GetAirAttributeOptions{}
	options.Params = Params

	return c.GetAirAttributeWithOptions(f_token, options)

}

func (c *Client) GetAirAttributeWithOptions(f_token string, options GetAirAttributeOptions) (resp GetAirAttributeResponse, err error) {
	if f_token == "" {
		return resp, errors.New("AirAttribute Function - access token must be specified")
	}

	req := getAirAttributeRequest{}
	if options.Params.CabinClass != "" {
		req.Options.Params = options.Params

		jsonBody, err := json.Marshal(req)

		if err != nil {
			return resp, err
		}

		err = c.Call("/api/search/amenities", jsonBody, &resp)
		return resp, err
	} else if options.Route.CabinClass != "" {
		req.Options.Route = options.Route

		jsonBody, err := json.Marshal(req)

		if err != nil {
			return resp, err
		}

		err = c.Call("/api/amenities/search/route", jsonBody, &resp)
		return resp, err
	}

	return resp, errors.New("Please check params of request")

}
