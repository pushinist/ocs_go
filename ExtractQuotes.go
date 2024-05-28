package main

import (
	"encoding/xml"
	"io"
	"net/http"
)

type xmlCurrencyIndex struct {
	Currencies []xmlCurrency `xml:"Cube>Cube>Cube"`
}

type xmlCurrency struct {
	Name string  `xml:"currency,attr"`
	Rate float64 `xml:"rate,attr"`
}

type rate struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func ExtractQuotes(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var s xmlCurrencyIndex
	xml.Unmarshal(body, &s)

	for _, currency := range s.Currencies {
		rates = append(rates, rate{currency.Name, currency.Rate})
	}
}
