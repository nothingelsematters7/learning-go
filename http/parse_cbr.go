package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
)

// ValCurs represents list of Valutes
type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

// Valute represent currency with its rate and nominal
type Valute struct {
	XMLName  xml.Name `xml:"Valute"`
	NumCode  string   `xml:"NumCode"`
	CharCode string   `xml:"CharCode"`
	Nominal  string
	Value    string
}

func (v Valute) String() string {
	return fmt.Sprintf("%4s: %s - %s\n", v.CharCode, v.Value, v.Nominal)
}

func main() {
	response, err := http.Get("http://cbr.ru/scripts/XML_daily.asp")
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(body)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	valCurs := ValCurs{}
	err = decoder.Decode(&valCurs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(valCurs)
}
