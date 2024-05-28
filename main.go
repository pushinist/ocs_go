package main

const URL string = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

func main() {
	ExtractQuotes(URL)
	routerInit()
}
