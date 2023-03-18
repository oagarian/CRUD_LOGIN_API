# GOlANG API CRUD

## How to run:
```sh
  cd src
  
  go mod tidy
<<<<<<< HEAD

  go run *.go 
  (or go run ./)
=======
  
  cd src
	
  go run main.go
>>>>>>> 6b18087c34b9f098dc1914c22f3bac4a0e6f880f
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
