package controllers

import (
	"employee-management-api/helpers"
	"employee-management-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var employeeHelper *helpers.EmployeeHelper

// Initialize DB for controllers
func InitDB(db *helpers.PGManager) {
	employeeHelper = helpers.NewEmployeeHelper(db)
}

func CreateEmployee(c *gin.Context) {
	var emp models.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := employeeHelper.Create(&emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create employee"})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func GetEmployee(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid employee id",
		})
		return
	}

	emp, err := employeeHelper.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "employee not found",
		})
		return
	}

	c.JSON(http.StatusOK, emp)
}

func ListEmployees(c *gin.Context) {
	employees, err := employeeHelper.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch employees",
		})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func UpdateEmployee(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
		return
	}

	var emp models.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emp.ID = id // 🔥 important

	updated, err := employeeHelper.Update(&emp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update employee"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func DeleteEmployee(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
		return
	}

	err = employeeHelper.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "employee deleted successfully",
	})
}
