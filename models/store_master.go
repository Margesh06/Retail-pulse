package models

import (
	"encoding/json"
	"os"
)

type StoreMaster struct {
	StoreID   string `json:"store_id"`
	StoreName string `json:"store_name"`
	AreaCode  string `json:"area_code"`
}

func LoadStoreMaster(file string) ([]StoreMaster, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var stores []StoreMaster
	err = json.Unmarshal(data, &stores)
	return stores, err
}
