package routes

import (
	"github.com/gorilla/mux"
	"github.com/mrDublionka/go-mysql-crud/pkg/controllers"
)

// adsasd
var RegisterBlogRoutes = func(router *mux.Router) {
	imageController := controllers.NewImageController()
	//POST ROUTES START
	router.HandleFunc("/posts/", imageController.CreatePost).Methods("POST")
	router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{postId}", controllers.GetPostById).Methods("GET")
	router.HandleFunc("/posts/{postId}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{postId}", controllers.DeletePost).Methods("DELETE")
	router.HandleFunc("/posts/like", controllers.LikePost).Methods("POST")

	//USER ROUTES START
	router.HandleFunc("/user/signin", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/user/profile", controllers.GetUserProfile).Methods("POST")
	router.HandleFunc("/posts-test", controllers.TestPost).Methods("GET")
	//USER ROUTES END

	router.HandleFunc("/upload-image", imageController.UploadImage).Methods("POST")

	//HEALTH
	router.HandleFunc("/health", controllers.Health).Methods("GET")

}
