CREATE TABLE USERS(
	USERNAME VARCHAR(50) NOT NULL,
    EMAIL VARCHAR(30) NOT NULL PRIMARY KEY,
    IS_ADMIN BOOLEAN NOT NULL,
    USER_PASSWORD VARCHAR(18) NOT NULL,
    CREATED_AT VARCHAR(12) NOT NULL
);