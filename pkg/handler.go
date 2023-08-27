package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type url struct {
	Path string
	URL  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if destURL, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, destURL, http.StatusMovedPermanently)
		}
		fallback.ServeHTTP(w, r)
	})
}

// YAMLHandler will parse the providedYAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func parseYAML(yml []byte) (map[string]string, error) {

	urls := make([]url, 0, 10)
	urlMap := make(map[string]string, 10)
	err := yaml.Unmarshal(yml, &urls)
	if err != nil {
		return nil, err
	}

	for _, element := range urls {
		urlMap[element.Path] = element.URL
	}

	return urlMap, nil
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	urlMap, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	return MapHandler(urlMap, fallback), nil

}
