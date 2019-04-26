package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func ensureProductTableExists() {
	if _, err := a.DB.Exec(productTableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearProductTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_product_id_seq RESTART WITH 1")
}

const productTableCreationQuery = `create table if not exists test_db.public.products
(
  product_id   serial      not null
    constraint products_pk
      primary key,
  product_name varchar(45) not null,
  price        integer     not null,
  category     varchar(50),
  product_url  varchar(200),
  is_active    boolean default true
);
`

func TestEmptyTable(t *testing.T) {
	clearProductTable()

	req, _ := http.NewRequest("GET", "/products", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentProduct(t *testing.T) {
	clearProductTable()

	req, _ := http.NewRequest("GET", "/products/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "sql: no rows in result set" {
		t.Errorf("Expected the 'error' key of the response to be set to 'sql: no rows in result set'. Got '%s'", m["error"])
	}
}
func TestCreateProduct(t *testing.T) {
	clearProductTable()

	payload := []byte(`{
	"product_name":"product",
	"price": 123,
	"category": "test_cat",
	"product_url": "www.test.com"
	}`)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["product_name"] != "product" {
		t.Errorf("Expected prouduct name to be 'product'. Got '%v'", m["product_name"])
	}

	if m["price"] != 123 {
		t.Errorf("Expected product price to be '123'. Got '%v'", m["price"])
	}

	if m["category"] != "test_cat" {
		t.Errorf("Expected product category to be 'test_cat'. Got '%v'", m["category"])
	}

	if m["product_url"] != "www.test.com" {
		t.Errorf("Expected product url to be 'www.test.com'. Got '%v'", m["product_url"])
	}
}

func TestGetProduct(t *testing.T) {
	clearProductTable()
	addProducts(1)

	req, _ := http.NewRequest("GET", "/products/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}
func TestUpdateProduct(t *testing.T) {
	clearProductTable()
	addProducts(1)

	req, _ := http.NewRequest("GET", "/products/1", nil)
	response := executeRequest(req)
	var originalProduct map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalProduct)

	payload := []byte(`{"product_name":"new name"}`)

	req, _ = http.NewRequest("PUT", "/products/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["product_id"] != originalProduct["product_id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalProduct["product_id"], m["product_id"])
	}

	if m["product_name"] == originalProduct["product_name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalProduct["product_name"], m["product_name"], m["product_name"])
	}
}

func addProducts(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		statement := fmt.Sprintf(`
		INSERT INTO products (
			product_name, price, category, product_url
		) VALUES (
			%s, %s, %s, %s
		)`, "'test'", "123", "'test_category'", "'www.test.com'")
		a.DB.Exec(statement)
	}
}
