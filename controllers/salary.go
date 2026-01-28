package controllers

import (
	"employee-management-api/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEmployeeSalary(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
		return
	}

	emp, err := employeeHelper.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}

	deduction, net := helpers.CalculateSalary(emp.Country, emp.Salary)

	c.JSON(http.StatusOK, gin.H{
		"employee_id":  emp.ID,
		"country":      emp.Country,
		"gross_salary": emp.Salary,
		"deduction":    deduction,
		"net_salary":   net,
	})
}

func GetSalaryMetrics(c *gin.Context) {
	metrics, err := employeeHelper.GetSalaryMetrics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch salary metrics",
		})
		return
	}

	c.JSON(http.StatusOK, metrics)
}

func GetSalaryByCountry(c *gin.Context) {
	country := c.Param("country")
	metrics, err := employeeHelper.GetSalaryByCountry(country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get metrics"})
		return
	}

	c.JSON(http.StatusOK, metrics)
}

func GetAvgSalaryByJob(c *gin.Context) {
	jobTitle := c.Param("jobtitle")
	avg, err := employeeHelper.GetAvgSalaryByJob(jobTitle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get metrics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"job_title":  jobTitle,
		"avg_salary": avg,
	})
}
