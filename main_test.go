package main

import (
	"fmt"
	"testing"
)

func TestGetSeatTypes(t *testing.T) {
	SetToken("test_904c384cb63468da5a6a4b20b4b353e45b142098")
	response := GetSeatTypes()
	fmt.Println(response)
	if response == "" {
		t.Errorf("aaa")
	}
}
