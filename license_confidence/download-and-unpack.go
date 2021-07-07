package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/paketo-buildpacks/packit/vacation"
)

type Dep struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

func main() {
	deps := []string{"bundler", "curl", "dotnet-aspnetcore", "dotnet-runtime", "dotnet-sdk", "go", "httpd", "icu", "nginx", "node", "php", "pip", "pipenv", "python", "ruby", "rust", "tini", "yarn"}

	for _, dep := range deps {
		resp, err := http.Get(fmt.Sprintf("https://api.deps.paketo.io/v1/dependency?name=%s", dep))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var d []Dep
		err = json.NewDecoder(resp.Body).Decode(&d)
		if err != nil {
			log.Fatal(err)
		}

		r, err := http.Get(d[0].Source)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()

		path := filepath.Join("output", dep)
		switch dep {
		case "bundler", "dotnet-aspnetcore", "dotnet-runtime", "dotnet-sdk":
			err = vacation.NewArchive(r.Body).Decompress(path)
			if err != nil {
				log.Fatal(fmt.Printf("%s failed", dep))
			}
		default:
			err = vacation.NewArchive(r.Body).StripComponents(1).Decompress(path)
			if err != nil {
				log.Fatal(fmt.Printf("%s failed", dep))
			}
		}

	}

}
