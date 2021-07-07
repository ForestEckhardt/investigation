package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/go-enry/go-license-detector/v4/licensedb"
	"github.com/go-enry/go-license-detector/v4/licensedb/filer"
)

type License struct {
	License    string  `json:"license"`
	Confidence float32 `json:"confidence"`
}

type Confidence struct {
	Name     string    `json:"name"`
	Licenses []License `json:"licenses"`
}

func main() {
	dirEntries, err := os.ReadDir("output")
	if err != nil {
		log.Fatal(err)
	}

	var confidence []Confidence

	for _, de := range dirEntries {
		filer, err := filer.FromDirectory(filepath.Join("output", de.Name()))
		if err != nil {
			log.Fatal(err)
		}

		licenses, err := licensedb.Detect(filer)
		if err != nil {
			log.Fatal(err)
		}

		var c Confidence
		c.Name = de.Name()

		for key, val := range licenses {
			c.Licenses = append(c.Licenses, License{License: key, Confidence: val.Confidence})
		}
		confidence = append(confidence, c)
	}

	file, err := os.Create("results.json")
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(confidence)
	if err != nil {
		log.Fatal(err)
	}
}
