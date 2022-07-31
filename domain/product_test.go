package domain

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestProduct(t *testing.T) {
	id := 12
	stock := 10
	name := "dress"
	body := "{\"id\":" + strconv.Itoa(id) + ",\"stock\":" + strconv.Itoa(stock) + ",\"name\":\"" + name + "\"}"
	product := Product{}
	err := json.Unmarshal([]byte(body), &product)

	if *product.ID != int64(id) {
		t.Errorf("Product: ID do not match, got = %v, want = %v", *product.ID, id)
		return
	}

	if int(*product.Stock) != stock {
		t.Errorf("Product: stock do not match, got = %v, want = %v", product.Stock, stock)
		return
	}

	if *product.Name != name {
		t.Errorf("Product: Name do not match, got = %v, want = %v", product.Name, name)
		return
	}

	if err != nil {
		t.Errorf("Product Struct has an error = %v", err)
		return
	}
}
