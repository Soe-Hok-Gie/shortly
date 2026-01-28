package repository

import (
	"context"
	"database/sql"
	"log"
	"shortly/model/domain"

	"github.com/go-sql-driver/mysql"
)

type userRepositoryImp struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) UserRepository {
	return &userRepositoryImp{DB: DB}
}

func (repository *userRepositoryImp) Register(ctx context.Context, user domain.User) (domain.User, error) {
	script := "INSERT INTO users (username,password) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, user.Username, user.Password)
	if err != nil {
		if IsDuplicateKeyError(err) {
			return user, err
		}
		return user, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	user.Id = id
	return user, nil
}

// Helper untuk cek duplicate key MySQL
func IsDuplicateKeyError(err error) bool {
	if me, ok := err.(*mysql.MySQLError); ok {
		return me.Number == 1062
	}
	return false
}

// login
func (repository *userRepositoryImp) Login(ctx context.Context, username string) (*domain.User, error) {
	user := &domain.User{}
	script := "SELECT id,username, password FROM users WHERE username = ?"
	row := repository.DB.QueryRowContext(ctx, script, username)

	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		log.Println("user not found:", err)
		return nil, err
	}
	return user, nil

}
