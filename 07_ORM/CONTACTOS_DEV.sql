CREATE DATABASE CONTACTOS_DEV
GO

USE CONTACTOS_DEV;
GO

-- DROP TABLE Alumno_modulos
-- DROP TABLE contacts
-- DROP TABLE Modulo
-- DROP TABLE Alumno
-- DROP TABLE profesor
-- sp_help contacts

INSERT INTO Alumno(Delegado_Id,Nombre,Apellido,FechaNacimiento)
VALUES 
(null,'Joel','Aguilar','2020-09-23 07:22:23.913'),
(null,'Andree','Galdos','2020-09-23 07:22:23.913'),
(1,'Gabriel','Meet','2020-09-23 07:22:23.913'),
(1,'Maria','Parra','2020-09-23 07:22:23.913'),
(1,'Diego','Linder','2020-09-23 07:22:23.913'),
(2,'Dorliska','Prim','2020-09-23 07:22:23.913'),
(2,'Jhonny','García','2020-09-23 07:22:23.913'),
(2,'Jhon','Doe','2020-09-23 07:22:23.913'),
(2,'Saúl','Arias','2020-09-23 07:22:23.913'),
(2,'Simón','Laura','2020-09-23 07:22:23.913');
GO

INSERT INTO Modulo(nombre)
VALUES
('Modulo 1'),
('Modulo 2'),
('Modulo 3'),
('Modulo 4'),
('Modulo 5');
GO

INSERT INTO alumno_modulos (alumno_id)
VALUES 
(1),
(2),
(3),
(4);
GO

INSERT INTO Profesor (rfc, nombres, direccion, telefono,Modulo_Codigo)
VALUES
(1,'Juan Santa Cruz','Av. tomas Marsano','79799779',1),
(2,'Omar Makensye','Av. Perú','10292029',1);

SELECT * FROM contacts
INSERT INTO contacts(nombre,edad,telefono, direccion, email, descripcion)
values 
('Julian Rengifo',70,'029929292','Av Pucalpa','julian@hotmail.com','Cliente');

-- CREATE TABLE books(
--     id INT IDENTITY(1,1) PRIMARY KEY NOT NULL,
--     title VARCHAR(100) NOT NULL,
--     author VARCHAR(50) NOT NULL,
--     year CHAR(4) NOt NULL
-- );
-- GO

-- INSERT INTO books(title, author, year) VALUES
-- ('Libro 1', 'Joel', 2020),
-- ('Libro 2', 'Chris', 1990),
-- ('Libro 3', 'Dante', 1980),
-- ('Libro 4', 'Viviana', 1950),
-- ('Libro 5', 'Lean', 2001);

-- SELECT * FROM Books
