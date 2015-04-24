package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	GoogleEndpoint = "http://maps.googleapis.com/maps/api/geocode/json?sensor=false&language=zh_CN&address="
)

type LatLng struct {
	Lat float64
	Lng float64
}

type Response struct {
	Address string
}

func Address(address string) {
	fmt.Println(address)

	response, error := http.Get(GoogleEndpoint+address)

	if error != nil {
		log.Fatal(error)
	}

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("%s", body)

	n := json.NewDecoder(response.Body)

	fmt.Printf("%s", n)
}

func main(){
	Address("广东省丽江花园");
}