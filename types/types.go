package types

import (
	"time"

	"github.com/google/uuid"
)

type UserStore interface {
	GetUserByEmailOrUsername(email string, username string) (*User, error)
	CreateUser(user User) error
}

type User struct {
	UserID      uuid.UUID  `json:"user_id" db:"user_id"`        // UUID primary key
	Username    string     `json:"username" db:"username"`      // Unique username
	PasswordHash string    `json:"password_hash" db:"password_hash"` // Hashed password
	Email       string     `json:"email" db:"email"`            // Unique email
	FullName    string     `json:"full_name,omitempty" db:"full_name"` // Optional full name
	RoleID      uuid.UUID  `json:"role_id" db:"role_id"`        // Foreign key for role (UUID)
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`  // Creation timestamp
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`  // Update timestamp
}

type RegisterUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	FullName string `json:"fullName"`
}