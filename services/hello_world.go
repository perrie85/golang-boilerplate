package services

import (
	"net/url"
)

func HelloWorld(params url.Values) []interface{} {
	var processedParams []interface{}

	for j, k := range params {
		processedParams = append(processedParams, j, k)
	}

	return processedParams
}
