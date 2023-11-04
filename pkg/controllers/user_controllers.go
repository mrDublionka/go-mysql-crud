package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mrDublionka/go-mysql-crud/pkg/models"
	"github.com/mrDublionka/go-mysql-crud/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

const SecretKey = "ion"

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	if err != nil {
		// Handle the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	UserToCreate := &models.User{}
	utils.ParseBody(r, UserToCreate)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(UserToCreate.UserPwd), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("Error hashing the password: %v\n", err)
		http.Error(w, "Error hashing the password", http.StatusInternalServerError)
		return
	}

	UserToCreate.UserPwd = string(hashedPassword)

	existingUser, _ := models.GetUserByEmail(UserToCreate.UserEmail)
	if existingUser != nil {
		log.Printf("User with the same email already exists\n")
		http.Error(w, "User with the same email already exists", http.StatusBadRequest)
		return
	}

	b := UserToCreate.CreateUser()
	if b == nil {
		log.Printf("Error creating user\n")
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	credentials := struct {
		UserEmail string `json:"user_email"`
		UserPwd   string `json:"user_pwd"`
	}{}

	utils.ParseBody(r, &credentials)

	user, _ := models.GetUserByEmail(credentials.UserEmail)

	if user == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.UserPwd), []byte(credentials.UserPwd))

	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 1 hour

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Printf("Error generating JWT token: %v\n", err)
		http.Error(w, "Error generating JWT token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"token": tokenString,
	}

	res, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetUserFromToken(r *http.Request) (*models.User, error) {
	tokenString := extractTokenFromRequest(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid token claims")
	}

	userID, ok := claims["user_id"].(float64) // Parse as float64

	if !ok {
		return nil, fmt.Errorf("Invalid user ID in token")
	}

	// Convert user ID to an integer
	userIDInt := int(userID)

	// Fetch the user from your data store or database
	user, err := models.GetUserByID(userIDInt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func extractTokenFromRequest(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	return tokenString
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	userJSON, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
