package client

import (
	"encoding/json"
	"errors"
)

type Cabin struct {
	Carrier        string   `json:"carrier"`
	CabinClass     string   `json:"cabin_class"`
	CabinClassNode string   `json:"cain_class_node"`
	FareClassCodes []string `json:"fare_class_codes"`
}

type Meta struct {
	After  string `json:"after"`
	Before string `json:"before"`
	Limit  int    `json:"limit"`
}

type GetCabinListResponse struct {
	Cabins []map[string][]Cabin `json:""`
}

type GetCabinResponseWithParams struct {
	Cabin map[string][]Cabin `json:""`
}

type getCabinRequest struct {
	FToken string `json:"f_token"`
}

func (c *Client) GetCabinList(f_token string) (resp GetCabinListResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/cabin-booking - FToken must be specified")
	}
	req := getCabinRequest{
		FToken: f_token,
	}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/api/cabin-booking", jsonBody, &resp)
	return resp, err
}

func (c *Client) GetCabinWithParams(f_token string, carrier string, cabin_class string) (resp GetCabinResponseWithParams, err error) {
	if f_token == "" {
		return resp, errors.New("/api/cabin-booking - FToken must be specified")
	}

	req := getCabinRequest{
		FToken: f_token,
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/api/cabin-booking/?carrier="+carrier+"&cabin_class"+cabin_class, jsonBody, &resp)
	return resp, err
}
