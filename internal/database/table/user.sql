CREATE TABLE "User" (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    user_email VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    create_at TIMESTAMP,
    delete_at TIMESTAMP,
    active BOOLEAN
);