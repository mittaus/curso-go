-- CREATE DATABASE LIBRERIA_DEV
-- GO

USE LIBRERIA_DEV;
GO


CREATE TABLE books(
    id INT IDENTITY(1,1) PRIMARY KEY NOT NULL,
    title VARCHAR(100) NOT NULL,
    author VARCHAR(50) NOT NULL,
    year CHAR(4) NOt NULL
);
GO

INSERT INTO books(title, author, year) VALUES
('Libro 1', 'Joel', 2020),
('Libro 2', 'Chris', 1990),
('Libro 3', 'Dante', 1980),
('Libro 4', 'Viviana', 1950),
('Libro 5', 'Lean', 2001);

SELECT * FROM Books