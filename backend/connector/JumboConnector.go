package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

const jumboAPIVersion = "v17"

type JumboConnector struct{}

func (jc *JumboConnector) SearchProducts(query string, page, size int) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://mobileapi.jumbo.com/%s/search", jumboAPIVersion)
	params := fmt.Sprintf("?offset=%d&limit=%d&q=%s", page*size, size, query)
	resp, err := http.Get(url + params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (jc *JumboConnector) GetProductDetails(productID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://mobileapi.jumbo.com/%s/products/%s", jumboAPIVersion, productID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (jc *JumboConnector) GetCategories() (map[string]interface{}, error) {
	url := fmt.Sprintf("https://mobileapi.jumbo.com/%s/categories", jumboAPIVersion)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (jc *JumboConnector) GetSubCategories(categoryID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://mobileapi.jumbo.com/%s/categories?id=%s", jumboAPIVersion, categoryID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (jc *JumboConnector) GetStores() (map[string]interface{}, error) {
	url := fmt.Sprintf("https://mobileapi.jumbo.com/%s/stores", jumboAPIVersion)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (jc *JumboConnector) GetPromotions() (map[string]interface{}, error) {
	url := fmt.Sprintf("https://mobileapi.jumbo.com/%s/promotion-overview", jumboAPIVersion)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (jc *JumboConnector) GetStorePromotions(storeID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://mobileapi.jumbo.com/%s/promotion-overview?store_id=%s", jumboAPIVersion, storeID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (jc *JumboConnector) SearchAllProducts(query string, size int) ([]map[string]interface{}, error) {
	var allProducts []map[string]interface{}
	page := 0
	for {
		products, err := jc.SearchProducts(query, page, size)
		if err != nil {
			return nil, err
		}
		data, ok := products["products"].(map[string]interface{})["data"].([]map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("missing data field in response")
		}
		allProducts = append(allProducts, data...)
		total, ok := products["products"].(map[string]interface{})["total"].(float64)
		if !ok {
			return nil, fmt.Errorf("missing total field in response")
		}
		totalPages := int(math.Ceil(total / float64(size)))
		if page+1 >= totalPages {
			break
		}
		page++
	}
	return allProducts, nil
}
