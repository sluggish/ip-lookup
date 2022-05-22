package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	figure "github.com/common-nighthawk/go-figure"
)

type results struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	Region      string  `json:"region"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func main() {
	var query string
	fmt.Print("Enter IP Address: ")
	fmt.Scanln(&query)
	res, err := http.Get("http://ip-api.com/json/"+query)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var result results
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println(err)
    }
	fmt.Print("\033[H\033[2J")
	
	banner := figure.NewFigure("FETCHED INFO", "alphabet", true)
	banner.Print()

	fmt.Printf("IP Address: %s\nCountry: %s\nRegion: %s\nCity: %s\nZip Code: %s\nISP: %s\nOrg: %s\nAS: %s", result.Query, result.Country, result.Region, result.City, result.Zip, result.Isp, result.Org, result.As)
	os.Exit(0)
}