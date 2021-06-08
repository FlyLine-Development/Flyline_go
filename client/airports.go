package client

import (
	"encoding/json"
	"errors"
)

type Airport struct {
	IataCode   string `json:"iata_code"`
	Name       string `json:"name"`
	Integrated bool   `json:"integrated"`
}

type Meta struct {
	After  string `json:"after"`
	Before string `json:"before"`
	Limit  int    `json:"limit"`
}

type GetAirportListResponse struct {
	Meta     Meta      `json:"meta"`
	Airports []Airport `json:"data"`
}

type GetAirportResponse struct {
	Airport Airport `json:"data"`
}

type getAirportRequest struct {
}

func (c *Client) GetAirportList(f_token string) (resp GetAirportListResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/airports - FToken must be specified")
	}
	req := getAirportRequest{}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/airports", f_token, jsonBody, &resp)
	return resp, err
}

func (c *Client) GetAirportByAirportIataCode(f_token string, iata_code string) (resp GetAirportResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/airports/iata_code - FToken must be specified")
	}

	if iata_code == "" {
		return c.GetAirportList(f_token)
	}

	req := getAirportRequest{}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/airports/"+iata_code, f_token, jsonBody, &resp)
	return resp, err
}

func (c *Client) GetAirportByCityIataCode(f_token string, iata_code string) (resp GetAirportResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/cities/iata_code/airports - FToken must be specified")
	}

	if iata_code == "" {
		return c.GetAirportList(f_token)
	}

	req := getAirportRequest{}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/api/cities/"+iata_code+"/airports", f_token, jsonBody, &resp)
	return resp, err
}
