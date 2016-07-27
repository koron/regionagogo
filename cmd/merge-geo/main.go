package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/kpawlik/geojson"
)

func loadFC(name string) (*geojson.FeatureCollection, error) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	c := new(geojson.FeatureCollection)
	err = json.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func saveFC(name string, c *geojson.FeatureCollection) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(c)
}

func main() {
	outfile := flag.String("out", "", "output filename")
	flag.Parse()
	if len(*outfile) == 0 {
		log.Fatal("-out {NAME} is required")
	}
	whole := geojson.NewFeatureCollection(nil)
	for _, infile := range flag.Args() {
		log.Printf("loading %s...", infile)
		c, err := loadFC(infile)
		if err != nil {
			log.Fatalf("failed to load feature collection: %s", err)
		}
		whole.AddFeatures(c.Features...)
	}
	log.Printf("writing %s...", *outfile)
	if err := saveFC(*outfile, whole); err != nil {
		log.Fatalf("failed to save feature collection: %s", err)
	}
}
