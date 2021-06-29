package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/alexstan12/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
// 	yaml := `
// - path: /urlshort
//   url: https://github.com/gophercises/urlshort
// - path: /urlshort-final
//   url: https://github.com/gophercises/urlshort/tree/solution
// `
var fileName string
flag.StringVar(&fileName, "fileName", "src/github.com/alexstan12/urlshort/paths.json", "paths values location")
 //pathsFile, err := os.OpenFile(fileName, os.O_RDWR, 0755)
 flag.Parse()
 pathsFile, err := ioutil.ReadFile(fileName)
 if err != nil {
	 fmt.Println(err)
 }else {
	 fmt.Println("Paths file successfully opened!")
 }
 
 
 	// yamlValues, err := yaml.Marshal(pathsFile)
	//  if err !=nil {
	// 	 fmt.Println("Could not Marshal pathsFile")
	//  }
	if strings.HasSuffix(fileName, "yaml"){
		yamlHandler, err := urlshort.YAMLHandler(pathsFile, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", yamlHandler)
	} else if strings.HasSuffix(fileName, "json"){
		jsonHandler, err := urlshort.JSONHandler(pathsFile, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", jsonHandler)
	}

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
