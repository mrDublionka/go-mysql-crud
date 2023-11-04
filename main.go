package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mrDublionka/go-mysql-crud/pkg/models"
	"github.com/mrDublionka/go-mysql-crud/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBlogRoutes(r)
	http.Handle("/", r)
	dsn := "root:@tcp(127.0.0.1:3306)/go_crud?parseTime=true"
	models.InitDB(dsn)

	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
