package main

import (
	"log"
	"net/http"
	"github.com/harsh/mongo-golang/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	// Create a new router
	router := httprouter.New()

	// Initialize UserController with a MongoDB session
	uc := controllers.NewUserController(getSession())

	// Define routes
	router.GET("/user/:id", uc.GetUser)     // Get a user by ID
	router.POST("/user", uc.CreateUser)    // Create a new user
	router.DELETE("/user/:id", uc.DeleteUser) // Delete a user by ID

	// Start the server
	log.Println("Server running on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

// Connect to MongoDB
func getSession() *mgo.Session {
	// Replace `mongodb://localhost:27017` with your MongoDB connection string if needed
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	return session
}
