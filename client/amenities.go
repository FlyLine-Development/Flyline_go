package client

import (
	"encoding/json"
	"errors"
)

type Meta struct {
	After  string `json:"after"`
	Before string `json:"before"`
	Limit  int    `json:"limit"`
}

type Seat struct {
	DisplayText string `json:"display_text"`
	Pitch       string `json:"pitch"`
	Width       string `json:"width"`
}

type GetSeatListResponse struct {
	Meta  Meta   `json:"meta"`
	Seats []Seat `json:"data"`
}

type getSeatListRequest struct{}

type Layout struct {
	DisplayText string `json:"display_text"`
}

type GetLayoutListResponse struct {
	Meta    Meta     `json:"meta"`
	Layouts []Layout `json:"data"`
}

type getLayoutListRequest struct{}

type Food struct {
	DisplayText string `json:"display_text"`
	Cost        string `json:"cost"`
}

type GetFoodListResponse struct {
	Meta  Meta   `json:"meta"`
	Foods []Food `json:"data"`
}

type getFoodListRequest struct{}

type Beverage struct {
	DisplayText      string `json:"display_text"`
	Type             string `json:"type"`
	AlcoholicCost    string `json:"alcoholic_cost"`
	NonalcoholicCost string `json:"nonalcoholic_cost"`
}

type GetBeverageListResponse struct {
	Meta      Meta       `json:"meta"`
	Beverages []Beverage `json:"data"`
}

type getBeverageListRequest struct{}

type Entertainment struct {
	DisplayText string `json:"display_text"`
}

type GetEntertainmentListResponse struct {
	Meta           Meta            `json:"meta"`
	Entertainments []Entertainment `json:"data"`
}

type getEntertainmentListRequest struct{}

type Wifi struct {
	DisplayText string `json:"display_text"`
	Quality     string `json:"quality"`
	Cost        string `json:"cost"`
}

type GetWifiListResponse struct {
	Meta  Meta   `json:"meta"`
	Wifis []Wifi `json:"data"`
}

type getWifiListRequest struct{}

type Power struct {
	DisplayText    string `json:"display_text"`
	MultipleatSeat string `jsong:"multiple_at_seat"`
	UsbPort        string `json:"usb_port"`
	PowerOutlet    string `json:"power_outlet"`
}

type GetPoweriListResponse struct {
	Meta   Meta    `json:"meta"`
	Powers []Power `json:"data"`
}

type getPowerListRequest struct{}

func (c *Client) GetSeatList(f_token string) (resp GetSeatListResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/seats - FToken must be specified")
	}

	req := getSeatListRequest{}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/api/seats", f_token, jsonBody, &resp)
	return resp, err
}

func (c *Client) GetLayoutList(f_token string) (resp GetLayoutListResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/layouts - FToken must be specified")
	}

	req := getLayoutListRequest{}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/api/layouts", f_token, jsonBody, &resp)
	return resp, err
}

func (c *Client) GetFoodList(f_token string) (resp GetFoodListResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/foods - FToken must be specified")
	}

	req := getFoodListRequest{}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/api/foods", f_token, jsonBody, &resp)
	return resp, err
}

func (c *Client) GetBeverageList(f_token string) (resp GetBeverageListResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/beverages - FToken must be specified")
	}

	req := getBeverageListRequest{}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/api/beverages", f_token, jsonBody, &resp)
	return resp, err
}

func (c *Client) GetEntertainmentList(f_token string) (resp GetEntertainmentListResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/entertainments - Fotken must be specified")
	}

	req := getEntertainmentListRequest{}
	jsonBody, err := json.Marshal(req)

	if err != nil {
		return resp, err
	}

	err = c.Call("/api/entertainments", f_token, jsonBody, &resp)
	return resp, err
}

func (c *Client) GetWifiList(f_token string) (resp GetWifiListResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/wifis - Fotken must be specified")
	}

	req := getWifiListRequest{}
	jsonBody, err := json.Marshal(req)

	if err != nil {
		return resp, err
	}

	err = c.Call("/api/wifis", f_token, jsonBody, &resp)
	return resp, err
}

func (c *Client) GetPowerList(f_token string) (resp GetPoweriListResponse, err error) {
	if f_token == "" {
		return resp, errors.New("/api/powers - Fotken must be specified")
	}

	req := getPowerListRequest{}
	jsonBody, err := json.Marshal(req)

	if err != nil {
		return resp, err
	}

	err = c.Call("/api/powers", f_token, jsonBody, &resp)
	return resp, err
}
