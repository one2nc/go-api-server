package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	connectDB()
	defer db.Close()

	// Set up the router
	router := gin.Default()

	// Define API routes
	router.POST("/students", addStudent)          // Add a new student
	router.GET("/students", getAllStudents)       // Get all students
	router.PUT("/students/:id", editStudent)      // Edit an existing student
	router.DELETE("/students/:id", deleteStudent) // Delete a student

	// Start the server on port 8080
	router.Run(":8080")
}
