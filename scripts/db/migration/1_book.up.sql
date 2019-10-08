CREATE TABLE books (
    id serial PRIMARY KEY,
    title VARCHAR (255) NOT NULL,
    author VARCHAR (255) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);