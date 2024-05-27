package main

import (
	"encoding/xml"
	"io"
	"net/http"
)

type CurrencyIndex struct {
	Currencies []Currency `xml:"Cube>Cube>Cube"`
}

type Currency struct {
	Name string  `xml:"currency,attr"`
	Rate float64 `xml:"rate,attr"`
}

type CurrencyJSON struct {
	Name string  `json:"currency,attr"`
	Rate float64 `json:"rate,attr"`
}

func ExtractQuotes(url string) CurrencyIndex {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var s CurrencyIndex
	xml.Unmarshal(body, &s)
	return s
}
