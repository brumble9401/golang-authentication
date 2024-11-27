package user

import (
	"database/sql"

	"github.com/brumble9401/golang-authentication/types"
	"github.com/gofiber/fiber/v2/log"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:db,
	}
}

func (s *Store) GetUserByEmailOrUsername(email string, username string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE username = $1 AND email = $2", username, email)
	if err != nil {
		log.Error("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()
	user := new(types.User)
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Email)
		if err != nil {
			log.Error("Error scanning row: %v", err)
			return nil, err
		}
    // Process user
	}
	if err = rows.Err(); err != nil {
		log.Error("Error with row iteration: %v", err)
		return nil, err
	}

	log.Info("No user found with username %s and email %s", username, email)
	return nil, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (userName, email, password_hash, full_name) VALUES ($1, $2, $3, $4, $5)", user.Username, user.Email, user.PasswordHash, user.FullName)
	if err != nil {
		return err
	}

	return nil
}