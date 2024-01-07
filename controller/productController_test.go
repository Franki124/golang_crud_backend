package controller

import (
	"awesomeProject/model"
	"bytes"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllProductsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCreateProduct(t *testing.T) {
	product := model.Product{Name: "New Product", Description: "New Description", Price: 100}
	productJSON, _ := json.Marshal(product)
	req, err := http.NewRequest("POST", "/product/create", bytes.NewBuffer(productJSON))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProductHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestUpdateProduct(t *testing.T) {
	objID, err := primitive.ObjectIDFromHex("507f191e810c19729de860ea")
	if err != nil {
		t.Fatal(err)
	}

	updatedProduct := model.Product{
		ID:          objID,
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       150,
	}

	productJSON, _ := json.Marshal(updatedProduct)
	req, err := http.NewRequest("PUT", "/product/update", bytes.NewBuffer(productJSON))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateProductHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteProduct(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/product/delete?id=some-product-id", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteProductHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
