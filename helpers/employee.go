package helpers

import (
	"employee-management-api/models"
	"errors"
)

type EmployeeHelper struct {
	DB *PGManager // Ensure PGManager is imported from helpers package
}

func NewEmployeeHelper(db *PGManager) *EmployeeHelper {
	return &EmployeeHelper{DB: db}
}

func (h *EmployeeHelper) Create(e *models.Employee) (*models.Employee, error) {
	query := `
		INSERT INTO employees (full_name, job_title, country, salary)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err := h.DB.DB.QueryRow(query, e.FullName, e.JobTitle, e.Country, e.Salary).Scan(&e.ID)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (h *EmployeeHelper) GetByID(id int64) (*models.Employee, error) {
	query := `
		SELECT id, full_name, job_title, country, salary
		FROM employees
		WHERE id = $1
	`

	var emp models.Employee

	err := h.DB.DB.QueryRow(query, id).Scan(
		&emp.ID,
		&emp.FullName,
		&emp.JobTitle,
		&emp.Country,
		&emp.Salary,
	)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (h *EmployeeHelper) GetAll() ([]models.Employee, error) {
	query := `
		SELECT id, full_name, job_title, country, salary
		FROM employees
		ORDER BY id DESC
	`

	rows, err := h.DB.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(
			&emp.ID,
			&emp.FullName,
			&emp.JobTitle,
			&emp.Country,
			&emp.Salary,
		)
		if err != nil {
			return nil, err
		}

		employees = append(employees, emp)
	}

	return employees, nil
}

func (h *EmployeeHelper) Update(e *models.Employee) (*models.Employee, error) {
	query := `
		UPDATE employees
		SET full_name = $1,
		    job_title = $2,
		    country = $3,
		    salary = $4
		WHERE id = $5
	`

	result, err := h.DB.DB.Exec(
		query,
		e.FullName,
		e.JobTitle,
		e.Country,
		e.Salary,
		e.ID,
	)

	if err != nil {
		return nil, err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return nil, errors.New("no employee found")
	}

	return e, nil
}

func (h *EmployeeHelper) Delete(id int64) error {
	query := `DELETE FROM employees WHERE id = $1`

	result, err := h.DB.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("no employee found")
	}

	return nil
}

func (h *EmployeeHelper) GetSalaryMetrics() (*models.SalaryMetrics, error) {
	query := `
		SELECT
		COALESCE(MIN(salary), 0),
		COALESCE(MAX(salary), 0),
		COALESCE(AVG(salary), 0),
		COUNT(*)
		FROM employees
		`

	row := h.DB.DB.QueryRow(query)

	var metrics models.SalaryMetrics
	err := row.Scan(
		&metrics.MinSalary,
		&metrics.MaxSalary,
		&metrics.AvgSalary,
		&metrics.Count,
	)

	if err != nil {
		return nil, err
	}

	return &metrics, nil
}

// GetSalaryByCountry returns min, max, avg salary for a specific country
func (h *EmployeeHelper) GetSalaryByCountry(country string) (*models.SalaryMetrics, error) {
	query := `
		SELECT 
			COALESCE(MIN(salary), 0),
			COALESCE(MAX(salary), 0),
			COALESCE(AVG(salary), 0),
			COUNT(*)
		FROM employees
		WHERE country = $1
	`

	row := h.DB.DB.QueryRow(query, country)

	var metrics models.SalaryMetrics
	err := row.Scan(
		&metrics.MinSalary,
		&metrics.MaxSalary,
		&metrics.AvgSalary,
		&metrics.Count,
	)
	if err != nil {
		return nil, err
	}
	return &metrics, nil
}

// GetAvgSalaryByJob returns average salary for a given job title
func (h *EmployeeHelper) GetAvgSalaryByJob(jobTitle string) (float64, error) {
	query := `
		SELECT COALESCE(AVG(salary), 0)
		FROM employees
		WHERE job_title = $1
	`

	var avg float64
	err := h.DB.DB.QueryRow(query, jobTitle).Scan(&avg)
	if err != nil {
		return 0, err
	}
	return avg, nil
}
