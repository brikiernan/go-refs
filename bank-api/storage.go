package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) (*Account, error)
	DeleteAccount(int) (string, error)
	UpdateAccount(*Account) (*Account, error)
	ListAccounts() ([]*Account, error)
	RetrieveAccount(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres password=gobank123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := ` create table if not exists account (
		id serial primary key,
		first_name varchar(96),
		last_name varchar(96),
		number serial,
		balance decimal(20, 2),
		created_at timestamp,
		updated_at timestamp
	); 
	
	create sequence if not exists acct_num_sequence
		start 100001
		increment 7;
	`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) (*Account, error) {
	query := `
	insert into account (first_name, last_name, number, balance, created_at, updated_at)
	values ($1, $2, nextval('acct_num_sequence'), $3, $4, $5)
	returning *
	`

	row := s.db.QueryRow(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Balance,
		acc.CreatedAt,
		acc.UpdatedAt,
	)

	account := &Account{}
	if err := row.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return account, nil
}

func (s *PostgresStore) DeleteAccount(int) (string, error) {
	return "", nil
}

func (s *PostgresStore) UpdateAccount(*Account) (*Account, error) {
	return &Account{}, nil
}

func (s *PostgresStore) ListAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account := &Account{}
		if err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
			&account.UpdatedAt,
		); err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil

}

func (s *PostgresStore) RetrieveAccount(int) (*Account, error) {
	return nil, nil
}
