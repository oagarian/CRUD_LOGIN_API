# GOlANG API CRUD

## How to run:
```sh
  	git clone https://github.com/oagarian/CRUD_LOGIN_API.git
	cd CRUD_LOGIN_API
	
  	go mod tidy
  	cd src
  	go run *.go 
```

### The API server will start on: 
http://localhost:8080

### MySQL database code:
```sh
CREATE DATABASE CRUD_LOGIN;
USE CRUD_LOGIN;

CREATE TABLE USERS(
	USERNAME VARCHAR(50) NOT NULL,
    EMAIL VARCHAR(30) NOT NULL PRIMARY KEY,
    USER_PASSWORD VARCHAR(18) NOT NULL
);
```
