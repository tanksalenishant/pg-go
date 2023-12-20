package repository

import (
	"context"
	"rest1/internal/domain"

	"github.com/jackc/pgx/v4"
	// type UserRepository struct {
	// 	Storage *[]domain.User
	// }
)

type UserRepo struct {
	conn *pgx.Conn
}

func NewUserRepo(conn *pgx.Conn) *UserRepo {
	return &UserRepo{conn: conn}
}

func (u *UserRepo) GetAll() ([]domain.User, error) {
	rows, err := u.conn.Query(context.Background(), "SELECT id, name, accountNo, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.AccountNo, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepo) GetByID(id string) (*domain.User, error) {
	var user domain.User
	err := u.conn.QueryRow(context.Background(), "SELECT id, name, accountNo, password FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.AccountNo, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) Create(user *domain.User) (*domain.User, error) {
	err := u.conn.QueryRow(context.Background(),
		"INSERT INTO users(name, accountNo, password) VALUES($1, $2, $3) RETURNING id",
		user.Name, user.AccountNo, user.Password).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) Withdraw(user *domain.User) error {
	// Implement withdrawal logic here
	return nil
}

func (u *UserRepo) Deposit(user *domain.User) error {
	// Implement deposit logic here
	return nil
}