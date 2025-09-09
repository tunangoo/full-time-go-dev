package model

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"users"`
	ID            int64     `json:"id" bun:"id,pk,autoincrement"`
	FirstName     string    `json:"first_name" bun:"first_name"`
	LastName      string    `json:"last_name" bun:"last_name"`
	Email         string    `json:"email" bun:"email"`
	Password      string    `json:"-" bun:"password"`
	Role          string    `json:"role" bun:"role"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bun:"updated_at"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=255"`
	LastName  string `json:"last_name" binding:"required,min=3,max=255"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=255"`
	Role      string `json:"role" binding:"required,oneof=admin user"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=255"`
	LastName  string `json:"last_name" binding:"required,min=3,max=255"`
	Email     string `json:"email" binding:"required,email"`
	Role      string `json:"role" binding:"required,oneof=admin user"`
}
