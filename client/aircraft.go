package client

import (
	"encoding/json"
	"errors"
)

type Aircraft struct {
	IataCode string `json:"iata_code"`
	Name     string `json:"name"`
	IcaoCode string `json:"icao_code"`
}

type Meta struct {
	After  string `json:"after"`
	Before string `json:"before"`
	Limit  int    `json:"limit"`
}

type GetAircraftListResponse struct {
	Meta      Meta       `json:"meta"`
	Aircrafts []Aircraft `json:"data"`
}

type GetAircraftResponse struct {
	Aircraft Aircraft `json:"data"`
}

type getAircraftRequest struct{}

func (c *Client) GetAircraftList(f_token string) (resp GetAircraftListResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/aircraft - FToken must be specified")
	}
	req := getAircraftRequest{}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/aircraft", f_token, jsonBody, &resp)
	return resp, err
}

func (c *Client) GetAircraftByIataCode(f_token string, iata_code string) (resp GetAircraftResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/aircraft/iata_code - FToken must be specified")
	}

	if iata_code == "" {
		return c.GetAircraftList(f_token)
	}

	req := getAircraftRequest{}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/aircraft/"+iata_code, f_token, jsonBody, &resp)
	return resp, err
}
