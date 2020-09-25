
-- CREATE DATABASE BLOG_DEV
-- GO

USE BLOG_DEV;
-- DROP TABLE blog_article
-- DROP TABLE blog_auth
-- DROP TABLE blog_tag

CREATE TABLE blog_article (
  [id] int NOT NULL PRIMARY KEY IDENTITY(1,1),
  [tag_id] int DEFAULT '0',
  [title] varchar(100) DEFAULT '',
  [desc] varchar(255) DEFAULT '',
  [content] text,
  [cover_image_url] varchar(255) DEFAULT '',
  [created_on] int DEFAULT '0',
  [created_by] varchar(100) DEFAULT '',
  [modified_on] int DEFAULT '0',
  [modified_by] varchar(255) DEFAULT '',
  [deleted_on] int DEFAULT '0',
  [state] tinyint DEFAULT '1',
)
GO

CREATE TABLE blog_auth (
  [id] int NOT NULL PRIMARY KEY IDENTITY(1,1),
  [username] varchar(50) DEFAULT '',
  [password] varchar(50) DEFAULT ''
)
GO

INSERT INTO blog_auth (username, password) VALUES ('admin', 'admin');
GO

CREATE TABLE blog_tag (
  [id] int NOT NULL PRIMARY KEY IDENTITY(1,1),
  [name] varchar(100) DEFAULT '',
  [created_on] int DEFAULT '0',
  [created_by] varchar(100) DEFAULT '',
  [modified_on] int DEFAULT '0',
  [modified_by] varchar(100) DEFAULT '',
  [deleted_on] int DEFAULT '0',
  [state] tinyint DEFAULT '1'
);

