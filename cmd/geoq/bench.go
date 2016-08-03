package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/akhenakh/regionagogo"
)

var (
	duration       = time.Second * 5
	latMin, latMax = 35.0, 37.0
	lngMin, lngMax = 138.0, 141.0
)

func dummyQuery(lat, lng float64) map[string]string {
	return nil
}

func genLatLng() (lat, lng float64) {
	lat = latMin + (latMax-latMin)*rand.Float64()
	lng = lngMin + (lngMax-lngMin)*rand.Float64()
	return lat, lng
}

func bench0(gs *regionagogo.GeoSearch) int {
	n := 0
	t := time.After(duration)
	for {
		select {
		case <-t:
			return n
		default:
			lat, lng := genLatLng()
			dummyQuery(lat, lng)
			n++
		}
	}
}

func bench1(gs *regionagogo.GeoSearch) int {
	n := 0
	t := time.After(duration)
	for {
		select {
		case <-t:
			return n
		default:
			lat, lng := genLatLng()
			gs.Query(lat, lng)
			n++
		}
	}
}

func runBench(gs *regionagogo.GeoSearch) {
	fmt.Println("executing bench0...")
	n0 := bench0(gs)
	fmt.Println("executing bench1...")
	n1 := bench1(gs)
	fmt.Printf("control: %d (%f QPS)\n", n0, float64(n0)/5.0)
	fmt.Printf("+query:  %d (%f QPS)\n", n1, float64(n1)/5.0)
}
