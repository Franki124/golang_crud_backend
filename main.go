package main

import (
	"awesomeProject/controller"
	"context"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	client, err := controller.ConnectToDB()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer client.Disconnect(context.TODO())

	controller.SetClient(client)

	http.HandleFunc("/product/create", controller.CreateProductHandler)
	http.HandleFunc("/product/update", controller.UpdateProductHandler)
	http.HandleFunc("/product/delete", controller.DeleteProductHandler)
	http.HandleFunc("/products", controller.GetAllProductsHandler)

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins in development
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(http.DefaultServeMux)

	log.Println("Starting server on localhost:8080")
	err = http.ListenAndServe("localhost:8080", handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
