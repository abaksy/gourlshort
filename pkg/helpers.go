package urlshort

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type URL struct {
	Path string
	URL  string
}

// Read data from a YAML file and return it as a byte slice
func ReadDataFile(fileName string) ([]byte, error) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func ParseData(data []byte, format string) (map[string]string, error) {

	urls := make([]URL, 0, 10)
	urlMap := make(map[string]string, 10)

	switch format {
	case "json":
		err := json.Unmarshal(data, &urls)
		if err != nil {
			return nil, err
		}

	case "yaml":
		err := yaml.Unmarshal(data, &urls)
		if err != nil {
			return nil, err
		}
	default:
		fmt.Println("Error! Invalid format supplied!")
		return nil, nil
	}

	for _, element := range urls {
		urlMap[element.Path] = element.URL
	}

	return urlMap, nil
}
