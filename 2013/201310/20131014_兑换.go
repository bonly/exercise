// http://play.golang.org/p/ebVIqodGPd

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
)

const (
	API = "http://api.fixer.io/"
)

var (
	val   	string
	base  	string
	convert float64
	date    string
)

func help() {
	fmt.Println("GoCurrency is used to quickly check the conversion rates")
	flag.PrintDefaults()
}

func exitHelp(code int) {
	help()
	os.Exit(code)
}

func init() {
	flag.Usage = help

	flag.StringVar(&base, "base", "USD", "The base currency to quote a currency against (default USD)")
	flag.Float64Var(&convert, "convert", 1, "Amount of currency to convert (default 1)")
	flag.StringVar(&date, "date", "", "Historical date to show (default today)")
	flag.Parse()

	if len(base) == 0 {
		fmt.Println("Base must be set")
		exitHelp(3)
	}

	if convert < 1 {
		fmt.Println("Convert must be 1 or greater")
		exitHelp(3)
	}

	val = flag.Arg(0)
}

func sendRequest() {
	v := url.Values{}
	if len(base) > 0 {
		v.Add("base", base)
	}
	if len(val) > 0 {
		v.Add("symbols", val)
	}
	params := v.Encode()

	if len(date) == 0 {
		date = "latest"
	}

	resp, err := http.Get(API + date + "?" + params)
	if err != nil {
		fmt.Println("Failed to get data", err)
		os.Exit(3)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Not 200 found")
		os.Exit(3)
	}

	handleResponse(resp.Body)
}

func sortDataKeys(m map[string]float64) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func convertCurrency(r float64, c float64) float64 {
	return r * c
}

type DataType struct {
	Base  string
	Date  string
	Rates map[string]float64
}

func handleResponse(r io.ReadCloser) {
	var data DataType

	err := json.NewDecoder(r).Decode(&data)
	if err != nil {
		fmt.Println("Error reading data", err)
		os.Exit(3)
	}

	fmt.Printf("\nBase currency: %s, Date: %s\n\n", data.Base, data.Date)

	if convert > 1 {
		fmt.Printf("| %-12s|-%-12s|-%-12s |\n", "------------", "------------", "------------")
		fmt.Printf("| %-12s| %-12s| %-12s |\n", "Symbol", "Rate", "Converted")
		fmt.Printf("| %-12s|-%-12s|-%-12s |\n", "------------", "------------", "------------")
	} else {
		fmt.Printf("| %-12s|-%-12s |\n", "------------", "------------")
		fmt.Printf("| %-12s| %-12s |\n", "Symbol", "Rate")
		fmt.Printf("| %-12s|-%-12s |\n", "------------", "------------")
	}


	keys := sortDataKeys(data.Rates)
	for _, k := range keys {
		if convert > 1 {
			fmt.Printf("| %-12s| %-12.3f| %-12.3f |\n", k, data.Rates[k], convertCurrency(data.Rates[k], convert))
		} else {
			fmt.Printf("| %-12s| %-12.3f |\n", k, data.Rates[k])
		}
	}

	if convert > 1 {
		fmt.Printf("| %-12s|-%-12s|-%-12s |\n", "------------", "------------", "------------")
	} else {
		fmt.Printf("| %-12s|-%-12s |\n", "------------", "------------")
	}
}

func main() {
	sendRequest()
}