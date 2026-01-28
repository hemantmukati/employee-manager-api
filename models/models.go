package models

type Employee struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	JobTitle string `json:"job_title"`
	Country  string `json:"country"`
	Salary   int64  `json:"salary"`
}

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SalaryMetrics struct {
	MinSalary int64   `json:"min_salary"`
	MaxSalary int64   `json:"max_salary"`
	AvgSalary float64 `json:"avg_salary"`
	Count     int64   `json:"total_employees"`
}
