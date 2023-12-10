package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mrDublionka/go-mysql-crud/pkg/controllers"
	"github.com/mrDublionka/go-mysql-crud/pkg/models"
	"github.com/mrDublionka/go-mysql-crud/pkg/routes"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	imageController := controllers.NewImageController()
	imageController.InitFirebaseStorage()

	corsOptions := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
	}

	// Use cors.New to create a cors.Cors instance with the provided options
	c := cors.New(corsOptions)

	routes.RegisterBlogRoutes(r)

	// Wrap your router with the CORS middleware
	handler := c.Handler(r)

	http.Handle("/", handler)
	dsn := "root:@tcp(127.0.0.1:3306)/go_crud?parseTime=true"
	models.InitDB(dsn)

	log.Fatal(http.ListenAndServe("localhost:9010", handler))
}
