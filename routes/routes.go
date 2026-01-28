package routes

import (
	"employee-management-api/controllers"
	"employee-management-api/helpers"
	"employee-management-api/middlewares"

	"github.com/gin-gonic/gin"
)

func StartGin() {
	r := gin.Default()

	// Connect DB and initialize controllers
	db := helpers.ConnectDB()
	controllers.InitDB(db)
	controllers.InitAuth(db)

	api := r.Group("/api/v1")
	{
		api.POST("/auth/register", controllers.Register)
		api.POST("/auth/login", controllers.Login)

		protected := api.Group("/")
		protected.Use(middlewares.JWTAuth())
		{
			// Employees CRUD
			protected.POST("/employees", controllers.CreateEmployee)
			protected.GET("/employees/:id", controllers.GetEmployee)
			protected.PUT("/employees/:id", controllers.UpdateEmployee)
			protected.DELETE("/employees/:id", controllers.DeleteEmployee)

			//Salary calculation
			protected.GET("/employees/:id/salary", controllers.GetEmployeeSalary)
			protected.GET("/employees/salary/metrics", controllers.GetSalaryMetrics)
			protected.GET("/employees/salary/country/:country", controllers.GetSalaryByCountry)
			protected.GET("/employees/salary/job/:jobtitle", controllers.GetAvgSalaryByJob)
		}
	}

	r.Run(":8080")
}
