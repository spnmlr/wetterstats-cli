package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type daten struct {
	FfUhrzeit   string `json:"ff_uhrzeit"`
	FfTempA     string `json:"ff_temp_a"`
	FfFeuchteA  string `json:"ff_feuchte_a"`
	FfLuftdruck string `json:"ff_luftdruck"`
	FfWind      string `json:"ff_wind"`
	FfRegenAkt  string `json:"ff_regen_akt"`
	FfTaupunkt  string `json:"ff_taupunkt"`
	FfSolar     string `json:"ff_solar"`
}

func main() {

	location := flag.String("endpoint", "", "a string")
	flag.Parse()

	resp, err := http.Get("https://www.wetterstats.de/echtzeitwerte/include/daten_temp/" + *location + ".php")
	if err != nil {
		log.Println("wetterstats - ", err)
		return
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {

		defer resp.Body.Close()

		Body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("wetterstats - ", err)
			return
		}

		var d daten

		json.Unmarshal(Body, &d)
		if err != nil {
			log.Println("wetterstats - ", err)
			return
		}

		fmt.Println("Uhrzeit: " + d.FfUhrzeit)
		fmt.Println("Temperatur: " + d.FfTempA)
		fmt.Println("Windstärke: " + d.FfWind)
		fmt.Println("Sonneneinstrahlung: " + d.FfSolar)
		fmt.Println("Luftfeuchtigkeit: " + d.FfFeuchteA)
		fmt.Println("Luftdruck: " + d.FfLuftdruck)
		fmt.Println("Taupunkt: " + d.FfTaupunkt)
		fmt.Println("Niederschlag: " + d.FfRegenAkt)
	} else {
		log.Println("wetterstats - HTTP Status ", resp.StatusCode, " - Webservice error for wetterstats.de")
	}
}
