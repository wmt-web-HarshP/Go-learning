package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/harsh/mongo-golang/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

// NewUserController creates a new UserController with a MongoDB session
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser retrieves a user by ID
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Check if ID is a valid ObjectId
	if !bson.IsObjectIdHex(id) {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	oid := bson.ObjectIdHex(id)

	u := models.User{}

	// Fetch the user from the database
	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Marshal user data to JSON and send response
	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uj)
}

// CreateUser creates a new user in the database
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	// Decode the request body into the User struct
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate a new ObjectId for the user
	u.Id = bson.NewObjectId()

	// Insert the user into the database
	if err := uc.session.DB("go-web-dev-db").C("users").Insert(u); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Marshal user data to JSON and send response
	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(uj)
}

// DeleteUser deletes a user by ID
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Check if ID is a valid ObjectId
	if !bson.IsObjectIdHex(id) {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	oid := bson.ObjectIdHex(id)

	// Delete the user from the database
	if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User deleted successfully")
}
