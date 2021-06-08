package client

import (
	"encoding/json"
	"errors"
)

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

type City struct {
	IataCode        string `json:"iata_code"`
	Name            string `json:"name"`
	IataCountryCode string `json:"iata_country_code"`
}

type Airport struct {
	Iatacode        string  `json:"iata_code"`
	Name            string  `json:"name"`
	IataCountryCode string  `json:"iata_country_code"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	IcaoCode        float64 `json:"icao_code"`
	Timezone        string  `json:"time_zone"`
	City            City    `json:"city"`
}

type Aircraft struct {
	IataCode string `json:"iata_code"`
	Name     string `json:"name"`
}

type Carrier struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Offer struct {
	Id            string      `json:"id"`
	BaseAmount    string      `json:"base_amount"`
	BaseCurrency  string      `json:"base_currency"`
	TaxAmount     string      `json:"tax_amount"`
	TaxCurrency   string      `json:"tax_currency"`
	TotalAmount   string      `json:"total_amount"`
	TotalCurrency string      `json:"total_currency"`
	Passengers    []Passenger `json:"passengers"`
	Owner         string      `json:"owner"`
}

type GetAirfareResponse struct {
	LiveMode  bool       `json:"live_mode"`
	Count     int        `json:"count"`
	Airports  []Airport  `json:"airports"`
	Aircrafts []Aircraft `json:"aircraft"`
	Carriers  []Carrier  `json:"carrier"`
	Offers    []Offer    `json:"offers"`
}
type GetAirfareOptions struct {
	CabinClass        string
	Slices            []RequestSlice
	Passengers        []Passenger
	PermittedCarriers []string
	MaxPrice          Price
}

type getAirfareRequestOptions struct {
	CabinClass        string         `json:"cabin_class, omitempty"`
	Slices            []RequestSlice `json:"slices, omitempty"`
	Passengers        []Passenger    `json:"passengers, omitempty"`
	PermittedCarriers []string       `json:"permitted_carriers, omitempty"`
	MaxPrice          Price          `json:"max_price"`
}
type getAirfareRequest struct {
	Options getAirfareRequestOptions `json:"options,omitempty"`
}

func (c *Client) GetAirfareByOneWay(f_token string, Slices []RequestSlice) (resp GetAirfareResponse, err error) {
	if len(Slices) == 1 {
		options := new GetAirfareOptions{
			Slices: Slices
		}
		return c.GetAirfareWithOptions(f_token, options)
	} else {
		return resp, errors.New("/api/flights/shop - Slice must be one specified")
	}
}

func (c *Client) GetAirfareByRoute(f_token string, Slices []RequestSlice) (resp GetAirfareResponse, err error) {
	if len(Slices) == 2 {
		options := new GetAirfareOptions{
			Slices: Slices
		}
		return c.GetAirfareWithOptions(f_token, options)
	} else {
		return resp, errors.New("/api/flights/shop - Slice must be two specified")
	}
}

func (c *Client) GetAirfareByAirlines(f_token string,  permittedCarrier []string) (resp GetAirfareResponse, err error) {
	if len(Slices) > 0 {
		options := new GetAirfareOptions{
			PermittedCarriers: permittedCarrier
		}
		return c.GetAirfareWithOptions(f_token, options)
	} else {
		return resp, errors.New("/api/flights/shop - Slice must be one specified")
	}
}

func (c *Client) GetAirfareByCabin(f_token string, CabinClass string) (resp GetAirfareResponse, err error) {
	if len(Slices) > 0 {
		options := new GetAirfareOptions{
			CabinClass: CabinClass
		}
		return c.GetAirfareWithOptions(f_token, options)
	} else {
		return resp, errors.New("/api/flights/shop - Slice must be one specified")
	}
}

func (c *Client) GetAirfareByDepartureTime(f_token string, PermittedTimes []PermittedTime) (resp GetAirfareResponse, err error) {
	if len(Slices) > 0 {
		options := new GetAirfareOptions{}
		options.Slices[0].Departure.PermittedTimes = PermittedTimes

		return c.GetAirfareWithOptions(f_token, options)
	} else {
		return resp, errors.New("/api/flights/shop - Slice must be one specified")
	}
}

func (c *Client) GetAirfareByArrivalTime(f_token string,  PermittedTimes []PermittedTime) (resp GetAirfareResponse, err error) {
	if len(Slices) > 0 {
		options := new GetAirfareOptions{}
		options.Slices[0].Arrival.PermittedTimes = PermittedTimes
		return c.GetAirfareWithOptions(f_token, options)
	} else {
		return resp, errors.New("/api/flights/shop - Slice must be one specified")
	}
}

func (c *Client) GetAirfareByDepartureDate(f_token string, PermittedDates []PermittedDate) (resp GetAirfareResponse, err error) {
	if len(Slices) > 0 {
		options := new GetAirfareOptions{}
		options.Slices[0].Departure.PermittedDates = PermittedDates
		return c.GetAirfareWithOptions(f_token, options)
	} else {
		return resp, errors.New("/api/flights/shop - Slice must be one specified")
	}
}

func (c *Client) GetAirfareByArrivalDate(f_token string,  PermittedDates []PermittedDate) (resp GetAirfareResponse, err error) {
	if len(Slices) == 1 {
		options := new GetAirfareOptions{}
		options.Slices[0].Arrival.PermittedDates = PermittedDates
		return c.GetAirfareWithOptions(f_token, options)
	} else {
		return resp, errors.New("/api/flights/shop - Slice must be one specified")
	}
}

func (c *Client) GetAirfareByFlightNumber(f_token string,  FlightNumbers []string) (resp GetAirfareResponse, err error) {
	if len(Slices) == 1 {
		options := new GetAirfareOptions{}
		options.Slices[0].FlightNumbers = FlightNumbers
		return c.GetAirfareWithOptions(f_token, options)
	} else {
		return resp, errors.New("/api/flights/shop - Slice must be one specified")
	}
}

func (c *Client) GetAirfareWithMaxPrice(f_token string,  MaxPrice Price) (resp GetAirfareResponse, err error) {
	if len(Slices) == 1 {
		options := new GetAirfareOptions{}
		options.MaxPrice = MaxPrice
		return c.GetAirfareWithOptions(f_token, options)
	} else {
		return resp, errors.New("/api/flights/shop - Slice must be one specified")
	}
}

func (c *Client) GetAirfareWithOptions(f_token string, options GetAirfareOptions) (resp GetAirfareResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/flights/shop - access token must be specified")
	}

	req := getAirfareRequest{}
	if options.CabinClass != "" {
		req.Options.CabinClass = options.CabinClass
	}
	if len(options.Slices) > 0 {
		req.Options.Slices = options.Slices
	}
	if len(options.Passengers) > 0 {
		req.Options.Passengers = options.Passengers
	}
	if len(options.PermittedCarriers) > 0 {
		req.Options.PermittedCarriers = options.PermittedCarriers
	}
	if options.MaxPrice.Price != "" {
		req.Options.MaxPrice = options.MaxPrice
	}

	jsonBody, err := json.Marshal(req)

	if err != nil {
		return resp, err
	}

	err = c.Call("/api/flights/shop", jsonBody, &resp)
	return resp, err
}
