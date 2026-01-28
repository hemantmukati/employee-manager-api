CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    job_title VARCHAR(255) NOT NULL,
    country VARCHAR(100) NOT NULL,
    salary BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);