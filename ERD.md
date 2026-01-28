# Entity Relationship Diagram (ERD)

This document represents the database schema for the Employee Management API.

## users

Stores application users for authentication (JWT-based).

| Column        | Type       |   Constraints          |
|---------------|------------|------------------------|
| id            | SERIAL     | Primary Key            |
| email         | VARCHAR    | UNIQUE, NOT NULL       |
| password_hash | TEXT       | NOT NULL               |
| created_at    | TIMESTAMP  | Default NOW()          |


## employees

Stores employee-related information.

| Column     | Type       |  Constraints     |
|------------|------------|------------------|
| id         | SERIAL     | Primary Key      |
| full_name  | VARCHAR    | NOT NULL         |
| job_title  | VARCHAR    | NOT NULL         |
| country    | VARCHAR    | NOT NULL         |
| salary     | BIGINT     | NOT NULL         |
| created_at | TIMESTAMP  | Default NOW()    |


## Relationships

- There is **no direct foreign key relationship** between `users` and `employees`.
- Authentication (`users`) and business data (`employees`) are kept separate by design.


## Notes

- Schema is designed for PostgreSQL.
- Managed via versioned SQL migrations.
- ERD aligns exactly with application models and migrations.