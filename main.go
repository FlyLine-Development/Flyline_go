package main

import (
	"fmt"
	"net/http"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	clientOptions := flyline.ClientOptions{
		&http.Client{},
	}

	client, err := flyline.NewClient(clientOptions)
	handleError(err)
	fmt.Println(instsResp.Institutions[0].Name, "has products:", instsResp.Institutions[0].Products)

}
