package user

import (
	"database/sql"
	"fmt"

	"auth-service/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.First,
		&user.Last,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM \"user\" WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == "" {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) GetUserByID(id string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM \"user\" WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == "" {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	sqlStatement := `
		INSERT INTO "user" (first, last, email, password, role)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := s.db.Exec(
		sqlStatement,
		user.First,
		user.Last,
		user.Email,
		user.Password,
		user.Role,
	)
	if err != nil {
		return err
	}
	// TODO return INSERT result
	return nil
}
