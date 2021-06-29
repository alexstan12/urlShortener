package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-yaml/yaml"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(rw http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if element,ok := pathsToUrls[path]; ok {
			http.Redirect(rw, r, element, http.StatusFound)
			return 
		} 
		fallback.ServeHTTP(rw, r)
		
	}
	}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	pd, err := parseYAML(yml)
	if err!= nil {
			return nil, err
	}
	
	return MapHandler(buildMap(pd),fallback), nil
	

}

func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error){

	pd, err := parseJson(jsonData)
	if err != nil {
		return nil, err
	}

	return MapHandler(buildMap(pd), fallback), nil
	
}

func buildMap(pd []pathData) map[string]string{
	pathsToUrls := make(map[string]string)
	for _, pathUnit := range pd {
			pathsToUrls[pathUnit.Path] = pathUnit.Url
	}
	return pathsToUrls
}

func parseYAML(yml []byte) ([]pathData,error){
	var pd []pathData 
	fmt.Println(string(yml))
	err := yaml.Unmarshal(yml, &pd)
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func parseJson(jsonData []byte)([]pathData, error){
	var pd []pathData
	err := json.Unmarshal(jsonData, &pd)
	if err != nil {
		return nil, err
	}
	return pd, nil
}

type pathData struct {
	 Path string `yaml:"path"`
	 Url string `yaml:"url"`
}
