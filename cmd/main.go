package main

import (
	"flag"
	"fmt"
	"net/http"

	urlshort "github.com/abaksy/gourlshort/pkg"
)

func main() {
	inputFile := flag.String("file", "test/test.yml", "Input filename to read data from")
	inputFormat := flag.String("format", "yaml", "Input format (currently supported: json,yaml)")
	flag.Parse()

	if *inputFormat != "yaml" && *inputFormat != "json" {
		panic("Supplied format must be 'json' or 'yaml'!")
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	mapHandler := urlshort.MapHandler(urlshort.DefaultURLMap, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	data, err := urlshort.ReadDataFile(*inputFile)
	if err != nil {
		panic(err)
	}

	handler := mapHandler
	switch *inputFormat {
	case "yaml":
		fmt.Println("Using YAML Handler")
		handler, err = urlshort.YAMLHandler(data, mapHandler)
		if err != nil {
			panic(err)
		}
	case "json":
		fmt.Println("Using JSON Handler")
		handler, err = urlshort.JSONHandler(data, mapHandler)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("Starting the server on %v\n", urlshort.Port)
	http.ListenAndServe(urlshort.Port, handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
