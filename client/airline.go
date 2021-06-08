package client

import (
	"encoding/json"
	"errors"
)

type Airline struct {
	IataCode   string `json:"iata_code"`
	Name       string `json:"name"`
	Integrated bool   `json:"integrated"`
}

type Meta struct {
	After  string `json:"after"`
	Before string `json:"before"`
	Limit  int    `json:"limit"`
}

type GetAirlineListResponse struct {
	Meta     Meta      `json:"meta"`
	Airlines []Airline `json:"data"`
}

type GetAirlineResponse struct {
	Airline Airline `json:"data"`
}

type getAirlineRequest struct {
	FToken string `json:"f_token"`
}

func (c *Client) GetAirlineList(f_token string) (resp GetAirlineListResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/airlines - FToken must be specified")
	}
	req := getAirlineRequest{
		FToken: f_token,
	}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/airlines", jsonBody, &resp)
	return resp, err
}

func (c *Client) GetAirlineByIataCode(f_token string, iata_code string) (resp GetAirlineResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/airlines/iata_code - FToken must be specified")
	}

	if iata_code == "" {
		return c.GetAirlineList(f_token)
	}

	req := getAirlineRequest{
		FToken: f_token,
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/airlines/"+iata_code, jsonBody, &resp)
	return resp, err
}
