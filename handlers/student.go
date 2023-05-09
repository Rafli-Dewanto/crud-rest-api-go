package handlers

import (
	"net/http"
	"strconv"

	"github.com/Rafli-Dewanto/go-rest/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentHandler struct {
	db *gorm.DB
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{db: db}
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var newStudent models.Student

	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := h.db.Create(&newStudent); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newStudent})
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var students []models.Student
	if result := h.db.Offset(offset).Limit(limit).Find(&students); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": students})
}

func (h *StudentHandler) GetStudentById(c *gin.Context) {
	studentIdParam := c.Param("id")
	id, err := strconv.Atoi(studentIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "fail when parsing student id"})
		return
	}
	var student models.Student;
	if result := h.db.Where("id = ?", id).First(&student); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": student})
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	var updatedStudent models.Student

	if err := c.ShouldBindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var existingStudent models.Student
	if result := h.db.First(&existingStudent, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	existingStudent.Name = updatedStudent.Name
	existingStudent.Age = updatedStudent.Age
	existingStudent.Address = updatedStudent.Address
	existingStudent.PhoneNumber = updatedStudent.PhoneNumber

	if result := h.db.Save(&existingStudent); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": existingStudent})
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var existingStudent models.Student
	if result := h.db.First(&existingStudent, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	if result := h.db.Delete(&existingStudent); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
