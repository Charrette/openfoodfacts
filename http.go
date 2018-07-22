package openfoodfacts

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://ssl-api.openfoodfacts.org/api/v0"
)

func GetProduct(code string) (*GetProductResponse, error) {
	url := fmt.Sprintf("%s/product/%s.json", baseURL, code)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	getProductResponse := &GetProductResponse{}
	if err := json.NewDecoder(resp.Body).Decode(getProductResponse); err != nil {
		return nil, err
	}

	return getProductResponse, nil
}
