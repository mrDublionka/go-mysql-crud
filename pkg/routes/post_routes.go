package routes

import (
	"github.com/gorilla/mux"
	"github.com/mrDublionka/go-mysql-crud/pkg/controllers"
)

// adsasd
var RegisterBlogRoutes = func(router *mux.Router) {
	//POST ROUTES START
	router.HandleFunc("/posts/", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{postId}", controllers.GetPostById).Methods("GET")
	router.HandleFunc("/posts/{postId}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{postId}", controllers.DeletePost).Methods("DELETE")
	//POST ROUTES END

	//USER ROUTES START
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/user/profile", controllers.GetUserProfile).Methods("POST")

	//USER ROUTES END
}
