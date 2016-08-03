package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/akhenakh/regionagogo"
	"github.com/peterh/liner"
)

func loop(gs *regionagogo.GeoSearch) {
	line := liner.NewLiner()
	defer line.Close()
	for {
		l, err := line.Prompt("Lat, Lng: ")
		if err != nil {
			if err == io.EOF {
				fmt.Println("EXIT")
				break
			}
			fmt.Println("ERROR: ", err)
			continue
		}
		if l == "exit" {
			break
		}
		var lat, lng float64
		_, err = fmt.Sscanf(l, "%f,%f", &lat, &lng)
		if err != nil {
			fmt.Println("ERROR: ", err)
			continue
		}
		ans := gs.Query(lat, lng)
		fmt.Printf("%f, %f -> %v\n", lat, lng, ans)
	}
}

var fields = []string{
	"N03_001",
	"N03_003",
	"N03_004",
	"N03_007",
}

func main() {
	geojson := flag.String("geojson", "", "geojson to import")
	bench := flag.Bool("bench", false, "benchmark mode")
	flag.Parse()
	if *geojson == "" {
		log.Fatal("require a -geojson {file}")
	}
	gs := regionagogo.NewGeoSearch()
	err := gs.ImportGeoJSON(*geojson, fields)
	if err != nil {
		log.Fatal("failed to import geojson: ", err)
	}
	if *bench {
		runBench(gs)
		return
	}
	loop(gs)
}
