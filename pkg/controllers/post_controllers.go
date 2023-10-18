package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mrDublionka/go-mysql-crud/pkg/models"
	"github.com/mrDublionka/go-mysql-crud/pkg/utils"
	"net/http"
	"strconv"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	CreatePost := &models.Post{}
	utils.ParseBody(r, CreatePost)
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
	vars := mux.Vars(r)
	postId := vars["postId"]
	ID, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	postDetails, _ := models.GetPostById(ID)
	res, _ := json.Marshal(postDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var updatePost = &models.Post{}
	utils.ParseBody(r, updatePost)
	vars := mux.Vars(r)
	postId := vars["postId"]
	ID, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	postDetails, db := models.GetPostById(ID)
	if updatePost.Title != "" {
		postDetails.Title = updatePost.Title
	}

	if updatePost.Content != "" {
		postDetails.Content = updatePost.Content
	}

	db.Save(&postDetails)
	res, _ := json.Marshal(postDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	ID, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	post := models.DeletePost(ID)
	res, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
