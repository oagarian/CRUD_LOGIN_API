// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import ()

type User struct {
	Username     string
	Email        string
	IsAdmin      bool
	UserPassword string
	CreatedAt    string
}
