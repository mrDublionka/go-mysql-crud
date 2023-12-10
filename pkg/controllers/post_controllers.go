package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mrDublionka/go-mysql-crud/pkg/config"
	"github.com/mrDublionka/go-mysql-crud/pkg/models"
	"github.com/mrDublionka/go-mysql-crud/pkg/utils"
	"io"
	"net/http"
	"strconv"
)

var db = config.GetDB()

func (ic *ImageControllers) CreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Extract form values
	title := r.FormValue("title")
	content := r.FormValue("content")
	topic := r.FormValue("topic")
	userIDString := r.FormValue("userID")
	userID, err := strconv.ParseUint(userIDString, 10, 32)

	CreatePost := &models.Post{}
	utils.ParseBody(r, CreatePost)
	println(CreatePost)

	file, handler, err := r.FormFile("image")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	defer file.Close()

	bucketName := "blog-next-php.appspot.com"
	objectName := handler.Filename

	wc := ic.storage.Bucket(bucketName).Object(objectName).NewWriter(ic.ctx)
	_, err = io.Copy(wc, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := wc.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	CreatePost.Title = title
	CreatePost.Content = content
	CreatePost.Topic = topic
	CreatePost.UserID = uint(userID)
	CreatePost.Image = "https://storage.googleapis.com/" + bucketName + "/" + objectName
	CreatePost.ImageName = objectName

	b := CreatePost.CreatePost()

	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllPosts()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	println(2)
	vars := mux.Vars(r)
	postId := vars["postId"]
	ID, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := models.GetPostById(ID)
	if err != nil {
		fmt.Println("Error while fetching post:", err)
		http.Error(w, "Error while fetching post", http.StatusInternalServerError)
		return
	}

	if post == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	println("hello 1")
	CreateLike := &models.Like{}
	// Attempt to retrieve the user from the request token.
	user, err := GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	println("hello 2")
	postId := r.URL.Query().Get("id")
	userId := fmt.Sprint(user.ID) // Assuming user.ID is the authenticated user's ID.

	if postId == "" || userId == "" {
		http.Error(w, "Post ID and User ID are required", http.StatusBadRequest)
		return
	}
	println("hello 3")
	// Parse the post ID to an unsigned integer.
	ID, err := strconv.ParseInt(postId, 0, 0)

	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	// Parse the user ID to an unsigned integer.
	userIDUint, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if !CreateLike.CheckIfIsLiked(uint(ID), uint(userIDUint)) {
		// The like does not exist, so create a new Like entry.
		newLike := models.Like{
			PostID: uint(ID),
			UserID: uint(userIDUint),
		}
		// Use the CreateLike method to create the Like record in the database.
		createdLike := newLike.CreateLike()
		if createdLike == nil {
			// If creating the like failed, respond with an internal server error.
			http.Error(w, "Could not create like", http.StatusInternalServerError)
			return
		}

		// Prepare the response with the new like's data.
		response := map[string]uint{
			"post_id": createdLike.PostID,
			"user_id": createdLike.UserID,
		}

		// Marshal the response map to JSON.
		res, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
			return
		}

		// Set the response header to 'Content-Type: application/json'.
		w.Header().Set("Content-Type", "application/json")
		// Write the response with status code 201 Created.
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	} else {
		// If the like already exists, respond with a status indicating no new resource was created.
		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	//var updatePost = &models.Post{}
	//utils.ParseBody(r, updatePost)
	//vars := mux.Vars(r)
	//postId := vars["postId"]
	//ID, err := strconv.ParseInt(postId, 0, 0)
	//if err != nil {
	//	fmt.Println("error while parsing")
	//}
	//postDetails, db := models.GetPostById(ID)
	//if updatePost.Title != "" {
	//	postDetails.Title = updatePost.Title
	//}
	//
	//if updatePost.Content != "" {
	//	postDetails.Content = updatePost.Content
	//}
	//
	//db.Save(&postDetails)
	//res, _ := json.Marshal(postDetails)
	//w.Header().Set("Content-Type", "pkglication/json")
	//w.WriteHeader(http.StatusOK)
	//w.Write(res)

}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	println(1)
	vars := mux.Vars(r)
	postId := vars["postId"]
	ID, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	// Pass the "db" instance as the first argument to DeletePost
	post := models.DeletePost(db, ID)

	res, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func TestPost(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"test": "Hello",
	}

	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Write the response with status code 201 Created.
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
