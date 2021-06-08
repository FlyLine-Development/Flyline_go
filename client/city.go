package client

import (
	"encoding/json"
	"errors"
)

type City struct {
	IataCode string `json:"iata_code"`
	Name     string `json:"name"`
	IcaoCode string `json:"icao_code"`
}

type Meta struct {
	After  string `json:"after"`
	Before string `json:"before"`
	Limit  int    `json:"limit"`
}

type GetCityListResponse struct {
	Meta   Meta   `json:"meta"`
	Cities []City `json:"data"`
}

type GetCityResponse struct {
	City City `json:"data"`
}

type getCityRequest struct {
	FToken string `json:"f_token"`
}

func (c *Client) GetCityList(f_token string) (resp GetCityListResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/cities - FToken must be specified")
	}
	req := getCityRequest{
		FToken: f_token,
	}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/cities", jsonBody, &resp)
	return resp, err
}

func (c *Client) GetCityByIataCode(f_token string, iata_code string) (resp GetCityResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/cities/iata_code - FToken must be specified")
	}

	if iata_code == "" {
		return c.GetCityList(f_token)
	}

	req := getCityRequest{
		FToken: f_token,
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/cities/"+iata_code, jsonBody, &resp)
	return resp, err
}
