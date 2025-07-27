-- PostgreSQL + ENUM => We have to create a custom type first
CREATE TYPE employment_status AS ENUM('employed', 'self-employed', 'unemployed');


CREATE TABLE users (
    first_name VARCHAR(200),
    yearlys_salary INT,
    current_status employment_status
    -- current_status ENUM('employed', 'self-employed', 'unemployed') -- MySQL syntax,

);