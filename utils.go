package geo_search

import (
	"encoding/json"
	"net/http"
)

func extractData(resp *http.Response) (*ServerData, error) {
	var data ServerData
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
