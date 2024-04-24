package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addStudent(c *gin.Context) {
	var newStudent Student
	if err := c.BindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO students (name, age, grade) VALUES ($1, $2, $3) RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, newStudent.Name, newStudent.Age, newStudent.Grade).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func getAllStudents(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Age, &s.Grade); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		students = append(students, s)
	}
	c.JSON(http.StatusOK, students)
}

func editStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updStudent Student
	if err := c.BindJSON(&updStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `UPDATE students SET name = $1, age = $2, grade = $3 WHERE id = $4`
	_, err := db.Exec(sqlStatement, updStudent.Name, updStudent.Age, updStudent.Grade, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func deleteStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := db.Exec("DELETE FROM students WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
